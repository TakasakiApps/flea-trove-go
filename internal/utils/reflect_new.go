package utils

import (
	"github.com/hanakogo/hanakoutilgo"
	"reflect"
)

func newValuePtrByType(p reflect.Type) any {
	return reflect.New(p).Interface()
}

func NewValuePtr[T any]() *T {
	ptrType := reflect.TypeOf((*T)(nil)).Elem()
	ptr := newValuePtrByType(ptrType)
	return hanakoutilgo.CastTo[*T](ptr)
}
