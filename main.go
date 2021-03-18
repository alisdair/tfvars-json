package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	ctyjson "github.com/zclconf/go-cty/cty/json"
)

func main() {
	args := os.Args[1:]

	if len(args) != 1 {
		fmt.Fprintf(os.Stderr, "Usage: tfvars-json file.tfvars\n")
		os.Exit(1)
	}

	filename := args[0]
	src, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading %q: %s", filename, err)
		os.Exit(1)
	}

	f, diags := hclsyntax.ParseConfig(src, filename, hcl.Pos{Line: 1, Column: 1})
	if diags.HasErrors() {
		fmt.Fprintf(os.Stderr, "Error parsing %q: %s", filename, diags.Error())
		os.Exit(1)
	}
	if f.Body == nil {
		fmt.Println("{}")
		os.Exit(0)
	}

	attrs, diags := f.Body.JustAttributes()
	if diags.HasErrors() {
		fmt.Fprintf(os.Stderr, "Error evaluating %q: %s", filename, diags.Error())
		os.Exit(1)
	}

	values := make(map[string]json.RawMessage, len(attrs))
	for name, attr := range attrs {
		value, diags := attr.Expr.Value(nil)
		if diags.HasErrors() {
			fmt.Fprintf(os.Stderr, "Error evaluating %q in %q: %s", name, filename, diags.Error())
			os.Exit(1)
		}
		buf, err := ctyjson.Marshal(value, value.Type())
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error converting %q in %q to JSON: %s", name, filename, err)
			os.Exit(1)
		}
		values[name] = json.RawMessage(buf)
	}

	output, err := json.MarshalIndent(values, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshalling %q to JSON: %s", filename, err)
		os.Exit(1)
	}
	fmt.Println(string(output))

	os.Exit(0)
}
