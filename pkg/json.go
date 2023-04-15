package pkg

import (
	"encoding/json"
	"fmt"

	"cloud.google.com/go/bigquery"
	"github.com/ahnsv/schemesis/types"
)

func CreateJsonSchema(schema bigquery.Schema) ([]byte, error) {
	var jsonSchema types.JSONSchema

	jsonSchema.Title = "Generated JSON Schema for BigQuery schema"
	jsonSchema.Type = "object"
	jsonSchema.Description = "This JSON schema is generated from a BigQuery schema."
	jsonSchema.Properties = make(map[string]interface{})
	var requiredFields []string

	for _, field := range schema {
		var fieldType string
		switch field.Type {
		case bigquery.StringFieldType:
			fieldType = "string"
		case bigquery.BytesFieldType:
			fieldType = "string"
		case bigquery.IntegerFieldType:
			fieldType = "integer"
		case bigquery.FloatFieldType:
			fieldType = "number"
		case bigquery.BooleanFieldType:
			fieldType = "boolean"
		case bigquery.TimestampFieldType:
			fieldType = "string"
		case bigquery.DateFieldType:
			fieldType = "string"
		case bigquery.TimeFieldType:
			fieldType = "string"
		case bigquery.DateTimeFieldType:
			fieldType = "string"
		case bigquery.RecordFieldType:
			fieldType = "object"
		case bigquery.NumericFieldType:
			fieldType = "number"
		}
		jsonSchema.Properties[field.Name] = map[string]interface{}{
			"type": fieldType,
		}
		if field.Required {
			requiredFields = append(requiredFields, field.Name)
		}
	}
	if len(requiredFields) > 0 {
		jsonSchema.Required = requiredFields
	}
	jsonBytes, err := json.Marshal(jsonSchema)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return jsonBytes, nil
}
