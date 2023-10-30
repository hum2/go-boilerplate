package golang

import (
	"github.com/dave/jennifer/jen"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"os"
	"path/filepath"
	"strings"
)

func getPackageName(tagName string) string {
	s := strings.Split(tagName, "/")
	return s[len(s)-1]
}

func getFileNamePrefix(tagName string) string {
	return strings.Replace(tagName, "/", "_", -1)
}

func getAppName(tagName string) string {
	return getPackageName(tagName) + "App"
}

func convTagName2CamelCase(tagName string) string {
	parts := strings.Split(tagName, "/")

	for i, part := range parts {
		parts[i] = cases.Title(language.Und, cases.NoLower).String(part)
	}

	return strings.Join(parts, "")
}

func output(f *jen.File, fileName string) error {
	if err := os.MkdirAll(filepath.Dir(fileName), 0755); err != nil {
		return err
	}
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	if err := f.Render(file); err != nil {
		return err
	}
	return nil
}
