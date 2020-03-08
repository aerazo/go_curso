package main

import (
	"fmt"
	"strconv"

	"errors"

	"github.com/aerazo/cursogo/maps"
	"github.com/aerazo/cursogo/structs"
)

// "github.com/aerazo/cursogo/arrays"

const holamundo string = "hola %s %s bienvenido al curso de go"

func main() {

	fmt.Println(maps.GetMap())
	fmt.Println(maps.GetKeyMap("edad2"))
	// x, y, z := getMultiplesVariables()

	var miVariable structs.Alejandro = "Hola"
	fmt.Println(miVariable)

	alejandro := structs.Persona{
		Nombre:    "Alejandro",
		Apellido:  "erazo",
		Documento: "123",
		Telefono:  []string{"123", "asdasd"},
		Direccion: "buin",
		Edad:      123,
	}

	fmt.Println(alejandro)

	ale := new(structs.Persona)
	ale.Nombre = "nolse"

	fmt.Println(ale)

	casa := structs.Casa{
		NumeroCasa: 1,
		Personas:   []structs.Persona{alejandro, *ale},
	}

	fmt.Println(casa)

	// name := getName()

	// lastname := "erazo"

	// var numero = suma(40, 65)

	// fmt.Printf(holamundo, name, lastname)
	// fmt.Println("")
	// fmt.Println(x, y, z, numero)

	// arrays.ImprimeArray()
	// arrays.ImprimeSlice()
	// multiplo5()
	// multiple5Switch()
	// bucles()
	// opersConStrings()
	casa.GetNumeroCasa()

	casa.GetPersonasCasa()

	numero, error := Suma("A", 2)

	if error != nil {
		// panic(error)
	} else {
		fmt.Println(numero)
	}

	canal := make(chan string)

	punteros()
	runHolaMundo(30, canal)

	for valor := range canal {
		fmt.Println(valor)
	}
}

func punteros() {
	x := 100

	var y *int

	y = &x

	fmt.Println(x, *y)

}

func holaMundo(i int, canal chan<- string) {
	// fmt.Println("hola el numero es: ", i)
	canal <- "hola mundo num:  " + strconv.Itoa(i)
}

func runHolaMundo(numHilos int, canal chan<- string) {
	for i := 0; i < numHilos; i++ {
		go holaMundo(i, canal)
	}
}

// func getName() string {
// 	var name string = "default"

// 	fmt.Print("ingresa tu nombre: ")
// 	fmt.Scanf("%s", &name)

// 	return name
// }

// func getMultiplesVariables() (int, int, int) {
// 	return 1, 2, 3
// }

// func suma(a int, b int) int {
// 	return a + b
// }

//Suma suma dos valores
func Suma(a interface{}, b interface{}) (int, error) {

	switch a.(type) {
	case string:
		return 0, errors.New("el valor a no es un int")
	}

	switch b.(type) {
	case string:
		return 0, errors.New("el valor b no es un int")
	}

	return a.(int) + b.(int), nil
}

// func multiplo5() {
// 	var numero = 0

// 	fmt.Println("ingresa un numero: ")
// 	fmt.Scanf("%d", &numero)

// 	if numero%5 == 0 {
// 		fmt.Println("es multiplo de 5", numero)
// 	} else {
// 		fmt.Println("no es ", numero)
// 	}

// }

// func multiple5Switch() {
// 	numero := 0

// 	fmt.Println("ingresa un numero: ")
// 	fmt.Scanf("%d", &numero)
// 	switch numero {
// 	case 2:
// 		fmt.Println("el numero es 2")
// 	default:
// 		fmt.Println("el nro no es dos")
// 	}
// }

// func bucles() {
// 	for i := 0; i < 3; i++ {
// 		fmt.Println("el indice es: ", i)
// 	}
// }

// func opersConStrings() {
// 	texto := "hola alejandro, hola a todos"

// 	fmt.Println(texto)

// 	fmt.Println(strings.ToUpper(texto))
// }
