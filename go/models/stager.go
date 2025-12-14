// generated boilerplate code
// edit the file for adding other stages
package models

import (
	"github.com/gin-gonic/gin"

	button_stack "github.com/fullstack-lang/gong/lib/button/go/stack"
	cursor_stack "github.com/fullstack-lang/gong/lib/cursor/go/stack"
	load_stack "github.com/fullstack-lang/gong/lib/load/go/stack"
	slider_stack "github.com/fullstack-lang/gong/lib/slider/go/stack"
	split_stack "github.com/fullstack-lang/gong/lib/split/go/stack"
	ssg_level1stack "github.com/fullstack-lang/gong/lib/ssg/go/level1stack"
	svg_stack "github.com/fullstack-lang/gong/lib/svg/go/stack"
	tone_stack "github.com/fullstack-lang/gong/lib/tone/go/stack"
	tree_stack "github.com/fullstack-lang/gong/lib/tree/go/stack"

	button "github.com/fullstack-lang/gong/lib/button/go/models"
	cursor "github.com/fullstack-lang/gong/lib/cursor/go/models"
	load "github.com/fullstack-lang/gong/lib/load/go/models"
	slider "github.com/fullstack-lang/gong/lib/slider/go/models"
	split "github.com/fullstack-lang/gong/lib/split/go/models"
	ssg "github.com/fullstack-lang/gong/lib/ssg/go/models"
	svg "github.com/fullstack-lang/gong/lib/svg/go/models"
	tone "github.com/fullstack-lang/gong/lib/tone/go/models"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

type Stager struct {
	parameter *Parameter

	cursor *cursor.Cursor

	phyllotaxymusicStage *Stage
	svgStage             *svg.Stage
	toneStage            *tone.Stage
	treeStage            *tree.Stage
	cursorStage          *cursor.Stage
	loadStage            *load.Stage
	sliderStage          *slider.Stage
	buttonStage          *button.Stage
	splitStage           *split.Stage
	ssgStage             *ssg.Stage

	tree *tree.Tree
}

func NewStager(r *gin.Engine, stage *Stage) (stager *Stager) {

	stager = new(Stager)

	// get the only diagram
	parameters := GetGongstructInstancesSet[Parameter](stage)
	var parameter *Parameter
	for parameter_ := range *parameters {
		if parameter_.Name == "Reference" {
			parameter = parameter_
		}
	}

	if parameter == nil {
		InitializeDatabase(stage)
		parameters := GetGongstructInstancesSet[Parameter](stage)
		for parameter_ := range *parameters {
			if parameter_.Name == "Reference" {
				parameter = parameter_
			}
		}
	}

	// temporary
	stager.parameter = parameter
	parameter.stager = stager

	name := stage.GetName()

	stager.phyllotaxymusicStage = stage

	// the root split name is "" by convention
	stager.splitStage = split_stack.NewStack(r, "", "", "", "", true, true).Stage

	stager.buttonStage = button_stack.NewStack(r, name, "", "", "", true, true).Stage
	stager.cursorStage = cursor_stack.NewStack(r, name, "", "", "", true, true).Stage
	stager.loadStage = load_stack.NewStack(r, name, "", "", "", true, true).Stage
	stager.ssgStage = ssg_level1stack.NewLevel1Stack(name, "", "", true, true).Stage
	stager.svgStage = svg_stack.NewStack(r, name, "", "", "", true, true).Stage
	stager.toneStage = tone_stack.NewStack(r, name, "", "", "", true, true).Stage
	stager.treeStage = tree_stack.NewStack(r, name, "", "", "", true, true).Stage
	stager.sliderStage = slider_stack.NewStack(r, name, "", "", "", true, true).Stage

	// connect parameter to cursor for start playing notification
	stager.cursor = new(cursor.Cursor).Stage(stager.cursorStage)
	stager.cursorStage.Commit()

	// Main View & Its Entire Layout
	split.StageBranch(stager.splitStage, &split.View{
		Name: "Main view",
		RootAsSplitAreas: []*split.AsSplitArea{
			(&split.AsSplitArea{
				Size:             50,
				ShowNameInHeader: false,
				AsSplit: (&split.AsSplit{
					Direction: split.Horizontal,
					AsSplitAreas: []*split.AsSplitArea{

						// Sidebar
						(&split.AsSplitArea{
							ShowNameInHeader: false,
							Size:             15,
							Tree: (&split.Tree{
								StackName: stager.treeStage.GetName(),
							}),
						}),

						// SVG area
						(&split.AsSplitArea{
							ShowNameInHeader: false,
							Size:             65,
							HasDiv:           true,
							DivStyle:         "position: relative; width: 100%; height: 100%;",
							Svg: (&split.Svg{
								StackName: stager.svgStage.GetName(),
								Style:     "position: absolute;top: 0;left: 0;width: 100%;height: 100%;z-index: 1;",
							}),
							Cursor: (&split.Cursor{
								StackName: stager.cursorStage.GetName(),
								Style:     "position: absolute;top: 0;left: 0;width: 100%;height: 100%;z-index: 2;pointer-events: none;",
							}),
						}),

						// Right area
						(&split.AsSplitArea{
							ShowNameInHeader: false,
							Size:             20,
							AsSplit: (&split.AsSplit{
								Direction: split.Vertical,
								AsSplitAreas: []*split.AsSplitArea{

									// Slider area
									(&split.AsSplitArea{
										ShowNameInHeader: false,
										Size:             60,
										Slider: (&split.Slider{
											StackName: stager.sliderStage.GetName(),
										}),
									}),

									// Button area
									(&split.AsSplitArea{
										ShowNameInHeader: false,
										Size:             20,
										Button: (&split.Button{
											StackName: stager.buttonStage.GetName(),
										}),
									}),

									// Tone area
									(&split.AsSplitArea{
										ShowNameInHeader: false,
										Size:             20,
										Tone: (&split.Tone{
											StackName: stager.toneStage.GetName(),
										}),
									}),

									// load stage is load but invisible (Size =0). It
									// is used to download generated Musescore files
									(&split.AsSplitArea{
										ShowNameInHeader: false,
										Size:             0,
										Load: (&split.Load{
											StackName: stager.loadStage.GetName(),
										}),
									}),
								},
							}),
						}),
					},
				}),
			}),
		},
	})

	split.StageBranch(stager.splitStage, &split.View{
		Name: "Probe",
		RootAsSplitAreas: []*split.AsSplitArea{
			(&split.AsSplitArea{
				Split: (&split.Split{
					StackName: stager.phyllotaxymusicStage.GetProbeSplitStageName(),
				}),
			}),
		},
	})

	split.StageBranch(stager.splitStage, &split.View{
		Name: "svg probe",
		RootAsSplitAreas: []*split.AsSplitArea{
			(&split.AsSplitArea{
				Split: (&split.Split{
					StackName: stager.svgStage.GetProbeSplitStageName(),
				}),
			}),
		},
	})

	split.StageBranch(stager.splitStage, &split.View{
		Name: "tree probe",
		RootAsSplitAreas: []*split.AsSplitArea{
			(&split.AsSplitArea{
				Split: (&split.Split{
					StackName: stager.treeStage.GetProbeSplitStageName(),
				}),
			}),
		},
	})

	split.StageBranch(stager.splitStage, &split.View{
		Name: "slider probe",
		RootAsSplitAreas: []*split.AsSplitArea{
			(&split.AsSplitArea{
				Split: (&split.Split{
					StackName: stager.sliderStage.GetProbeSplitStageName(),
				}),
			}),
		},
	})

	split.StageBranch(stager.splitStage, &split.View{
		Name: "button probe",
		RootAsSplitAreas: []*split.AsSplitArea{
			(&split.AsSplitArea{
				Split: (&split.Split{
					StackName: stager.buttonStage.GetProbeSplitStageName(),
				}),
			}),
		},
	})

	split.StageBranch(stager.splitStage, &split.View{
		Name: "tone probe",
		RootAsSplitAreas: []*split.AsSplitArea{
			(&split.AsSplitArea{
				Split: (&split.Split{
					StackName: stager.toneStage.GetProbeSplitStageName(),
				}),
			}),
		},
	})

	split.StageBranch(stager.splitStage, &split.View{
		Name: "cursor probe",
		RootAsSplitAreas: []*split.AsSplitArea{
			(&split.AsSplitArea{
				Split: (&split.Split{
					StackName: stager.cursorStage.GetProbeSplitStageName(),
				}),
			}),
		},
	})

	split.StageBranch(stager.splitStage, &split.View{
		Name: "split probe",
		RootAsSplitAreas: []*split.AsSplitArea{
			(&split.AsSplitArea{
				Split: (&split.Split{
					StackName: stager.splitStage.GetProbeSplitStageName(),
				}),
			}),
		},
	})

	split.StageBranch(stager.splitStage, &split.View{
		Name: "ssg probe",
		RootAsSplitAreas: []*split.AsSplitArea{
			(&split.AsSplitArea{
				Split: (&split.Split{
					StackName: stager.ssgStage.GetProbeSplitStageName(),
				}),
			}),
		},
	})

	stager.splitStage.Commit()

	return
}

func (stager *Stager) GetButtonsStage() (buttonStage *button.Stage) {
	return stager.buttonStage
}

func (stager *Stager) GetSliderStage() (sliderStage *slider.Stage) {
	return stager.sliderStage
}

func (stager *Stager) GetGongtreeStage() (treeStage *tree.Stage) {
	return stager.treeStage
}
