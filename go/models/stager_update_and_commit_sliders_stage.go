package models

import (
	m "github.com/fullstack-lang/gong/lib/slider/go/models"
)

func (stager *Stager) UpdateAndCommitSlidersStage() {
	stager.sliderStage.Reset()

	layout := new(m.Layout).Stage(stager.sliderStage)

	group1 := new(m.Group).Stage(stager.sliderStage)
	group1.Percentage = 25
	layout.Groups = append(layout.Groups, group1)

	group2 := new(m.Group).Stage(stager.sliderStage)
	group2.Percentage = 25
	layout.Groups = append(layout.Groups, group2)

	group3 := new(m.Group).Stage(stager.sliderStage)
	group3.Percentage = 50
	layout.Groups = append(layout.Groups, group3)

	{
		group1.Sliders = append(
			group1.Sliders,
			m.NewSlider(
				stager,
				"Side Length",
				5,
				200,
				5,
				&stager.parameter.SideLength,
			),
		)

		group1.Sliders = append(
			group1.Sliders,
			m.NewSlider(
				stager,
				"Origin X",
				-100,
				600,
				10,
				&stager.parameter.OriginX,
			),
		)

		group1.Sliders = append(
			group1.Sliders,
			m.NewSlider(
				stager,
				"Origin Y",
				0,
				3000,
				50,
				&stager.parameter.OriginY,
			),
		)

		group1.Sliders = append(
			group1.Sliders,
			m.NewSlider(
				stager,
				"Spiral Initial Radius",
				0.0,
				5.0,
				0.015,
				&stager.parameter.SpiralRadiusRatio,
			),
		)
	}

	{

		group2.Sliders = append(
			group2.Sliders,
			m.NewSlider(
				stager,
				"Inside Angle",
				60,
				120,
				1,
				&stager.parameter.InsideAngle,
			),
		)

		group2.Sliders = append(
			group2.Sliders,
			m.NewSlider(
				stager,
				"Bezier Strength",
				0,
				4,
				0.01,
				&stager.parameter.BezierControlLengthRatio,
			),
		)

		group2.Sliders = append(
			group2.Sliders,
			m.NewSlider(
				stager,
				"M",
				1,
				20,
				1,
				&stager.parameter.M,
			),
		)

		group2.Sliders = append(
			group2.Sliders,
			m.NewSlider(
				stager,
				"N",
				1,
				20,
				1,
				&stager.parameter.N,
			),
		)

		group2.Sliders = append(
			group2.Sliders,
			m.NewSlider(
				stager,
				"Z",
				1,
				120,
				1,
				&stager.parameter.Z,
			),
		)

		group2.Sliders = append(
			group2.Sliders,
			m.NewSlider(
				stager,
				"S Bezier Strength",
				0,
				10,
				0.1,
				&stager.parameter.SpiralBezierStrength,
			),
		)

		group2.Sliders = append(
			group2.Sliders,
			m.NewSlider(
				stager,
				"Nb Interpol Points",
				1,
				50,
				1,
				&stager.parameter.NbInterpolationPoints,
			),
		)

		group2.Sliders = append(
			group2.Sliders,
			m.NewSlider(
				stager,
				"Stack Height",
				1,
				8,
				1,
				&stager.parameter.StackHeight,
			),
		)
	}
	{

		group3.Sliders = append(
			group3.Sliders,
			m.NewSlider(
				stager,
				"Pitch Height",
				0,
				0.3,
				0.001,
				&stager.parameter.PitchHeight,
			),
		)

		group3.Sliders = append(
			group3.Sliders,
			m.NewSlider(
				stager,
				"Nb Beats in Theme",
				1,
				64,
				1,
				&stager.parameter.NbOfBeatsInTheme,
			),
		)

		group3.Sliders = append(
			group3.Sliders,
			m.NewSlider(
				stager,
				"BeatsPerSecond",
				0,
				20,
				0.05,
				&stager.parameter.BeatsPerSecond,
			),
		)

		group3.Sliders = append(
			group3.Sliders,
			m.NewSlider(
				stager,
				"1st voice X",
				-1,
				1,
				0.01,
				&stager.parameter.FirstVoiceShiftX,
			),
		)

		group3.Sliders = append(
			group3.Sliders,
			m.NewSlider(
				stager,
				"1st voice Y",
				-1,
				4,
				0.01,
				&stager.parameter.FirstVoiceShiftY,
			),
		)

		group3.Sliders = append(
			group3.Sliders,
			m.NewSlider(
				stager,
				"2nd voice pitch diff",
				-12,
				24,
				1,
				&stager.parameter.PitchDifference,
			),
		)

		group3.Sliders = append(
			group3.Sliders,
			m.NewSlider(
				stager,
				"Level",
				0,
				20,
				0.1,
				&stager.parameter.Level,
			),
		)

		group3.Sliders = append(
			group3.Sliders,
			m.NewSlider(
				stager,
				"Actual Beats Shift",
				0,
				20,
				1,
				&stager.parameter.ActualBeatsTemporalShift,
			),
		)

		group3.Checkboxes = append(
			group3.Checkboxes,
			m.NewCheckbox(
				stager,
				"Scale",
				"Minor",
				"Major",
				&stager.parameter.IsMinor,
			),
		)

	}

	stager.sliderStage.Commit()
}
