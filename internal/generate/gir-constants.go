package generate

import "github.com/dave/jennifer/jen"

type Constants []*Constant

func (cc Constants) init(ns *Namespace) {
	for _, constant := range cc {
		constant.init(ns)
	}
}

func (cc Constants) generateSys(f *jen.File) {
	for _, constant := range cc {
		constant.generateSys(f)
	}
}
