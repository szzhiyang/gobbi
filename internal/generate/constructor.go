package generate

import (
	"github.com/dave/jennifer/jen"
)

type Constructor struct {
	*Function
}

func (c *Constructor) init(ns *Namespace, record *Record) {
	// Some constructors defined in gir files reference a subclass
	// of their real type.
	// Ensure that all constructors return a pointer to their
	// record/class type.
	c.ReturnValue.Type.CType = record.CType + "*"
	c.ReturnValue.Type.Name = record.Name

	c.Function.init(ns, nil)
	c.GoName = record.GoName + makeExportedGoName(c.Name)

	if record.Version != "" && c.Version == "" {
		c.Version = record.Version
	}
}

func (c *Constructor) generate(g *jen.Group, version *Version) {
	supported, reason := c.supported()
	if !supported {
		g.Commentf("Unsupported : %s", reason)
		g.Line()
		return
	}

	if !supportedByVersion(c, version) {
		return
	}

	c.Function.generate(g, version)
}
