package server

import (
	"log"
	"net/http"

	"github.com/austinmarner/system_admin_backend/top"
)

var topInfo top.Top

func responseUtilData(w http.ResponseWriter, r *http.Request) {
	topInfo.Reset()
	err := topInfo.RetriveInfo()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		// topInfo.PrintData()
		temp := topInfo.GetTop()
		bytes, err := top.FormatData(temp)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Fatal(err.Error())
		}
		w.WriteHeader(http.StatusAccepted)
		w.Write(bytes)
	}

}

//InitiateServer ...
func InitiateServer(topInfo top.Top) {
	http.HandleFunc("/topData", responseUtilData)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err.Error())
	}
}
