package content

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"text/template"
)

func Args(args ...any) []any {
	return args
}

func IntToIcons(order int) any {
	icons := ""
	for _, c := range strings.Split(strconv.Itoa(order), "") {
		switch c {
		case "0":
			icons = fmt.Sprintf("%s%s", icons, ":material-numeric-0:{.order-icon}")
		case "1":
			icons = fmt.Sprintf("%s%s", icons, ":material-numeric-1:{.order-icon}")
		case "2":
			icons = fmt.Sprintf("%s%s", icons, ":material-numeric-2:{.order-icon}")
		case "3":
			icons = fmt.Sprintf("%s%s", icons, ":material-numeric-3:{.order-icon}")
		case "4":
			icons = fmt.Sprintf("%s%s", icons, ":material-numeric-4:{.order-icon}")
		case "5":
			icons = fmt.Sprintf("%s%s", icons, ":material-numeric-5:{.order-icon}")
		case "6":
			icons = fmt.Sprintf("%s%s", icons, ":material-numeric-6:{.order-icon}")
		case "7":
			icons = fmt.Sprintf("%s%s", icons, ":material-numeric-7:{.order-icon}")
		case "8":
			icons = fmt.Sprintf("%s%s", icons, ":material-numeric-8:{.order-icon}")
		case "9":
			icons = fmt.Sprintf("%s%s", icons, ":material-numeric-9:{.order-icon}")
		default:
			return ""
		}
	}

	return icons
}

/*
* Render templates with a given data and export them to files or stdout
* Params:
*   t: templates loaded
*   data: data used to fill the templates
*   templateName: template name to render
*   dest: destination file
 */
func Render(t *template.Template, templateName, dest string, data any) error {
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
