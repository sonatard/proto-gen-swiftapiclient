package printer

import (
	"strings"

	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/iancoleman/strcase"
	"github.com/jhump/protoreflect/desc"
	"github.com/sonatard/proto-gen-swiftapiclient/str"
)

func valueName(field *desc.FieldDescriptor) string {
	return str.ToAbbreviation(strcase.ToLowerCamel(field.GetName()))
}

var swiftTypeMaps = map[descriptor.FieldDescriptorProto_Type]string{
	1:  "Double",
	2:  "Float",
	3:  "Int64",
	4:  "Uint64",
	5:  "Int32",
	6:  "Double",
	7:  "Float",
	8:  "Bool",
	9:  "String",
	10: "TYPE_GROUP",
	11: "TYPE_MESSAGE",
	12: "TYPE_BYTES",
	13: "Uint32",
	14: "TYPE_ENUM",
	15: "Float",
	16: "Double",
	17: "Int32",
	18: "Int64",
}

func typeName(packageName string, field *desc.FieldDescriptor) string {
	var typeName string
	fieldType := field.GetType()
	switch fieldType {
	case descriptor.FieldDescriptorProto_TYPE_GROUP:
		panic("group not supported")
	case descriptor.FieldDescriptorProto_TYPE_MESSAGE:
		typeName = field.GetMessageType().GetName()
	case descriptor.FieldDescriptorProto_TYPE_ENUM:
		typeName = strings.Trim(field.GetEnumType().GetFullyQualifiedName(), packageName)
	default:
		typeName = swiftTypeMaps[field.GetType()]
	}

	typeName = str.ToAbbreviation(typeName)

	if field.IsRepeated() {
		return "[" + typeName + "]"
	}

	return typeName
}
