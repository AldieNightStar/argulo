package main

import (
	"fmt"

	"github.com/AldieNightStar/argulo"
)

func main() {
	a := argulo.New("tool").
		RequiredParam("name", "Specifies the name").
		Sample("name Ihor").
		Sample("name Xander").
		Param("age", "Set's the age").
		Sample("-age 18").
		Sample("-age 32").
		RequiredParam("option", "What to do").
		Sample("-option create").
		Sample("-option add").
		Sample("-option remove").
		Sample("-option kill").
		Sample("-option list").
		Build().
		ParseOs()
	if a.ValidateOk() {
		doOp(a)
	}
}

func doOp(a *argulo.Argulo) {
	fmt.Printf("%s is %s years old with option %s",
		a.GetFirstOr("name", "user"), a.GetFirstOr("age", "18"), a.GetFirstOr("option", "list"),
	)
}
