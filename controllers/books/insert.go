package books

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"kanolibrary/models"
	"kanolibrary/mongodb"
	"kanolibrary/utils"
	"time"
)

//insert data to database
func Insert(collection string, query models.Books) (string, error) {
	ctx := context.Background()
	db, err := mongodb.Connect()
	utils.ErrorHandler(err)
	query.Id = primitive.NewObjectID()
	query.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
	query.UpdatedAt = primitive.NewDateTimeFromTime(time.Now())
	_, err = db.Collection(collection).InsertOne(ctx, query)
	utils.ErrorHandler(err)
	return "Insert success", err
}
