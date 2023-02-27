package books

import (
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"kanolibrary/models"
	"kanolibrary/mongodb"
	"kanolibrary/utils"
)

//read data from database
//find one data from database
func FindOne(collection string, query map[string]interface{}) ([]byte, error) {

	ctx := context.Background()
	db, err := mongodb.Connect()
	utils.ErrorHandler(err)

	csr, err := db.Collection(collection).Find(ctx, query)
	utils.ErrorHandler(err)
	defer csr.Close(ctx)

	result := make([]models.Books, 0)
	for csr.Next(ctx) {
		var row models.Books
		err := csr.Decode(&row)
		utils.ErrorHandler(err)
		result = append(result, row)
	}
	data, err := json.Marshal(result)
	utils.ErrorHandler(err)
	return data, err
}

//FindMany
func FindMany(collection string, query map[string]interface{}) ([]byte, error) {

	ctx := context.Background()
	db, err := mongodb.Connect()
	utils.ErrorHandler(err)

	csr, err := db.Collection(collection).Find(ctx, query)
	utils.ErrorHandler(err)
	defer csr.Close(ctx)

	result := make([]models.Books, 0)
	for csr.Next(ctx) {
		var row models.Books
		err := csr.Decode(&row)
		utils.ErrorHandler(err)

		result = append(result, row)

	}
	//iterate over the result with title and author
	data, err := json.Marshal(result)
	utils.ErrorHandler(err)
	return data, err

}

func FindAll(collection string) ([]byte, error) {

	ctx := context.Background()
	db, err := mongodb.Connect()
	utils.ErrorHandler(err)

	csr, err := db.Collection(collection).Find(ctx, bson.M{})
	utils.ErrorHandler(err)
	defer csr.Close(ctx)

	result := make([]models.Books, 0)
	for csr.Next(ctx) {
		var row models.Books
		err := csr.Decode(&row)
		utils.ErrorHandler(err)

		result = append(result, row)

	}
	data, err := json.Marshal(result)
	utils.ErrorHandler(err)
	return data, err
}
