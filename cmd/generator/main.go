package main

import (
	"fmt"
	"github.com/hum2/backend/internal/infrastructure/code_generator/golang"
	"github.com/hum2/backend/internal/infrastructure/code_generator/yaml"
	"github.com/spf13/cobra"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"net/http"
	"os"
	"strings"
)

const (
	// レイヤー名
	controllerLayer = "controller"
	domainLayer     = "domain"
	repositoryLayer = "repository"

	// コントローラ関連設定
	goModulePath         = "github.com/hum2/backend/internal"
	goCtrFileFormat      = "internal/interface/controller/%s/%s_controller.go"
	goCtrGenPathFormat   = "%s/interface/controller/%s/gen"
	goUsecasePathFormat  = "github.com/hum2/backend/internal/usecase/%s"
	goUsecaseFileFormat  = "internal/interface/usecase/%s/%s_usecase.go"
	domainPathFormat     = "github.com/hum2/backend/internal/domain/%s"
	domainFileFormat     = "internal/domain/%s/%s.go"
	domainRepoFileFormat = "internal/domain/%s/%s_repository.go"
	infraRepoFileFormat  = "internal/infrastructure/repository/%s/%s_repository.go"
	// OAPI関連設定
	errFilenameFormat      = "/schema/response/_shared/ErrorResponse.yaml"
	assetsPathPrefix       = "openapi/assets"
	pathFilenameFormat     = "%s/path/%s/%s.yaml"
	responseFilenameFormat = "%s/schema/response/%s/%sResponse.yaml"
	requestFilenameFormat  = "%s/schema/request/%s/%sRequest.yaml"
)

func main() {
	var methods []string
	var rootCmd = &cobra.Command{
		Use:   "generator <layer name> <tag name>",
		Short: "Generator is a CLI tool for generating code",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			run(args, methods)
		},
	}

	rootCmd.PersistentFlags().StringSliceVar(
		&methods,
		"method",
		[]string{http.MethodGet},
		"HTTP method")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run(args []string, methods []string) {
	layerName := args[0]
	tagName := args[1]

	gg := golang.NewGenerator(
		tagName,
		convTagName2CamelCase(tagName),
		goModulePath,
		goCtrFileFormat,
		goCtrGenPathFormat,
		goUsecasePathFormat,
		goUsecaseFileFormat,
		domainPathFormat,
		domainFileFormat,
		domainRepoFileFormat,
		infraRepoFileFormat)
	yg := yaml.NewGenerator(
		tagName,
		convTagName2CamelCase(tagName),
		assetsPathPrefix,
		pathFilenameFormat,
		responseFilenameFormat,
		requestFilenameFormat,
		errFilenameFormat)
	switch layerName {
	case controllerLayer:
		fmt.Println("generate go files")
		generatedFiles, err := gg.GenerateController(methods)
		if err != nil {
			panic(err)
		}
		for _, createdFile := range generatedFiles {
			fmt.Printf("\tgenerated: %s\n", createdFile)
		}

		fmt.Println("generate yaml files")
		generatedFiles, err = yg.GenerateOAPI(methods)
		if err != nil {
			panic(err)
		}
		for _, createdFile := range generatedFiles {
			fmt.Printf("\tgenerated: %s\n", createdFile)
		}
	case domainLayer:
		gg.GenerateDomain()
	case repositoryLayer:
		gg.GenerateRepository()
	default:
		help()
	}
}

func help() {
	fmt.Println("Usage: generator <tag name>")
	os.Exit(1)
}

func convTagName2CamelCase(tagName string) string {
	parts := strings.Split(tagName, "/")

	for i, part := range parts {
		parts[i] = cases.Title(language.Und, cases.NoLower).String(part)
	}

	return strings.Join(parts, "")
}
