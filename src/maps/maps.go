package main

import "fmt"

// Vertex struct contains lat,long float coordinate points
type Vertex struct {
	Lat, Long float64
}

func main() {
	var hashmap map[string]Vertex
	hashmap = make(map[string]Vertex)

	hashmap["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	hashmap["office"] = Vertex{
		40.750810, -73.973540,
	}

	var map2 = map[string]Vertex{
		"Google": Vertex{
			37.42202, -122.08408,
		},
		"Google_w/o_type_declaration": {37.42202, -122.08408},
	}

	var o = hashmap["office"]
	fmt.Println("should be nyc", o)
	fmt.Println(hashmap["Bell Labs"])

	var tx = "Bell Labs"
	delete(hashmap, tx)
	fmt.Println("should be nil", hashmap[tx])

	fmt.Println(hashmap["office"])
	fmt.Println(map2, map2["Google"], map2["Google_w/o_type_declaration"])

	e, ok := hashmap[tx]
	fmt.Println(e,ok)
	e, ok = hashmap["office"]
	fmt.Println(e,ok)
	
}
