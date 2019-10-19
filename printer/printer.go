package printer

import (
	"fmt"

	"github.com/jhump/protoreflect/desc"
)

/*
example
import Foundation

enum UserService {
		// exec following function
		printMethod(packageName, service, method)
}
*/
func Print(files []*desc.FileDescriptor) {
	fmt.Printf("import Foundation\n")
	for _, file := range files {
		packageName := file.GetPackage()
		services := file.GetServices()

		if len(services) > 0 {
			fmt.Printf("\n")
		}

		for _, service := range services {
			fmt.Printf("enum %v {\n", service.GetName())

			for _, method := range service.GetMethods() {
				fmt.Printf("\n")
				printMethod(packageName, service, method)
			}

			fmt.Printf("}\n")
		}
	}
}

/*
example
	struct CreateDoctor: AppRequest {

		typealias Request = CreateUserRequest

		typealias Response = CreateUserResponse

		// exec following function
		printProperties(packageName, fields)

		var method: String {
			return "UserService/CreateUser"
		}

		// exec following function
		printRequestParameters(fields)
	}
*/
func printMethod(packageName string, service *desc.ServiceDescriptor, method *desc.MethodDescriptor) {
	fmt.Printf("\tstruct %v: AppRequest {\n", method.GetName())
	fmt.Printf("\n")
	fmt.Printf("\t\ttypealias Request = %v\n", method.GetInputType().GetName())
	fmt.Printf("\n")
	fmt.Printf("\t\ttypealias Response = %v\n", method.GetOutputType().GetName())

	fields := method.GetInputType().GetFields()
	if len(fields) > 0 {
		fmt.Printf("\n")
	}
	printProperties(packageName, fields)
	fmt.Printf("\n")

	fmt.Printf("\t\tvar method: String {\n")
	fmt.Printf("\t\t\treturn \"%v/%v\"\n", service.GetName(), method.GetName())
	fmt.Printf("\t\t}\n")
	fmt.Printf("\n")

	printRequestParameters(fields)
	fmt.Printf("\t}\n")
}

/*
example
	let name: String
	let age: Int64
*/
func printProperties(packageName string, fields []*desc.FieldDescriptor) {
	for _, field := range fields {
		fmt.Printf("\t\tlet %v: %v\n", valueName(field), typeName(packageName, field))
	}
}

/*
	var input: Request {
		return Request.with {
			$0.name = name
			$0.age = age
		}
	}
*/
func printRequestParameters(fields []*desc.FieldDescriptor) {
	fmt.Printf("\t\tvar input: Request {\n")

	if len(fields) == 0 {
		fmt.Printf("\t\t\treturn Request()\n")
	} else {
		fmt.Printf("\t\t\treturn Request.with {\n")

		for _, field := range fields {
			valueName := valueName(field)
			fmt.Printf("\t\t\t\t$0.%v = %v\n", valueName, valueName)
		}
		fmt.Printf("\t\t\t}\n")
	}

	fmt.Printf("\t\t}\n")
}