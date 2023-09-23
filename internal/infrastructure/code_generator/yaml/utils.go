package yaml

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

func createMethodEntity(tag, operationID, responseFileName, requestFileName, errFileName string) *Method {
	var parameters []*Ref
	var requestBody *RequestBody
	responses := map[string]*Response{
		"500": {
			Description: "internal error",
			Content: map[string]*Content{
				"application/json": {
					Schema: &Schema{
						Ref: errFileName,
					},
				},
			},
		},
	}
	if responseFileName != "" {
		responses["200"] = &Response{
			Description: "Success",
			Content: map[string]*Content{
				"application/json": {
					Schema: &Schema{
						Type: "array",
						Items: &Ref{
							Ref: responseFileName,
						},
					},
				},
			},
		}
	} else {
		responses["201"] = &Response{Description: "Success"}
	}
	if requestFileName != "" {
		requestBody = &RequestBody{
			Description: "TODO: Change me",
			Required:    true,
			Content: map[string]*Content{
				"application/json": {
					Schema: &Schema{
						Ref: requestFileName,
					},
				},
			},
		}
	}
	return &Method{
		Summary:     "TODO: Change me",
		OperationID: operationID,
		Description: "TODO: Change me",
		Tags:        []string{tag},
		Parameters:  parameters,
		RequestBody: requestBody,
		Responses:   responses,
	}
}

func generateRequestYAML(filename string) {
	content := Req{
		Properties: Properties{
			"id": Property{
				Type:        "string",
				Description: "ID",
			},
		},
		Required: []string{"id"},
	}
	if err := writeYAML(filename, content); err != nil {
		panic(fmt.Sprintf("Failed to write file %s: %v\n", filename, err))
	}
	fmt.Printf("\tgenerated: %s\n", filename)
}

func generateResponseYAML(filename string) {
	content := Res{
		Properties: Properties{
			"id": Property{
				Type:        "string",
				Description: "ID",
			},
		},
		Required: []string{"id"},
	}
	if err := writeYAML(filename, content); err != nil {
		panic(fmt.Sprintf("Failed to write file %s: %v\n", filename, err))
	}
	fmt.Printf("\tgenerated: %s\n", filename)
}

func writeYAML(filename string, content interface{}) error {
	dir := filepath.Dir(filename)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := yaml.NewEncoder(file)
	encoder.SetIndent(2)
	if err := encoder.Encode(content); err != nil {
		return err
	}

	return nil
}
