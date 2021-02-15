package matriz

import ("fmt")

type Product struct {
	Nombre string//Indice
	Codigo int // Departamentos puede que cambie a int
	Descripcion string //datos
}

type nodo struct {
	//estos atributos son especificos para la matriz
	x,y int //Saber en que cabecera estoy
	matriz *Product //tipo de objeto
	izquierdo, derecho, arriba, abajo *nodo //nodos para desplazarnos dentro de la matriz
	//estos atributos son especificos para la lista
	header int //tipo interno de cabecera
	siguiente,anterior *nodo //nodo con los que nos vamos a desplazar dentro de las listas
}

type lista struct {
	first, last *nodo
}
type Matriz struct {
	list_H,list_V *lista
}

//creamos un constructor para una nueva matriz
func NodoMatriz(x int ,y int, matriz *Product)*nodo  {
	return &nodo{x,y,matriz,nil,nil,nil,nil,nil,nil,nil}

}
//constructor para la lista de listas
func NodoLista(header int) *nodo  {
	return &nodo{nil,nil,nil,nil,nil,nil,nil,header,nil,nil}
}

//para setear

func (n *nodo) HeaderX() int {
	return n.x
}

func (n *nodo) HeaderY() int {
	return n.y
}
//metodo
func (n *nodo) toString() string  {
	return "Nombre: "+ n.matriz.Nombre + "\n Descripcion: "+n.matriz.Descripcion
}

//var nuevo = &lista{nil,nil}

func (l *lista) Ordenar(nuevo *nodo)  {
	aux:=l.first
	for aux!=nil{
		if nuevo.header>aux.header{
			aux =aux.siguiente
		}else{
			if aux==l.first{
				nuevo.siguiente=aux
				aux.anterior=nuevo
				l.first= nuevo
			}else{
				nuevo.anterior=aux.anterior
				aux.anterior.siguiente=nuevo
				nuevo.siguiente=aux
				aux.anterior=nuevo
			}
			return
		}
	}
	l.last.siguiente = nuevo
	nuevo.anterior=l.last
	l.last=nuevo
}
func (l *lista) Insertar(header int)  {
	nuevo := NodoLista(header)
	if l.first==nil{
		l.first=nuevo
		l.last=nuevo
	}else {
		l.Ordenar(nuevo)
	}
}

func (l *lista) Buscar(header int) *nodo {
	temp:=l.first
	for temp!=nil{
		if temp.header==header{
			return temp
		}
		temp=temp.siguiente
	}
	fmt.Println("Error, no se encontro")
	return nil
	//return temp
}

func (l *lista) Imprimir()  {
	temporal:=l.first
	for temporal!=nil{
		fmt.Println("Cabecera: ",temporal.header)
		temporal=temporal.siguiente
	}
}