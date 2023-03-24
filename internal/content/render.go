package content

import (
	"log"
	"os"
	"text/template"
)

func Args(args ...interface{}) []interface{} {
	return args
}

func IntToStringRepresentation(i int) string {
	var s string
	switch {
	case i == 0:
		s = "zero"
	case i == 1:
		s = "one"
	case i == 2:
		s = "two"
	case i == 3:
		s = "three"
	case i == 4:
		s = "four"
	case i == 5:
		s = "five"
	case i == 6:
		s = "six"
	case i == 7:
		s = "seven"
	case i == 8:
		s = "eight"
	case i == 9:
		s = "nine"
	}

	return s
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
