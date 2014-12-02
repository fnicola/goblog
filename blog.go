package main

import (
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type Entry struct {
	Title string
	Body  string
	Time  string
	Tag   string
}

func getFileListFromPath(path string) ([]os.FileInfo, error) {
	dirFiles, err := ioutil.ReadDir(path)

	if err != nil {
		return nil, err
	}
	return dirFiles, nil
}

func getagFromPath(path string) (string, error) {
	titleTagged := strings.Split(path, ".")[0]
	tags := strings.Split(titleTagged, "_")[1]

	return tags, nil
}

func getEntries(dirFiles []os.FileInfo, path string, tag string) ([]Entry, error) {

	var entries []Entry

	for a := range dirFiles {
		if filepath.Ext(dirFiles[a].Name()) == ".md" {
			//read file content
			byArr, err := ioutil.ReadFile(path + dirFiles[a].Name())

			if err != nil {
				return nil, err
			}

			titleTagged := strings.Split(dirFiles[a].Name(), ".")[0]
			title := strings.Split(titleTagged, "_")[0]
			tags := strings.Split(titleTagged, "_")[1]
			const layout = "2006-Jan-02 at 3:04pm"
			time_file := dirFiles[a].ModTime().Format(layout)
			tag_entry, _ := getagFromPath(dirFiles[a].Name())

			if tag_entry == tag || tag == "" {
				entry := Entry{Title: title, Body: string(byArr), Time: time_file,
					Tag: tags}
				entries = append(entries, entry)
			}
		}
	}
	return entries, nil
}

func renderTemplateIndex(w http.ResponseWriter, tmpl string, i []Entry) {

	t, _ := template.ParseFiles(tmpl + ".html")
	t.Execute(w, i)
}

func renderTemplateError(w http.ResponseWriter, tmpl string) {

	t, _ := template.ParseFiles(tmpl + ".html")
	t.Execute(w, "error")
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	path := "content/"
	tag := r.URL.Path[len("/view/"):]
	dirFiles, err := getFileListFromPath(path)

	if err != nil {
		renderTemplateError(w, "view")
	}
	entries, err := getEntries(dirFiles, path, tag)

	if err != nil {
		renderTemplateError(w, "view")
	}
	renderTemplateIndex(w, "view", entries)
}

func main() {
	http.HandleFunc("/view/", viewHandler)
	http.ListenAndServe(":8080", nil)
}
