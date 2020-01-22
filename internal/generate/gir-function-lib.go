package generate

import "C"
import (
	"github.com/blang/semver"
	"github.com/dave/jennifer/jen"
)

func (f *Function) generateLib(fi *jen.File, version semver.Version) {
	if supported, reason := f.isSupported(); !supported {
		fi.Commentf("UNSUPPORTED : %s : %s", f.CIdentifier, reason)
		fi.Line()
		return
	}

	// TODO
	if f.Throws {
		fi.Commentf("UNSUPPORTED : %s : throws", f.CIdentifier)
		fi.Line()
		return
	}

	// TODO
	for _, p := range f.Parameters {
		if p.Array != nil {
			fi.Commentf("UNSUPPORTED : %s : has array param, %s", f.CIdentifier, p.Name)
			fi.Line()
			return
		}

		if !p.isIn() {
			fi.Commentf("UNSUPPORTED : %s : has [in]out param, %s", f.CIdentifier, p.Name)
			fi.Line()
			return
		}
	}

	if f.MovedTo != "" || f.ShadowedBy != "" {
		return
	}

	if !generateEntityForVersion(version, f.version) {
		return
	}

	fi.Commentf("%s wraps the C function %s.", f.goName, f.CIdentifier)
	docVersion(fi, f.Version)

	// func Fn_some_function(...) [return type] {...}
	fi.
		Func().
		Do(f.generateLibReceiverDeclaration).
		Id(f.goName).
		ParamsFunc(f.generateLibParamsDeclaration).
		ParamsFunc(f.generateLibReturnTypeDeclaration).
		BlockFunc(f.generateLibBody)

	fi.Line()
}

func (f *Function) generateLibReceiverDeclaration(s *jen.Statement) {
	if f.InstanceParameter != nil {
		s.Id("recv").Add(jen.Id(f.InstanceParameter.Type.Name))

	}
}

func (f *Function) generateLibParamsDeclaration(g *jen.Group) {
	for _, param := range f.Parameters {
		if param.isVarargsOrValist() {
			continue
		}

		if !param.isIn() {
			continue
		}

		goType := param.libParamGoType()

		g.Id(param.goVarName).Add(goType)
	}

	//if f.Throws {
	//	g.Id("error").Add(jenUnsafePointer())
	//}
}

func (f *Function) generateLibReturnTypeDeclaration(g *jen.Group) {
	if f.ReturnValue.isVoid() {
		return
	}

	// TODO below belongs in Type.sysParamGoType ???
	typ := f.ReturnValue.Type
	if typ.isStruct() {
		g.Op("*").Add(typ.idOrQual())
		return
	}

	g.Add(f.ReturnValue.sysParamGoType())
}

func (f *Function) generateLibBody(g *jen.Group) {
	for _, p := range f.Parameters {
		if p.isVarargsOrValist() {
			continue
		}

		if !p.isIn() {
			continue
		}

		argVarName := "sys_" + p.goVarName
		p.generateLibArg(g, argVarName)
	}

	// call sys function
	g.
		Do(func(s *jen.Statement) {
			if !f.ReturnValue.isVoid() {
				s.Id("retSys").Op(":=")
			}
		}).
		Qual(f.namespace.goFullSysPackageName, f.sysName).CallFunc(f.generateLibCallParams)

	f.generateMarhsalReturnValue(g)
	f.generateLibReturn(g)
}

func (f *Function) generateLibCallParams(g *jen.Group) {
	//if f.InstanceParameter != nil {
	//	g.Id("cValueInstance")
	//}

	for _, param := range f.Parameters {
		if param.isVarargsOrValist() {
			continue
		}

		if !param.isIn() {
			continue
		}

		g.Id("sys_" + param.goVarName)
	}

	//if f.Throws {
	//	g.Id("cError")
	//}
}

func (f *Function) generateMarhsalReturnValue(g *jen.Group) {
	if f.ReturnValue.isVoid() {
		return
	}

	if f.ReturnValue.Type.isStruct() {
		if f.ReturnValue.Type.isQualifiedName() {
			ctorName := f.ReturnValue.Type.foreignName + "NewFromC"
			g.Id("ret").Op(":=").Qual(f.ReturnValue.Type.foreignNamespace.goFullPackageName, ctorName).Call(jen.Id("retSys"))
			return
		}

		ctorName := f.ReturnValue.Type.Name + "NewFromC"
		g.Id("ret").Op(":=").Id(ctorName).Call(jen.Id("retSys"))
		return
	}

	g.Id("ret").Op(":=").Id("retSys")
}

func (f *Function) generateLibReturn(g *jen.Group) {
	if f.ReturnValue.isVoid() {
		return
	}

	g.Line()

	g.Return().Do(func(s *jen.Statement) {
		if !f.ReturnValue.isVoid() {
			s.Id("ret")
		}
	})
}

//func (f *Function) generateLibVarArgsCFunction(fi *jen.File) {
//	params := []string{}
//	args := []string{}
//
//	if f.InstanceParameter != nil {
//		if f.InstanceParameter.Type != nil {
//			params = append(params, f.InstanceParameter.Type.CType+" "+f.InstanceParameter.Name)
//			args = append(args, f.InstanceParameter.Name)
//		} else if f.InstanceParameter.Array != nil {
//			params = append(params, f.InstanceParameter.Array.CType+" "+f.InstanceParameter.Name)
//			args = append(args, f.InstanceParameter.Name)
//		} else {
//			panic(fmt.Sprintf("Do not know how to generate instance parameter %s for %s", f.InstanceParameter.Name, f.CIdentifier))
//		}
//	}
//
//	for _, param := range f.Parameters {
//		if param.isVarargsOrValist() {
//			args = append(args, "NULL")
//		} else if param.Type != nil {
//			params = append(params, param.Type.CType+" "+param.Name)
//			args = append(args, param.Name)
//		} else if param.Array != nil {
//			params = append(params, param.Array.CType+" "+param.Name)
//			args = append(args, param.Name)
//		} else {
//			panic(fmt.Sprintf("Do not know how to generate parameter %s for %s", param.Name, f.CIdentifier))
//		}
//	}
//
//	fnDecl := fmt.Sprintf(`
//static %s c_%s(%s) {
//    return %s(%s);
//}
//`,
//		f.ReturnValue.Type.CType,
//		f.CIdentifier,
//		strings.Join(params, ", "),
//		f.CIdentifier,
//		strings.Join(args, ", "),
//	)
//
//	fi.CgoPreamble(fnDecl)
//}