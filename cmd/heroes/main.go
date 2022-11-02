package main

import (
	"context"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jirawat-rackz/golang-gin-tour-of-heroes/feature/heroes"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	credential := options.Credential{
		Username: "root",
		Password: "password",
	}

	// Connect Mongo with Credential and URI
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://mongodb:27017/").SetAuth(credential))
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(err)
	}
	// Connect DB
	mongodb := client.Database("tour-of-heroes")

	router := gin.Default()

	heroService := heroes.HeroHandler{}
	heroRepository := heroes.HeroRepository{
		DBconn: mongodb,
	}

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET"},
		AllowHeaders: []string{"Content-Type,access-control-allow-origin, access-control-allow-headers"},
	}))

	router.GET("/heroes", heroService.GetAllHero(heroRepository))

	router.Run(":8080")
}
