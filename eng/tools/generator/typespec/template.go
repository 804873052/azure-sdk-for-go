package typespec

import (
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"time"
)

func ParseTypeSpecTemplates(templateDir, outputDir string, data map[string]any, funcMap template.FuncMap) error {
	if data["releaseDate"] == "" {
		data["releaseDate"] = time.Now().Format("2006-01-02")
	}

	tpl := template.New("parse.tpl").Funcs(funcMap)
	tpl, err := tpl.ParseGlob(filepath.Join(templateDir, "*.tpl"))
	if err != nil {
		return err
	}
	for _, t := range tpl.Templates() {
		fName, _ := strings.CutSuffix(t.Name(), ".tpl")
		w, err := os.OpenFile(filepath.Join(outputDir, fName), os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
		if err != nil {
			return err
		}
		defer w.Close()

		err = tpl.ExecuteTemplate(w, t.Name(), data)
		if err != nil {
			return err
		}
	}

	return nil
}
