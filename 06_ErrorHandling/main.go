package main

import (
	"fmt"
	"log"

	"github.com/juju/errors"
)

func foo() error {
	return errors.Annotate(bar(), "failed to attempt foo")
}

func bar() error {
	return errors.New("failed to attempt bar")
}

func main() {
	fmt.Println(foo())

	log.Println(errors.Details(foo()))

	fmt.Printf("%s\n", errors.ErrorStack(foo()))
}
