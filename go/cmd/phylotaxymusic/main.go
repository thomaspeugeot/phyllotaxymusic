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
	parameterImpl.parameter.ComputeShapes(parameterImpl.phylotaxymusicStage)
	parameterImpl.parameter.GenerateSvg(parameterImpl.gongsvgStage)
	parameterImpl.parameter.GenerateNotes(parameterImpl.gongtoneStage)
	parameterImpl.tree.Generate(parameterImpl.parameter)
	parameterImpl.phylotaxymusicStage.Commit()

	// get path of the hour handle
	mapFrontCurveStack := *phylotaxymusic_models.GetGongstructInstancesMap[phylotaxymusic_models.FrontCurveStack](
		parameterImpl.phylotaxymusicStage)

	hourHandleStack := mapFrontCurveStack["Hour Handle"]
	hourHandlePath := hourHandleStack.FrontCurves[0]

	minuteHandleStack := mapFrontCurveStack["Minute Handle"]
	minuteHandlePath := minuteHandleStack.FrontCurves[0]

	mapSpiralCircle := *phylotaxymusic_models.GetGongstructInstancesMap[phylotaxymusic_models.SpiralCircle](
		parameterImpl.phylotaxymusicStage)

	hourHandleMarker := mapSpiralCircle["Hour Marker"]
	minuteHandleMarker := mapSpiralCircle["Minute Marker"]

	err := generateIndexHTML(
		minuteHandlePath.Path,
		minuteHandleMarker.Path,
		hourHandlePath.Path,
		hourHandleMarker.Path,
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

}

func generateIndexHTML(string1a, string1b, string2a, string2b string) error {
	const htmlTemplate = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Responsive Interactive Clock SVG</title>
    <meta http-equiv="refresh" content="30">
    <style>
        body, html {
            margin: 0;
            padding: 0;
            width: 100%;
            height: 100%;
        }
        svg {
            width: 100%;
            height: calc(100% - 50px); /* Adjust height to accommodate slider */
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
    </style>
</head>
<body>
    <!-- Responsive SVG -->
    <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 1000 1000" preserveAspectRatio="xMidYMid meet">
      <path id="background" d="
        M 200,400
        C 80,400 0,280 0,180
        C 0,80 80,0 200,0
        C 320,0 400,80 400,180
        C 400,280 320,400 200,400
        Z
      " 
      fill="none" 
      fill-rule="evenodd"/>

      <path id="minute-handle" d="
      {{.String1a}}
      {{.String1b}}
      " 
      fill="lightblue" 
      fill-rule="evenodd"/>

      <path id="hour-handle" d="
      {{.String2a}}
      {{.String2b}}
      " fill="lightgreen" fill-rule="evenodd" />
    </svg>

    <!-- Slider Container -->
    <div id="slider-container">
        <input type="range" min="0" max="12" step="0.01" value="0" id="time-slider">
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

            // Initialize clock
            updateClock();

            // Update clock when slider value changes
            timeSlider.addEventListener('input', updateClock);
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
		String1a string
		String1b string
		String2a string
		String2b string
	}{
		String1a: string1a,
		String1b: string1b,
		String2a: string2a,
		String2b: string2b,
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
