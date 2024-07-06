package models

var GeneratorSingloton Generator

type Generator struct {
	Impl GeneratorInterface
}

func (generator *Generator) Generate() {
	if generator.Impl != nil {
		generator.Impl.Generate()
	}
}

type GeneratorInterface interface {
	Generate()
}
