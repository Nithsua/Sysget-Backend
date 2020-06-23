package server

import (
	"fmt"
	"net/http"

	"github.com/austinmarner/system_admin_backend/deen"
	"github.com/austinmarner/system_admin_backend/top"
)

var topInfo top.SystemTopStruct

func responseUtilData(w http.ResponseWriter, r *http.Request) {
	temp, err := top.GetCPUUtil()
	topInfo.CPUUtil = temp
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		fmt.Printf("CPU Utilization: %.2f\n", topInfo.CPUUtil)
		bytes, err := deen.FormatData(topInfo)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
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
