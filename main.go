package main

import (
	"fmt"
	"project/src"
)

func main() {
	svc := src.Service.New(models.stub.stubStore)

	fmt.Print(svc.store.get_dataset("cpih"))
	fmt.Print(svc.store.get_datasets())

}
