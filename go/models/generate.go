package models

var GeneratorSingloton Generator

type Generator struct {
	Impl GeneratorInterface
}

func (generator *Generator) Generate() {

}

type GeneratorInterface interface {
	Generate()
}
