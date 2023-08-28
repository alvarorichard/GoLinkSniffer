package db

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func VistedLink(link string) bool {
	client, ctx := getConnection()
	defer client.Disconnect(ctx)

	c := client.Database("gocrawler").Collection("links")

	opts := options.Count().SetMaxTime(1)

	n, err := c.CountDocuments(
		context.TODO(),
		bson.D{{"link", link}},
		opts,
	)
	if err != nil {
		panic(err)
	}
	return n > 0

}
