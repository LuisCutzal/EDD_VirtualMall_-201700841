package matriz

import ("fmt")

type Tienda struct {
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
	matriz *Tienda //tipo de objeto
	izquierdo, derecho, arriba, abajo *nodo //nodos para desplazarnos dentro de la matriz

	//estos atributos son especificos para la lista
	header int //tipo interno de cabecera
	dato string
	siguiente,anterior *nodo //nodo con los que nos vamos a desplazar dentro de las listas
}

type lista struct {
	first, last *nodo
}
type Matriz struct {
	listH,listV *lista
}

//creamos un constructor para una nueva matriz
func NodoMatriz(x int ,y int, matriz *Tienda)*nodo  {
	return &nodo{x,y,matriz,nil,nil,nil,nil,0,"",nil, nil}

}
//constructor para la lista de listas
func NodoLista(header int, dato string) *nodo  {
	return &nodo{0,0,nil,nil,nil,nil,nil,header,dato,nil,nil}
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

func (l *lista) Insertar(header int, dato string)  {
	nuevo := NodoLista(header, dato)
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
	//fmt.Println("Error, no se encontro",header)
	return nil
	//return temp
}
func (l *lista) Imprimir()  {
	temporal:=l.first
	for temporal!=nil{
		fmt.Printf("Cabecera: %d Nombre: %s \n",temporal.header, temporal.dato)
		temporal=temporal.siguiente
	}
}



func (m *Matriz) Insertar(producto Tienda, x int, y int, dato string) {
	horizontal:=m.listH.Buscar(x)
	v:=m.listV.Buscar(y)
	if horizontal==nil && v==nil{
		m.NoExiste(producto,x,y,dato)
	}else if horizontal==nil && v!=nil{
		m.ExisteVertical(producto,x,y,dato)
	}else if horizontal!=nil && v==nil{
		m.ExisteHorizontal(producto,x,y, dato)
	}else{
		m.Existe(producto,x,y)
	}

}

func (m *Matriz)NoExiste(producto Tienda, x int, y int, dato string) {
	m.listH.Insertar(x,dato) //insertamos en la lista que emula la cabecera horizontal
	m.listV.Insertar(y,dato) //insertamos en la lista que emula la cabecera vertical
	h:=m.listH.Buscar(x)//buscamos el nodo que acabamos de incertar para poder enlazar
	v:=m.listV.Buscar(y)

	nuevo:=NodoMatriz(x,y,&producto) //creamos nuevo nodo tipo matriz
	h.abajo=nuevo//enlazamos el nodo horizontal hacia abajo
	nuevo.arriba=h //enlazamos el nuevo nodo hacia arriba
	v.derecho=nuevo //enlazamos el nodo vertical hacia la derecha
	nuevo.izquierdo=v //enlazamos el nuevo nodo hacia la izquierda
}

func (m *Matriz) ExisteVertical(producto Tienda, x int, y int, dato string) {
	m.listH.Insertar(x,dato) //insertamos en la lista que emula la cabecera horizontal
	h:=m.listH.Buscar(x)
	v:=m.listV.Buscar(y)
	var agregado =false
	nuevo:=NodoMatriz(x,y,&producto)
	auxiliar:=v.derecho
	var headerx int
	for auxiliar!=nil{
		headerx=auxiliar.header
		if headerx < x{
			auxiliar=auxiliar.derecho
		}else{
			nuevo.derecho=auxiliar
			nuevo.izquierdo=auxiliar.izquierdo
			auxiliar.izquierdo.derecho=nuevo
			auxiliar.izquierdo=nuevo
			agregado=true
		}
	}
	if agregado==false {
		auxiliar=v.derecho
		for auxiliar.derecho!=nil{
			auxiliar=auxiliar.derecho
		}
		nuevo.izquierdo=auxiliar
		auxiliar.derecho=nuevo
	}
	h.abajo=nuevo
	nuevo.arriba=h

}

func (m *Matriz) ExisteHorizontal(producto Tienda, x int, y int,dato string) {
	m.listV.Insertar(y,dato)
	h:=m.listH.Buscar(x)
	v:=m.listV.Buscar(y)
	var agregado =false
	nuevo:=NodoMatriz(x,y,&producto)
	auxiliar:=h.abajo
	for auxiliar!=nil{
		headery:=auxiliar.HeaderY()
		if headery<y {
			auxiliar=auxiliar.abajo
		}else{
			nuevo.abajo=auxiliar
			nuevo.arriba=auxiliar.arriba
			auxiliar.arriba.abajo=nuevo
			auxiliar.arriba=nuevo
			agregado=true
		}
	}
	if agregado ==false{
		auxiliar=h.abajo
		for auxiliar.abajo!=nil{
			auxiliar=auxiliar.abajo
		}
		nuevo.arriba=auxiliar
		auxiliar.abajo=nuevo
	}
	v.derecho=nuevo
	nuevo.izquierdo=v
}

func (m *Matriz) Existe(producto Tienda, x int, y int) {
	h:=m.listH.Buscar(x)
	v:=m.listV.Buscar(y)
	nuevo:=NodoMatriz(x,y,&producto)
	var agregado =false
	auxiliar:=h.abajo
	//var header_y int
	for auxiliar!=nil{
		headery:=auxiliar.HeaderY()
		if headery<y {
			auxiliar=auxiliar.abajo
		}else{
			nuevo.abajo=auxiliar
			nuevo.arriba=auxiliar.arriba
			auxiliar.arriba.abajo=nuevo
			auxiliar.arriba=nuevo
			agregado=true
		}
	}
	if agregado ==false {
		auxiliar=h.abajo
		for auxiliar.abajo!=nil {
			auxiliar=auxiliar.abajo
		}
		nuevo.arriba=auxiliar
		auxiliar.abajo=nuevo
	}
	auxiliar=v.derecho
	for auxiliar!=nil{
		headery:=auxiliar.HeaderX()
		if headery < x{
			auxiliar=auxiliar.derecho
		}else{
			nuevo.derecho=auxiliar
			nuevo.izquierdo=auxiliar.izquierdo
			auxiliar.izquierdo.derecho=nuevo
			auxiliar.izquierdo=nuevo
			agregado=true
		}
	}
	if agregado==false {
		auxiliar=v.derecho
		for auxiliar.derecho!=nil {
			auxiliar=auxiliar.derecho
		}
		nuevo.izquierdo=nuevo
		auxiliar.derecho=nuevo
	}
}

