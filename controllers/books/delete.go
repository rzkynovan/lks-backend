package books

import (
	"context"
	"kanolibrary/mongodb"
	"kanolibrary/utils"
)

//delete data from database collection books

func DeleteOne(collection string, query map[string]interface{}) (string, error) {
	ctx := context.Background()
	db, err := mongodb.Connect()
	utils.ErrorHandler(err)

	_, err = db.Collection(collection).DeleteOne(ctx, query)

	utils.ErrorHandler(err)

	return "Delete success", err
}
