package main

import "fmt"

func main() {
    
    var contry map[string]string //declare a map

	contry = make(map[string]string) // initialize the map

	contry["IN"] = "india"
	contry["US"] = "unitedd state"

	fmt.Println(contry) // Output: map[IN:India US:United States]


	//short 
	state := map[string]string{
		"GJ" : "guajrat",
		"MP" : "maharastra",
	}
    //modify
	state["MP"] = "madhayapradesh"
	fmt.Println(state)

	//Checking If a Key Exists
	countries := map[string]string{
        "IN": "India",
        "US": "United States",
    }

    value, exists := countries["UK"] // Check if "UK" exists
    if exists {
        fmt.Println("Found:", value)
    } else {
        fmt.Println("Key not found")
    }

	//Deleting a Key
	delete(countries, "US") // Remove key "US"
    fmt.Println(countries) // Output: map[IN:India]

	//Iterating Over a Map
	for key, value := range countries {
        fmt.Println(key, "->", value)
    }
}
