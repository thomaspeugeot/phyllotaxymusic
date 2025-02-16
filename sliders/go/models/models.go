package models

type Layout struct {
	Name string

	Groups []*Group
}

type Group struct {
	Name       string
	Sliders    []*Slider
	Checkboxes []*Checkbox
}

type Slider struct {
	Name string

	IsFloat64 bool
	IsInt     bool

	MinInt   int
	MaxInt   int
	StepInt  int
	ValueInt int

	MinFloat64   float64
	MaxFloat64   float64
	StepFloat64  float64
	ValueFloat64 float64
}

type Checkbox struct {
	Name string
}
