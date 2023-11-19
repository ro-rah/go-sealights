package main

import (
	"fmt"
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

func main() {
	myService := &MyService{}

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

	fmt.Println("Server listening on :8080")
	http.ListenAndServe(":8080", nil)
}
