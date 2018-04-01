package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func checkUpdates(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	queryParams := mux.Vars(req)
	version := queryParams["v"]
	currentVersion := "1.0.0"
	if version != currentVersion {
		json.NewEncoder(w).Encode(true)
	}
}

func sendUpdate(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/octet-stream")
	os.Chdir("/home/shahab/go/src/github.com/shahab96/UbiAPI/src/")
	file, err := os.Open("crud")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	fileinfo, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return
	}

	filesize := fileinfo.Size()
	buffer := make([]byte, filesize)

	// bytesread, err := file.Read(buffer)

	if err != nil {
		fmt.Println(err)
		return
	}

	written, err := fmt.Fprintf(w, string(buffer))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(written)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/update_check", checkUpdates).Methods("GET")
	router.HandleFunc("/update", sendUpdate).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}
