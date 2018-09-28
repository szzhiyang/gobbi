package generate

import (
	"fmt"
	"github.com/dave/jennifer/jen"
)

type ReturnValue struct {
	Namespace *Namespace

	TransferOwnership string `xml:"transfer-ownership,attr"`
	Nullable          string `xml:"nullable,attr"`
	Doc               *Doc   `xml:"doc"`
	Type              *Type  `xml:"type"`
	Array             *Array `xml:"array"`
}

func (r *ReturnValue) init(ns *Namespace) {
	r.Namespace = ns

	if r.Type != nil {
		r.Type.init(ns)
	}
}

func (r *ReturnValue) mergeAddenda(addenda *ReturnValue) {
	if addenda.TransferOwnership != "" {
		r.TransferOwnership = addenda.TransferOwnership
	}

	if addenda.Type != nil {
		if addenda.Type.Name != "" {
			r.Type.Name = addenda.Type.Name
		}
		if addenda.Type.CType != "" {
			r.Type.CType = addenda.Type.CType
		}
	}
}

func (r *ReturnValue) isSupported() (bool, string) {
	if r.Type == nil {
		return false, "no return type"
	}
	if r.Type.generator == nil {
		return false, "no return generator"
	}

	if supported, reason := r.Type.generator.isSupportedAsReturnValue(); !supported {
		return false, fmt.Sprintf("return type : %s", reason)
	}

	return true, ""
}

func (r *ReturnValue) generateFunctionDeclaration(g *jen.Group) {
	r.Type.generator.generateReturnFunctionDeclaration(g)
}

func (r *ReturnValue) generateCToGo(g *jen.Group, cVarName string, goVarName string) {
	pkg := ""
	if r.Type.qname.ns != r.Namespace {
		pkg = r.Type.qname.ns.fullGoPackageName
	}

	r.Type.generator.generateReturnCToGo(g, false, cVarName, goVarName, pkg, r.TransferOwnership)
}
