package main

import (
	"fmt"

	"Go-Project/src"
	"Go-Project/stores"
)

func main() {
	store := &stores.Stub{}

	svc := src.New(store)

	dataset := svc.Store.GetDataset("cpih")
	fmt.Print(dataset)
}
