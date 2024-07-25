package main

import (
	"flag"
	"log"
	"strconv"

	phylotaxymusic_models "github.com/thomaspeugeot/phylotaxymusic/go/models"
	phylotaxymusic_tree "github.com/thomaspeugeot/phylotaxymusic/go/tree"

	phylotaxymusic_stack "github.com/thomaspeugeot/phylotaxymusic/go/stack"
	phylotaxymusic_static "github.com/thomaspeugeot/phylotaxymusic/go/static"

	gongsvg_models "github.com/fullstack-lang/gongsvg/go/models"
	gongsvg_stack "github.com/fullstack-lang/gongsvg/go/stack"

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

	// get the only diagram
	parameters := phylotaxymusic_models.GetGongstructInstancesMap[phylotaxymusic_models.Parameter](phylotaxymusicStack.Stage)

	if len(*parameters) == 0 {
		log.Println("")
		// log.Fatalln("")
	}

	parameter := (*parameters)["Reference"]

	parameterImpl := new(ParameterImpl)
	parameterImpl.parameter = parameter
	parameterImpl.gongsvgStage = gongsvg_stack.Stage
	parameterImpl.phylotaxymusicStage = phylotaxymusicStack.Stage
	parameter.Impl = parameterImpl

	phylotaxymusic_models.GeneratorSingloton.Impl = parameterImpl
	parameterImpl.Generate()

	tree := new(phylotaxymusic_tree.Tree)
	tree.TreeStack = gongtree_stack
	tree.Generate(parameter)

	log.Printf("Server ready serve on localhost:" + strconv.Itoa(*port))
	err := r.Run(":" + strconv.Itoa(*port))
	if err != nil {
		log.Fatalln(err.Error())
	}
}

type ParameterImpl struct {
	gongsvgStage        *gongsvg_models.StageStruct
	phylotaxymusicStage *phylotaxymusic_models.StageStruct
	parameter           *phylotaxymusic_models.Parameter
}

// Generate implements models.GeneratorInterface.
func (parameterImpl *ParameterImpl) Generate() {
	parameterImpl.parameter.ComputeShapes(parameterImpl.phylotaxymusicStage)
	parameterImpl.parameter.GenerateSvg(parameterImpl.gongsvgStage)
	parameterImpl.phylotaxymusicStage.Commit()
}

func (parameterImpl *ParameterImpl) OnUpdated(updatedParameter *phylotaxymusic_models.Parameter) {

	log.Println("OnUpdated", parameterImpl.parameter.InsideAngle, parameterImpl.parameter.SideLength)
	// phylotaxymusic_svg.GenerateSvg(parameterImpl.gongsvgStage, parameterImpl.phylotaxymusicStage)

	updatedParameter.ComputeShapes(parameterImpl.phylotaxymusicStage)
	parameterImpl.parameter.GenerateSvg(parameterImpl.gongsvgStage)
}
