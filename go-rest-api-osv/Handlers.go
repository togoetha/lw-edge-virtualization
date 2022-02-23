package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var Todos []Todo = []Todo{}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
	if err := json.NewEncoder(w).Encode(Todos); err != nil {
		panic(err)
	}
}

func TodoById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId, _ := strconv.Atoi(vars["todoId"])

	if todoId < len(Todos) {
		todoObj := Todos[todoId]

		if err := json.NewEncoder(w).Encode(todoObj); err != nil {
			panic(err)
		}
	} else {
		fmt.Fprintln(w, "Todo not found:", todoId)
	}
}

var largeDataStr []byte

func GetLargeData(w http.ResponseWriter, r *http.Request) {
	//if err := json.NewEncoder(w).Encode(largeDataStr); err != nil {
	//	panic(err)
	//}
	w.Write(largeDataStr)
}

func TodoCreate(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var todo Todo
	if err := decoder.Decode(&todo); err != nil {
		panic(err)
	}
	Todos = append(Todos, todo)
}

func PostBubbleSort(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var numbers []int
	if err := decoder.Decode(&numbers); err != nil {
		panic(err)
	}

	output := BubbleSort(numbers)
	if err := json.NewEncoder(w).Encode(output); err != nil {
		panic(err)
	}
}

func BubbleSort(numbers []int) []int {
	n := len(numbers)

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if numbers[j] > numbers[j+1] {
				numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
			}
		}
	}
	return numbers
}
