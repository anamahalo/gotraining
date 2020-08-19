// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare and make a map of integer values with a string as the key. Populate the
// map with five values and iterate over the map to display the key/value pairs.
package main

// Add imports.
import "fmt"

func main() {

	// Declare and make a map of integer type values.
    notableWomen := make(map[string]int)

	// Initialize some data into the map.
    notableWomen["Hedy Lamarr"] = 1914
    notableWomen["Grace Hopper"] = 1906
    notableWomen["Kitty O'Brien Joyner"] = 1952
    notableWomen["Barbara Jean Paulson"] = 1928
    notableWomen["Margaret Hamilton"] = 1936
	// Display each key/value pair.

    for k,v := range notableWomen {
        fmt.Printf("Notable Woman: %s | Birth Year: %d\n", k, v)
    }
}
