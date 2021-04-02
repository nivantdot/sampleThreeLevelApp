package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sampleThreeLevelApp/utils"

	"github.com/gorilla/mux"
)

func getConsumers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "consumer side visible only to consumer and admin")

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("consumerList")
}

func getProducer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "producer side visible only to consumer and admin")

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("consumerList")
}

func getAdmin(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "visible only to admin")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("consumerList")
}

func welcome(w http.ResponseWriter, r *http.Request) {

	str := fmt.Sprintf("welcome to 3 level sample. " +
		"Admin and two connecting sides")

	json.NewEncoder(w).Encode(str)
}

// Controller ...
func Controller() {
	localip, err := utils.GetLocalIP()
	if err != nil {
		log.Println("ip error occurred ", err)
		return
	}
	log.Println(localip)

	r := mux.NewRouter()

	r.HandleFunc("/", welcome).Methods("GET", "PUT", "POST", "DELETE")
	r.HandleFunc("/api/consumers", getConsumers).Methods("GET")
	r.HandleFunc("/api/producers", getProducer).Methods("GET")
	r.HandleFunc("/api/admin", getAdmin).Methods("GET")

	http.ListenAndServe(localip+":12345", r)
}
