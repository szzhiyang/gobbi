package generate

import (
	"fmt"
	"strings"

	"github.com/dave/jennifer/jen"
)

type ParameterTypeString struct {
	param         *Parameter
	cTypeName     string
	indirectLevel int
}

func ParameterTypeStringNew(param *Parameter) *ParameterTypeString {
	return &ParameterTypeString{
		param:         param,
		cTypeName:     param.Type.cTypeName,
		indirectLevel: param.Type.indirectLevel,
	}
}

func (pt *ParameterTypeString) isSupported() (supported bool, reason string) {
	if pt.param.Direction != "out" && pt.indirectLevel > 1 {
		return false, fmt.Sprintf("in for string with indirection level of %d", pt.indirectLevel)
	}

	return true, ""
}

func (pt *ParameterTypeString) generateFunctionDeclaration(g *jen.Group) {
	g.
		Id(pt.param.goVarName).
		Id(pt.param.goType)
}

func (pt *ParameterTypeString) generateCallArgument(g *jen.Group) {
	g.Id(pt.param.cVarName)
}

func (pt *ParameterTypeString) generateOutCallArgument(g *jen.Group) {
	g.
		Op("&").
		Id(pt.param.cVarName)
}

func (pt *ParameterTypeString) generateCVar(g *jen.Group) {
	g.
		Id(pt.param.cVarName).
		Op(":=").
		Qual("C", "CString").
		Call(jen.Id(pt.param.goVarName))

	if pt.param.TransferOwnership != "none" {
		// ownership is transferred (to the library) so we should not
		// free the string memory
		return
	}

	g.
		Defer().
		Qual("C", "free").
		Call(jen.
			Qual("unsafe", "Pointer").
			Call(jen.Id(pt.param.cVarName)))
}

func (pt *ParameterTypeString) generateOutCVar(g *jen.Group) {
	g.
		Var().
		Id(pt.param.cVarName).
		Op(strings.Repeat("*", pt.indirectLevel-1)).
		Qual("C", pt.cTypeName)
}
