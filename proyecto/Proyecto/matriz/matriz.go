package matriz

import ("fmt")

type Product struct {
	//Nombre string//
	//Calificacion int //
	//Descripcion string //datos


	Nombre string
	Descripcion string
	Contacto string
	Calificacion int
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
	return &nodo{x,y,matriz,nil,nil,nil,nil,0,nil,nil}

}
//constructor para la lista de listas
func NodoLista(header int) *nodo  {
	return &nodo{0,0,nil,nil,nil,nil,nil,header,nil,nil}
}

func NuevaLista() *lista {
	return &lista{nil,nil}
}

//constructor para la matriz
func NuevaMatriz() *Matriz  {
	return &Matriz{NuevaLista(),NuevaLista()}
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

func (m *Matriz) Insertar(producto *Product, x int, y int) {
	horizontal:=m.list_H.Buscar(x)
	v:=m.list_V.Buscar(y)
	if horizontal==nil && v==nil{
		m.NoExiste(producto,x,y)
	}else if horizontal==nil && v!=nil{
		m.ExisteVertical(producto,x,y)
	}else if horizontal!=nil && v==nil{
		m.ExisteHorizontal(producto,x,y)
	}else{
		m.Existe(producto,x,y)
	}

}

func (m *Matriz)NoExiste(producto *Product, x int, y int) {
	m.list_H.Insertar(x) //insertamos en la lista que emula la cabecera horizontal
	m.list_V.Insertar(y) //insertamos en la lista que emula la cabecera vertical
	h:=m.list_H.Buscar(x)//buscamos el nodo que acabamos de incertar para poder enlazar
	v:=m.list_V.Buscar(y)

	nuevo:=NodoMatriz(x,y,producto) //creamos nuevo nodo tipo matriz
	h.abajo=nuevo//enlazamos el nodo horizontal hacia abajo
	nuevo.arriba=h //enlazamos el nuevo nodo hacia arriba
	v.derecho=nuevo //enlazamos el nodo vertical hacia la derecha
	nuevo.izquierdo=v //enlazamos el nuevo nodo hacia la izquierda
}

func (m *Matriz) ExisteVertical(producto *Product, x int, y int) {
	m.list_H.Insertar(x) //insertamos en la lista que emula la cabecera horizontal
	m.list_V.Insertar(y) //insertamos en la lista que emula la cabecera vertical
	h:=m.list_H.Buscar(x)//buscamos el nodo que acabamos de incertar para poder enlazar
	v:=m.list_V.Buscar(y)

	nuevo:=NodoMatriz(x,y,producto) //creamos nuevo nodo tipo matriz
	h.abajo=nuevo//enlazamos el nodo horizontal hacia abajo
	nuevo.arriba=h //enlazamos el nuevo nodo hacia arriba
	v.derecho=nuevo //enlazamos el nodo vertical hacia la derecha
	nuevo.izquierdo=v //enlazamos el nuevo nodo hacia la izquierda
}

func (m *Matriz) ExisteHorizontal(producto *Product, x int, y int) {
	m.list_H.Insertar(x) //insertamos en la lista que emula la cabecera horizontal
	m.list_V.Insertar(y) //insertamos en la lista que emula la cabecera vertical
	h:=m.list_H.Buscar(x)//buscamos el nodo que acabamos de incertar para poder enlazar
	v:=m.list_V.Buscar(y)

	nuevo:=NodoMatriz(x,y,producto) //creamos nuevo nodo tipo matriz
	h.abajo=nuevo//enlazamos el nodo horizontal hacia abajo
	nuevo.arriba=h //enlazamos el nuevo nodo hacia arriba
	v.derecho=nuevo //enlazamos el nodo vertical hacia la derecha
	nuevo.izquierdo=v //enlazamos el nuevo nodo hacia la izquierda
}

func (m *Matriz) Existe(producto *Product, x int, y int) {
	m.list_H.Insertar(x) //insertamos en la lista que emula la cabecera horizontal
	m.list_V.Insertar(y) //insertamos en la lista que emula la cabecera vertical
	h:=m.list_H.Buscar(x)//buscamos el nodo que acabamos de incertar para poder enlazar
	v:=m.list_V.Buscar(y)

	nuevo:=NodoMatriz(x,y,producto) //creamos nuevo nodo tipo matriz
	h.abajo=nuevo//enlazamos el nodo horizontal hacia abajo
	nuevo.arriba=h //enlazamos el nuevo nodo hacia arriba
	v.derecho=nuevo //enlazamos el nodo vertical hacia la derecha
	nuevo.izquierdo=v //enlazamos el nuevo nodo hacia la izquierda
}

