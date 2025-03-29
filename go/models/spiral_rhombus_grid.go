package models

import gongsvg_models "github.com/fullstack-lang/gong/lib/svg/go/models"

type SpiralRhombusGrid struct {
	Name string

	Shape

	SpiralRhombuses []*SpiralRhombus
}

func (spiralRhombusGrid *SpiralRhombusGrid) Draw(gongsvgStage *gongsvg_models.Stage, layer *gongsvg_models.Layer, p *Parameter) {

	for _, sr := range spiralRhombusGrid.SpiralRhombuses {
		sr.Draw(gongsvgStage, layer, p)
	}

}
