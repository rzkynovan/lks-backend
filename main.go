package main

import (
	"fmt"
	"kanolibrary/routes"
	"kanolibrary/utils"
	"net/http"
)

//connect

func main() {
	//books.Insert("books", models.FindAllBook{
	//	Title:     "Seni Mereog",
	//	Author:    "Annisa Nur Isnaeni",
	//	Available: true,
	//	CreatedAt: primitive.NewDateTimeFromTime(time.Now()),
	//})
	//books.FindMany("books", bson.M{"_id": "ObjectId('62b5a4549587967663b278a6')"})
	//books.DeleteOne("books", bson.M{"name": "The Moles"}) //done!
	//
	//books.FindMany("books", bson.M{"name": "Lord of the Rings"})
	//books.UpdateOne("books", bson.M{"name": "The Mole"}, bson.M{"$set": bson.M{"available": false}}) //done! but little bit buggy
	//even id the name is wrong it will still print update success

	//books.FindMany("books", bson.M{"ID": "62b599109e2471cde91762f5"}) //findmany by id stiil error
	//controllers.FindAll("books")

	//objectId, err := primitive.ObjectIDFromHex("62bc10ca42ba1645a909d5ab")
	//result, err := books.FindMany("books", bson.M{"_id": objectId})
	//utils.ErrorHandler(err)
	//fmt.Println(result)

	http.HandleFunc("/books", routes.FindAllBook)
	http.HandleFunc("/books/find", routes.FindByIdBook)
	http.HandleFunc("/books/insert", routes.InsertBook)
	http.HandleFunc("/books/update", routes.UpdateBook)
	http.HandleFunc("/books/delete", routes.DeleteBook)
	fmt.Println("starting web server at http://localhost:8080/")
	err := http.ListenAndServe(":8080", nil)
	utils.ErrorHandler(err)
}
