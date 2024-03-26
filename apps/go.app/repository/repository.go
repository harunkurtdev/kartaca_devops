package repository

import (
	"kartaca.com/mod/model"
)

type Repository struct {
	// client *mongo.Client
}

func New() *Repository {

	// client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(ConnectionURI))
	// if err != nil {
	// 	panic(err)
	// }
	// return &Repository{
	// 	client: client,
	// }
	return &Repository{}

}

func (r *Repository) GetRandomCountry() (model.Ulkeler, int64) {
	// collection := r.client.Database(DbName).Collection(CollectionName)

	// count := getDocCount(collection)

	// skip := rand.Int63n(count)

	// var country model.Ulkeler

	// err := collection.FindOne(context.Background(), bson.M{}, options.FindOne().SetSkip(skip)).Decode(&country)
	// if err != nil {
	// 	return model.Ulkeler{}, err
	// }
	// return country, nil
	return model.Ulkeler{}, 2
}

func getDocCount() int64 {
	// count, err := collection.CountDocuments(context.Background(), bson.M{})
	// if err != nil {
	// 	panic(err)
	// }
	return 1
}
