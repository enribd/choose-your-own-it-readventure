package content

import (
	"log"
	"os"
	"text/template"
)

func Args(args ...interface{}) []interface{} {
	return args
}

/*
* Render templates with a given data and export them to files or stdout
* Params:
*   t: templates loaded
*   data: data used to fill the templates
*   templateName: template name to render
*   dest: destination file
 */
func Render(t *template.Template, templateName, dest string, data interface{}) error {
	var file *os.File
	var err error

	if dest == "stdout" {
		file = os.Stdout
	} else {
		// Create destination file
		file, err = os.Create(dest)
		if err != nil {
			log.Fatalln("create file: ", err)
			return err
		}
	}

	// Render template
	err = t.ExecuteTemplate(file, templateName, data)
	if err != nil {
		log.Fatalln(err)
		return err
	}

	return nil
}
