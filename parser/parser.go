package parser

import (
	"github.com/wzshiming/gotype"
	"go/parser"
	"os"
)

type Parser struct {
	importer *gotype.Importer
	filter   func(info os.FileInfo) bool
}

type Option func(*Parser)

func WithFileFilter(filter func(info os.FileInfo) bool) Option {
	return func(p *Parser) {
		p.filter = filter
	}
}

func NewParser(opts ...Option) (p *Parser) {
	p = &Parser{importer: gotype.NewImporter()}
	for _, opt := range opts {
		opt(p)
	}
	return
}

func (p *Parser) Dir(fp string) (pkgsTypes map[string]gotype.Type, err error) {
	pkgsTypes = make(map[string]gotype.Type)
	pkgs, parserErr := parser.ParseDir(p.importer.FileSet(), fp, p.filter, parser.ParseComments)
	if parserErr != nil {
		err = parserErr
		return
	}
	for k, v := range pkgs {
		pkg, pkgErr := p.importer.ImportPackage(k, v)
		if pkgErr != nil {
			err = pkgErr
			return
		}
		pkgsTypes[k] = pkg
	}
	return
}

func (p *Parser) File(fp string) (fileType gotype.Type, err error) {
	file, fErr := parser.ParseFile(p.importer.FileSet(), fp, nil, parser.ParseComments)
	if fErr != nil {
		err = fErr
		return
	}
	return p.importer.ImportFile(file.Name.Name, file)
}
