package main

import (
	"flag"
	"log"
	"strconv"

	phylotaxymusic_models "github.com/thomaspeugeot/phylotaxymusic/go/models"
	phylotaxymusic_svg "github.com/thomaspeugeot/phylotaxymusic/go/svg"

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

	// setup phylotaxymusicStack
	phylotaxymusicStack := phylotaxymusic_stack.NewStack(r,
		phylotaxymusic_models.Phylotaxy.ToString(), *unmarshallFromCode, *marshallOnCommit, "", *embeddedDiagrams, true)
	phylotaxymusicStack.Probe.Refresh()
	phylotaxymusicStack.Stage.Checkout()

	gongsvg_stack := gongsvg_stack.NewStack(r, phylotaxymusic_models.GongsvgStackName.ToString(), "", "", "", true, true)

	phylotaxymusic_svg.GenerateSvg(gongsvg_stack.Stage, phylotaxymusicStack.Stage)

	// get the only diagram
	diagrams := phylotaxymusic_models.GetGongstructInstancesMap[phylotaxymusic_models.Parameter](phylotaxymusicStack.Stage)

	if len(*diagrams) == 0 {
		log.Println("")
		// log.Fatalln("")
	}

	diagram := (*diagrams)["Reference"]
	_ = diagram

	impl := new(DiagramImpl)
	impl.diagram = diagram
	impl.gongsvgStage = gongsvg_stack.Stage
	impl.phylotaxymusicStage = phylotaxymusicStack.Stage
	diagram.Impl = impl

	log.Printf("Server ready serve on localhost:" + strconv.Itoa(*port))
	err := r.Run(":" + strconv.Itoa(*port))
	if err != nil {
		log.Fatalln(err.Error())
	}
}

type DiagramImpl struct {
	gongsvgStage        *gongsvg_models.StageStruct
	phylotaxymusicStage *phylotaxymusic_models.StageStruct
	diagram             *phylotaxymusic_models.Parameter
}

func (diagramImpl *DiagramImpl) OnDiagramUpdated(updatedDiagram *phylotaxymusic_models.Parameter) {

	log.Println("", diagramImpl.diagram.DiamondAngle)
	phylotaxymusic_svg.GenerateSvg(diagramImpl.gongsvgStage, diagramImpl.phylotaxymusicStage)
}
