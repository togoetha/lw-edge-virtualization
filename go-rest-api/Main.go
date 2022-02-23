package main

import (
	//	"C"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// "C"

func main() {
	router := NewRouter()

	fmt.Println(time.Now().UnixNano())

	largeData := []int{}
	for i := 0; i < 50000; i++ {
		largeData = append(largeData, i%10)
	}
	largeDataStr, _ = json.Marshal(largeData)
	/*fmt.Print("[")
	for x := 20000; x > 0; x-- {
		fmt.Printf("%d,", x)
	}
	fmt.Println("]")*/

	log.Fatal(http.ListenAndServe(":8080", router))

}

//export GoMain
/*func GoMain() {
	router := NewRouter()

	largeData := []int{}
	for i := 0; i < 50000; i++ {
		largeData = append(largeData, i%10)
	}
	largeDataStr, _ = json.Marshal(largeData)

	log.Fatal(http.ListenAndServe(":8080", router))
}*/
