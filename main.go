package main

import (
    "fmt"
    "log"
    "net/http"
    "html/template"
    "io/ioutil"
    "regexp"
)

type PageInfo struct { 
        Title string
        Filenames[] string   
}   

func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}

func vidya(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "video games are SHIT")
	fmt.Println("vidya")
}

func images(w http.ResponseWriter, r *http.Request){
    tmpl, err := template.ParseFiles("index.html")
    files, err := ioutil.ReadDir(".")
    if err != nil {
        log.Fatal(err)
    }
    data := PageInfo{Title: "index"}
    for _, file := range files {
        matched, _ := regexp.Match(".(jpeg|jpg|gif|png)", []byte(file.Name()))
        if(matched == true) {
            data.Filenames = append(data.Filenames, file.Name())
        }
    }
    tmpl.Execute(w, data)
}

func folder(w http.ResponseWriter, r *http.Request){
    tmpl, err := template.ParseFiles("folder.html")
    files, err := ioutil.ReadDir(".")
    if err != nil {
        log.Fatal(err)
    }
    data := PageInfo{Title: "Folder"}
    for _, file := range files {
        matched, _ := regexp.Match(".(jpeg|jpg|gif|png)", []byte(file.Name()))
        if(matched == true) {
            data.Filenames = append(data.Filenames, file.Name())
        }
    }
    tmpl.Execute(w, data)
}

func handleRequests() {
	http.HandleFunc("/v", vidya)
    http.HandleFunc("/images", images)
    http.HandleFunc("/folder", folder)
    fs := http.FileServer(http.Dir("."))
    http.Handle("/", fs)
    log.Fatal(http.ListenAndServe(":80", nil))
}

func main() {
    handleRequests()
}