package icobra

import (
	"fmt"
	"go/format"
	"os"
	"path/filepath"

	"github.com/najeal/clipper/icobra/specs"
	"github.com/najeal/clipper/icobra/templater"
)

func NewCliGenerator() *CliGenerator {
	return &CliGenerator{}
}

type CliGenerator struct{}

func (*CliGenerator) Generate(inputFile []byte, packageName string) (outputFiles []byte, err error) {
	specs, err := specs.Unmarshal(inputFile)
	if err != nil {
		return nil, err
	}
	_ = specs
	dataWrapper := transform(packageName, specs)
	rootData := dataWrapper.RootData
	content, err := templater.GenerateRoot(rootData)
	if err != nil {
		return nil, fmt.Errorf("failed to apply template on specs: %w", err)
	}
	formatedContent, err := format.Source(content)
	if err != nil {
		return nil, fmt.Errorf("failed to format generated cli go code: %w", err)
	}

	if err := os.MkdirAll("cmd", os.FileMode(0o755)); err != nil {
		return nil, fmt.Errorf("failed to create cmd directory: %s", err.Error())
	}
	err = os.WriteFile(
		filepath.Join("cmd", rootData.Command.CommandName+".go"),
		formatedContent,
		os.FileMode(0o755),
	)
	if err != nil {
		return nil, fmt.Errorf("cannot write file: %s", err.Error())
	}

	for _, command := range dataWrapper.Commands {
		content, err := templater.GenerateCommand(command)
		if err != nil {
			return nil, fmt.Errorf("failed to apply template on command spec: %w", err)
		}
		formatedContent, err := format.Source(content)
		if err != nil {
			return nil, fmt.Errorf("failed to format generated cli command go code: %w", err)
		}
		err = os.WriteFile(
			filepath.Join("cmd", command.CommandName+".go"),
			formatedContent,
			os.FileMode(0o755),
		)
		if err != nil {
			return nil, fmt.Errorf("cannot write file: %s", err.Error())
		}
	}

	return formatedContent, nil
}
