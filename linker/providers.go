package linker

import (
	"github.com/Prepodavan/gombok/fields"
	"github.com/Prepodavan/gombok/gen"
	"github.com/Prepodavan/gombok/gen/bodies"
	"go/types"
)

type ReceiverProvider interface {
	Provide(typ *types.Type) gen.NamedHasGenerator
}

type NamingProvider interface {
	Provide(field *fields.Info) gen.HasGenerator
}

type BodyBuilderProvider interface {
	Provide(receiver bodies.Receiver, field *fields.Info) gen.HasGenerator
}
