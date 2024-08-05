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
}

func (parameterImpl *ParameterImpl) OnUpdated(updatedParameter *phylotaxymusic_models.Parameter) {

	log.Println("OnUpdated", parameterImpl.parameter.InsideAngle, parameterImpl.parameter.SideLength)
	// phylotaxymusic_svg.GenerateSvg(parameterImpl.gongsvgStage, parameterImpl.phylotaxymusicStage)

	updatedParameter.ComputeShapes(parameterImpl.phylotaxymusicStage)
	updatedParameter.GenerateSvg(parameterImpl.gongsvgStage)
	parameterImpl.tree.Generate(updatedParameter)
	updatedParameter.GenerateNotes(parameterImpl.gongtoneStage)

}
