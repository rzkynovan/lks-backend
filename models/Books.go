package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Books struct {
	Id          primitive.ObjectID `bson:"_id"`
	Title       string             `bson:"title" validate:"required"`
	Author      string             `bson:"author" validate:"required"`
	ReleaseDate string             `bson:"releaseDate" validate:"required"`
	Synopsis    string             `bson:"synopsis" validate:"required"`
	Available   bool               `bson:"available"`
	CreatedAt   primitive.DateTime `bson:"createdAt"`
	UpdatedAt   primitive.DateTime `bson:"updatedAt"`
}
