package global

import (
	"context"
	"errors"
	"fmt"
	"github.com/daydreme/classcharts-server-mock/pkg/global/models/responses"
	"github.com/daydreme/classcharts-server-mock/pkg/student/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

var DB *mongo.Database
var Students *mongo.Collection

type StudentDB struct {
	Id        int
	Name      string
	FirstName string
	LastName  string

	DOB  string
	Code string
}

func InitDB() {
	fmt.Println("Connecting to MongoDB...")
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017").SetConnectTimeout(time.Second * 3)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		panic(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		panic(err)
	}

	DB = client.Database("classcharts")
	Students = DB.Collection("students")

	fmt.Println("Connected to MongoDB")
}

func GetStudentByID(id int) StudentDB {
	var student StudentDB
	filter := responses.Object{
		"id": id,
	}

	err := Students.FindOne(context.TODO(), filter).Decode(&student)
	if err != nil {
		panic(err)
	}

	return student
}

func GetStudents() []StudentDB {
	var students []StudentDB

	cursor, err := Students.Find(context.TODO(), responses.Object{})
	if err != nil {
		panic(err)
	}

	err = cursor.All(context.TODO(), &students)
	if err != nil {
		panic(err)
	}

	return students
}

func CreateStudent(student StudentDB) models.StudentUser {
	_, err := DB.Collection("students").InsertOne(context.TODO(), student)
	if err != nil {
		panic(err)
	}

	return student.ToStudentUser()
}

func GetNextID() int {
	var student StudentDB
	opts := options.FindOne().SetSort(responses.Object{"id": -1})
	err := Students.FindOne(context.TODO(), responses.Object{}, opts).Decode(&student)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return 1
		}
		panic(err)
	}

	return student.Id + 1
}

func (s StudentDB) ToStudentUser() models.StudentUser {
	user := models.NewUser()
	user.Id = s.Id
	user.Name = s.Name
	user.FirstName = s.FirstName
	user.LastName = s.LastName

	return user
}
