package render

import (
	"bytes"
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/ryutaro-asada/go-practice/pkg/config"
	"github.com/ryutaro-asada/go-practice/pkg/models"
)

var app *config.AppConfig

func NewTemplate(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {

	var tc map[string]*template.Template
	if app.UseCache {
		log.Println("000000000000000000000")
		// get the template chache from the app config
		tc = app.TemplateCache
	} else {
		log.Println(")))))))))))))))))))))")
		tc, _ = CreateTemplateCache()
	}

	// get requested template from chache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("can not get template from chache")
	}

	buf := new(bytes.Buffer)

	td = AddDefaultData(td)

	err := t.Execute(buf, td)
	if err != nil {
		log.Println(err)
	}

	// render template
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}

}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	// get all of the files named *.page.tmpl from ./templates
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}
	log.Println("---------pages:  ", pages)

	// reange through all files ending with  *.pate.tmpl
	for _, page := range pages {
		log.Println("page:  ", page)

		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		log.Println("ts after new template", ts)

		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}
		log.Println("matches after filepath.Glob()", matches)

		if len(matches) > 0 {
			log.Println("in len(match)")

			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			log.Println("ts in len(match)", ts)
			if err != nil {
				return myCache, err
			}
		}
		log.Println("ts after matches", ts)

		myCache[name] = ts
		log.Println("myCache every iterate", myCache)
		log.Println("-------------- end of loop -------------------------------")

	}
	return myCache, nil
}
