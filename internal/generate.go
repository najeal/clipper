package pkg

import (
	"fmt"
	"go/format"

	"github.com/najeal/clipper/internal/specs"
	"github.com/najeal/clipper/internal/templater"
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
	content, err := templater.GenerateContent(templateRoot)
	if err != nil {
		return nil, fmt.Errorf("failed to apply template on specs: %w", err)
	}
	formatedContent, err := format.Source(content)
	if err != nil {
		return nil, fmt.Errorf("failed to format generated cli go code: %w", err)
	}
	return formatedContent, nil
}
