package main

import (
	"flag"
	"log"
	"strconv"

	// insertion point for models import
	phyllotaxymusic_level1stack "github.com/thomaspeugeot/phyllotaxymusic/go/level1stack"
	m "github.com/thomaspeugeot/phyllotaxymusic/go/models"
	// static
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
	log.SetFlags(log.Lmicroseconds)

	// parse program arguments
	flag.Parse()

	// setup phyllotaxymusic_stack
	stack := phyllotaxymusic_level1stack.NewLevel1Stack(
		m.Phylotaxy.ToString(), *unmarshallFromCode, *marshallOnCommit, true, *embeddedDiagrams)

	stager := m.NewStager(stack.R, stack.Stage)

	stager.UpdateAllStages()

	log.Println("Server ready serve on localhost:" + strconv.Itoa(*port))
	err := stack.R.Run(":" + strconv.Itoa(*port))
	if err != nil {
		log.Fatalln(err.Error())
	}
}
