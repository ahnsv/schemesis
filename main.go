package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"cloud.google.com/go/bigquery"
	"github.com/brianvoe/gofakeit/v6"
)

type JSONSchema struct {
	Title       string                 `json:"title"`
	Type        string                 `json:"type"`
	Description string                 `json:"description,omitempty"`
	Properties  map[string]interface{} `json:"properties"`
	Required    []string               `json:"required,omitempty"`
}

func GenerateFakeData(schemaBytes []byte, length int) ([]map[string]interface{}, error) {
	var schema JSONSchema
	err := json.Unmarshal(schemaBytes, &schema)
	if err != nil {
		return nil, err
	}

	var data []map[string]interface{}
	for i := 0; i < length; i++ {
		item := make(map[string]interface{})
		for propertyName, propertySchema := range schema.Properties {
			propertyType := propertySchema.(map[string]interface{})["type"].(string)
			switch propertyType {
			case "string":
				item[propertyName] = gofakeit.Word()
			case "number":
				item[propertyName] = gofakeit.Float64()
			case "integer":
				item[propertyName] = gofakeit.Int64()
			case "boolean":
				item[propertyName] = gofakeit.Bool()
			}
		}
		data = append(data, item)
	}

	return data, nil
}

func generate(schema []byte, length int) {
	data, err := GenerateFakeData(schema, length)
	if err != nil {
		fmt.Println(err)
		return
	}
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(jsonBytes))
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go [input file]")
		return
	}
	filename := os.Args[1]
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	var schema bigquery.Schema
	err = json.Unmarshal(bytes, &schema)
	if err != nil {
		fmt.Println(err)
		return
	}
	var jsonSchema JSONSchema
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
		return
	}
	fmt.Println(string(jsonBytes))

	generate(jsonBytes, 3)
}
