package main

import (
	"flag"
	"log"
	"os"
	"strconv"
	"text/template"

	phylotaxymusic_models "github.com/thomaspeugeot/phylotaxymusic/go/models"

	phylotaxymusic_stack "github.com/thomaspeugeot/phylotaxymusic/go/stack"
	phylotaxymusic_static "github.com/thomaspeugeot/phylotaxymusic/go/static"

	gongsvg_models "github.com/fullstack-lang/gongsvg/go/models"
	gongsvg_stack "github.com/fullstack-lang/gongsvg/go/stack"

	gongtone_models "github.com/fullstack-lang/gongtone/go/models"
	gongtone_stack "github.com/fullstack-lang/gongtone/go/stack"

	gongtree_stack "github.com/fullstack-lang/gongtree/go/stack"
)

var (
	logGINFlag = flag.Bool("logGIN", false, "log mode for gin")

	unmarshallFromCode = flag.String("unmarshallFromCode", "", "unmarshall data from go file and '.go' (must be lowercased without spaces), If unmarshallFromCode arg is '', no unmarshalling")
	marshallOnCommit   = flag.String("marshallOnCommit", "", "on all commits, marshall staged data to a go file with the marshall name and '.go' (must be lowercased without spaces). If marshall arg is '', no marshalling")

	diagrams         = flag.Bool("diagrams", true, "parse/analysis go/models and go/diagrams")
	embeddedDiagrams = flag.Bool("embeddedDiagrams", false, "parse/analysis go/models and go/embeddedDiagrams")

	port = flag.Int("port", 8080, "port server")
)

func main() {

	log.SetPrefix("phylotaxymusic: ")
	log.SetFlags(0)

	// parse program arguments
	flag.Parse()

	// setup the static file server and get the controller
	r := phylotaxymusic_static.ServeStaticFiles(*logGINFlag)

	// setup phylotaxymusicStack
	phylotaxymusicStack := phylotaxymusic_stack.NewStack(r,
		phylotaxymusic_models.Phylotaxy.ToString(), *unmarshallFromCode, *marshallOnCommit, "", *embeddedDiagrams, true)
	phylotaxymusicStack.Probe.Refresh()
	phylotaxymusicStack.Stage.Checkout()

	gongsvg_stack := gongsvg_stack.NewStack(r, phylotaxymusic_models.GongsvgStackName.ToString(), "", "", "", true, true)
	gongtree_stack := gongtree_stack.NewStack(r, phylotaxymusic_models.SidebarTree.ToString(), "", "", "", true, true)
	gongtone_stack := gongtone_stack.NewStack(r, phylotaxymusic_models.GongtoneStackName.ToString(), "", "", "", true, true)

	// get the only diagram
	parameters := phylotaxymusic_models.GetGongstructInstancesMap[phylotaxymusic_models.Parameter](phylotaxymusicStack.Stage)

	if len(*parameters) == 0 {
		log.Println("")
		// log.Fatalln("")
	}

	f4 := new(gongtone_models.Freqency).Stage(gongtone_stack.Stage)
	f4.Name = "F4"

	notef4 := new(gongtone_models.Note).Stage(gongtone_stack.Stage)
	notef4.Frequencies = append(notef4.Frequencies, f4)
	notef4.Start = 0
	notef4.Duration = 1
	notef4.Velocity = 1

	gongtone_stack.Stage.Commit()

	parameter := (*parameters)["Reference"]

	tree := new(phylotaxymusic_models.Tree)
	tree.TreeStack = gongtree_stack
	tree.Stage = phylotaxymusicStack.Stage

	parameterImpl := new(ParameterImpl)
	parameterImpl.parameter = parameter
	parameterImpl.gongsvgStage = gongsvg_stack.Stage
	parameterImpl.gongtoneStage = gongtone_stack.Stage
	parameterImpl.phylotaxymusicStage = phylotaxymusicStack.Stage
	parameterImpl.tree = tree

	parameter.Impl = parameterImpl

	phylotaxymusic_models.GeneratorSingloton.Impl = parameterImpl
	parameterImpl.Generate()

	tree.Generate(parameter)

	log.Printf("Server ready serve on localhost:" + strconv.Itoa(*port))
	err := r.Run(":" + strconv.Itoa(*port))
	if err != nil {
		log.Fatalln(err.Error())
	}
}

type ParameterImpl struct {
	gongsvgStage        *gongsvg_models.StageStruct
	gongtoneStage       *gongtone_models.StageStruct
	phylotaxymusicStage *phylotaxymusic_models.StageStruct
	parameter           *phylotaxymusic_models.Parameter
	tree                *phylotaxymusic_models.Tree
}

// Generate implements models.GeneratorInterface.
func (parameterImpl *ParameterImpl) Generate() {
	p := parameterImpl.parameter

	p.ComputeShapes(parameterImpl.phylotaxymusicStage)
	p.GenerateSvg(parameterImpl.gongsvgStage)
	p.GenerateNotes(parameterImpl.gongtoneStage)
	parameterImpl.tree.Generate(p)
	parameterImpl.phylotaxymusicStage.Commit()

	// get path of the hour handle
	parameterImpl.generateHTMLindex(p)
}

func (parameterImpl *ParameterImpl) generateHTMLindex(p *phylotaxymusic_models.Parameter) {
	mapFrontCurveStack := *phylotaxymusic_models.GetGongstructInstancesMap[phylotaxymusic_models.FrontCurveStack](
		parameterImpl.phylotaxymusicStage)

	hourHandleStack := mapFrontCurveStack["Hour Handle"]
	hourHandlePath := hourHandleStack.FrontCurves[0]

	minuteHandleStack := mapFrontCurveStack["Minute Handle"]
	minuteHandlePath := minuteHandleStack.FrontCurves[0]

	BackendHandleStack := mapFrontCurveStack["Backend curve"]
	BackendHandlePath := BackendHandleStack.FrontCurves[0]

	mapSpiralCircle := *phylotaxymusic_models.GetGongstructInstancesMap[phylotaxymusic_models.SpiralCircle](
		parameterImpl.phylotaxymusicStage)

	hourHandleMarker := mapSpiralCircle["Hour Marker"]
	minuteHandleMarker := mapSpiralCircle["Minute Marker"]
	backendHandleMarker := mapSpiralCircle["Backend Marker"]

	err := generateIndexHTML(
		minuteHandlePath.Path,
		minuteHandleMarker.Path,
		p.MinuteColor,

		hourHandlePath.Path,
		hourHandleMarker.Path,
		p.HourColor,

		BackendHandlePath.Path,
		backendHandleMarker.Path,
		p.BackendColor,
	)
	if err != nil {
		panic(err)
	}
}

func (parameterImpl *ParameterImpl) OnUpdated(updatedParameter *phylotaxymusic_models.Parameter) {

	log.Println("OnUpdated", parameterImpl.parameter.InsideAngle, parameterImpl.parameter.SideLength)
	// phylotaxymusic_svg.GenerateSvg(parameterImpl.gongsvgStage, parameterImpl.phylotaxymusicStage)

	updatedParameter.ComputeShapes(parameterImpl.phylotaxymusicStage)
	updatedParameter.GenerateSvg(parameterImpl.gongsvgStage)
	parameterImpl.tree.Generate(updatedParameter)
	updatedParameter.GenerateNotes(parameterImpl.gongtoneStage)
	parameterImpl.generateHTMLindex(updatedParameter)

}

func generateIndexHTML(
	MinutePath,
	MinuteMarkerPath,
	MinuteColor,

	HourPath,
	HourMarkerPath,
	HourColor,

	BackendPath,
	BackendMarkerPath,
	BackendColor string) error {
	const htmlTemplate = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Responsive Interactive Clock SVG</title>
    <meta http-equiv="refresh" content="300">
    <style>
        body, html {
            margin: 0;
            padding: 0;
            width: 100%;
            height: 100%;
        }
        svg {
            width: 100%;
            height: calc(100% - 150px); /* Adjusted height to accommodate slider and larger buttons */
            display: block;
        }
        #slider-container {
            width: 100%;
            height: 50px;
            padding: 10px;
            box-sizing: border-box;
            background-color: #f0f0f0;
        }
        #time-slider {
            width: 100%;
        }
        /* Styles for the buttons and rank display */
        #button-container {
            width: 100%;
            padding: 10px;
            box-sizing: border-box;
            background-color: #e0e0e0;
            text-align: center;
        }
        .color-button {
            padding: 20px 40px; /* Doubled padding */
            font-size: 32px;     /* Doubled font size */
            margin: 0 10px;      /* Added margin between buttons */
        }
        .color-rank {
            font-size: 32px;     /* Match font size with buttons */
            margin: 0 10px;
            vertical-align: middle;
            display: inline-block;
        }
    </style>
</head>
<body>
    <!-- Responsive SVG -->
    <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 1000 1000" preserveAspectRatio="xMidYMid meet">

    <path id="background" d="
        {{.BackendPath}}
        " 
        fill="{{.BackendColor}}" 
        />

	<path id="background" d="
        {{.BackendMarkerPath}}
        " 
        fill="red" 
        />
        
      <path id="minute-handle" d="
      {{.MinutePath}}
      {{.MinuteMarkerPath}}
      " 
      fill="{{.MinuteColor}}" 
      fill-rule="evenodd"/>

      <path id="hour-handle" d="
      {{.HourPath}}
      {{.HourMarkerPath}}
      " fill="{{.HourColor}}" fill-rule="evenodd" />


    </svg>

    <!-- Slider Container -->
    <div id="slider-container">
        <input type="range" min="0" max="12" step="0.01" value="0" id="time-slider">
    </div>

    <!-- Button Container -->
    <div id="button-container">
        <button id="color-button-backward" class="color-button">Previous Color</button>
        <span id="color-rank" class="color-rank">1 / 7</span>
        <button id="color-button-forward" class="color-button">Next Color</button>
    </div>

    <!-- JavaScript code for rotation -->
    <script>
        (function() {
            var hourHandle = document.getElementById('hour-handle');
            var minuteHandle = document.getElementById('minute-handle');
            var timeSlider = document.getElementById('time-slider');

            function updateClock() {
                var sliderValue = parseFloat(timeSlider.value); // Hours added
                var now = new Date();
                var midnight = new Date(now.getFullYear(), now.getMonth(), now.getDate());
                var elapsedSeconds = (now - midnight) / 1000; // Seconds since midnight

                // Add slider value in seconds
                elapsedSeconds += sliderValue * 3600;

                var durationHour = 43200; // 12 hours in seconds (360 degrees)
                var durationMinute = 3600; // 1 hour in seconds (360 degrees)

                // Degrees per second
                var anglePerSecondHour = 360 / durationHour; 
                var anglePerSecondMinute = 360 / durationMinute; 

                // Calculate the current angle based on elapsed time
                var currentAngleHour = (elapsedSeconds % durationHour) * anglePerSecondHour;
                var currentAngleMinute = (elapsedSeconds % durationMinute) * anglePerSecondMinute;

                // Apply rotation using SVG transforms
                hourHandle.setAttribute('transform', 'rotate(' + currentAngleHour + ' 500 500)');
                minuteHandle.setAttribute('transform', 'rotate(' + currentAngleMinute + ' 500 500)');
            }

            var backgroundElement = document.getElementById('background');
            var colorButtonForward = document.getElementById('color-button-forward');
            var colorButtonBackward = document.getElementById('color-button-backward');
            var colorRankElement = document.getElementById('color-rank');

            var colorSets = [
    ['#D1C5B4', '#8FA382', '#536C87'],
    ['#F06A49', '#D6A539', '#006D6F'], // Original set
                ['#B0A18F', '#6F8D75', '#4E5E6A'], // Muted earth tones
                ['#E0D7C6', '#A8A59D', '#6B6E70'], // Soft neutrals with depth
                ['#C7B299', '#8E7C6D', '#5F4B3D'], // Warm, desaturated hues
    ['#0C5C60', '#A6591F', '#F08271'], // New set 3 - inverted roles
    ['#F1F1F1', '#5A5A5A', '#1E1E1E'], 
    ['#5A5A5A', '#F1F1F1', '#1E1E1E'], 
    ['#2B3A42', '#A8B7C7', '#E3E7E9'], // Darker background with clearer hour handle (blueish tones)
    ['#174A32', '#68A691', '#B9DEC1'], // Green variations
    ['#7D4964', '#B787A3', '#E7C2D4'], // Rose variations
	    ['red', 'white', 'blue'], // French colors
            ];

            var colorSetIndex = 0;

            function updateColors(delta) {
                colorSetIndex = (colorSetIndex + delta + colorSets.length) % colorSets.length;
                var colors = colorSets[colorSetIndex];
                backgroundElement.setAttribute('fill', colors[0]);
                minuteHandle.setAttribute('fill', colors[1]);
                hourHandle.setAttribute('fill', colors[2]);

                // Update the color rank display
                colorRankElement.textContent = (colorSetIndex + 1) + ' / ' + colorSets.length;
            }

            // Initialize clock and colors
            updateClock();
            updateColors(0);

            // Update clock when slider value changes
            timeSlider.addEventListener('input', updateClock);

            // Add event listeners to color buttons
            colorButtonForward.addEventListener('click', function() { updateColors(1); });
            colorButtonBackward.addEventListener('click', function() { updateColors(-1); });
        })();
    </script>
</body>
</html>
`

	// Create or open the index.html file
	file2, err := os.Create("../../../../myclock/index.html")
	if err != nil {
		return err
	}
	defer file2.Close()

	// Create or open the index.html file
	file, err := os.Create("index.html")
	if err != nil {
		return err
	}
	defer file.Close()

	// Data to be injected into the template
	data := struct {
		BackendPath       string
		BackendMarkerPath string
		BackendColor      string

		MinutePath       string
		MinuteMarkerPath string
		MinuteColor      string

		HourPath       string
		HourMarkerPath string
		HourColor      string
	}{
		BackendPath:       BackendPath,
		BackendMarkerPath: BackendMarkerPath,
		BackendColor:      BackendColor,

		MinutePath:       MinutePath,
		MinuteMarkerPath: MinuteMarkerPath,
		MinuteColor:      MinuteColor,

		HourPath:       HourPath,
		HourMarkerPath: HourMarkerPath,
		HourColor:      HourColor,
	}

	// Parse and execute the template with the provided data
	tmpl, err := template.New("index").Parse(htmlTemplate)
	if err != nil {
		return err
	}

	err = tmpl.Execute(file, data)
	if err != nil {
		return err
	}

	err = tmpl.Execute(file2, data)
	if err != nil {
		return err
	}

	return nil
}
