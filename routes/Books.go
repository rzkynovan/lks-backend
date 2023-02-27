package routes

import (
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"kanolibrary/controllers/books"
	"kanolibrary/models"
	"kanolibrary/utils"
	"net/http"
	"time"
)

//get all books
func FindAllBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		result, err := books.FindAll("books")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		//if string(result) == "[]" {
		//	//make response with json
		//	errMsg := map[string]interface{}{"Error": "Document not found"}
		//	errStr := utils.ErrorHandlerToJson(errMsg)
		//	http.Error(w, string(errStr), http.StatusNotFound)
		//	return
		//}

		w.Write(result)
		return
	}
	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}

func FindByIdBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		requestId := r.URL.Query().Get("id")
		objId, err := primitive.ObjectIDFromHex(requestId)
		result, err := books.FindOne("books", bson.M{"_id": objId})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Println(len(result))
		fmt.Println(string(result))
		//if result is empty array, then not found
		if string(result) == "[]" {
			//make response with json
			errMsg := map[string]interface{}{"Error": "Collection is not found"}
			errStr := utils.ErrorHandlerToJson(errMsg)
			http.Error(w, string(errStr), http.StatusNotFound)
			return
		}
		w.Write(result)
		return
	}
	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}

//insert book
func InsertBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		var book models.Books
		err := json.NewDecoder(r.Body).Decode(&book)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		result, err := books.Insert("books", book)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write([]byte(result))
		return
	}
	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}

//update book
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "PUT" {
		var book models.Books
		err := json.NewDecoder(r.Body).Decode(&book)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		requestId := r.URL.Query().Get("id")

		objId, err := primitive.ObjectIDFromHex(requestId)
		utils.ErrorHandler(err)
		result, err := books.UpdateOne("books", bson.M{"_id": objId}, bson.M{"$set": bson.M{"title": book.Title, "author": book.Author, "available": book.Available, "releaseDate": book.ReleaseDate, "synopsis": book.Synopsis, "updatedAt": primitive.NewDateTimeFromTime(time.Now())}})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write([]byte(result))
		return
	}
	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "DELETE" {
		requestId := r.URL.Query().Get("id")

		objId, err := primitive.ObjectIDFromHex(requestId)
		result, err := books.DeleteOne("books", bson.M{"_id": objId})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write([]byte(result))
		return

	}
	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}
