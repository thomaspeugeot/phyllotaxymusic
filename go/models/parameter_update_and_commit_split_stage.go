package models

import (
	split "github.com/fullstack-lang/gong/lib/split/go/models"
)

func (parameter *Parameter) UpdateAndCommitSplitStage() {
	parameter.splitStage.Reset()

	mainView := (&split.View{
		Name: "Main view",
	}).Stage(parameter.splitStage)

	phyllotaxtProbeView := (&split.View{
		Name: "Probe view",
	}).Stage(parameter.splitStage)

	splitProbeArea := (&split.AsSplitArea{
		Name: "Probe area",
	}).Stage(parameter.splitStage)
	phyllotaxtProbeView.RootAsSplitAreas = append(phyllotaxtProbeView.RootAsSplitAreas, splitProbeArea)

	splitProbe := (&split.Split{
		Name:      "Probe area",
		StackName: "phyllotaxymusic-probe",
	}).Stage(parameter.splitStage)
	splitProbeArea.Split = splitProbe

	topSplitArea := (&split.AsSplitArea{
		Name:             "Top",
		Size:             50,
		ShowNameInHeader: false,
	}).Stage(parameter.splitStage)
	mainView.RootAsSplitAreas = append(mainView.RootAsSplitAreas, topSplitArea)

	horizontalSplit := (&split.AsSplit{
		Name:      "Top, sidebar, table & form",
		Direction: split.Horizontal,
	}).Stage(parameter.splitStage)
	topSplitArea.AsSplits = append(topSplitArea.AsSplits, horizontalSplit)

	sidebarArea := (&split.AsSplitArea{
		Name:             "sidebar tree",
		ShowNameInHeader: false,
		Size:             15,
	}).Stage(parameter.splitStage)
	horizontalSplit.AsSplitAreas = append(horizontalSplit.AsSplitAreas, sidebarArea)

	tree := (&split.Tree{
		Name:      "Sidebar",
		StackName: SidebarTree.ToString(),
		TreeName:  Sidebar.ToString(),
	}).Stage(parameter.splitStage)
	sidebarArea.Tree = tree

	svgArea := (&split.AsSplitArea{
		Name:             "svg area",
		ShowNameInHeader: false,
		Size:             65,
	}).Stage(parameter.splitStage)
	horizontalSplit.AsSplitAreas = append(horizontalSplit.AsSplitAreas, svgArea)

	svg := (&split.Svg{
		Name:      "svg",
		StackName: GongsvgStackName.ToString(),
	}).Stage(parameter.splitStage)
	svgArea.Svg = svg

	rightArea := (&split.AsSplitArea{
		Name:             "right area",
		ShowNameInHeader: false,
		Size:             20,
	}).Stage(parameter.splitStage)
	horizontalSplit.AsSplitAreas = append(horizontalSplit.AsSplitAreas, rightArea)

	verticalAsSplit := (&split.AsSplit{
		Name:      "vertical",
		Direction: split.Vertical,
	}).Stage(parameter.splitStage)
	rightArea.AsSplits = append(rightArea.AsSplits, verticalAsSplit)

	sliderArea := (&split.AsSplitArea{
		Name:             "slider area",
		ShowNameInHeader: false,
		Size:             60,
	}).Stage(parameter.splitStage)
	verticalAsSplit.AsSplitAreas = append(verticalAsSplit.AsSplitAreas, sliderArea)

	slider := (&split.Slider{
		Name:      "slider",
		StackName: GongLibSliderStackName.ToString(),
	}).Stage(parameter.splitStage)
	sliderArea.Slider = slider

	buttonArea := (&split.AsSplitArea{
		Name:             "button area",
		ShowNameInHeader: false,
		Size:             20,
	}).Stage(parameter.splitStage)
	verticalAsSplit.AsSplitAreas = append(verticalAsSplit.AsSplitAreas, buttonArea)

	button := (&split.Button{
		Name:      "button",
		StackName: GongLibButtonStackName.ToString(),
	}).Stage(parameter.splitStage)
	buttonArea.Button = button

	toneArea := (&split.AsSplitArea{
		Name:             "tone area",
		ShowNameInHeader: false,
		Size:             20,
	}).Stage(parameter.splitStage)
	verticalAsSplit.AsSplitAreas = append(verticalAsSplit.AsSplitAreas, toneArea)

	tone := (&split.Tone{
		Name:      "tone",
		StackName: GongtoneStackName.ToString(),
	}).Stage(parameter.splitStage)
	toneArea.Tone = tone

	parameter.splitStage.Commit()
}
