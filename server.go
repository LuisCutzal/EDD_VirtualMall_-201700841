package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)
type Dato struct {
	Nombre string
}
func main() {
	request()
}

func request(){
	ruta:=mux.NewRouter().StrictSlash(true)//es para enlazar diferentes peticones, get, post, put, delete
	ruta.HandleFunc("/", home)
	ruta.HandleFunc("/getArreglo",getArreglo).Methods("GET")//se quedara esperando una respuesta, si no la tiene dara error
	ruta.HandleFunc("/metodopost", metodopost).Methods("POST")//no se queda esperando una respuesta, solo mandamos informacion

	log.Fatal(http.ListenAndServe(":3000",ruta))

}
func home(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprint(w,"Servidor en go")
}

func getArreglo(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprint(w,"[1,2,3,4,5]")// lo q tenemos que mandar esta en formato json

}
func metodopost(w http.ResponseWriter, r *http.Request){
	body, _ := ioutil.ReadAll(r.Body)
	var re Dato
	json.Unmarshal(body, &re)
	fmt.Println(re)
}

