package mongorepo

import (
	"context"
	"log"
	"time"

	"github.com/moises-ba/ms-dynamic-qrcode/model/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	collection_qrcode_name = "qrcode"
)

type repo struct {
	client           *mongo.Client
	database         *mongo.Database
	collectionQrcode *mongo.Collection
}

func NewRepository(database *mongo.Database) Repository {
	return &repo{
		client:           database.Client(),
		database:         database,
		collectionQrcode: database.Collection(collection_qrcode_name),
	}
}

func (r *repo) Insert(qrcode *domain.QRCodeModel) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if result, err := r.collectionQrcode.InsertOne(ctx, qrcode); err == nil {
		log.Println(result)
		return nil
	} else {
		return err
	}

}

func (r *repo) FindQRCodes(pFilter *domain.QRCodeFilter) ([]*domain.QRCodeModel, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	filter := bson.D{}
	if pFilter.User != "" {
		filter = append(filter, primitive.E{Key: "user", Value: pFilter.User})
	}

	cur, err := r.collectionQrcode.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	qrcodes := make([]*domain.QRCodeModel, 0, 10)

	for cur.Next(ctx) {
		var result *domain.QRCodeModel
		err := cur.Decode(&result)
		if err != nil {
			return nil, err
		}

		qrcodes = append(qrcodes, result)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}

	return qrcodes, nil
}
