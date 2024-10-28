package models

import (
	"os"
	"text/template"
)

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
            height: calc(100% - 300px); /* Adjusted height to accommodate larger slider and buttons */
            display: block;
        }
        #slider-container {
            width: 100%;
            height: 150px; /* Tripled the height */
            padding: 10px;
            box-sizing: border-box;
            background-color: #f0f0f0;
        }
        #time-slider {
            width: 100%;
            height: 100px; /* Adjusted height for the slider */
            -webkit-appearance: none; /* Remove default appearance */
            appearance: none;
            margin-top: 25px; /* Center the slider vertically */
        }
        /* Style for the slider track */
        #time-slider::-webkit-slider-runnable-track {
            width: 100%;
            height: 20px; /* Height of the track */
            background: #ddd;
            border-radius: 10px;
        }
        #time-slider::-moz-range-track {
            width: 100%;
            height: 20px;
            background: #ddd;
            border-radius: 10px;
        }
        /* Style for the slider thumb */
        #time-slider::-webkit-slider-thumb {
            -webkit-appearance: none;
            appearance: none;
            width: 40px;
            height: 40px;
            background: "grey";
            border-radius: 50%;
            cursor: pointer;
            margin-top: -10px; /* Align thumb with track */
        }
        #time-slider::-moz-range-thumb {
            width: 40px;
            height: 40px;
            background: "grey";
            border-radius: 50%;
            cursor: pointer;
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
            padding: 20px 40px;
            font-size: 32px;
            margin: 0 10px;
        }
        .color-rank {
            font-size: 32px;
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

func GenerateIndexHTML(
	filePath,
	MinutePath,
	MinuteMarkerPath,
	MinuteColor,

	HourPath,
	HourMarkerPath,
	HourColor,

	BackendPath,
	BackendMarkerPath,
	BackendColor string) error {

	// Create or open the index.html file
	file2, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file2.Close()

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

	err = tmpl.Execute(file2, data)
	if err != nil {
		return err
	}

	return nil
}
