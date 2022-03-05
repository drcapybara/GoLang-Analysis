package main

import (
	"fmt"
	"sort"
)

//instantiate a new type
type animalLanguage struct {
	animal string
}

func main() {

	//kinda generic and good enough in some situations
	sortOfGeneric := make(map[string]animalLanguage)

	sortOfGeneric["python"] = animalLanguage{animal: "snake"}
	sortOfGeneric["rust"] = animalLanguage{animal: "rustacean"}
	sortOfGeneric["go"] = animalLanguage{animal: "gopher"}
	sortByString(sortOfGeneric)
}

func sortByString(sortMe map[string]animalLanguage) {

	keys := make([]string, 0, len(sortMe))
	for k := range sortMe {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Println(k, sortMe[k])
	}

}
