# Demo: Terraform Provider Generation

Using the HashiCorp Terraform Provider Generator.

## Usage

### Scaffolding

``` sh
# Setup the repository and the tooling
make setup

# Generate the provider boilerplate code
make scaffold

# Start coding
edit main.go
```

### Develop

``` sh
# Build Terraform Provider
make build
```

## References

- [Demo Workflow](https://developer.hashicorp.com/terraform/plugin/code-generation/workflow-example)
- [Framework Generator](https://developer.hashicorp.com/terraform/plugin/code-generation/framework-generator#installation)
