package main

import "fmt"
import "reflect"

func main() {
	var name string = "Lemon"
	var age = 18 // tipagem implicita por inferencia
	var idade float32 = 18.5
	var salario float64 = 50000.50
	gender := "M" // tipagem explicita som sugar sintaze

	fmt.Println("Hello, world!", name)
	fmt.Println("Age: ", age)
	fmt.Println("Idade: ", idade)
	fmt.Println("Salario: ", salario)
	fmt.Println("Sexo: ", gender)
	fmt.Println(reflect.TypeOf(name))
	fmt.Println(reflect.TypeOf(age))
}
