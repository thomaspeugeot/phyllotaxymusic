package models

import (
	m "github.com/thomaspeugeot/phyllotaxymusic/sliders/go/models"
)

func (parameter *Parameter) UpdateAndCommitSlidersStage() {
	parameter.slidersStage.Reset()

	layout := new(m.Layout).Stage(parameter.slidersStage)

	group1 := new(m.Group).Stage(parameter.slidersStage)
	group1.Percentage = 25
	layout.Groups = append(layout.Groups, group1)

	group2 := new(m.Group).Stage(parameter.slidersStage)
	group2.Percentage = 25
	layout.Groups = append(layout.Groups, group2)

	group3 := new(m.Group).Stage(parameter.slidersStage)
	group3.Percentage = 50
	layout.Groups = append(layout.Groups, group3)

	{
		group1.Sliders = append(
			group1.Sliders,
			NewSlider(
				parameter,
				"Side Length",
				5,
				200,
				5,
				&parameter.SideLength,
			),
		)

		group1.Sliders = append(
			group1.Sliders,
			NewSlider(
				parameter,
				"Origin X",
				-100,
				600,
				10,
				&parameter.OriginX,
			),
		)

		group1.Sliders = append(
			group1.Sliders,
			NewSlider(
				parameter,
				"Origin Y",
				0,
				3000,
				50,
				&parameter.OriginY,
			),
		)

		group1.Sliders = append(
			group1.Sliders,
			NewSlider(
				parameter,
				"Spiral Initial Radius",
				0.0,
				5.0,
				0.015,
				&parameter.SpiralRadiusRatio,
			),
		)
	}

	{

		group2.Sliders = append(
			group2.Sliders,
			NewSlider(
				parameter,
				"Inside Angle",
				60,
				120,
				1,
				&parameter.InsideAngle,
			),
		)

		group2.Sliders = append(
			group2.Sliders,
			NewSlider(
				parameter,
				"Bezier Strength",
				0,
				4,
				0.01,
				&parameter.BezierControlLengthRatio,
			),
		)

		group2.Sliders = append(
			group2.Sliders,
			NewSlider(
				parameter,
				"M",
				1,
				20,
				1,
				&parameter.M,
			),
		)

		group2.Sliders = append(
			group2.Sliders,
			NewSlider(
				parameter,
				"N",
				1,
				20,
				1,
				&parameter.N,
			),
		)

		group2.Sliders = append(
			group2.Sliders,
			NewSlider(
				parameter,
				"Z",
				1,
				120,
				1,
				&parameter.Z,
			),
		)

		group2.Sliders = append(
			group2.Sliders,
			NewSlider(
				parameter,
				"S Bezier Strength",
				0,
				10,
				0.1,
				&parameter.SpiralBezierStrength,
			),
		)

		group2.Sliders = append(
			group2.Sliders,
			NewSlider(
				parameter,
				"Nb Interpol Points",
				1,
				50,
				1,
				&parameter.NbInterpolationPoints,
			),
		)

		group2.Sliders = append(
			group2.Sliders,
			NewSlider(
				parameter,
				"Stack Height",
				1,
				8,
				1,
				&parameter.StackHeight,
			),
		)
	}
	{

		group3.Sliders = append(
			group3.Sliders,
			NewSlider(
				parameter,
				"Pitch Height",
				0,
				0.1,
				0.001,
				&parameter.PitchHeight,
			),
		)

		group3.Sliders = append(
			group3.Sliders,
			NewSlider(
				parameter,
				"Nb Beats in Theme",
				1,
				64,
				1,
				&parameter.NbOfBeatsInTheme,
			),
		)

		group3.Sliders = append(
			group3.Sliders,
			NewSlider(
				parameter,
				"BeatsPerSecond",
				0,
				20,
				0.05,
				&parameter.BeatsPerSecond,
			),
		)

		group3.Sliders = append(
			group3.Sliders,
			NewSlider(
				parameter,
				"1st voice X",
				-1,
				1,
				0.01,
				&parameter.FirstVoiceShiftX,
			),
		)

		group3.Sliders = append(
			group3.Sliders,
			NewSlider(
				parameter,
				"1st voice Y",
				-1,
				4,
				0.01,
				&parameter.FirstVoiceShiftY,
			),
		)

		group3.Sliders = append(
			group3.Sliders,
			NewSlider(
				parameter,
				"2nd voice pitch diff",
				-12,
				24,
				1,
				&parameter.PitchDifference,
			),
		)

		group3.Sliders = append(
			group3.Sliders,
			NewSlider(
				parameter,
				"Level",
				0,
				20,
				0.1,
				&parameter.Level,
			),
		)

		group3.Sliders = append(
			group3.Sliders,
			NewSlider(
				parameter,
				"Actual Beats Shift",
				0,
				20,
				1,
				&parameter.ActualBeatsTemporalShift,
			),
		)

	}

	parameter.slidersStage.Commit()
}
