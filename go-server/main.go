package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request){
	if err:= r.ParseForm(); err!=nil{
		fmt.Fprintf(w,"Praseform() err: %v",err)
		return
	}
	fmt.Fprintln(w,"POST request successful")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w,"UserName = %s \n",name)
	fmt.Fprintf(w,"adderess = %s \n", address)
}

func helloHandler(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/hello"{
		http.Error(w,"404 not found BRO", http.StatusNotFound)
		return
	}
	if r.Method != "GET"{
		http.Error(w,"method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprint(w,"Hello there this is Giridharan")
}

func main(){
	fileServer := http.FileServer(http.Dir("./server"))
	http.Handle("/",fileServer)
	http.HandleFunc("/form",formHandler)
	http.HandleFunc("/hello",helloHandler)

	fmt.Printf("Starting server at port 8080\n")
	if err:= http.ListenAndServe(":8080",nil); err != nil {
		log.Fatal(err)
	}
}