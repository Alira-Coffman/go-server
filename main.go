package main

import(
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter,r *http.Request){
	//error checking
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "hello!")
}

func formHandler(w http.ResponseWriter,r *http.Request){
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	fmt.Fprintf(w, "Post request successful!\n");
	name := r.FormValue("name")
	address := r.FormValue("address")

	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}


func main(){
	//check out static directory
	fileServer := http.FileServer(http.Dir("./static")) //:= short form. declares and defines
	http.Handle("/", fileServer) //handles emtpy route
	http.HandleFunc("/form", formHandler) //deals with form
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting Server..... Connection at port 8080........\n")
	if err := http.ListenAndServe(":8080", nil); err!= nil {
		log.Fatal(err) //uses log package
	}

}