package main

import "fmt"

func main() {

	fmt.Println("This is demo on Maps")

	//Follwing line creates the Map

	myMap := make(map[string]string)

	fmt.Println("my map data : ", myMap)
	fmt.Printf("my map type :%T\n ", myMap)
	fmt.Println("my map len is:", len(myMap))

	myMap["frnd1"] = "one"
	myMap["frnd2"] = "two"
	myMap["frnd3"] = "three"
	myMap["frnd4"] = "four"

	fmt.Println("my map data after adding: ", myMap)
	fmt.Println("my map len is:", len(myMap))
	myMap["frnd5"] = "five"

	fmt.Println("my map data after adding another: ", myMap)
	fmt.Println("my map len is:", len(myMap))

	//delete the specified value
	delete(myMap, "frnd3")

	fmt.Println("my map data after deleting: ", myMap)
	fmt.Println("my map len is:", len(myMap))

	fmt.Println("Looping over my map data")
	//Loop over the maps
	for key, value := range myMap {
		fmt.Println("Key\t:", key, "value\t:", value)

	}

}
