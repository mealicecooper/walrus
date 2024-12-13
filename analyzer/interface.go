package analyzer

import (
	"fmt"
	"walrus/ast"
	"walrus/errgen"
	"walrus/utils"
)

func checkInterfaceTypeDecl(interfaceName string, interfaceNode ast.InterfaceType, env *TypeEnvironment) Interface {

	methods := make([]InterfaceMethodType, 0)

	for _, method := range interfaceNode.Methods {

		fnEnv := NewTypeENV(env, FUNCTION_SCOPE, method.Identifier.Name, env.filePath)

		params := make([]FnParam, 0)

		for _, param := range method.Parameters {
			fnParam := FnParam{
				Name:       param.Identifier.Name,
				IsOptional: param.IsOptional,
				Type:       evaluateTypeName(param.Type, fnEnv),
			}

			//check if the parameter is already declared
			if utils.Some(params, func(p FnParam) bool {
				return p.Name == fnParam.Name
			}) {
				errgen.AddError(env.filePath, param.Identifier.Start.Line, param.Identifier.End.Line, param.Identifier.Start.Column, param.Identifier.End.Column,
					fmt.Sprintf("parameter '%s' is already defined for method '%s'", fnParam.Name, method.Identifier.Name)).ErrorLevel(errgen.CRITICAL)
			}

			params = append(params, fnParam)
		}

		returns := evaluateTypeName(method.ReturnType, fnEnv)

		//check if the method already exists
		if utils.Some(methods, func(m InterfaceMethodType) bool {
			return m.Name == method.Identifier.Name
		}) {
			errgen.AddError(env.filePath, method.Identifier.Start.Line, method.Identifier.End.Line, method.Identifier.Start.Column, method.Identifier.End.Column,
				fmt.Sprintf("method '%s' already exists in interface '%s'", method.Identifier.Name, interfaceName)).ErrorLevel(errgen.CRITICAL)
		}

		methods = append(methods, InterfaceMethodType{
			Name: method.Identifier.Name,
			Method: Fn{
				DataType:      FUNCTION_TYPE,
				Params:        params,
				Returns:       returns,
				FunctionScope: *fnEnv,
			},
		})
	}

	return Interface{
		DataType:      INTERFACE_TYPE,
		InterfaceName: interfaceName,
		Methods:       methods,
	}
}
