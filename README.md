# Extend your PATH

Add to your `~/.bashrc` or `~/.zshrc`

	export PATH=$PATH:$(go env GOPATH)/bin

Then reload zsh/bash

	source ~/.zshrc

# Installation

[Reference](https://golang.org/doc/code.html)

```bash
go get github.com/hashicorp/terraform
go get github.com/terraform-providers/terraform-provider-template

go install github.com/hashicorp/terraform
go install github.com/terraform-providers/terraform-provider-template

go install .

go build -o terraform-provider-netbox
```
