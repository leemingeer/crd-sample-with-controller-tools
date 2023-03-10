package main

import (
	"context"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/tools/clientcmd"
	"log"
)


func main(){
	config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		log.Fatalln(err)
	}

	dclient, err := dynamic.NewForConfig(config)
	if err != nil{
		log.Fatalln(err)
	}
	gvr := schema.GroupVersionResource{Group: "ming.k8s.io", Version: "v1", Resource: "bars"}
	res, err := dclient.Resource(gvr).Namespace("default").Get(context.TODO(), "test-crd", metav1.GetOptions{})
    if err != nil{
		log.Fatalln(err)
	}

	fmt.Printf("bars resource: %#v", res)
}