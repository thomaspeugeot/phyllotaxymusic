package main

import (
	"flag"
	"log"
	"strconv"

	phylotaxymusic_models "github.com/thomaspeugeot/phylotaxymusic/go/models"
	phylotaxymusic_stack "github.com/thomaspeugeot/phylotaxymusic/go/stack"
	phylotaxymusic_static "github.com/thomaspeugeot/phylotaxymusic/go/static"

	gongsvg_models "github.com/fullstack-lang/gongsvg/go/models"
	gongsvg_stack "github.com/fullstack-lang/gongsvg/go/stack"
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

	// setup stack
	stack := phylotaxymusic_stack.NewStack(r, "phylotaxymusic", *unmarshallFromCode, *marshallOnCommit, "", *embeddedDiagrams, true)
	stack.Probe.Refresh()

	gongsvg_stack := gongsvg_stack.NewStack(r, phylotaxymusic_models.GongsvgStackName.ToString(), "", "", "", true, true)

	svg := (&gongsvg_models.SVG{Name: `SVG`}).Stage(gongsvg_stack.Stage)
	layer := new(gongsvg_models.Layer).Stage(gongsvg_stack.Stage)
	svg.Layers = append(svg.Layers, layer)

	circle := new(gongsvg_models.Circle).Stage(gongsvg_stack.Stage)
	layer.Circles = append(layer.Circles, circle)

	circle.Name = `C1a`
	circle.CX = 100.000000
	circle.CY = 100.000000
	circle.Radius = 50.000000
	circle.Color = `greenlight`
	circle.FillOpacity = 0.800000
	circle.Stroke = `blue`
	circle.StrokeOpacity = 0.000000
	circle.StrokeWidth = 1.000000
	circle.StrokeDashArray = ``
	circle.StrokeDashArrayWhenSelected = ``
	circle.Transform = ``

	gongsvg_stack.Stage.Commit()

	log.Printf("Server ready serve on localhost:" + strconv.Itoa(*port))
	err := r.Run(":" + strconv.Itoa(*port))
	if err != nil {
		log.Fatalln(err.Error())
	}
}
