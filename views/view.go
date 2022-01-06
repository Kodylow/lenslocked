package views

import (
	"html/template"
	"net/http"
	"path/filepath"
)

const (
	layoutDir 		string = "views/layouts/"
	templateDir		string = "views/"
	templateExt 	string = ".gohtml"
)

//NewView creates a view instance with proper templating
func NewView(layout string, files ...string) *View {

	addTemplatePath(files)
	addTemplateExt(files)

	files = append(files, layoutFiles()...)

	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}

	return &View{
		Template: t,
		Layout: layout,
	}
}

//View has a template and layout for serving up http
type View struct {
	Template	*template.Template
	Layout		string

}

func (v *View) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := v.Render(w, nil); err != nil {
		panic(err)
	}
}

//Render renders a View with the predefined layout
func (v *View) Render(w http.ResponseWriter, data interface{}) error {
	w.Header().Set("Content-Type", "text/html")
	return v.Template.ExecuteTemplate(w, v.Layout, data)
}

//layoutFiles returns a slice of strings representing hte layout files used in the applicaton
func layoutFiles() []string {
 files, err := filepath.Glob(layoutDir + "*" + templateExt)
 if err != nil {
	 panic(err)
 }

 return files
}

//addTemplatePath prepends the TemplateDir to each string in a slice
func addTemplatePath(files []string) {
	for i, f := range files {
		files[i] = templateDir + f
	}
}

//addTemplateExt appends the TemplateExt to each string in a slice
func addTemplateExt(files []string) {
	for i, f := range files {
		files[i] = f + templateExt
	}
}