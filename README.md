# terraform-provider-administration
The Terraform provider administration is a plugin for Terraform that allows some administration configuration of quortex solution.
This provider is maintained internally by the Quortex team.

## Documentation

Full, comprehensive documentation is available on the Terraform website:

https://registry.terraform.io/providers/quortex/administration/latest/docs

## Build provider

Run the following command to build the provider

```shell
$ go build -o terraform-provider-administraition
```

## Test sample configuration

First, build and install the provider.

```shell
$ make install
```

Then, navigate to the `examples/development` directory.

```shell
$ cd examples/development
```

Then update the following variables according to your needs in the provider "adminisatration", for instance in your `main.tf`
```jsonc
provider "administration" {
  auth_server = "https://example.auth.server" //can be omitted to use default
  host        = "http://localhost:8000" //can be omitted to use default
  client_id     = "my_client_id"
  client_secret = "my_client_secret"
}

```
Run the following command to initialize the workspace and apply the sample configuration.

```shell
$ terraform init && terraform apply
```

## Tests acceptance

Test was inspired from this repo : https://github.com/hashicorp/terraform-provider-scaffolding-framework/tree/main

