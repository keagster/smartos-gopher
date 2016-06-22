package main

// Once again large chunk from rest5 i added to go-snippets


import (
	"net/http"
	"time"
	"os"
	"fmt"
	"encoding/json"
	_ "net"
	_ "os/exec"
)
/*
 #####################################
 # HomeHandler: Root Handler for "/" #
 #####################################
 This handler gives the application information about the host of the service
*/
func HomeHandler(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	type HomeResponse struct {
		Hostname string
		SystemUser string
		Time string
	}

	// Set all live vars before marshal in such a way that they update each request
	time_get := time.Now()
	user_get := os.Getenv("USER")
	hostname_get, hostname_get_err := os.Hostname()
	if hostname_get_err != nil {
		fmt.Println("Error: HomeHandler: Error getting hostname")
	}

	response_data := HomeResponse{hostname_get, user_get, time_get.String()}
	HomeResponse_json, err := json.Marshal(response_data)

	if err == nil {
		w.Write(HomeResponse_json)
	}
}