package db

import(
	"context"
	"log"
)

func InsertLink(collection string, data interface{}) error{
	client, ctx := getConnection()
	defer client.Disconnect(ctx)
	c := client.Database("crawler").Collection(collection)
	_, err := c.InsertOne(context.Background(), data)
	log.Println("Inserido no banco de dados")
	return err
}