package main

import "fmt"

func main() {
	hola := "Hola"
	mundo := "Mundo"

	fmt.Println(hola)
	fmt.Println(mundo)

	nombre := "Javier"
	edad := 28
	fmt.Printf("Hola , %s y su edad es %d \n", nombre, edad)

	fmt.Printf("Hola , %v y su edad es %v \n", nombre, edad)

	mensaje := fmt.Sprintf("Hola , %s y su edad es %d \n", nombre, edad)

	fmt.Println(mensaje)
	//tipo de dato de una variable
	fmt.Printf("nombre: %T \n", edad)
	//
	fmt.Print("Ingrese otro nombre : ")
	fmt.Scanln(&nombre)
	fmt.Println("Otro nombre", nombre)
}
