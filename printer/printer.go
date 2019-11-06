package printer

import (
	"fmt"
	"strings"

	"github.com/jhump/protoreflect/desc"
)

/*
example
import Foundation
import Entity

public enum UserService {
		// exec following function
		printMethod(packageName, service, method)
}
*/
func Print(files []*desc.FileDescriptor) {
	fmt.Printf("import Foundation\n")
	fmt.Printf("import Entity\n")
	for _, file := range files {
		packageName := file.GetPackage()
		services := file.GetServices()

		if len(services) > 0 {
			fmt.Printf("\n")
		}

		for _, service := range services {
			fmt.Printf("public enum %v {\n", service.GetName())

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
	public struct CreateDoctor: ProtoRequest {

		public typealias Request = CreateUserRequest

		public typealias Response = CreateUserResponse

		// exec following function
		printProperties(packageName, fields)

		public var method: String {
			return "UserService/CreateUser"
		}

		// exec following function
		printRequestParameters(fields)
	}
*/
func printMethod(packageName string, service *desc.ServiceDescriptor, method *desc.MethodDescriptor) {
	fmt.Printf("    public struct %v: ProtoRequest {\n", method.GetName())
	fmt.Printf("\n")
	fmt.Printf("        public typealias Request = %v\n", method.GetInputType().GetName())
	fmt.Printf("\n")
	fmt.Printf("        public typealias Response = %v\n", method.GetOutputType().GetName())

	fields := method.GetInputType().GetFields()
	if len(fields) > 0 {
		fmt.Printf("\n")
	}
	printProperties(packageName, fields)
	fmt.Printf("\n")
	printInit(packageName, fields)
	fmt.Printf("\n")

	fmt.Printf("        public var method: String {\n")
	fmt.Printf("            return \"%v/%v\"\n", service.GetName(), method.GetName())
	fmt.Printf("        }\n")
	fmt.Printf("\n")

	printRequestParameters(fields)
	fmt.Printf("    }\n")
}

/*
example
	let name: String
	let age: Int64
*/
func printProperties(packageName string, fields []*desc.FieldDescriptor) {
	for _, field := range fields {
		fmt.Printf("        public let %v: %v\n", valueName(field), typeName(packageName, field))
	}
}

/*
example
    public init(name: String, age: Int64) {
        self.name = name
		self.age = age
    }
*/
func printInit(packageName string, fields []*desc.FieldDescriptor) {
	init := "        public init("
	for _, field := range fields {
		init += fmt.Sprintf("%v: %v, ", valueName(field), typeName(packageName, field))
	}
	init = strings.TrimSuffix(init, ", ")
	init += ") {\n"
	for _, field := range fields {
		init += fmt.Sprintf("            self.%v = %v\n", valueName(field), valueName(field))
	}
	init += "        }"
	fmt.Println(init)
}

/*
	public var input: Request {
		return Request.with {
			$0.name = name
			$0.age = age
		}
	}
*/
func printRequestParameters(fields []*desc.FieldDescriptor) {
	fmt.Printf("        public var input: Request {\n")

	if len(fields) == 0 {
		fmt.Printf("            return Request()\n")
	} else {
		fmt.Printf("            return Request.with {\n")

		for _, field := range fields {
			valueName := valueName(field)
			fmt.Printf("                $0.%v = %v\n", valueName, valueName)
		}
		fmt.Printf("            }\n")
	}

	fmt.Printf("        }\n")
}
