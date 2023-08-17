package render

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func RenderTemplateTest(w http.ResponseWriter, tmpl string) {
	parseTemp, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")
	err := parseTemp.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template:", err)
		return
	}
}

var tc = make(map[string]*template.Template)

func RenderTemplate(w http.ResponseWriter, t string) {
	var tmpl *template.Template
	var err error

	//Check to see if we already have the TMPL
	_, inMap := tc[t]

	if !inMap {
		log.Println("Creating template and adding to cache")
		err = createTemplateCache(t)
		if err != nil {
			log.Println(err)
		}
	} else {
		log.Println("using cache template")
	}

	tmpl = tc[t]
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}
}

func createTemplateCache(t string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", t),
		"./templates/base.layout.tmpl",
	}

	//parse the template
	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}

	//add template to cache map
	tc[t] = tmpl

	return nil
}
