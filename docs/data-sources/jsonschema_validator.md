# jsonschema_validator Data Source

The `jsonschema_validator` data source validates a json document using [json-schema](https://json-schema.org/).

## Example Usage

```hcl-terraform
data "jsonschema_validator" "values" {
  document = file("/path/to/document.json")
  schema = "/path/to/schema.json"
}
```

## Argument Reference

List arguments this data source takes:

* `document` &mdash; (Required) Content of a json document.
* `schema` &mdash; (Required) File path or file URL to a [json-schema](https://json-schema.org/) document.

## Attributes Reference

List attributes that this data source exports:

* `validated` &mdash; equivalent to `document` argument.
