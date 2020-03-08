package structs

import "fmt"

type Alejandro string

type Persona struct {
	Nombre    string
	Apellido  string
	Documento string
	Telefono  []string
	Direccion string
	Edad      int
}

type Casa struct {
	NumeroCasa int
	Personas   []Persona
}

func (c Casa) GetNumeroCasa() {
	fmt.Println("el numero de la casa es: ", c.NumeroCasa)
}

func (c Casa) GetPersonasCasa() {
	defer SoyDefer()
	fmt.Println("las personas de la casa son: ", c.Personas)
}

func SoyDefer() {
	fmt.Println("soy defer, esto se va a ejecutar al final")
}
