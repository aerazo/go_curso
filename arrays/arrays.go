package arrays

import "fmt"

//ImprimeArray  muestra como se labura
func ImprimeArray() {
	var array1 [2]string
	array1[0] = "hola"
	array1[1] = "mundo"

	array2 := [4]int{1, 2, 3, 4}

	fmt.Println(array1)
	fmt.Println(len(array1))
	fmt.Println(array2)

	var matriz [2][2]string
	matriz[0][0] = "hola"
	matriz[0][1] = "mundo"
	matriz[1][0] = "curso"
	matriz[1][1] = "go"

	fmt.Println(matriz)

}

//ImprimeSlice como se trabaja
func ImprimeSlice() {
	var slice1 []string
	slice1 = append(slice1, "hola", "slice")
	slice1 = append(slice1, "ne eleme")
	fmt.Println(slice1)
}
