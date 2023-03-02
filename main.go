package main

import(
	"fmt"
	"log"
	"net/http"
	"html/template"
)

func helloHandler(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/hello" {
		http.Error(w,"404 not found",http.StatusNotFound)
		return
	}

	if r.Method != "GET"{
		http.Error(w,"method is not supported", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w,"hello! bruh")

}



func formHandler(w http.ResponseWriter, r *http.Request){
	fmt.Println("method:", r.Method) 
	if r.Method == "GET"{
		tmpl,_ := template.ParseFiles("./static/form.html")
		tmpl.Execute(w,nil)
		return
	}

	r.ParseForm()

	fmt.Println("name: ",r.Form["name"])
	fmt.Println("address: ",r.Form["address"])

}


func main(){
	
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/",fileServer)
	http.HandleFunc("/form",formHandler)
	http.HandleFunc("/hello",helloHandler)


	fmt.Println("starting server at port 8080")
	if err := http.ListenAndServe(":8080",nil) ; err != nil{
		log.Fatal(err)
	}


}
