package commands

import (
	"errors"
	"fmt"
	"path/filepath"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/argoproj/argo-workflows/v3/cmd/argo/commands/transpiler"
)

func lengthError() error {
	return errors.New("length of filename should always be greater than the length of the extension")
}

func invalidFileError(inputFileExt string) error {
	return fmt.Errorf("invalid file extension %s, only common workflow language (.cwl) files are allowed", inputFileExt)
}
func extractNoExtFileName(filename string, ext string) (string, error) {
	if len(filename) <= len(ext) {
		return "", lengthError()
	}
	name := filename[0 : len(filename)-len(ext)]
	return name, nil
}

func processFile(inputFile string, inputsFile string, locationsFile *string) {

	ext := filepath.Ext(inputFile)
	if ext != ".cwl" {
		log.Fatalf("%+v", invalidFileError(ext))
	}
	name, err := extractNoExtFileName(inputFile, ext)
	if err != nil {
		log.Fatalf("%+v", err)
	}

	// TODO If your file is in a sub dir, this won't work
	newName := fmt.Sprintf("%s_argo.yaml", name)
	log.Infof("Transpiling file %s with extension %s and ext free name %s to %s", inputFile, ext, name, newName)
	err = transpiler.TranspileFile(inputFile, inputsFile, locationsFile, newName)
	if err != nil {
		log.Fatalf("%+v", err)
	}

}

func NewTranspileCommand() *cobra.Command {

	command := &cobra.Command{
		Use:   "transpile [CWL...]",
		Short: "Transpile common workflow language file to Argo Workflow yaml",
		Example: `# Wait on a workflow:

		argo transpile my-cwl-file.cwl

		`,
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) != 2 && len(args) != 3 {
				return errors.New("transpile accepts at least two arguments <WORKFLOW.cwl> and <INPUTS.(yml|json|cwl)> and optionally <LOCATION.(yml.json)>")
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			var locationsFile *string
			locationsFile = nil

			if len(args) == 3 {
				locationsFile = &args[2]
			}
			processFile(args[0], args[1], locationsFile)
		},
	}

	return command
}
