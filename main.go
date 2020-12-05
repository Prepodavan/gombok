package main

import (
	"bytes"
	"flag"
	"github.com/Prepodavan/gombok/config"
	"github.com/Prepodavan/gombok/gen"
	"github.com/Prepodavan/gombok/parser"
	"github.com/wzshiming/gotype"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var (
	wholePkg     = flag.Bool("pkg", false, "process package instead of file")
	ptrReceiver  = flag.Bool("ptr", true, "uses pointer method receiver if true")
	getterPrefix = flag.String("g", "", "defines prefix for getters")
	setterPrefix = flag.String("s", "Set", "defines prefix for setters")
)

const suffix = "gombok"

func output(typ gotype.Type) (s *string) {
	p := typ.PkgPath()
	i := strings.LastIndex(p, "/")
	if i != -1 {
		p = p[i+1:]
	}
	p = p + "_" + suffix + ".go"
	s = &p
	return
}

func main() {
	flag.Parse()
	conf := &config.Config{}
	conf.SetPtrReceiver(*ptrReceiver)
	conf.SetGetterPrefix(*getterPrefix)
	conf.SetSetterPrefix(*setterPrefix)
	var outputFN *string
	gofile := os.Getenv("GOFILE")
	p := parser.NewParser(parser.WithFileFilter(func(info os.FileInfo) bool {
		return !strings.Contains(info.Name(), "_"+suffix+".go")
	}))
	if !*wholePkg {
		typ, err := p.File(gofile)
		if err != nil {
			panic(err)
		}
		g := gen.NewGenerator(typ, conf)
		buf := &bytes.Buffer{}
		_, err = g.WriteTo(buf)
		if err != nil {
			panic(err)
		}
		outputFN = output(typ)
		err = ioutil.WriteFile(*outputFN, buf.Bytes(), 0664)
		if err != nil {
			panic(err)
		}
	} else {
		fp, err := filepath.Abs(gofile)
		if err != nil {
			panic(err)
		}
		pkgs, err := p.Dir(filepath.Dir(fp))
		if err != nil {
			panic(err)
		}
		buf := &bytes.Buffer{}
		var g *gen.Generator
		for _, typ := range pkgs {
			if outputFN == nil {
				outputFN = output(typ)
			}
			if g == nil {
				g = gen.NewGenerator(typ, conf)
			} else {
				g = gen.NewGenerator(typ, conf, gen.WithJenFile(g.File()))
			}
			_, err = g.WriteTo(buf)
			if err != nil {
				panic(err)
			}
		}
		err = ioutil.WriteFile(*outputFN, buf.Bytes(), 0664)
		if err != nil {
			panic(err)
		}
	}
}
