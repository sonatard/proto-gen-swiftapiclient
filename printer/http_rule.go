package printer

import (
	"fmt"

	"github.com/golang/protobuf/proto"
	"github.com/jhump/protoreflect/desc"
	"golang.org/x/xerrors"
	"google.golang.org/genproto/googleapis/api/annotations"
)

func httpRule(method *desc.MethodDescriptor) (*annotations.HttpRule, error) {
	opts := method.GetOptions()

	if !proto.HasExtension(opts, annotations.E_Http) {
		return nil, fmt.Errorf("http rule not found")
	}

	ext, err := proto.GetExtension(opts, annotations.E_Http)
	if err != nil {
		return nil, xerrors.Errorf(": %w", err)
	}

	rule, ok := ext.(*annotations.HttpRule)
	if !ok {
		return nil, xerrors.New("annotation extension assertion error")
	}

	return rule, nil
}

type endpoint struct {
	method string
	path   string
}

func newEndpoint(rule *annotations.HttpRule) (*endpoint, error) {
	var e *endpoint
	switch opt := rule.GetPattern().(type) {
	case *annotations.HttpRule_Get:
		e = &endpoint{"GET", opt.Get}
	case *annotations.HttpRule_Put:
		e = &endpoint{"PUT", opt.Put}
	case *annotations.HttpRule_Post:
		e = &endpoint{"POST", opt.Post}
	case *annotations.HttpRule_Delete:
		e = &endpoint{"DELETE", opt.Delete}
	case *annotations.HttpRule_Patch:
		e = &endpoint{"PATCH", opt.Patch}
	default:
		return nil, xerrors.New("annotation http rule method dose not support type")
	}

	return e, nil
}
