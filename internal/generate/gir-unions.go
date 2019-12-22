package generate

import (
	"github.com/blang/semver"
	"github.com/dave/jennifer/jen"
)

type Unions []*Union

func (uu Unions) init(ns *Namespace) {
	for _, union := range uu {
		union.init(ns)
	}
}

func (uu Unions) byName(name string) (*Union, bool) {
	for _, u := range uu {
		if u.Name == name {
			return u, true
		}
	}

	return nil, false
}

func (uu Unions) generateSys(f *jen.File, version semver.Version) {
	f.Comment("unions")

	for _, u := range uu {
		u.generateSys(f, version)
	}

	f.Line()
}