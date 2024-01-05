package main


import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"
	"context"
	"os"
	"os/signal"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/thedevsaddam/renderer"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var rnd *renderer.Render
var db *mgo.Database

const (
	hostName		string = "localhost:27017"
	dbName			string = "todo_demo"
	collectionName	string = "todos"
	port			string = ":9000"
)

type(
	todoModel struct {
		ID			bson.ObjectId	`bson:"_id, omitempty"`
		Title		string			`bson:"title"`
		Completed	bool			`bson:"completed"`
		CreatedAt	time.Time		`bson:"created_at"`
	}
	todo struct {
		ID			string `json:"id"`
		Title		string `json:"title"`
		Completed	string `json:"completed"`
		CreatedAt	string `json:"created_at"`
 	}
)


func init() {
	rnd = renderer.New()
	sess, err := mgo.Dial(hostName)
	checkErr(err)
	db = sess.DB(dbName)
}
