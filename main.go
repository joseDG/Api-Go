package main

import (
    "encoding/json"
    "net/http"
    "log"
    "github.com/gorilla/mux"
)

type info struct {
    Universidad string `json:"Universidad"` 
	Curso    string `json:"Curso"`
	Alumno    string `json:"Alumno"`
	Periodo    string `json:"Periodo"`
	LenguajeProgramacionPreferido    string `json:"LenguajeProgramacionPreferido"`
	AspiraciónPostGraduación    string `json:"AspiraciónPostGraduación"`
}

type allInfo []info

var informacion = allInfo{
	{
		Universidad: "UTPL",
		Curso:    "Procesos de Ingeniería de Software",
		Alumno:    "José Luis Díaz G",
		Periodo: "Abr/Ago 2021",
		LenguajeProgramacionPreferido: "Python",
		AspiraciónPostGraduación: "Especializarme en Machine Learning",
	},
}

func getInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(informacion)
}


func main() {
    router := mux.NewRouter().StrictSlash(true)

    router.HandleFunc("/", getInfo).Methods("GET")

   
    log.Fatal(http.ListenAndServe(":3000", router))
}