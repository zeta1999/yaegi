package main

import (
	"fmt"
	"net/http"
)

type Fromage struct {
	*http.Server
}

func main() {
	a := Fromage{}
	fmt.Println(a.Server)
}

// Output:
// <nil>
