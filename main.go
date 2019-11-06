package main

import (
	"fmt"
	"github.com/shrikar007/01-mongo-example/database"
	"html/template"
	"log"
	"net/http"
)

func main(){
	var db database.Database
//aws:=types.Request{
//	Name:"shrikar vaitala",
//	Timestamp:time.Now().Format("01-02-2006 15:04:05"),
//	Prioriy:"critical",
//	ProjectName:"aws",
//	Service:"ec2",
//	Region:"us",
//	Details:"demo",
//	Rperiod:"demo",
//	ApprovedBy:"yash",
//}

	db = &database.MongoDB{
		Host: "localhost", Port: 27017,
	}

	err := db.Open()
	if err != nil {
		log.Fatal(err)
	}
  //err = db.Save(aws)
	//if err != nil {
	//	log.Fatal(err)
	//}

  res,err1:=db.Retrieve()
	if err1 != nil {
		log.Fatal(err1)
	}
	http.HandleFunc("/display", func(writer http.ResponseWriter, request *http.Request) {
		tmpl := template.Must(template.ParseFiles("display.html"))
		err = tmpl.Execute(writer,res)
		if err != nil {
			fmt.Println("executing template:", err)
		}
	})
	fmt.Println(http.ListenAndServe(":8080",nil))

  err=db.Close()
	if err != nil {
		log.Fatal(err)
	}
}
