package main

import (
	"country_capital3/src/controller"
	"country_capital3/src/storage"
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4/middleware"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	DB *mongo.Client
)

const (
	host     = "192.168.187.129"
	port     = 5432
	user     = "postgres"
	password = "number1admin!"
	name     = "postgres"
)

type Capitals struct {
	id      int       `json:"id"`
	capital string    `json:"capital"`
	addtime time.Time `json:"addtime"`
}

func getPostgresConnect() (*sql.DB, error) {
	dbinfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, name)
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	fmt.Println("postgres open")

	var id int
	var country string
	var insertTime string
	rows, err := db.Query("SELECT * FROM test1")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	postgresList := make(map[int]string)
	for rows.Next() {
		err := rows.Scan(&id, &country, &insertTime)
		if err != nil {
			panic(err)
		}
		postgresList[id] = strconv.Itoa(id) + " " + country + " " + insertTime
		fmt.Println(id, country, insertTime)
	}

	return db, err
}

// func getMongoConnect() (*mongo.Client, error) {
// 	clientOptions := options.Client().ApplyURI("mongodb://192.168.187.129:27017")
// 	client, err := mongo.Connect(context.TODO(), clientOptions)
// 	if err != nil {
// 		panic(err)
// 	}
// 	collection := client.Database("testDB").Collection("test1")

// 	cur, err := collection.Find(context.TODO(), bson.D{{}})
// 	if err != nil {
// 		panic(err)
// 	}
// 	for cur.Next(context.TODO()) {
// 		var elem bson.M
// 		err := cur.Decode(&elem)
// 		if err != nil {
// 			panic(err)
// 		}
// 		fmt.Println(&elem)
// 	}

// 	cur.Close(context.TODO())

// 	DB = client
// 	return DB, err
// }

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func main() {
	e := echo.New()
	storage.NewDB()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/", controller.GetCountry)
	e.GET("/hello", hello)
	e.Logger.Fatal(e.Start(":924"))

	// mgdb, err := getMongoConnect()
	// if err != nil {
	// 	fmt.Println("Mongo Error")
	// 	fmt.Println(err)
	// }
	// fmt.Println(mgdb)

	post, _ := getPostgresConnect()
	// if err != nil {
	// 	fmt.Println("Postgres Error")
	// 	fmt.Println(err)
	// }
	fmt.Println(post)

	// e.POST("/", func(c echo.Context) error {
	// 	ca := new(Capitals)
	// 	if err := c.Bind(ca); err != nil {
	// 		return err
	// 	}
	// 	return c.JSON(http.StatusCreated, ca)
	// })
}
