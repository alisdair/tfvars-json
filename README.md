# tfvars-json

`tfvars-json` reads [Terraform](https://github.com/hashicorp/terraform) [`.tfvars` variable definition files](https://www.terraform.io/docs/language/values/variables.html#variable-definitions-tfvars-files) and outputs the same values in Terraform JSON syntax.

## Usage

```shellsession
$ tfvars-json file.tfvars > file.tfvars.json
```

`tfvars-json` processes a single file and outputs to stdout. You can use a JSON scripting tool like [`jq`](https://stedolan.github.io/jq/) to filter or post-process this output.

## Installation

```shellsession
$ go get github.com/alisdair/tfvars-json
```
