package main

import "fmt"

func main() {
  fmt.Print("Ingrese un texto: ")
  var txt string
  fmt.Scanf("%s", &txt)

  fmt.Println(txt)
}

/*EJERCICIO: Cambie el anterior programa para que en vez de capturar
	un decimal, capture un texto y lo imprima.*/