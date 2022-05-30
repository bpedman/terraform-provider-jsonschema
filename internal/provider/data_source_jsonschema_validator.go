package provider

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/santhosh-tekuri/jsonschema/v5"
)

func dataSourceJsonschemaValidator() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceJsonschemaValidatorRead,

		Schema: map[string]*schema.Schema{
			"document": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "body of a json document to validate as a string",
			},

			"schema": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "file path or url to a schema document to use for validation",
			},

			"validated": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceJsonschemaValidatorRead(d *schema.ResourceData, m interface{}) error {
	var (
		err error = nil
	)

	document := d.Get("document").(string)
	schemaPathOrUrl := d.Get("schema").(string)

	compiledSchema, err := jsonschema.Compile(schemaPathOrUrl)
	if err != nil {
		return fmt.Errorf("error parsing schema definition: %v", err)
	}

	var parsedDocument interface{}
	if err := json.Unmarshal([]byte(document), &parsedDocument); err != nil {
		return fmt.Errorf("error parsing provided document as json: %v", err)
	}

	err = compiledSchema.Validate(parsedDocument)
	if err != nil {
		return fmt.Errorf("document is not valid:\n%#v", err)
	}

	err = d.Set("validated", document)
	if err != nil {
		return fmt.Errorf("internal error setting validated document: %v", err)
	}

	d.SetId(hash(document))
	return nil
}

func hash(s string) string {
	sha := sha256.Sum256([]byte(s))
	return hex.EncodeToString(sha[:])
}
