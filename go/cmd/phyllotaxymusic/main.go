package main

import (
	"flag"
	"log"
	"strconv"

	m "github.com/thomaspeugeot/phyllotaxymusic/go/models"
	phyllotaxymusic_stack "github.com/thomaspeugeot/phyllotaxymusic/go/stack"
	phyllotaxymusic_static "github.com/thomaspeugeot/phyllotaxymusic/go/static"

	gongsvg_stack "github.com/fullstack-lang/gong/lib/svg/go/stack"

	gongtone_stack "github.com/fullstack-lang/gong/lib/tone/go/stack"

	gongtree_stack "github.com/fullstack-lang/gong/lib/tree/go/stack"

	cursor_models "github.com/fullstack-lang/gong/lib/cursor/go/models"
	cursor_stack "github.com/fullstack-lang/gong/lib/cursor/go/stack"

	load_stack "github.com/fullstack-lang/gong/lib/load/go/stack"

	slider_stack "github.com/fullstack-lang/gong/lib/slider/go/stack"

	button_stack "github.com/fullstack-lang/gong/lib/button/go/stack"

	split_stack "github.com/fullstack-lang/gong/lib/split/go/stack"
)

var (
	logGINFlag = flag.Bool("logGIN", false, "log mode for gin")

	unmarshallFromCode = flag.String("unmarshallFromCode", "", "unmarshall data from go file and '.go' (must be lowercased without spaces), If unmarshallFromCode arg is '', no unmarshalling")
	marshallOnCommit   = flag.String("marshallOnCommit", "", "on all commits, marshall staged data to a go file with the marshall name and '.go' (must be lowercased without spaces). If marshall arg is '', no marshalling")

	embeddedDiagrams = flag.Bool("embeddedDiagrams", false, "parse/analysis go/models and go/embeddedDiagrams")

	port = flag.Int("port", 8080, "port server")

	genmusicxml = flag.Bool("genmusicxml", false, "parse/analysis go/models and go/diagrams")
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

	button_stack := button_stack.NewStack(r, m.ButtonStackName.ToString(), "", "", "", false, true)
	cursor_stack := cursor_stack.NewStack(r, m.CursorStackName.ToString(), "", "", "", false, true)
	load_stack := load_stack.NewStack(r, m.LoadStackName.ToString(), "", "", "", false, true)
	gongsvg_stack := gongsvg_stack.NewStack(r, m.GongsvgStackName.ToString(), "", "", "", true, true)
	gongtone_stack := gongtone_stack.NewStack(r, m.ToneStackName.ToString(), "", "", "", true, true)
	gongtree_stack := gongtree_stack.NewStack(r, m.SidebarTree.ToString(), "", "", "", true, true)
	sliders_stack := slider_stack.NewStack(r, m.GongLibSliderStackName.ToString(), "", "", "", false, true)
	split_stack := split_stack.NewStack(r, m.RootSplitStackName.ToString(), "", "", "", false, true)

	// get the only diagram
	parameters := m.GetGongstructInstancesMap[m.Parameter](phyllotaxymusic_stack.Stage)
	parameter, ok := (*parameters)["Reference"]
	if !ok {
		log.Fatal("no Reference parameter on stage")
	}
	m.NewStager(r, phyllotaxymusic_stack.Stage)

	// parameter is used for coordinating all updates to the
	// phyllotaxy stage and compute all updates to the all stages
	parameter.SetSvgStage(gongsvg_stack.Stage)
	parameter.SetPhyllotaxymusicStage(phyllotaxymusic_stack.Stage)
	parameter.SetToneStage(gongtone_stack.Stage)
	parameter.SetCursorStage(cursor_stack.Stage)
	parameter.SetLoadStage(load_stack.Stage)
	parameter.SetSlidersStage(sliders_stack.Stage)
	parameter.SetButtonsStage(button_stack.Stage)
	parameter.SetSplitsStage(split_stack.Stage)
	parameter.SetGongtreeStage(gongtree_stack.Stage)

	// connect parameter to cursor for start playing notification
	cursor := new(cursor_models.Cursor).Stage(cursor_stack.Stage)
	cursor_stack.Stage.Commit()
	parameter.SetCursor(cursor)

	parameter.UpdateAndCommitSplitStage()
	parameter.UpdateAllStages()

	if *genmusicxml {
		parameter.UpdatePhyllotaxyStage()
		parameter.GenerateMusicXMLFile()
	}
	parameter.CommitPhyllotaxymusicStage()

	log.Printf("%s", "Server ready serve on localhost:"+strconv.Itoa(*port))
	err := r.Run(":" + strconv.Itoa(*port))
	if err != nil {
		log.Fatalln(err.Error())
	}

}
