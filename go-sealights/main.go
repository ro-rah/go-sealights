package main

import (
	"fmt"
	"encoding/json"
	"net/http"
)

type MyService struct{}

func (s *MyService) MethodA(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Method A called")
}
func (s *MyService) MethodB(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Method B called")
}
func (s *MyService) MethodC(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Method C called")
}
func (s *MyService) MethodD(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Method D called")
}	
func (s *MyService) MethodE(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Method E called")
}
func (s *MyService) MethodF(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Method F called")
}
func (s *MyService) MethodG(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Method G called")
}
func (s *MyService) MethodH(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Method H called")
}
func (s *MyService) MethodI(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Method I called")
}
func (s *MyService) MethodJ(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Method J called")
}


// New method added for AngularJS frontend
func (s *MyService) GetEndpointNames(w http.ResponseWriter, r *http.Request) {
	// You might want to dynamically fetch endpoint names based on your Go API configuration
	endpointNames := []string{"MethodA", "MethodB", "MethodC", "MethodD", "MethodE", "MethodF", "MethodG", "MethodH", "MethodI", "MethodJ"}

	// Create a struct to represent the JSON response
	responseStruct := struct {
		Endpoints []string `json:"endpoints"`
	}{
		Endpoints: endpointNames,
	}

	// Marshal the struct into JSON
	jsonResponse, err := json.Marshal(responseStruct)
	if err != nil {
		http.Error(w, "Failed to marshal JSON response", http.StatusInternalServerError)
		return
	}

	// Set content type and write the JSON response
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}



func main() {
	myService := &MyService{}
	
	// Serve static files from the /app directory
	http.Handle("/", http.FileServer(http.Dir("/app")))

	http.HandleFunc("/methodA", myService.MethodA)
	http.HandleFunc("/methodB", myService.MethodB)
	http.HandleFunc("/methodC", myService.MethodC)
    http.HandleFunc("/methodD", myService.MethodD)
    http.HandleFunc("/methodE", myService.MethodE)
    http.HandleFunc("/methodF", myService.MethodF)
    http.HandleFunc("/methodG", myService.MethodG)
    http.HandleFunc("/methodH", myService.MethodH)
    http.HandleFunc("/methodI", myService.MethodI)
    http.HandleFunc("/methodJ", myService.MethodJ)

	// Add a new endpoint for AngularJS to get endpoint names
	http.HandleFunc("/getEndpointNames", myService.GetEndpointNames)

	fmt.Println("Server listening on :8080")
	http.ListenAndServe(":8080", nil)
}
