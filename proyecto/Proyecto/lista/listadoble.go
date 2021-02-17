package lista

import "fmt"
//nodos de informacion
type nodo struct {
	siguiente , anterior *nodo
	contenido string
}
//estructura para almacenar nodos de informacion
type Lista struct {
	inicio, ultimo *nodo
	tam int
}
//creamos una nueva lista
func NuevaLista() *Lista  {
	return &Lista{nil,nil,0}
}

type Tienda struct {
	Nombre       string `json:Nombre`
	Descripcion  string `json:Descripcion`
	Contacto     string `json:Contacto`
	Calificacion int `json:Calificacion`
}


//insertar un nodo
func (m *Lista) Insertar(dato string)  {
	nuevo := &nodo{nil,nil,dato}
	if m.inicio == nil{
		m.inicio=nuevo
		m.ultimo=nuevo
	}else {
		m.ultimo.siguiente=nuevo
		nuevo.anterior=m.ultimo
		m.ultimo=nuevo
	}
	m.tam++

}

/*
func (m *Lista) Insertar(dato string)  {
	nuevo := &nodo{nil,nil,dato}
	if m.inicio == nil{
		m.inicio=nuevo
		m.ultimo=nuevo
	}else {
		m.ultimo.siguiente=nuevo
		nuevo.anterior=m.ultimo
		m.ultimo=nuevo
	}
	m.tam++
}

*/



//imprimimos la lista
func (m *Lista)Imprimir()  {
	aux:=m.inicio
	for aux!=nil{
		//fmt.Print("<-[",aux.contenido,"]->")
		fmt.Printf("<-[ %s ]->",aux.contenido)
		aux=aux.siguiente
	}
	fmt.Println()
	fmt.Println("Tamaño lista: ", m.tam)
}

//buscar elemento dentro de lista
func (m *Lista)Buscar(valor string) *nodo  {
	//func, recive una lista, nombre de la funcion, resive como parametro un valor string, y retorna un nodo (o una referencia)
	aux:=m.inicio
	for aux!=nil{
		if aux.contenido ==valor{
			//fmt.Println("Si se encontro el nodo", valor)
			fmt.Printf("Si se encontro el nodo %s \n",valor)
			return aux
		}
		aux = aux.siguiente
	}
	//fmt.Println("No se encontro el nodo", valor)
	fmt.Printf("No se encontro el nodo %s \n",valor)
	return aux
}

//eliminar un nodo (cualquier posicion) de la lista
func (m *Lista)Eliminar(valor string)  {
	aux:=m.Buscar(valor) //le asignaremos el nodo que nosotros encontremos y le mandamos como parametro el valor
	//vamos a verificar varios casos, ya sea eliminar al principio, al final o en cualquier otra posicion de la lista
	if m.tam==1{
		fmt.Println("No se puede eliminar el elemento")
	}else if aux==nil || aux.siguiente==nil || aux.anterior==nil{
		fmt.Println("No existe el nodo")
	}else{
		if m.inicio==aux{
			m.inicio=aux.siguiente
			aux.siguiente.anterior=nil //para eliminar la referencia
			aux.siguiente=nil //para eliminar la referencia y que este nodo desaparezca
		}else if m.ultimo==aux {
			m.ultimo=aux.anterior
			aux.anterior.siguiente=nil
			aux.anterior=nil
		}else{ //navegamos dentro de la lista hasta encontrar el elemento que queremos eliminar
			aux.anterior.siguiente=aux.siguiente //movemos el puntero hacia otro nodo
			aux.siguiente.anterior=aux.anterior //
			aux.anterior=nil //eliminamos
			aux.siguiente=nil
		}
		m.tam-- //para eliminar el espacio del contador
	}

}