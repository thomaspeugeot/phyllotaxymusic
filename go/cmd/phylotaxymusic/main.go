package main

import (
	"flag"
	"log"
	"strconv"

	phylotaxymusic_models "github.com/thomaspeugeot/phylotaxymusic/go/models"
	phylotaxymusic_svg "github.com/thomaspeugeot/phylotaxymusic/go/svg"

	phylotaxymusic_stack "github.com/thomaspeugeot/phylotaxymusic/go/stack"
	phylotaxymusic_static "github.com/thomaspeugeot/phylotaxymusic/go/static"

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

	// setup phylotaxymusicStack
	phylotaxymusicStack := phylotaxymusic_stack.NewStack(r, "phylotaxymusic", *unmarshallFromCode, *marshallOnCommit, "", *embeddedDiagrams, true)
	phylotaxymusicStack.Probe.Refresh()
	phylotaxymusicStack.Stage.Checkout()

	gongsvg_stack := gongsvg_stack.NewStack(r, phylotaxymusic_models.GongsvgStackName.ToString(), "", "", "", true, true)

	phylotaxymusic_svg.GenerateSvg(gongsvg_stack.Stage, phylotaxymusicStack.Stage)

	log.Printf("Server ready serve on localhost:" + strconv.Itoa(*port))
	err := r.Run(":" + strconv.Itoa(*port))
	if err != nil {
		log.Fatalln(err.Error())
	}
}
