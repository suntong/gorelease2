package main

import (
	"fmt"

	shared "github.com/suntong/gorelease2"
)

func main() {
	fmt.Printf("From mainB(), V = '%s'\n", shared.V)
}
