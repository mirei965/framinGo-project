package framingo

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"maps"
	"net/http"
	"path"
	"path/filepath"
)

func (f *Framingo) ReadJSON(w http.ResponseWriter, r *http.Request, data any) error {
	maxBytes := 1024 * 1024 // 1MB
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	dec := json.NewDecoder(r.Body)
	err := dec.Decode(data)
	if err != nil {
		return err
	}
	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		return errors.New("body must only contain a single JSON object")
	}
	return nil
}

func (f *Framingo) WriteJSON(w http.ResponseWriter, status int, data any, headers ...http.Header) error {
	out, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	if len(headers) > 0 {
		maps.Copy(w.Header(), headers[0])
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	_, err = w.Write(out)
	if err != nil {
		return err
	}
	return nil
}

// WriteXML writes the data to the response writer as XML

func (f *Framingo) WriteXML(w http.ResponseWriter, status int, data any, headers ...http.Header) error {
	out, err := xml.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	if len(headers) > 0 {
		maps.Copy(w.Header(), headers[0])
	}

	w.Header().Set("Content-Type", "application/xml")
	w.WriteHeader(status)
	_, err = w.Write(out)
	if err != nil {
		return err
	}
	return nil
}

func (f *Framingo) DownloadFile(w http.ResponseWriter, r *http.Request, pathToFile, fileName string) error {
	fp := path.Join(pathToFile, fileName)
	fileToServe := filepath.Clean(fp)
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", fileToServe))
	http.ServeFile(w, r, fileToServe)

	return nil
}

func (f *Framingo) Error404(w http.ResponseWriter, r *http.Request) {
	f.ErrorStatus(w, http.StatusNotFound)
}

func (f *Framingo) Error500(w http.ResponseWriter, r *http.Request) {
	f.ErrorStatus(w, http.StatusInternalServerError)
}
func (f *Framingo) ErrorUnauthorized(w http.ResponseWriter, r *http.Request) {
	f.ErrorStatus(w, http.StatusUnauthorized)
}
func (f *Framingo) ErrorForbidden(w http.ResponseWriter, r *http.Request) {
	f.ErrorStatus(w, http.StatusForbidden)
}

func (f *Framingo) ErrorStatus(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)

}
