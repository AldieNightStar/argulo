package argulo

import (
	"fmt"
	"os"
	"strings"
)

type ArguloBuilder struct {
	name     string
	descs    map[string]string
	samples  map[string][]string
	lastname string
	required []string
}

func New(name string) *ArguloBuilder {
	return &ArguloBuilder{
		name:     name,
		descs:    make(map[string]string),
		samples:  make(map[string][]string),
		required: make([]string, 0, 8),
	}
}

func (b *ArguloBuilder) Param(name, description string) *ArguloBuilder {
	b.descs[name] = description
	b.lastname = name
	return b
}

func (b *ArguloBuilder) RequiredParam(name, description string) *ArguloBuilder {
	b.descs[name] = description + " (required)"
	b.lastname = name
	b.required = append(b.required, name)
	return b
}

func (b *ArguloBuilder) Sample(str string) *ArguloBuilder {
	b.samples[b.lastname] = append(b.samples[b.lastname], str)
	return b
}

func (b *ArguloBuilder) usageString() string {
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("%s Usage:\n", b.name))
	for name, desc := range b.descs {
		sb.WriteString(fmt.Sprintf("\t%s - %s\n", name, desc))
		samples, isPresent := b.samples[name]
		if isPresent {
			for _, sample := range samples {
				sb.WriteString(fmt.Sprintf("\t\t%s\n", sample))
			}
		}
	}
	return sb.String()
}

func (b *ArguloBuilder) Build() *Argulo {
	a := &Argulo{
		usage:    b.usageString(),
		required: b.required,
	}
	return a
}

// ======================

type Argulo struct {
	mp       argsmap
	usage    string
	required []string
}

func (a *Argulo) Parse(args []string) *Argulo {
	a.mp = toMap(args)
	return a
}

func (a *Argulo) ParseOs() *Argulo {
	a.mp = toMap(os.Args[1:])
	return a
}

func (a *Argulo) Get(name string) ([]string, bool) {
	s, ok := a.mp[name]
	return s, ok
}

func (a *Argulo) GetFirst(name string) (string, bool) {
	arr, ok := a.mp[name]
	if !ok {
		return "", false
	}
	return arr[0], true
}

func (a *Argulo) GetFirstOr(name string, def string) string {
	arr, ok := a.mp[name]
	if !ok {
		return def
	}
	return arr[0]
}

func (a *Argulo) IsPresent(name string) bool {
	_, ok := a.mp[name]
	return ok
}

func (a *Argulo) Usage() string {
	return a.usage
}

func (a *Argulo) PrintUsage() {
	fmt.Println(a.usage)
}

func (a *Argulo) IsRequiredParamsOk() (ok bool, absents []string) {
	for _, param := range a.required {
		if !a.IsPresent(param) {
			absents = append(absents, param)
		}
	}
	return len(absents) < 1, absents
}

func (a *Argulo) ValidateOk() bool {
	if a.IsPresent("help") {
		a.PrintUsage()
		return false
	}
	allPresent, absents := a.IsRequiredParamsOk()
	if !allPresent {
		fmt.Printf("[!] Required parameters missing: %s\n\n", strings.Join(absents, ", "))
		a.PrintUsage()
		return false
	}
	return true
}
