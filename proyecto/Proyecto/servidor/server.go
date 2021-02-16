package servidor
//package main

import (
	"../lista"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"../ll"
)

type Tienda struct {
	Nombre       string `json:Nombre`
	Descripcion  string `json:Descripcion`
	Contacto     string `json:Contacto`
	Calificacion int `json:Calificacion`
}
type Departamento struct {
	Nombre  string `json:Nombre`
	Tiendas []Tienda `json:Tiendas`
}
type Dato struct {
	Indice string `json:Indice`
	Departamentos []Departamento `json:Departamentos`
}
type Sobre struct {
	Datos []Dato
}

//func main(){
//	Request()
//}

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
		Recorrido()
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

	//for _,recorrido:=range re.Datos{

		//json.NewEncoder(w).Encode(recorrido)
		/*
		//fmt.Println("Indice:",recorrido.Indice)
		fmt.Fprintf(w,"Indice: %s\n",recorrido.Indice)
		for _,recorrido2:=range recorrido.Departamentos{
			//fmt.Println("Departamento:",recorrido2.Nombre)
			//fmt.Fprintf(w,"Departamento: %s\n",recorrido2.Nombre)
			json.NewEncoder(w).Encode(recorrido2.Nombre)
			for _,recorrido3:=range recorrido2.Tiendas{
				//fmt.Println("Nombre de la Tienda:",recorrido3.Nombre)
				fmt.Fprintf(w,"Nombre de la Tienda: %s\n",recorrido3.Nombre)
				json.NewEncoder(w).Encode(recorrido3.Nombre)
				//fmt.Println("Descripcion:",recorrido3.Descripcion)
				//fmt.Fprintf(w,"Descripcion: %s\n",recorrido3.Descripcion)
				json.NewEncoder(w).Encode(recorrido3.Descripcion)
				//fmt.Println("Contacto:",recorrido3.Contacto)
				//fmt.Fprintf(w,"Contacto: %s\n",recorrido3.Contacto)
				json.NewEncoder(w).Encode(recorrido3.Contacto)
				//fmt.Println("Calificacion:",recorrido3.Calificacion)
				//fmt.Fprintf(w,"Calificacion: %d\n",recorrido3.Calificacion)
				json.NewEncoder(w).Encode(recorrido3.Calificacion)
				//fmt.Println("Tipo de dato de calificacion:",reflect.TypeOf(recorrido3.Calificacion))
			}
		}
		*/
	//}
	Recorrido()
}

func Recorrido()  {
	lisVertical:=lista.NuevaLista()
	lisHorizontal:=lista.NuevaLista()
	otro:=ll.NuevaLista()
	//ma:=matriz.NuevaMatriz() //tengo que agregar lista 1 y lista 2
	//li:=matriz.NuevaLista() //la lista es para agregar horizontal y vertical

	for _,recorrido:=range re.Datos{
		//fmt.Println("Indice:",recorrido.Indice)
		lisVertical.Insertar(recorrido.Indice)





		/*
		crear primero una matriz para el indice (filas) y departamentos(columnas)
		luego llenar la lista dobremente enlazada con la informacion
		correspondiente de las tiendas por medio de la calificacion
		la lista de las tiendas va ordenada

		cada 20 posiciones en la matriz (indice * departamentos) creamos una imagen
				*/

		for _,recorrido2:=range recorrido.Departamentos{
			//fmt.Println("Departamento:",recorrido2.Nombre)

			for index,recorrido3:=range recorrido2.Tiendas{
				//fmt.Println("Nombre de la Tienda:",recorrido3.Nombre)
				//fmt.Println("Descripcion:",recorrido3.Descripcion)
				//fmt.Println("Contacto:",recorrido3.Contacto)
				//fmt.Println("Calificacion:",recorrido3.Calificacion)
				//
				//fmt.Println("Tipo de dato de calificacion:",reflect.TypeOf(recorrido3.Calificacion))


				fmt.Println("Lista de listas")
				otro.InsertarListaListas(index,recorrido3.Nombre)
				//otro.Print()
			}
			fmt.Println("Horizontal")
			lisHorizontal.Insertar(recorrido2.Nombre)
			lisHorizontal.Imprimir()
		}

	}
	fmt.Println("Vertical")
	//lisVertical.Imprimir()

	//li.Imprimir()
}

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