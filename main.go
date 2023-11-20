package main

import (

	// "html/template"
	"bytes"
	"io"
	"net/http"
	"strconv"

	"github.com/iqra-shams/chi/cmd"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	// start := time.Now()
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// Index handler
	r.Post("/", HandlerPostReq)


	http.ListenAndServe(":3333", r)

}


func HandlerPostReq(w http.ResponseWriter, r *http.Request) {

    err:=r.ParseMultipartForm(1000<<20) 
    if err!=nil {
        http.Error(w,"fail to pasrse form",http.StatusInternalServerError)
    }
    Sroutines:= r.FormValue("routines")


file , _ ,err := r.FormFile("file")
if err!=nil{
    http.Error(w,"fail to get  file",http.StatusInternalServerError)
    return
}

defer file.Close()


var BufFile bytes.Buffer

  _ , err= io.Copy(&BufFile, file)
if err!=nil{
    http.Error(w,"fail to read file",http.StatusInternalServerError)
    return
}
   


    routines,_:=strconv.Atoi(Sroutines)

    cmd.ProcessFile(BufFile.String(),routines)

   
}




