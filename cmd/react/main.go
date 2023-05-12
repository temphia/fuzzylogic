package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/temphia/fuzzylogic/fuzzyreact"
)

func main() {

	promot := strings.Join(os.Args[1:], " ")

	fmt.Println("running_react_start |>", promot)

	fr := fuzzyreact.NewReAct(os.Getenv("OPENAPI_KEY"))
	fmt.Println(fr.Execute(promot, 5))
	fmt.Println("running_react_end")

}
