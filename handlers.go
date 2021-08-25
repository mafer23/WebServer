package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)



// Handler para manejar la ruta principal
//Parametros: escritor w del tipo http.ResponseWriter y objeto request
func HandleRoot(w http.ResponseWriter, r *http.Request)  {
	//Impresion en el navegador
	//parametros: escritor-objeto encargado de responder al cliente 
	//y mensaje escrito a travez del escritor
	fmt.Fprintf(w, "Hello World from handlers")
}

func HandleHome(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "This is the API Endpoint")
}

func PostRequest(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var metadata MetaData
	err := decoder.Decode(&metadata)

	if err != nil {
		fmt.Fprintf(w, "error %v", err)
		return
	}


   
	fmt.Fprintf(w,"Payload %v\n", metadata)

}

func UserPostRequest(w http.ResponseWriter, r *http.Request)  {
	decoder := json.NewDecoder(r.Body)
	var user User
	err := decoder.Decode(&user)
	if err != nil {
		fmt.Fprintf(w, "error: %v", err)
		return
	}
	response, err := user.ToJson()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
