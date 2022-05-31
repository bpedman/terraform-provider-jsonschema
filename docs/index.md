# Json Schema Provider

Terraform provider for validating json files using [json-schema](https://json-schema.org/).

## Example Usage

```hcl-terraform
#: Validate values file
data "jsonschema_validator" "values" {
  document = file("/path/to/document.json")
  schema = "/path/to/schema.json"
}

#: Install a helm release with the validated json
resource "helm_release" "service_overview" {
  values = [
    data.jsonschema_validator.values.validated,
  ]
}
```
