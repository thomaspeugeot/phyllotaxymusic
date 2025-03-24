package main

import (
	"flag"
	"log"
	"strconv"

	m "github.com/thomaspeugeot/phyllotaxymusic/go/models"
	phyllotaxymusic_stack "github.com/thomaspeugeot/phyllotaxymusic/go/stack"
	phyllotaxymusic_static "github.com/thomaspeugeot/phyllotaxymusic/go/static"
)

var (
	logGINFlag = flag.Bool("logGIN", false, "log mode for gin")

	unmarshallFromCode = flag.String("unmarshallFromCode", "", "unmarshall data from go file and '.go' (must be lowercased without spaces), If unmarshallFromCode arg is '', no unmarshalling")
	marshallOnCommit   = flag.String("marshallOnCommit", "", "on all commits, marshall staged data to a go file with the marshall name and '.go' (must be lowercased without spaces). If marshall arg is '', no marshalling")

	embeddedDiagrams = flag.Bool("embeddedDiagrams", false, "parse/analysis go/models and go/embeddedDiagrams")

	port = flag.Int("port", 8080, "port server")
)

func main() {

	log.SetPrefix("phyllotaxymusic: ")
	log.SetFlags(0)

	// parse program arguments
	flag.Parse()

	// setup the static file server and get the controller
	r := phyllotaxymusic_static.ServeStaticFiles(*logGINFlag)

	// setup phyllotaxymusic_stack
	phyllotaxymusic_stack := phyllotaxymusic_stack.NewStack(r,
		m.Phylotaxy.ToString(), *unmarshallFromCode, *marshallOnCommit, "", *embeddedDiagrams, true)
	phyllotaxymusic_stack.Probe.Refresh()
	phyllotaxymusic_stack.Stage.Checkout()

	m.NewStager(r, phyllotaxymusic_stack.Stage).UpdateAllStages()

	log.Printf("%s", "Server ready serve on localhost:"+strconv.Itoa(*port))
	err := r.Run(":" + strconv.Itoa(*port))
	if err != nil {
		log.Fatalln(err.Error())
	}

}
