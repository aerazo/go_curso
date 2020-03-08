package maps

//GetMap devuelve un tipo de dato mapa
func GetMap() map[string]int {
	// myMap := make(map[string]int)

	myMap := map[string]int{
		"Ale":   29,
		"Erazo": 30,
	}
	myMap["edad1"] = 18
	myMap["edad2"] = 19

	delete(myMap, "edad1")

	return myMap
}

//GetKeyMap devuelve el valor segun key
func GetKeyMap(key string) int {

	myMap := map[string]int{
		"Ale":   29,
		"Erazo": 30,
	}
	myMap["edad1"] = 18
	myMap["edad2"] = 19

	delete(myMap, "edad1")

	return myMap[key]
}
