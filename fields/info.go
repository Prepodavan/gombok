package fields

import "go/types"

type Info struct {
	*types.Var
	tag string
	str *types.Struct
}

func (info *Info) Struct() *types.Struct {
	return info.str
}

func (info *Info) Tag() string {
	return info.tag
}
