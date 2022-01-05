package views

import (
	"html/template"
	"net/http"
	"path/filepath"
)

var (
	LayoutDir string = "views/layouts/"
	TemplateExt string = ".gohtml"
)

func NewView(layout string, files ...string) *View {
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

type View struct {
	Template	*template.Template
	Layout		string

}

//Render renders a View with the predefined layout
func (v *View) Render(w http.ResponseWriter, data interface{}) error {
	return v.Template.ExecuteTemplate(w, v.Layout, data)
}

//layoutFiles returns a slice of strings representing hte layout files used in the applicaton
func layoutFiles() []string {
 files, err := filepath.Glob(LayoutDir + "*" + TemplateExt)
 if err != nil {
	 panic(err)
 }

 return files
}
