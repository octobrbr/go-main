package mongo

import (
	"GoNews/pkg/storage"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	databaseName   = "news"
	collectionPost = "posts"
)

type Store struct {
	db *mongo.Client
}

// Конструктор объекта хранилища.
func New(constr string) (*Store, error) {
	mongoOpts := options.Client().ApplyURI(constr)

	client, err := mongo.Connect(context.Background(), mongoOpts)

	if err != nil {
		log.Fatal(err)
	}

	s := Store{
		db: client,
	}
	return &s, nil
}

// Отключение от БД
func (s *Store) Disconnect() {
	s.db.Disconnect(context.Background())
}

// Получение списка всех статей из БД,
func (s *Store) Posts() ([]storage.Post, error) {
	collection := s.db.Database(databaseName).Collection(collectionPost)
	filter := bson.D{}
	cur, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())
	var posts []storage.Post
	for cur.Next(context.Background()) {
		var l storage.Post
		err := cur.Decode(&l)
		if err != nil {
			return nil, err
		}
		posts = append(posts, l)
	}

	return posts, cur.Err()
}

// Добавление статьи в БД
func (s *Store) AddPost(post storage.Post) error {
	collection := s.db.Database(databaseName).Collection(collectionPost)

	_, err := collection.InsertOne(context.Background(), post)
	if err != nil {
		return err
	}

	return nil
}

// Обновление статьи в БД
func (s *Store) UpdatePost(post storage.Post) error {
	collection := s.db.Database(databaseName).Collection(collectionPost)

	_, err := collection.UpdateOne(
		context.Background(),
		bson.M{"id": post.ID},
		bson.D{
			{Key: "$set", Value: bson.D{
				{Key: "title", Value: post.Title},
				{Key: "content", Value: post.Content},
				{Key: "authorid", Value: post.AuthorID},
				{Key: "authorname", Value: post.AuthorName},
				{Key: "createdat", Value: post.CreatedAt},
				{Key: "publishedat", Value: post.PublishedAt}}},
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Print(err)
	}
	return nil
}

// Удаление статьи из БД
func (s *Store) DeletePost(post storage.Post) error {
	collection := s.db.Database(databaseName).Collection(collectionPost)

	_, err := collection.DeleteOne(context.Background(), post)
	if err != nil {
		return err
	}

	return nil
}
