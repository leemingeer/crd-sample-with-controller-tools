// +groupName=ming.k8s.io
package v1

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var (
	Scheme = runtime.NewScheme()
	GroupVersion = schema.GroupVersion{Group: "ming.k8s.io", Version: "v1"}
)

func init(){
	Scheme.AddKnownTypes(GroupVersion, &Bar{}, &BarList{})
}