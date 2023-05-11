package main

import (
	"fmt"

	"github.com/temphia/fuzzylogic/fuzzyreact"
)

func main() {

	fmt.Println("@hello")
	fr := fuzzyreact.New("")

	fmt.Println(fr.Execute("tell me about kathmandu", "", 5))

}
