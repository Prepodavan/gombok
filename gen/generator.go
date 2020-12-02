package gen

import "github.com/dave/jennifer/jen"

// +gen slice:"Where"
type Generator func(statement *jen.Statement)

type HasGenerator interface {
	Generator() Generator
}

type NamedHasGenerator interface {
	HasGenerator
	Name() string
}

type Generators []Generator

func (gs Generators) Add(others ...Generator) Generators {
	return append(gs, others...)
}
