package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"operation"
)

type CalcHandler struct {
	History []operation.Operation
}

func (history *CalcHandler) OperationsHandler(w http.ResponseWriter, r *http.Request) {
	a, b, operationType := ParseParams(r)
	op := operation.Operation{A: a, B: b, Operation: operationType}
	
	if result, err := op.Calculate() ; err != nil {
		message := fmt.Sprintf("Error: %v", err)
		http.Error(w, message, http.StatusBadRequest)
	} else {
		history.History = append(history.History, op)
		fmt.Fprintf(w, "%f\n", result)
	}
}

func (history *CalcHandler) HistoryHandler(w http.ResponseWriter, r *http.Request) {
	historyJson, _ := json.Marshal(history.History)
	fmt.Fprintf(w, string(historyJson))
}

func ParseParams(r *http.Request) (float64, float64, string) {
	vars := mux.Vars(r)
	a, _ := strconv.ParseFloat(vars["a"], 64)
	b, _ := strconv.ParseFloat(vars["b"], 64)
	op := vars["operation"]

	return a, b, op
}

func main() {
	router := mux.NewRouter()
	handler := &CalcHandler{History: []operation.Operation{}}

	router.HandleFunc("/calc/{operation}/{a}/{b}", handler.OperationsHandler)
	router.HandleFunc("/calc/history", handler.HistoryHandler)
	http.Handle("/", router)

	log.Fatal(http.ListenAndServe(":8080", router))
}
