package models

import (
	button "github.com/fullstack-lang/gong/lib/button/go/models"
	slider "github.com/fullstack-lang/gong/lib/slider/go/models"
	split "github.com/fullstack-lang/gong/lib/split/go/models"
	svg "github.com/fullstack-lang/gong/lib/svg/go/models"
	tone "github.com/fullstack-lang/gong/lib/tone/go/models"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
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
									StackName: parameter.treeStage.GetName(),
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
									StackName: parameter.svgStage.GetName(),
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
													StackName: parameter.sliderStage.GetName(),
												}).Stage(parameter.splitStage),
											}).Stage(parameter.splitStage),

											// Button area
											(&split.AsSplitArea{
												Name:             "button area",
												ShowNameInHeader: false,
												Size:             20,
												Button: (&split.Button{
													Name:      "button",
													StackName: parameter.buttonStage.GetName(),
												}).Stage(parameter.splitStage),
											}).Stage(parameter.splitStage),

											// Tone area
											(&split.AsSplitArea{
												Name:             "tone area",
												ShowNameInHeader: false,
												Size:             20,
												Tone: (&split.Tone{
													Name:      "tone",
													StackName: parameter.toneStage.GetName(),
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

	(&split.View{
		Name: "Probe",
		RootAsSplitAreas: []*split.AsSplitArea{
			(&split.AsSplitArea{
				Name: "Probe area",
				Split: (&split.Split{
					Name:      "Probe area",
					StackName: parameter.phyllotaxymusicStage.GetName() + ProbeSplitSuffix,
				}).Stage(parameter.splitStage),
			}).Stage(parameter.splitStage),
		},
	}).Stage(parameter.splitStage)

	(&split.View{
		Name: "svg probe",
		RootAsSplitAreas: []*split.AsSplitArea{
			(&split.AsSplitArea{
				Name: "svg probe area",
				Split: (&split.Split{
					Name:      "svg probe area",
					StackName: parameter.svgStage.GetName() + svg.ProbeSplitSuffix,
				}).Stage(parameter.splitStage),
			}).Stage(parameter.splitStage),
		},
	}).Stage(parameter.splitStage)

	(&split.View{
		Name: "tree probe",
		RootAsSplitAreas: []*split.AsSplitArea{
			(&split.AsSplitArea{
				Name: "tree probe area",
				Split: (&split.Split{
					Name:      "tree probe area",
					StackName: parameter.treeStage.GetName() + tree.ProbeSplitSuffix,
				}).Stage(parameter.splitStage),
			}).Stage(parameter.splitStage),
		},
	}).Stage(parameter.splitStage)

	(&split.View{
		Name: "slider probe",
		RootAsSplitAreas: []*split.AsSplitArea{
			(&split.AsSplitArea{
				Name: "slider probe area",
				Split: (&split.Split{
					Name:      "slider probe area",
					StackName: parameter.sliderStage.GetName() + slider.ProbeSplitSuffix,
				}).Stage(parameter.splitStage),
			}).Stage(parameter.splitStage),
		},
	}).Stage(parameter.splitStage)

	(&split.View{
		Name: "button probe",
		RootAsSplitAreas: []*split.AsSplitArea{
			(&split.AsSplitArea{
				Name: "button probe area",
				Split: (&split.Split{
					Name:      "button probe area",
					StackName: parameter.buttonStage.GetName() + button.ProbeSplitSuffix,
				}).Stage(parameter.splitStage),
			}).Stage(parameter.splitStage),
		},
	}).Stage(parameter.splitStage)

	(&split.View{
		Name: "tone probe",
		RootAsSplitAreas: []*split.AsSplitArea{
			(&split.AsSplitArea{
				Name: "tone probe area",
				Split: (&split.Split{
					Name:      "tone probe area",
					StackName: parameter.toneStage.GetName() + tone.ProbeSplitSuffix,
				}).Stage(parameter.splitStage),
			}).Stage(parameter.splitStage),
		},
	}).Stage(parameter.splitStage)

	(&split.View{
		Name: "split probe",
		RootAsSplitAreas: []*split.AsSplitArea{
			(&split.AsSplitArea{
				Name: "split probe area",
				Split: (&split.Split{
					Name:      "split probe area",
					StackName: parameter.splitStage.GetName() + split.ProbeSplitSuffix,
				}).Stage(parameter.splitStage),
			}).Stage(parameter.splitStage),
		},
	}).Stage(parameter.splitStage)

	parameter.splitStage.Commit()
}
