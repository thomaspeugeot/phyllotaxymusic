package main

import (
	"time"

	"github.com/thomaspeugeot/phyllotaxymusic/sliders/go/models"

	// injection point for ident package import declaration{{ImportPackageDeclaration}}
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var _ time.Time

// _ point for meta package dummy declaration{{ImportPackageDummyDeclaration}}

// When parsed, those maps will help with the renaming process
var _ map[string]any = map[string]any{
	// injection point for docLink to identifiers{{EntriesDocLinkStringDocLinkIdentifier}}
}

// function will stage objects
func _(stage *models.StageStruct) {

	// Declaration of instances to stage

	__Checkbox__000000_Orientation := (&models.Checkbox{}).Stage(stage)

	__Group__000000_Group_1 := (&models.Group{}).Stage(stage)
	__Group__000001_Group_2 := (&models.Group{}).Stage(stage)

	__Layout__000000_Layout := (&models.Layout{}).Stage(stage)

	__Slider__000000_Value_1_with_int_type := (&models.Slider{}).Stage(stage)
	__Slider__000001_Value_2_int_type := (&models.Slider{}).Stage(stage)
	__Slider__000002_Value_3_Float := (&models.Slider{}).Stage(stage)

	// Setup of values

	__Checkbox__000000_Orientation.Name = `Orientation`
	__Checkbox__000000_Orientation.ValueBool = false
	__Checkbox__000000_Orientation.LabelForTrue = `Is Minor`
	__Checkbox__000000_Orientation.LabelForFalse = `Is Major`

	__Group__000000_Group_1.Name = `Group 1`

	__Group__000001_Group_2.Name = `Group 2`

	__Layout__000000_Layout.Name = `Layout`

	__Slider__000000_Value_1_with_int_type.Name = `Value 1 with int type`
	__Slider__000000_Value_1_with_int_type.IsFloat64 = false
	__Slider__000000_Value_1_with_int_type.IsInt = true
	__Slider__000000_Value_1_with_int_type.MinInt = -15
	__Slider__000000_Value_1_with_int_type.MaxInt = 20
	__Slider__000000_Value_1_with_int_type.StepInt = 5
	__Slider__000000_Value_1_with_int_type.ValueInt = -10
	__Slider__000000_Value_1_with_int_type.MinFloat64 = 0.000000
	__Slider__000000_Value_1_with_int_type.MaxFloat64 = 0.000000
	__Slider__000000_Value_1_with_int_type.StepFloat64 = 0.000000
	__Slider__000000_Value_1_with_int_type.ValueFloat64 = 0.000000

	__Slider__000001_Value_2_int_type.Name = `Value 2 int type`
	__Slider__000001_Value_2_int_type.IsFloat64 = false
	__Slider__000001_Value_2_int_type.IsInt = true
	__Slider__000001_Value_2_int_type.MinInt = -100
	__Slider__000001_Value_2_int_type.MaxInt = 200
	__Slider__000001_Value_2_int_type.StepInt = 20
	__Slider__000001_Value_2_int_type.ValueInt = -60
	__Slider__000001_Value_2_int_type.MinFloat64 = 0.000000
	__Slider__000001_Value_2_int_type.MaxFloat64 = 0.000000
	__Slider__000001_Value_2_int_type.StepFloat64 = 0.000000
	__Slider__000001_Value_2_int_type.ValueFloat64 = 0.000000

	__Slider__000002_Value_3_Float.Name = `Value 3 Float`
	__Slider__000002_Value_3_Float.IsFloat64 = true
	__Slider__000002_Value_3_Float.IsInt = false
	__Slider__000002_Value_3_Float.MinInt = 0
	__Slider__000002_Value_3_Float.MaxInt = 0
	__Slider__000002_Value_3_Float.StepInt = 0
	__Slider__000002_Value_3_Float.ValueInt = 0
	__Slider__000002_Value_3_Float.MinFloat64 = -50.000000
	__Slider__000002_Value_3_Float.MaxFloat64 = 50.000000
	__Slider__000002_Value_3_Float.StepFloat64 = 2.500000
	__Slider__000002_Value_3_Float.ValueFloat64 = -42.500000

	// Setup of pointers
	__Group__000000_Group_1.Sliders = append(__Group__000000_Group_1.Sliders, __Slider__000000_Value_1_with_int_type)
	__Group__000000_Group_1.Sliders = append(__Group__000000_Group_1.Sliders, __Slider__000001_Value_2_int_type)
	__Group__000000_Group_1.Sliders = append(__Group__000000_Group_1.Sliders, __Slider__000002_Value_3_Float)
	__Group__000000_Group_1.Checkboxes = append(__Group__000000_Group_1.Checkboxes, __Checkbox__000000_Orientation)
	__Layout__000000_Layout.Groups = append(__Layout__000000_Layout.Groups, __Group__000000_Group_1)
	__Layout__000000_Layout.Groups = append(__Layout__000000_Layout.Groups, __Group__000001_Group_2)
}
