package gen

import "github.com/dave/jennifer/jen"

type CodeGen func(statement *jen.Statement)

type CodeGens []CodeGen

func NewCodeGens() (gs *CodeGens) {
	_gs := make(CodeGens, 0)
	gs = &_gs
	return
}

func (gs *CodeGens) Add(others ...CodeGen) *CodeGens {
	*gs = append(*gs, others...)
	return gs
}
