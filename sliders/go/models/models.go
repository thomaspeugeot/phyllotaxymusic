package models

type Group struct {
	Name       string
	Sliders    []*Slider
	Checkboxes []*Checkbox
}

type Slider struct {
	Name string
}

type Checkbox struct {
	Name string
}
