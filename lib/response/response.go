package response

import (
	"fmt"
	"net/http"
	"text/template"
	"viry_sun/lib/request"
)

func RenderTemplate(w *http.ResponseWriter, r *http.Request, tplFile string, data any) {
	tplDir := request.GetTemplateDir(r)
	tmpl, err := template.ParseFiles(tplDir+tplFile, tplDir+"/common.html")
	if err != nil {
		panic(fmt.Sprintf("Template Parse Fatal: %v", err))
	}

	tmpl.Execute(*w, data)
}

func WriteJson(w http.ResponseWriter, bt []byte) {
	w.Header().Set("Content-Type", "application-json")
	w.Write(bt)
}
