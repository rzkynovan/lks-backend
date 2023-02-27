package books

import (
	"context"
	"kanolibrary/mongodb"
	"kanolibrary/utils"
)

//update by id
func UpdateOne(collection string, query map[string]interface{}, update map[string]interface{}) (string, error) {

	ctx := context.Background()
	db, err := mongodb.Connect()
	utils.ErrorHandler(err)

	_, err = db.Collection(collection).UpdateOne(ctx, query, update)

	utils.ErrorHandler(err)

	return "Update success", err
}
