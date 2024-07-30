package transpiler

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/tidwall/pretty"
	"gopkg.in/yaml.v3"
)

const (
	CWLVersion = "v1.2"
)

func TranspileCommandlineTool(cl CommandlineTool, inputs map[string]CWLInputEntry, locations FileLocations, outputFile string) error {
	err := TypeCheckCommandlineTool(&cl, inputs)
	if err != nil {
		return err
	}
	wf, err := EmitCommandlineTool(&cl, inputs, locations)
	if err != nil {
		return err
	}
	// HACK: yaml Marshalling doesn't marshal correctly
	// therefore we turn the Workflow to map[string]interface and marshal that
	data, err := json.Marshal(wf)
	if err != nil {
		return err
	}

	data = pretty.Pretty(data)
	fmt.Println(string(data))

	m := make(map[string]interface{})
	err = json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	data, err = yaml.Marshal(m)
	if err != nil {
		return err
	}
	return os.WriteFile(outputFile, data, 0644)
}

func TranspileCWLWorkflow(workflow Workflow, inputs map[string]CWLInputEntry, locations FileLocations, outputFile string) error {

	// Check the workflow provided
	err := TypeCheckWorkflow(&workflow, inputs)
	if err != nil {
		return err
	}

	// Convert the CWL Workflow to Argo Workflows
	wf, err := EmitWorkflow(&workflow, inputs, locations)
	if err != nil {
		return err
	}

	data, err := json.Marshal(wf)
	if err != nil {
		return err
	}

	data = pretty.Pretty(data)
	fmt.Println(string(data))

	m := make(map[string]interface{})
	err = json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	data, err = yaml.Marshal(m)
	if err != nil {
		return err
	}

	return os.WriteFile(outputFile, data, 0644)
}

func TranspileFile(inputFile string, inputsFile string, locationsFile *string, outputFile string) error {

	log.Warn("Currently the transpiler expects preprocessed CWL input, sbpack is the recommended way of preprocessing, use cwlpack from here https://github.com/rabix/sbpack/tree/84bd7867a0630a826280a702db715377aa879f6a")

	var cwl map[string]interface{}
	var inputs map[string]CWLInputEntry
	var fileLocations FileLocations

	def, err := os.ReadFile(inputFile)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(def, &cwl)
	if err != nil {
		return err
	}

	data, err := os.ReadFile(inputsFile)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(data, &inputs)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(data, &inputs)
	if err != nil {
		return err
	}

	if locationsFile != nil {
		data, err := os.ReadFile(*locationsFile)
		if err != nil {
			return err
		}
		err = json.Unmarshal(data, &fileLocations)
		if err != nil {
			return err
		}
	}

	class, ok := cwl["class"]
	if !ok {
		return errors.New("<class> expected")
	}

	// Add support for Transpiling workflows
	if class == "CommandLineTool" {
		var cliTool CommandlineTool
		err := yaml.Unmarshal(def, &cliTool)
		if err != nil {
			return err
		}

		return TranspileCommandlineTool(cliTool, inputs, fileLocations, outputFile)
	} else if class == "Workflow" {

		var workflow Workflow
		err := yaml.Unmarshal(def, &workflow)
		if err != nil {
			return err
		}

		fmt.Printf("inputsFile %+v\n", inputsFile)
		fmt.Printf("inputs %+v\n", inputs)

		return TranspileCWLWorkflow(workflow, inputs, fileLocations, outputFile)
	} else {
		return fmt.Errorf("%s is not supported as of yet", class)
	}

}