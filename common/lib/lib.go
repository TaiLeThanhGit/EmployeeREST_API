package Lib

import (
	Config "employee/config"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func JsonResponse(w http.ResponseWriter, i interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(i)
}

func ShowDebugInfo(i ...interface{}) {
	if (len(os.Args) > 1 && os.Args[1] == "true") || Config.DEBUG == true {
		fmt.Println(i...)
	}
}
