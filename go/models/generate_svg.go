package models

import (
	"fmt"
	"os"
	"text/template"
)

const svgTemplate = `
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

`

func GenerateSVG(filePath string,
	MinutePath,
	MinuteMarkerPath,
	MinuteColor,

	HourPath,
	HourMarkerPath,
	HourColor,

	BackendPath,
	BackendMarkerPath,
	BackendColor string) error {

	// Create or open the file at the specified path
	file, err := os.Create(filePath)
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
	tmpl, err := template.New("index").Parse(svgTemplate)
	if err != nil {
		return err
	}

	err = tmpl.Execute(file, data)
	if err != nil {
		return err
	}

	fmt.Println("SVG file generated at:", filePath)
	return nil
}
