package store

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoConf _
type MongoConf struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

var defaultMongoConf MongoConf

// Init _
func Init(conf MongoConf) {
	defaultMongoConf = conf
}

func getConn(do func(*mongo.Client)) {
	cfg := defaultMongoConf
	opts := options.Client().ApplyURI(fmt.Sprintf(
		"mongodb://%s:%s@%s:%d", cfg.Username, cfg.Password, cfg.Host, cfg.Port))
	client, err := mongo.NewClient(opts)
	if err != nil {
		fmt.Printf("Fail create mongo client: %v", err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		fmt.Printf("Fail connect to mongo: %v", err)
		return
	}
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			return
		}
	}()

	do(client)
}

// AddMapData _
func AddMapData(coll string, data map[string]interface{}) (added map[string]interface{}) {
	db := defaultMongoConf.Database
	getConn(func(client *mongo.Client) {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		c := client.Database(db).Collection(coll)
		res, err := c.InsertOne(ctx, bson.M(data))
		if err != nil {
			fmt.Printf("Fail to InsertOne: %v", err)
			return
		}
		data["_id"] = res.InsertedID
		added = data
	})

	return
}

// Find _
func Find(coll string, where map[string]interface{}) (data []interface{}) {
	return FindWithSort(coll, where)
}

// FindWithSort _
func FindWithSort(coll string, where map[string]interface{}, sortOpts ...bson.E) (data []interface{}) {
	db := defaultMongoConf.Database
	getConn(func(client *mongo.Client) {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		c := client.Database(db).Collection(coll)

		// set sort options
		var cursor *mongo.Cursor
		var err error
		if len(sortOpts) > 0 {
			opts := options.Find().SetSort(bson.D(sortOpts))
			cursor, err = c.Find(ctx, mapToD(where), opts)
		} else {
			cursor, err = c.Find(ctx, mapToD(where))
		}
		if err != nil {
			fmt.Printf("Fail to Find: %v", err)
			return
		}

		// fetch all data
		var results []bson.M
		if err = cursor.All(context.TODO(), &results); err != nil {
			fmt.Printf("Fail to call cursor.All: %v", err)
			return
		}
		for i := range results {
			data = append(data, results[i])
		}
	})

	return
}

func mapToD(m map[string]interface{}) (d bson.D) {
	d = make(primitive.D, 0, len(m))
	for k := range m {
		d = append(d, bson.E{k, m[k]})
	}
	return
}
