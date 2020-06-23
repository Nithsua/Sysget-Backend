package server

import (
	"log"
	"net/http"

	"github.com/austinmarner/system_admin_backend/top"
)

var topInfo top.SystemTopStruct

func responseUtilData(w http.ResponseWriter, r *http.Request) {
	err := topInfo.RetriveInfo()
	// fmt.Println(topInfo.GetTop())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		topInfo.PrintData()
		temp := topInfo.GetTop()
		bytes, err := top.FormatData(temp)
		// fmt.Println(bytes)
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
	http.HandleFunc("/", responseUtilData)
	http.ListenAndServe(":8080", nil)
}
