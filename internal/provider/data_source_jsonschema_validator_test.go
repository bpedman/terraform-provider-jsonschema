package provider

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

func Test_dataSourceJsonschemaValidatorRead(t *testing.T) {
	var cases = []struct {
		document      string
		schema        string
		errorExpected bool
	}{
		{"asd asdasd: ^%^*&^%", "{}", true},
		{"{}", schemaValid, true},
		{`{"test": "test"}`, schemaValid, false},
	}

	for _, tt := range cases {
		t.Run(fmt.Sprintf("%s with error expected %t", tt.document, tt.errorExpected), func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck:          func() { testAccPreCheck(t) },
				ProviderFactories: providerFactories,
				Steps: []resource.TestStep{
					{
						Config: makeDataSource(t, tt.document, tt.schema),
						Check: resource.ComposeAggregateTestCheckFunc(
							resource.TestCheckResourceAttr("data.jsonschema_validator.test", "validated", fmt.Sprintf("%s\n", tt.document)),
						),
					},
				},
				ErrorCheck: func(err error) error {
					if tt.errorExpected {
						if err == nil {
							return fmt.Errorf("error expected")
						} else {
							return nil
						}
					}

					return err
				},
			})
		})
	}
}

func makeDataSource(t *testing.T, document string, schema string) string {
	tempDir := t.TempDir()
	schemaFile := fmt.Sprintf("%s/schema.json", tempDir)
	err := os.WriteFile(schemaFile, []byte(schema), 0644)
	if err != nil {
		t.Fatal(err)
	}

	return fmt.Sprintf(`
data "jsonschema_validator" "test" {
  document = <<EOF
%s
EOF
  schema   = "%s"
}
`, document, schemaFile)
}

var schemaValid = `{
  "$schema": "http://json-schema.org/draft-07/schema#",
  "x-$id": "https://example.com",
  "type": "object",
  "required": ["test"]
}`
