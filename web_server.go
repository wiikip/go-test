package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
)

func upload(w http.ResponseWriter, r *http.Request) {

    if r.Method != "POST" {
        fmt.Fprintf(w, "Bad Request: method %s not authorized", r.Method)
        return
    }
    r.ParseMultipartForm(50 << 20)
    file, header, err := r.FormFile("go")
    if err != nil {
        fmt.Fprintf(w, "Failed while trying to upload file : %s", err)
        return
    }
    data, err := ioutil.ReadAll(file)
    if err != nil{
        fmt.Fprintf(w, "Failed while trying to read file : %s", err)
        return
    }
    fileMod := fs.FileMode(0644)
    filepath := path.Join(os.Getenv("FILE_STORAGE"),header.Filename)

    err = ioutil.WriteFile(filepath,data, fileMod)
        if err != nil {
        fmt.Println(err)
        fmt.Fprintf(w, "File not uploaded:  %s",  err)
        return
    }
    fmt.Fprint(w, "Successfuly uploaded file !")
    defer file.Close()
}

func main() {
    http.HandleFunc("/upload", upload)
    log.Fatal(http.ListenAndServe(":8080", nil))
}

