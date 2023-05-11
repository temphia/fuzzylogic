package main

import (
	"fmt"
	"os"

	"github.com/temphia/fuzzylogic/fuzzyreact"
)

func main() {

	fmt.Println("@hello")
	fr := fuzzyreact.New(os.Getenv("OPENAPI_KEY"))

	fmt.Println(fr.Execute("tell me about nepal", 5))

}
