package golang

import (
	"fmt"
	. "github.com/dave/jennifer/jen"
	"strings"
)

const (
	ginPath = "github.com/gin-gonic/gin"

	controllerCommentFormat = `//go:generate oapi-codegen -include-tags {{.TagName}} -generate types -o ./gen/types.gen.go  -package gen $PROJECT_DIR/openapi/oapi-codegen.gen.yaml
//go:generate oapi-codegen -include-tags {{.TagName}} -generate gin   -o ./gen/server.gen.go -package gen $PROJECT_DIR/openapi/oapi-codegen.gen.yaml`
	mockgenComment          = `//go:generate mockgen -package=$GOPACKAGE -source=$GOFILE -destination=mock_$GOFILE`
	domainRepositoryComment = `//go:generate mockgen -package=$GOPACKAGE -source=$GOFILE -destination=mock_$GOFILE
//go:generate goimports mock_$GOFILE
//go:generate gofmt -w mock_$GOFILE`
)

type Generator interface {
	GenerateController(methods []string) (generatedFiles []string, err error)
	GenerateDomain()
	GenerateRepository()
}

type generator struct {
	tagName              string
	tagNameCamelCase     string
	modulePath           string
	ctrFileFormat        string
	genPathFormat        string
	usecasePathFormat    string
	usecaseFileFormat    string
	domainPathFormat     string
	domainFileFormat     string
	domainRepoFileFormat string
	infraRepoFileFormat  string
}

func NewGenerator(
	tagName string,
	tagNameCamelCase string,
	modulePath string,
	ctrFileFormat string,
	genPathFormat string,
	usecasePathFormat string,
	usecaseFileFormat string,
	domainPathFormat string,
	domainFileFormat string,
	domainRepoFileFormat string,
	infraRepoFileFormat string,
) Generator {
	return &generator{
		tagName:              tagName,
		tagNameCamelCase:     tagNameCamelCase,
		modulePath:           modulePath,
		ctrFileFormat:        ctrFileFormat,
		genPathFormat:        genPathFormat,
		usecasePathFormat:    usecasePathFormat,
		usecaseFileFormat:    usecaseFileFormat,
		domainPathFormat:     domainPathFormat,
		domainFileFormat:     domainFileFormat,
		domainRepoFileFormat: domainRepoFileFormat,
		infraRepoFileFormat:  infraRepoFileFormat,
	}
}

func (g *generator) GenerateController(methods []string) (generatedFiles []string, err error) {
	generatedFile, err := g.generateController(methods)
	if err != nil {
		return nil, err
	}
	generatedFiles = append(generatedFiles, generatedFile)

	generatedFile, err = g.generateUsecase()
	if err != nil {
		return nil, err
	}
	generatedFiles = append(generatedFiles, generatedFile)
	return
}

func (g *generator) GenerateDomain() {
	g.generateDomain()
	g.generateDomainRepository()
}

func (g *generator) GenerateRepository() {
	g.generateRepository()
}

func (g *generator) generateController(methods []string) (generatedFile string, err error) {
	packageName := getPackageName(g.tagName)
	appName := getAppName(g.tagName)
	pathName := getFileNamePrefix(g.tagName)
	generatedFile = fmt.Sprintf(g.ctrFileFormat, g.tagName, pathName)
	genPath := fmt.Sprintf(g.genPathFormat, g.modulePath, g.tagName)
	usecasePath := fmt.Sprintf(g.usecasePathFormat, g.tagName)

	f := NewFile(packageName)
	f.HeaderComment(strings.Replace(controllerCommentFormat, "{{.TagName}}", g.tagName, -1))
	f.ImportName(ginPath, "gin")
	f.ImportName(genPath, "gen")
	f.ImportAlias(usecasePath, appName)

	f.Type().Id("Controller").Struct(
		Id("usecase").Op("*").Qual(usecasePath, "usecase"),
	)
	f.Comment("New is a constructor of Controller").Line().
		Func().
		Id("New").Params(Id("u").Op("*").Qual(usecasePath, "usecase")).
		Qual(genPath, "ServerInterface").
		Block(
			Return(Op("&").Id("Controller").Values(
				Dict{
					Id("usecase"): Id("u"),
				},
			)),
		).Line()

	// Create controller methods
	for _, method := range methods {
		methodName := method + g.tagNameCamelCase
		f.Commentf("%s is a controller for %s /%s", methodName, strings.ToUpper(method), g.tagName).Line().
			Func().
			Params(Id("c").Op("*").Id("Controller")).
			Id(methodName).Params(Id("ctx").Op("*").Qual(ginPath, "Context")).
			Block(
				Comment("TODO: write your code"),
				Id("ctx").Dot("JSON").Call(
					Qual("net/http", "StatusOK"),
					Qual(ginPath, "H").Values(
						Dict{
							Lit("message"): Lit("pong"),
						}),
				),
			).Line()
	}

	if err := output(f, generatedFile); err != nil {
		return "", err
	}
	return
}

func (g *generator) generateUsecase() (string, error) {
	packageName := getPackageName(g.tagName)
	pathName := getFileNamePrefix(g.tagName)
	fileName := fmt.Sprintf(g.usecaseFileFormat, g.tagName, pathName)
	txPath := g.modulePath + "/usecase/shared/transaction"

	f := NewFile(packageName)
	f.HeaderComment(mockgenComment)

	f.ImportName("context", "context")
	f.ImportName(txPath, "transaction")

	f.Comment("Usecase is a usecase of user").Line().
		Type().Id("Usecase").Interface(
		Id("FindAll").Params(Id("ctx").
			Qual("context", "Context")).
			Params(String(), Error()),
		Id("FindByID").Params(Id("ctx").
			Qual("context", "Context"), Id("id").String()).
			Params(String(), Error()),
		Id("Create").Params(Id("ctx").
			Qual("context", "Context"), Id("id").String()).
			Params(String(), Error()),
		Id("Update").Params(Id("ctx").
			Qual("context", "Context"), Id("id").String()).
			Error(),
		Id("Delete").Params(Id("ctx").
			Qual("context", "Context"), Id("id").String()).
			Error(),
	)
	f.Type().Id("usecase").Struct(
		Id("tx").Qual(txPath, "Handler"),
		Comment("TODO: Inject repository"),
		Comment("repo domain.Repository"),
	)
	f.Func().Id("New").Params(Id("tx").
		Qual(txPath, "Handler")).Id("Usecase").
		Block(
			Return(Op("&").Id("usecase").Values(Dict{
				Id("tx"): Id("tx"),
			})),
		)

	methods := []struct {
		name       string
		parameters []Code
		results    []Code
		body       []Code
	}{
		{
			"FindAll",
			[]Code{Id("ctx").Qual("context", "Context")},
			[]Code{String(), Error()},
			[]Code{Return(Lit("FindAll"), Nil())},
		},
		{
			"FindByID",
			[]Code{Id("ctx").Qual("context", "Context"), Id("id").String()},
			[]Code{String(), Error()},
			[]Code{Return(Lit("FindByID"), Nil())},
		},
		{
			"Create",
			[]Code{Id("ctx").Qual("context", "Context"), Id("id").String()},
			[]Code{String(), Error()},
			[]Code{
				List(
					Id("_"), Id("err")).Op(":=").Id("u").Dot("tx").
					Dot("Transaction").
					Call(
						Id("ctx"),
						Func().Params(Id("ctx").
							Qual("context", "Context")).Params(Interface(), Error()).
							Block(Return(Nil(), Nil())),
					),
				Return(Lit("Create"), Id("err")),
			},
		},
		{
			"Update",
			[]Code{Id("ctx").
				Qual("context", "Context"), Id("id").String()},
			[]Code{Error()},
			[]Code{
				List(Id("_"), Id("err")).Op(":=").Id("u").Dot("tx").
					Dot("Transaction").
					Call(
						Id("ctx"),
						Func().Params(Id("ctx").Qual("context", "Context")).Params(Interface(), Error()).
							Block(Return(Nil(), Nil())),
					),
				Return(Id("err")),
			},
		},
		{
			"Delete",
			[]Code{Id("ctx").Qual("context", "Context"), Id("id").String()},
			[]Code{Error()},
			[]Code{
				List(Id("_"), Id("err")).Op(":=").Id("u").Dot("tx").
					Dot("Transaction").
					Call(
						Id("ctx"),
						Func().Params(Id("ctx").
							Qual("context", "Context")).Params(Interface(), Error()).
							Block(Return(Nil(), Nil())),
					),
				Return(Id("err")),
			},
		},
	}

	for _, method := range methods {
		c := append([]Code{Comment("TODO: write your code")}, method.body...)
		f.Func().Params(Id("u").Op("*").Id("usecase")).Id(method.name).Params(method.parameters...).Params(method.results...).Block(
			c...,
		).Line()
	}

	if err := output(f, fileName); err != nil {
		return "", err
	}
	return fileName, nil
}

func (g *generator) generateDomain() {
	packageName := getPackageName(g.tagName)
	fileNamePrefix := getFileNamePrefix(g.tagName)
	fileName := fmt.Sprintf(g.domainFileFormat, g.tagName, fileNamePrefix)

	f := NewFile(packageName)

	f.Var().Id("ErrNotFound").Op("=").
		Qual("errors", "New").Call(Lit("not found")).Line()
	f.Comment(fmt.Sprintf("%s is a domain model.", g.tagNameCamelCase)).Line().
		Type().Id(g.tagNameCamelCase).Struct(Comment("TODO: write your code"))
	f.Comment(fmt.Sprintf("New is a constructor of %s.", g.tagNameCamelCase)).Line().
		Comment("TODO: write your code").Line().
		Func().Id("New").Params().Parens(List(Op("*").Id(g.tagNameCamelCase), Error())).
		Block(
			Comment("TODO: write validation"),
			Return(Op("&").Id(g.tagNameCamelCase).Values(), Nil()),
		)

	if err := output(f, fileName); err != nil {
		panic(err)
	}
}

func (g *generator) generateDomainRepository() {
	packageName := getPackageName(g.tagName)
	fileNamePrefix := getFileNamePrefix(g.tagName)
	fileName := fmt.Sprintf(g.domainRepoFileFormat, g.tagName, fileNamePrefix)
	domainName := convTagName2CamelCase(g.tagName)

	f := NewFile(packageName)
	f.HeaderComment(domainRepositoryComment)

	f.ImportName("context", "context")

	f.Comment(fmt.Sprintf("Repository is a repository interface of %s.", domainName)).Line().
		Type().Id("Repository").Interface(
		Id("FindAll").Params(Id("ctx").Qual("context", "Context")).Parens(List(Index().Op("*").Id(domainName), Error())),
		Id("FindByID").Params(Id("ctx").Qual("context", "Context"), Id("id").String()).Parens(List(Op("*").Id(domainName), Error())),
		Id("Create").Params(Id("ctx").Qual("context", "Context"), Id("d").Op("*").Id(domainName)).Error(),
		Id("Update").Params(Id("ctx").Qual("context", "Context"), Id("d").Op("*").Id(domainName)).Error(),
		Id("Delete").Params(Id("ctx").Qual("context", "Context"), Id("id").String()).Error(),
	)

	if err := output(f, fileName); err != nil {
		panic(err)
	}
}

func (g *generator) generateRepository() {
	packageName := getPackageName(g.tagName)
	fileNamePrefix := getFileNamePrefix(g.tagName)
	fileName := fmt.Sprintf(g.infraRepoFileFormat, g.tagName, fileNamePrefix)
	domainName := convTagName2CamelCase(g.tagName)
	domainAliasName := strings.ToLower(domainName[:1]) + domainName[1:] + "Domain"
	domainPath := g.modulePath + "/domain/" + g.tagName
	entPath := g.modulePath + "/infrastructure/db/ent"

	f := NewFile(packageName)
	f.ImportName("context", "context")
	f.ImportAlias(domainPath, domainAliasName)
	f.ImportName(entPath, "ent")

	f.Comment(fmt.Sprintf("Repository is a repository of %s", domainName)).Line().
		Type().Id("Repository").
		Struct(
			Id("dbHandler").Qual(entPath, "Handler"),
		)
	f.Comment("New is a constructor of Repository").Line().
		Func().Id("New").Params(Id("dbHandler").
		Qual(entPath, "Handler")).
		Qual(domainPath, "Repository").Block(
		Return(&Statement{
			Op("&").Id("Repository").Values(Dict{
				Id("dbHandler"): Id("dbHandler"),
			}),
		}),
	)

	varName := strings.ToLower(domainName[:1]) + domainName[1:]
	typeName := domainAliasName + "." + domainName

	f.Comment(fmt.Sprintf("FindAll is a method to find all %s", typeName)).Line().
		Func().Params(Id("r").Op("*").Id("Repository")).Id("FindAll").
		Params(Id("ctx").Qual("context", "Context")).
		Params(Index().Op("*").Id(typeName), Error()).
		Block(
			Id("client").Op(":=").Id("r.dbHandler.GetClient()").Line(),
			List(Id(varName+"s"), Id("err")).Op(":=").Id(fmt.Sprintf("client.%s.Query().All(ctx)", domainName)),
			If(Id("err").Op("!=").Nil()).Block(
				Return(Nil(), Id("err")),
			),
			Id(varName+"Domains").Op(":=").Make(Index().Op("*").Id(typeName), Len(Id(varName+"s"))),
			For(Id("i, o").Op(":=").Range().Id(varName+"s")).Block(
				Id(varName+"Domains[i]").Op("=").Id(domainAliasName+".NewFromRepository(o.ID)"),
			),
			Line(),
			Return(Id(varName+"Domains"), Nil()),
		)

	f.Comment(fmt.Sprintf("FindByID is a method to find %s by id", typeName)).Line().
		Func().Params(Id("r").Op("*").Id("Repository")).Id("FindByID").
		Params(
			Id("ctx").Qual("context", "Context"),
			Id("id").Id("string"),
		).
		Params(Op("*").Id(typeName), Error()).
		Block(
			Id("client").Op(":=").Id("r.dbHandler.GetClient()").Line(),
			List(Id(varName), Id("err")).Op(":=").Id(fmt.Sprintf("client.%s.Query().Get(ctx, id)", domainName)),
			If(Id("err").Op("!=").Nil()).Block(
				Return(Nil(), Id("err")),
			),
			Line(),
			Return(Id(fmt.Sprintf("%s.NewFromRepository(%s.ID)", domainAliasName, varName)), Nil()),
		)

	f.Comment(fmt.Sprintf("Create is a method to create %s", typeName)).Line().
		Func().Params(Id("r").Op("*").Id("Repository")).Id("Create").
		Params(
			Id("ctx").Qual("context", "Context"),
			Id(varName).Op("*").Id(typeName),
		).
		Params(Error()).
		Block(
			Id("client").Op(":=").Id("r.dbHandler.GetClient()").Line(),
			List(Id(varName), Id("err")).Op(":=").
				Id(fmt.Sprintf("client.%s.Create()", domainName)).Op(".").Line().
				Id(fmt.Sprintf("SetID(%s.ID())", varName)).Op(".").Line().
				Id("Save(ctx)").Line().
				If(Id("err").Op("!=").Nil()).Block(
				Return(Nil(), Id("err")),
			),
			Line(),
			Return(Id(fmt.Sprintf("%s.NewFromRepository(%s.ID)", domainAliasName, varName)), Nil()),
		)

	if err := output(f, fileName); err != nil {
		panic(err)
	}
}
