package main

import "fmt"
		

func main() {
	
monica:=Persona{"Monica"}
	fmt.Println(monica.correr())
}


type Persona struct{
	nombre string
}
	
func (per *Persona) correr() string {
	x:=per.nombre
	x+=" Corra!!!!!"
	return x
}
	
	
