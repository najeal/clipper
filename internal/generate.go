package pkg

import (
	"fmt"

	"github.com/najeal/go-clipper/internal/specs"
	"github.com/najeal/go-clipper/internal/templater"
)

func NewCliGenerator() *CliGenerator {
	return &CliGenerator{}
}

type CliGenerator struct{}

func (*CliGenerator) Generate(inputFile []byte, packageName string) (outputFile []byte, err error) {
	specs, err := specs.Unmarshal(inputFile)
	if err != nil {
		return nil, err
	}
	templateRoot := transform(packageName, specs)
	outputFile, err = templater.GenerateContent(templateRoot)
	if err != nil {
		return nil, fmt.Errorf("failed to apply template on specs: %w", err)
	}
	return outputFile, nil
}
