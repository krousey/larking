// lark
package main

import (
	"flag"
	"log"

	"github.com/emcfarlane/larking"
	"github.com/emcfarlane/starlarkrepl"
	"go.starlark.net/repl"
	"go.starlark.net/starlark"
	"go.starlark.net/starlarkstruct"
)

func run() error {
	flag.Parse()

	mux, err := larking.NewMux()
	if err != nil {
		return err
	}
	thread := &starlark.Thread{Load: repl.MakeLoad()}
	globals := starlark.StringDict{
		"struct": starlark.NewBuiltin("struct", starlarkstruct.Make),
		"grpc":   larking.NewModule(mux),
	}
	options := starlarkrepl.Options{AutoComplete: true}

	return starlarkrepl.Run(thread, globals, options)
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}