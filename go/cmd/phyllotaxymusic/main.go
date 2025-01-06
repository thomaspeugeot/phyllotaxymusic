package main

import (
	"flag"
	"log"
	"strconv"

	phyllotaxymusic_models "github.com/thomaspeugeot/phyllotaxymusic/go/models"
	"github.com/thomaspeugeot/phyllotaxymusic/go/models/musicxml"
	phyllotaxymusic_stack "github.com/thomaspeugeot/phyllotaxymusic/go/stack"
	phyllotaxymusic_static "github.com/thomaspeugeot/phyllotaxymusic/go/static"

	gongsvg_stack "github.com/fullstack-lang/gongsvg/go/stack"

	gongtone_stack "github.com/fullstack-lang/gongtone/go/stack"

	gongtree_stack "github.com/fullstack-lang/gongtree/go/stack"

	cursor_models "github.com/thomaspeugeot/phyllotaxymusic/cursor/go/models"
	cursor_stack "github.com/thomaspeugeot/phyllotaxymusic/cursor/go/stack"
)

var (
	logGINFlag = flag.Bool("logGIN", false, "log mode for gin")

	unmarshallFromCode = flag.String("unmarshallFromCode", "", "unmarshall data from go file and '.go' (must be lowercased without spaces), If unmarshallFromCode arg is '', no unmarshalling")
	marshallOnCommit   = flag.String("marshallOnCommit", "", "on all commits, marshall staged data to a go file with the marshall name and '.go' (must be lowercased without spaces). If marshall arg is '', no marshalling")

	diagrams         = flag.Bool("diagrams", true, "parse/analysis go/models and go/diagrams")
	embeddedDiagrams = flag.Bool("embeddedDiagrams", false, "parse/analysis go/models and go/embeddedDiagrams")

	port = flag.Int("port", 8080, "port server")

	genmusicxml = flag.Bool("genmusicxml", true, "parse/analysis go/models and go/diagrams")
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
		phyllotaxymusic_models.Phylotaxy.ToString(), *unmarshallFromCode, *marshallOnCommit, "", *embeddedDiagrams, true)
	phyllotaxymusic_stack.Probe.Refresh()
	phyllotaxymusic_stack.Stage.Checkout()

	gongsvg_stack := gongsvg_stack.NewStack(r, phyllotaxymusic_models.GongsvgStackName.ToString(), "", "", "", true, true)
	gongtree_stack := gongtree_stack.NewStack(r, phyllotaxymusic_models.SidebarTree.ToString(), "", "", "", true, true)
	gongtone_stack := gongtone_stack.NewStack(r, phyllotaxymusic_models.GongtoneStackName.ToString(), "", "", "", true, true)
	cursorStack := cursor_stack.NewStack(r, cursor_models.Cursorstakcname.ToString(), "", "", "", false, false)

	// get the only diagram
	parameters := phyllotaxymusic_models.GetGongstructInstancesMap[phyllotaxymusic_models.Parameter](phyllotaxymusic_stack.Stage)
	parameter, ok := (*parameters)["Reference"]
	if !ok {
		log.Fatal("no Reference parameter on stage")
	}

	// parameter is used for coordinating all updates to the
	// phyllotaxy stage and compute all updates to the all stages
	parameter.SetGongsvgStage(gongsvg_stack.Stage)
	parameter.SetPhyllotaxymusicStage(phyllotaxymusic_stack.Stage)
	parameter.SetGongtoneStage(gongtone_stack.Stage)
	parameter.SetCursorStage(cursorStack.Stage)
	parameter.SetGongtreeStage(gongtree_stack.Stage)
	parameter.SetTreeProxy()

	// connect parameter to cursor for start playing notification
	cursor := new(cursor_models.Cursor).Stage(cursorStack.Stage)
	cursorStack.Stage.Commit()
	parameter.SetCursor(cursor)

	parameter.UpdatePhyllotaxyStage()
	parameter.UpdateAndCommitCursorStage()
	parameter.UpdateAndCommitSVGStage()
	parameter.UpdateAndCommitToneStage()
	parameter.UpdateAndCommitTreeStage()
	parameter.CommitPhyllotaxymusicStage()

	if *genmusicxml {
		musicxml.GenerateMusicXMLFile()
	}

	log.Printf("%s", "Server ready serve on localhost:"+strconv.Itoa(*port))
	err := r.Run(":" + strconv.Itoa(*port))
	if err != nil {
		log.Fatalln(err.Error())
	}

}
