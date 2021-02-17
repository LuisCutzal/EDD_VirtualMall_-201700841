package ll

import ("fmt"
"../lista"
)
//nodos de informacion
type nodo struct {
	lst *lista.Lista
	siguiente, anterior *nodo
	contenido int
}
func NuevoNodo(index int)*nodo  {
	return &nodo{lista.NuevaLista(),nil,nil,index}
}
//estructura para almacenar nodos de informacion
type ListaDelistas struct {
	inicio, ultimo *nodo
	tam int
}
//creamos una lista de listas
func NuevaLista() *ListaDelistas  {
	return &ListaDelistas{nil,nil,0}
}
func (m *ListaDelistas) InsertarListaListas(index int, dato string)  {
	nuevo:=NuevoNodo(index)
	if m.inicio==nil{
		m.inicio=nuevo
		m.ultimo=nuevo
		nuevo.lst.Insertar(dato)
	}else{
		auxiliar:=m.inicio
		for auxiliar!=nil{
			if auxiliar.contenido==index{
				auxiliar.lst.Insertar(dato)
				return
			}
			auxiliar=auxiliar.siguiente
		}
		nuevo.lst.Insertar(dato)
		m.ultimo.siguiente=nuevo
		nuevo.anterior=m.ultimo
		m.ultimo=nuevo
	}
	m.tam++
}
func (m *ListaDelistas) Print()  {
	auxiliar:=m.inicio
	for auxiliar!=nil {
		fmt.Println("------index",auxiliar.contenido,"------")
		auxiliar.lst.Imprimir()
		auxiliar=auxiliar.siguiente
	}
}