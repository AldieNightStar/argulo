# Argulo

### Framework for your argument-based CLI

# Usage
```go
a := argulo.New("name").
	Param("param", "This is my param").
	Sample("-param a").
	Sample("-param b").
	RequiredParam("param2", "This is required param").
	Sample("-param2 data").
	Build().
	ParseOs() // or Parse(args)

if a.ValidateOk() {
	// Do something here
}
```

# Sample
```go
package main

import (
	"fmt"

	"github.com/AldieNightStar/argulo"
)

func main() {
    // Don't forget to put '.' at the end to make a chain

    // Create new argument parser for "tool"
    // a - is new parser for arguments
	a := argulo.New("tool").
        // This param is required
        // We add param and samples to it
		RequiredParam("name", "Specifies the name").
		Sample("name Ihor").
		Sample("name Xander").
        // Second param is NOT required
        // We add param and samples to it
		Param("age", "Set's the age").
		Sample("-age 18").
		Sample("-age 32").
        // This param is required
        // We add param and a lot of samples to it
		RequiredParam("option", "What to do").
		Sample("-option create").
		Sample("-option add").
		Sample("-option remove").
		Sample("-option kill").
		Sample("-option list").
        // Now we done with parser creation
		Build().
        // Now we need to get argument to parse
        // ParseOs()   - parse from Operation system
        // Parse(args) - parse from array of strings
		ParseOs()

    // If all variables are present and no -help - it will call doOp
	// Otherwise it will show error and print the usage
	if a.ValidateOk() {
		doOp(a)
	}
}

func doOp(a *argulo.Argulo) {
	fmt.Printf("%s is %s years old with option %s",
		a.GetFirstOr("name", "user"), a.GetFirstOr("age", "18"), a.GetFirstOr("option", "list"),
	)
}

```