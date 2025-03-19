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
				Size:             50,
				ShowNameInHeader: false,
				AsSplits: []*split.AsSplit{
					(&split.AsSplit{
						Direction: split.Horizontal,
						AsSplitAreas: []*split.AsSplitArea{

							// Sidebar
							(&split.AsSplitArea{
								ShowNameInHeader: false,
								Size:             15,
								Tree: (&split.Tree{
									StackName: parameter.treeStage.GetName(),
									TreeName:  Sidebar.ToString(),
								}).Stage(parameter.splitStage),
							}).Stage(parameter.splitStage),

							// SVG area
							(&split.AsSplitArea{
								ShowNameInHeader: false,
								Size:             65,
								Svg: (&split.Svg{
									StackName: parameter.svgStage.GetName(),
								}).Stage(parameter.splitStage),
							}).Stage(parameter.splitStage),

							// Right area
							(&split.AsSplitArea{
								ShowNameInHeader: false,
								Size:             20,
								AsSplits: []*split.AsSplit{
									(&split.AsSplit{
										Direction: split.Vertical,
										AsSplitAreas: []*split.AsSplitArea{

											// Slider area
											(&split.AsSplitArea{
												ShowNameInHeader: false,
												Size:             60,
												Slider: (&split.Slider{
													StackName: parameter.sliderStage.GetName(),
												}).Stage(parameter.splitStage),
											}).Stage(parameter.splitStage),

											// Button area
											(&split.AsSplitArea{
												ShowNameInHeader: false,
												Size:             20,
												Button: (&split.Button{
													StackName: parameter.buttonStage.GetName(),
												}).Stage(parameter.splitStage),
											}).Stage(parameter.splitStage),

											// Tone area
											(&split.AsSplitArea{
												ShowNameInHeader: false,
												Size:             20,
												Tone: (&split.Tone{
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
				Split: (&split.Split{
					StackName: parameter.phyllotaxymusicStage.GetName() + ProbeSplitSuffix,
				}).Stage(parameter.splitStage),
			}).Stage(parameter.splitStage),
		},
	}).Stage(parameter.splitStage)

	(&split.View{
		Name: "svg probe",
		RootAsSplitAreas: []*split.AsSplitArea{
			(&split.AsSplitArea{
				Split: (&split.Split{
					StackName: parameter.svgStage.GetName() + svg.ProbeSplitSuffix,
				}).Stage(parameter.splitStage),
			}).Stage(parameter.splitStage),
		},
	}).Stage(parameter.splitStage)

	(&split.View{
		Name: "tree probe",
		RootAsSplitAreas: []*split.AsSplitArea{
			(&split.AsSplitArea{
				Split: (&split.Split{
					StackName: parameter.treeStage.GetName() + tree.ProbeSplitSuffix,
				}).Stage(parameter.splitStage),
			}).Stage(parameter.splitStage),
		},
	}).Stage(parameter.splitStage)

	(&split.View{
		Name: "slider probe",
		RootAsSplitAreas: []*split.AsSplitArea{
			(&split.AsSplitArea{
				Split: (&split.Split{
					StackName: parameter.sliderStage.GetName() + slider.ProbeSplitSuffix,
				}).Stage(parameter.splitStage),
			}).Stage(parameter.splitStage),
		},
	}).Stage(parameter.splitStage)

	(&split.View{
		Name: "button probe",
		RootAsSplitAreas: []*split.AsSplitArea{
			(&split.AsSplitArea{
				Split: (&split.Split{
					StackName: parameter.buttonStage.GetName() + button.ProbeSplitSuffix,
				}).Stage(parameter.splitStage),
			}).Stage(parameter.splitStage),
		},
	}).Stage(parameter.splitStage)

	(&split.View{
		Name: "tone probe",
		RootAsSplitAreas: []*split.AsSplitArea{
			(&split.AsSplitArea{
				Split: (&split.Split{
					StackName: parameter.toneStage.GetName() + tone.ProbeSplitSuffix,
				}).Stage(parameter.splitStage),
			}).Stage(parameter.splitStage),
		},
	}).Stage(parameter.splitStage)

	(&split.View{
		Name: "split probe",
		RootAsSplitAreas: []*split.AsSplitArea{
			(&split.AsSplitArea{
				Split: (&split.Split{
					StackName: parameter.splitStage.GetName() + split.ProbeSplitSuffix,
				}).Stage(parameter.splitStage),
			}).Stage(parameter.splitStage),
		},
	}).Stage(parameter.splitStage)

	parameter.splitStage.Commit()
}
