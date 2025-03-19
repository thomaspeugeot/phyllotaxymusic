package models

import (
	split "github.com/fullstack-lang/gong/lib/split/go/models"
)

func (parameter *Parameter) UpdateAndCommitSplitStage() {

	parameter.splitStage.Reset()

	// Main View & Its Entire Layout
	(&split.View{
		Name: "Main view",
		RootAsSplitAreas: []*split.AsSplitArea{
			(&split.AsSplitArea{
				Name:             "Top",
				Size:             50,
				ShowNameInHeader: false,
				AsSplits: []*split.AsSplit{
					(&split.AsSplit{
						Name:      "Top, sidebar, table & form",
						Direction: split.Horizontal,
						AsSplitAreas: []*split.AsSplitArea{

							// Sidebar
							(&split.AsSplitArea{
								Name:             "sidebar tree",
								ShowNameInHeader: false,
								Size:             15,
								Tree: (&split.Tree{
									Name:      "Sidebar",
									StackName: SidebarTree.ToString(),
									TreeName:  Sidebar.ToString(),
								}).Stage(parameter.splitStage),
							}).Stage(parameter.splitStage),

							// SVG area
							(&split.AsSplitArea{
								Name:             "svg area",
								ShowNameInHeader: false,
								Size:             65,
								Svg: (&split.Svg{
									Name:      "svg",
									StackName: GongsvgStackName.ToString(),
								}).Stage(parameter.splitStage),
							}).Stage(parameter.splitStage),

							// Right area
							(&split.AsSplitArea{
								Name:             "right area",
								ShowNameInHeader: false,
								Size:             20,
								AsSplits: []*split.AsSplit{
									(&split.AsSplit{
										Name:      "vertical",
										Direction: split.Vertical,
										AsSplitAreas: []*split.AsSplitArea{

											// Slider area
											(&split.AsSplitArea{
												Name:             "slider area",
												ShowNameInHeader: false,
												Size:             60,
												Slider: (&split.Slider{
													Name:      "slider",
													StackName: GongLibSliderStackName.ToString(),
												}).Stage(parameter.splitStage),
											}).Stage(parameter.splitStage),

											// Button area
											(&split.AsSplitArea{
												Name:             "button area",
												ShowNameInHeader: false,
												Size:             20,
												Button: (&split.Button{
													Name:      "button",
													StackName: GongLibButtonStackName.ToString(),
												}).Stage(parameter.splitStage),
											}).Stage(parameter.splitStage),

											// Tone area
											(&split.AsSplitArea{
												Name:             "tone area",
												ShowNameInHeader: false,
												Size:             20,
												Tone: (&split.Tone{
													Name:      "tone",
													StackName: GongtoneStackName.ToString(),
												}).Stage(parameter.splitStage),
											}).Stage(parameter.splitStage),
										},
									}).Stage(parameter.splitStage),
								},
							}).Stage(parameter.splitStage),
						},
					}).Stage(parameter.splitStage),
				},
			}).Stage(parameter.splitStage),
		},
	}).Stage(parameter.splitStage)

	// Probe View & Its Root Split
	(&split.View{
		Name: "Probe view",
		RootAsSplitAreas: []*split.AsSplitArea{
			(&split.AsSplitArea{
				Name: "Probe area",
				Split: (&split.Split{
					Name:      "Probe area",
					StackName: "phyllotaxymusic-probe",
				}).Stage(parameter.splitStage),
			}).Stage(parameter.splitStage),
		},
	}).Stage(parameter.splitStage)

	parameter.splitStage.Commit()
}
