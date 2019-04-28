package main

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func answerResponse(w http.ResponseWriter, r *http.Request) {

	data := []byte(`[
    {
      "action": "talk",
      "text": "Please wait while we connect you."
    },
    {
      "action": "connect",
      "timeout": 20,
      "from": "14155493310",
      "endpoint": [
        {
          "type": "phone",
          "number": "15629006541"
        }
      ]
    }
  ]`)

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)

}

func eventResponse(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)

	defer r.Body.Close()

	if body == nil {
        return
    }


	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var msg interface{}

	err = json.Unmarshal(body, &msg)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	output, err := json.Marshal(msg)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}


	fmt.Println(string(output))

	w.WriteHeader(http.StatusOK)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/answer", answerResponse).Methods("GET")
	r.HandleFunc("/event", eventResponse).Methods("POST")
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
