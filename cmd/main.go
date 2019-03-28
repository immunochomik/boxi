package main

import (
	"boxi"
	"fmt"
)

func main() {
	fmt.Printf("sanity!!!!:\n")
	s, err := boxi.Secret()
	if err != nil {
		panic(err)
	}
	k, err := boxi.PubKey()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", s)
	fmt.Printf("%s\n", k)

}
