package db

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"time"

	"github.com/CommunityCharts/CCModels/school"
	"github.com/CommunityCharts/CCModels/shared"
	"github.com/CommunityCharts/CCModels/student"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

var Schools *mongo.Collection
var Students *mongo.Collection

var jwtSecret = []byte("secret")

type Claims struct {
	StudentID int `json:"student_id"`
	jwt.StandardClaims
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

	Schools = DB.Collection("schools")
	Students = DB.Collection("students")

	fmt.Println("Connected to MongoDB")
}

func GetSchoolByID(id int) (school.School, error) {
	var s school.School
	filter := shared.Object{
		"id": id,
	}

	err := Schools.FindOne(context.TODO(), filter).Decode(&s)
	if err != nil {
		return s, err
	}

	return s, nil
}

func GetStudentByID(id int) student.DBStudentUser {
	var dbStudentUser student.DBStudentUser
	filter := shared.Object{
		"studentuser.id": id,
	}

	err := Students.FindOne(context.TODO(), filter).Decode(&dbStudentUser)
	if err != nil {
		panic(err)
	}

	return dbStudentUser
}

func GetStudents() []student.DBStudentUser {
	var students []student.DBStudentUser

	cursor, err := Students.Find(context.TODO(), shared.Object{})
	if err != nil {
		panic(err)
	}

	err = cursor.All(context.TODO(), &students)
	if err != nil {
		panic(err)
	}

	return students
}

func CreateStudent(student student.DBStudentUser) int {
	_, err := DB.Collection("students").InsertOne(context.TODO(), student)
	if err != nil {
		panic(err)
	}

	return student.StudentUser.Id
}

func UpdateStudent(student student.DBStudentUser) {
	filter := shared.Object{
		"studentuser.id": student.StudentUser.Id,
	}

	_, err := Students.ReplaceOne(context.TODO(), filter, student)
	if err != nil {
		panic(err)
	}
}

func GetStudentJWTForLogin(student student.DBStudentUser) *string {
	expirationTime := time.Now().Add(3 * time.Minute)
	claims := &Claims{
		StudentID: student.StudentUser.Id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		panic(err)
	}

	return &tokenString
}

func GetNextID() int {
	var dbStudentUser student.DBStudentUser
	opts := options.FindOne().SetSort(shared.Object{"id": -1})
	err := Students.FindOne(context.TODO(), shared.Object{}, opts).Decode(&dbStudentUser)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return 1
		}
		panic(err)
	}

	return dbStudentUser.StudentUser.Id + 1
}

func CreateSchool(school school.School) {
	_, err := Schools.InsertOne(context.TODO(), school)
	if err != nil {
		panic(err)
	}
}
