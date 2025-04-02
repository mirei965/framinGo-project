package render

import (
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/CloudyKit/jet/v6"
	"github.com/alexedwards/scs/v2"
	"github.com/justinas/nosurf"
)

type Render struct {
	Renderer   string
	RootPath   string
	Secure     bool
	Port       string
	ServerName string
	JetViews   *jet.Set
	Session    *scs.SessionManager
}

type TemplateData struct {
	IsAuthenticated bool
	InData          map[string]int
	StringMap       map[string]string
	FloatMap        map[string]float32
	Data            map[string]interface{}
	CSRFToken       string
	Port            string
	ServerName      string
	Secure          bool
	Error           string
	Flash           string

}

func (f *Render) defaultData(td *TemplateData, r *http.Request) *TemplateData {
	td.Secure = f.Secure
	td.ServerName = f.ServerName
	td.CSRFToken = nosurf.Token(r)
	td.Port = f.Port

	if f.Session.Exists(r.Context(), "userID") {
		td.IsAuthenticated = true
	}
	td.Error = f.Session.PopString(r.Context(), "error")
	td.Flash = f.Session.PopString(r.Context(), "flash")

	return td
}

func (f *Render) Page(w http.ResponseWriter, r *http.Request, view string, variables, data any) error {
	switch strings.ToLower(f.Renderer) {
	case "go":
		return f.GoPage(w, r, view, data)
	case "jet":
		return f.JetPage(w, r, view, variables, data)
	default:

	}

	return errors.New("No rendering engine specified")
}

// GoPage renders a go template
func (f *Render) GoPage(w http.ResponseWriter, r *http.Request, view string, data interface{}) error {
	tmpl, err := template.ParseFiles(fmt.Sprintf("%s/views/%s.page.tmpl", f.RootPath, view))
	if err != nil {
		return err
	}
	// td = template data
	td := &TemplateData{}
	if data != nil {
		td = data.(*TemplateData)
	}

	err = tmpl.Execute(w, &td)
	if err != nil {
		return err
	}

	return nil
}

// JetPage renders a jet template
func (f *Render) JetPage(w http.ResponseWriter, r *http.Request, templateName string, variables, data any) error {
	var vars jet.VarMap

	if variables == nil {
		vars = make(jet.VarMap)
	} else {
		vars = variables.(jet.VarMap)
	}

	td := &TemplateData{}
	if data != nil {
		td = data.(*TemplateData)
	}

	td = f.defaultData(td, r)

	t, err := f.JetViews.GetTemplate(fmt.Sprintf("%s.jet", templateName))
	if err != nil {
		log.Println(err)
		return err
	}

	err = t.Execute(w, vars, td)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
