package yaml

import (
	"fmt"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"net/http"
	"strings"
)

type Generator interface {
	GenerateOAPI(methods []string) (createdFiles []string, err error)
}

type generator struct {
	tagName                string
	tagNameCamelCase       string
	assetsPathPrefix       string
	pathFilenameFormat     string
	responseFilenameFormat string
	requestFilenameFormat  string
	errFilenameFormat      string
}

func NewGenerator(
	tagName string,
	tagNameCamelCase string,
	assetsPathPrefix string,
	pathFilenameFormat string,
	responseFilenameFormat string,
	requestFilenameFormat string,
	errFilenameFormat string,
) Generator {
	return &generator{
		tagName:                tagName,
		tagNameCamelCase:       tagNameCamelCase,
		assetsPathPrefix:       assetsPathPrefix,
		pathFilenameFormat:     pathFilenameFormat,
		responseFilenameFormat: responseFilenameFormat,
		requestFilenameFormat:  requestFilenameFormat,
		errFilenameFormat:      errFilenameFormat,
	}
}

// GenerateOAPI は指定されたコントローラ名に対応したOAPI定義ファイルを作成します
// 作成するファイルは以下の通りです
// - エンドポイント情報を記述したroot.yamlファイルを作成
// - 指定されたHTTPメソッド（`GET`, `PUT`, `POST`, `DELETE`）ごとに`{HTTPメソッド}.yaml`ファイルを生成
// - TODO：パスパラメータが設定されていたら、パスパラメータを別ファイルで定義
// - `PUT`, `POST`, `DELETE`のリクエストファイルを生成
// - `GET`, `PUT`, `POST`の正常系レスポンスファイルを生成
func (g *generator) GenerateOAPI(methods []string) (createdFiles []string, err error) {
	_r := strings.Repeat("../", strings.Count(g.tagName, "/")+2)
	relativePath := _r[:len(_r)-1]
	errFileName := relativePath + g.errFilenameFormat
	root := &Root{
		Methods: make(map[string]*Ref),
	}
	for _, method := range methods {
		lowerMethod := strings.ToLower(method)
		pathFilename := fmt.Sprintf(g.pathFilenameFormat, g.assetsPathPrefix, g.tagName, lowerMethod)
		operationID := cases.Title(language.Und, cases.NoLower).String(lowerMethod)
		filenamePrefix := operationID + g.tagNameCamelCase

		var mEntity *Method
		switch method {
		case http.MethodGet:
			resFilepath := fmt.Sprintf(g.responseFilenameFormat, g.assetsPathPrefix, g.tagName, filenamePrefix)
			generateResponseYAML(resFilepath)
			createdFiles = append(createdFiles, resFilepath)

			resRelativePath := fmt.Sprintf(g.responseFilenameFormat, relativePath, g.tagName, filenamePrefix)
			mEntity = createMethodEntity(g.tagName, operationID, resRelativePath, "", errFileName)
		case http.MethodPost, http.MethodPut:
			reqFilepath := fmt.Sprintf(g.requestFilenameFormat, g.assetsPathPrefix, g.tagName, filenamePrefix)
			generateRequestYAML(reqFilepath)
			createdFiles = append(createdFiles, reqFilepath)

			resFilepath := fmt.Sprintf(g.responseFilenameFormat, g.assetsPathPrefix, g.tagName, filenamePrefix)
			generateResponseYAML(resFilepath)
			createdFiles = append(createdFiles, resFilepath)

			reqRelativePath := fmt.Sprintf(g.requestFilenameFormat, relativePath, g.tagName, filenamePrefix)
			resRelativePath := fmt.Sprintf(g.responseFilenameFormat, relativePath, g.tagName, filenamePrefix)
			mEntity = createMethodEntity(g.tagName, operationID, resRelativePath, reqRelativePath, errFileName)
		case http.MethodDelete:
			reqFilepath := fmt.Sprintf(g.requestFilenameFormat, g.assetsPathPrefix, g.tagName, filenamePrefix)
			generateRequestYAML(reqFilepath)
			createdFiles = append(createdFiles, reqFilepath)

			reqRelativePath := fmt.Sprintf(g.requestFilenameFormat, relativePath, g.tagName, filenamePrefix)
			mEntity = createMethodEntity(g.tagName, operationID, "", reqRelativePath, errFileName)
		}
		if err := writeYAML(pathFilename, mEntity); err != nil {
			return nil, fmt.Errorf("failed to write file %s: %w", pathFilename, err)
		}
		createdFiles = append(createdFiles, pathFilename)

		root.Methods[lowerMethod] = &Ref{Ref: fmt.Sprintf("./%s.yaml", lowerMethod)}
	}
	rootFileName := fmt.Sprintf(g.pathFilenameFormat, g.assetsPathPrefix, g.tagName, "root")
	if err := writeYAML(rootFileName, root); err != nil {
		return nil, fmt.Errorf("failed to write file %s: %w", rootFileName, err)
	}
	createdFiles = append(createdFiles, rootFileName)
	return
}
