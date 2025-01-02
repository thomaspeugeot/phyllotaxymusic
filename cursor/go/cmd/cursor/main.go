package main

import (
	"flag"
	"log"
	"strconv"

	cursor_models "github.com/thomaspeugeot/phyllotaxymusic/cursor/go/models"
	cursor_stack "github.com/thomaspeugeot/phyllotaxymusic/cursor/go/stack"
	cursor_static "github.com/thomaspeugeot/phyllotaxymusic/cursor/go/static"
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

	log.SetPrefix("cursor: ")
	log.SetFlags(0)

	// parse program arguments
	flag.Parse()

	// setup the static file server and get the controller
	r := cursor_static.ServeStaticFiles(*logGINFlag)

	// setup stack
	stack := cursor_stack.NewStack(r, cursor_models.Cursorstakcname.ToString(), *unmarshallFromCode, *marshallOnCommit, "", *embeddedDiagrams, true)
	stack.Probe.Refresh()

	log.Printf("Server ready serve on localhost:" + strconv.Itoa(*port))
	err := r.Run(":" + strconv.Itoa(*port))
	if err != nil {
		log.Fatalln(err.Error())
	}
}
