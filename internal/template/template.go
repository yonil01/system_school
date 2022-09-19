package template

import (
	"bytes"
	"fmt"
	"foro-hotel/internal/logger"
	"html/template"
)

func GenerateTemplateMail(param map[string]string, path string, cas int64) (string, error) {

	bf := &bytes.Buffer{}
	tpl := &template.Template{}

	if cas == 1 {
		tpl = template.Must(template.New("").ParseGlob("templates/*.gohtml"))
		err := tpl.ExecuteTemplate(bf, path, &param)
		if err != nil {
			fmt.Println(err)
			logger.Error.Printf("couldn't generate template body email: %v", err)
			return "", err
		}
		return bf.String(), err
	}
	if cas == 2 {
		tpl = template.Must(template.New("").ParseGlob("template-reservacion/*.gohtml"))
		err := tpl.ExecuteTemplate(bf, path, &param)
		if err != nil {
			fmt.Println(err)
			logger.Error.Printf("couldn't generate template body email: %v", err)
			return "", err
		}
		return bf.String(), err
	}
	return "", nil

}
