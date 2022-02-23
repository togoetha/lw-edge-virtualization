package main

import (
        "C"
	"encoding/json"
	"log"
	"net/http"
)

func main() {
}

//export GoMain
func GoMain() {
	router := NewRouter()

	largeData := []int{}
	for i := 0; i < 50000; i++ {
		largeData = append(largeData, i%10)
	}
	largeDataStr, _ = json.Marshal(largeData)

	log.Fatal(http.ListenAndServe(":8080", router))
}
