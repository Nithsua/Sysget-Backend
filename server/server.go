package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/austinmarner/system_admin_backend/top"
)

var topInfo top.SystemTopStruct

func responseUtilData(w http.ResponseWriter, r *http.Request) {
	topInfo.Reset()
	err := topInfo.RetriveInfo()
	// fmt.Println(topInfo)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		topInfo.PrintData()
		temp := topInfo.GetTop()
		fmt.Println(temp)
		bytes, err := top.FormatData(temp)
		fmt.Println(string(bytes))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Fatal(err.Error())
		}
		w.WriteHeader(http.StatusAccepted)
		w.Write(bytes)
	}

}

//InitiateServer ...
func InitiateServer(topInfo top.SystemTopStruct) {
	http.HandleFunc("/topData", responseUtilData)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err.Error())
	}
}
