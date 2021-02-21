package servidor
//package main

import (
	"../lista"
	"../ll"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

type Tienda struct {
	Nombre       string `json:Nombre`
	Descripcion  string `json:Descripcion`
	Contacto     string `json:Contacto`
	Calificacion int `json:Calificacion`
}
type Departamento struct {
	Nombre  string `json:Nombre`
	Tiendas []Tienda
}
type Dato struct {
	Indice string `json:Indice`
	Departamentos []Departamento `json:Departamentos`
}
type Sobre struct {
	Datos []Dato
}


func Request() {
	ruta := mux.NewRouter().StrictSlash(true) //es para enlazar diferentes peticones, get, post, put, delete
	ruta.HandleFunc("/", Home)
	ruta.HandleFunc("/GetArreglo", GetArreglo).Methods("GET")  //se quedara esperando una respuesta, si no la tiene dara error
	ruta.HandleFunc("/CargarTienda", MetodoPost).Methods("POST") //no se queda esperando una respuesta, solo mandamos informacion
	//ruta.HandleFunc("/Eliminar",metodoEliminar).Methods("DELETE")
	ruta.HandleFunc("/TiendaEspecifica", MetodoBusqueda).Methods("POST")
	log.Fatal(http.ListenAndServe(":3000", ruta))
}
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Servidor en Go")
}

var re Sobre

func MetodoPost(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err !=nil{
		fmt.Fprintf(w,"Error al cargar el archivo")
	}else{
		json.Unmarshal(body, &re)
	}
	json.NewEncoder(w).Encode(re) //mostrá los datos en postman

	Recorrido()
}

func GetArreglo(w http.ResponseWriter, r *http.Request) {
	/*
		Grafico completo del arreglo
		Este reporte permitirá visualizar el arreglo completo con sus listas que se generará
		al momento de aplicar la técnica correcta de linealización, permitiendo de tal
		manera visualizar el gráfico de la lista doblemente enlazada, y así comprobar que la
		estructura creada, corresponde correctamente a lo que se requiere.
		Para lograr esto se debe generar un gráfico del arreglo con sus listas y almacenarse
		en la misma carpeta del proyecto.
		El endpoint que permitirá ejecutar dicha acción será de tipo GET, y la dirección o
		URL exacta es: localhost:puerto/getArreglo
	*/
	//fmt.Fprint(w, "[1,2,3,4,5]") // lo q tenemos que mandar esta en formato json
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(re)


}
var ListaY = lista.NuevaLista() //lista doblemente enlazada
var ListaX=lista.NuevaLista() //lista doblemente enlazada
var ListaDatos=lista.NuevaLista() //lista doblemente enlazada de las tiendas
var listaDeListas=ll.NuevaLista() //lista de listas para los datos que estaran guardados de las tiendas
func Recorrido()  {
	/*
		crear primero una matriz para el indice (filas) y departamentos(columnas)
		luego llenar la lista dobremente enlazada con la informacion
		correspondiente de las tiendas por medio de la calificacion
		la lista de las tiendas va ordenada
		cada 20 posiciones en la matriz (indice * departamentos) creamos una imagen
	*/
	//var datosTienda Tienda
	var columnas Departamento
	var filas Dato
	var indiceY, indiceX int
	contadorx :=0
	contadory:=0
	contadorx2:=0
	var arreglo [10]int
	for indiceY,filas=range re.Datos{
		fmt.Println("indice y",indiceY)
		if filas.Indice!=""{
			ListaY.Insertar(re.Datos[indiceY].Indice)
			contadory++
			for indiceX, columnas = range filas.Departamentos{
				fmt.Println("indicex",indiceX)
				if columnas.Nombre!=""{
					ListaX.Insertar(columnas.Nombre)
					contadorx+=5
					arreglo[indiceY]=indiceX
				}else{
					fmt.Println("Error, no existe un departamento")
					contadorx2++
				}
			}
		}else{
			fmt.Println("Error, no existe un indice")
		}
	}
	fmt.Println("arreglo",arreglo)

	//for i:=0;i<=len(arreglo);i++{
	//	if arreglo[i]<=arreglo[i]{
	//		fmt.Println("arreglo", arreglo[i])
	//	}
	//}

	for i:=0;i<=indiceY;i++{
		for j:=0;j<=(contadorx);j++{
			linealizacion:=i*(contadorx)+(j)
			listaDeListas.InsertarListaListas(linealizacion,"dsdf")
			fmt.Println("filas",i)
			fmt.Println("columnas",j)
			fmt.Println("Linealizacion",linealizacion)

		}

	}
	fmt.Println("Contadorx: ",contadorx)
	fmt.Println("Lista de listas")
	listaDeListas.Print()
	fmt.Println("contadorx2")

}
var re2 Dato
func MetodoBusqueda(w http.ResponseWriter,r *http.Request){
	/*
	Para este reporte el usuario realizará una solicitud especificando los tres
	parámetros necesarios para mostrar la información de una tienda que se encuentre
	dentro de la estructura del arreglo, de tal manera que se enviaran tres parámetros
	en formato json, estos parámetros son el departamento, el nombre de la tienda
	y la calificación de la tienda, con estos datos la aplicación devolverá la
	información de la tienda específica que cumpla con todos los parámetros enviados.

	En el momento que la aplicación reciba este json debe retornar toda la información
	de la tienda.
	El endpoint que permitirá ejecutar dicha acción será de tipo POST, el contenido es
	como el ejemplo mostrado con anterioridad, y la dirección o URL exacta es :
	localhost:puerto/TiendaEspecifica.
	*/
	body, err := ioutil.ReadAll(r.Body)
	if err !=nil{
		fmt.Fprintf(w,"Error al cargar el archivo")
	}else{
		json.Unmarshal(body, &re2)
	}
	//mostrá los datos en postman
	//bandera:=false
	//for i,indices:=range re2.Datos{
	//	for d,departamentos:=range indices.Departamentos{
	//		bandera=true
	//		if bandera==true{
	//			if re.Datos[i].Departamentos[d].Nombre==departamentos.Nombre{
	//				//ListaX.Buscar(indices.Departamentos[s].Nombre)
	//				fmt.Println("dep",departamentos.Nombre)
	//			}
	//
	//		}
	//
	//	}
	//}

	for _,departamentos:=range re2.Departamentos{
		if re.Datos[0].Departamentos[0].Nombre==departamentos.Nombre{
			fmt.Println("dep",departamentos.Nombre)
		}
	}
}


func MetodoEliminar()  {
	/*
		Esta función nos permite eliminar una tienda que ya no necesitamos mandando a
		una url el Nombre, Categoría y Calificación. Con los datos mandados se
		procederá a realizar el cálculo en la posición que debería de ir la tienda y eliminar
		en esa posición.
		Ejemplo:
		localhost:puerto/Eliminar.

	*/
}


func MetodoBusquedaLinealizado()  {
	/*
	Para este reporte el usuario podrá realizar una petición de tipo GET hacia nuestro
	servidor donde se le enviará como parámetros en la URL la posición dentro de la
	lista, dando como retorno a la petición toda la información de las tiendas
	perteneciente a esta posición, en dado caso no exista ninguna tiene en esta posición
	deberá retorna un mensaje donde explique que no hay registro de tiendas en dicho
	índice.
	Ejemplo:
	◆ localhost:puerto:/id/:numero
	Nota: número es el índice que se desea retornar.

	*/
}

func GuardarDatos()  {
	/*
	Esta función del servidor permite almacenar la información que contiene la matriz
	de tiendas en un nuevo archivo con formato JSON, este archivo debe ser parecido al
	archivo de entrada.
	Nota: Se va a cargar el archivo generado por el estudiante para comprobar que el
	formato es el adecuado.

	*/
}