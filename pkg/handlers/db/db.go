package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "./db.sqlite")
	if err != nil {
		panic(err)
	}

	// Create a table if it doesn't exist
	createTableSQL := `CREATE TABLE IF NOT EXISTS students (
        "id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        "name" TEXT,
        "first_name" TEXT,
        "last_name" TEXT,
        "dob" TEXT,
    	"code" TEXT
    );`

	_, err = DB.Exec(createTableSQL)
	if err != nil {
		panic(err)
	}
}

func GetStudentByID(id int) (StudentDB, error) {
	var student StudentDB

	row := DB.QueryRow("SELECT * FROM students WHERE id = ?", id)
	err := row.Scan(&student.Id, &student.Name, &student.FirstName, &student.LastName, &student.DOB, &student.Code)
	if err != nil {
		return student, err
	}

	return student, nil
}

func GetStudents() []StudentDB {
	var students []StudentDB

	rows, err := DB.Query("SELECT * FROM students")
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var student StudentDB
		err = rows.Scan(&student.Id, &student.Name, &student.FirstName, &student.LastName, &student.DOB, &student.Code)
		if err != nil {
			panic(err)
		}

		students = append(students, student)
	}

	return students
}

func CreateStudent(student StudentDB) {
	_, err := DB.Exec("INSERT INTO students (name, first_name, last_name, dob, code) VALUES (?, ?, ?, ?, ?)", student.Name, student.FirstName, student.LastName, student.DOB, student.Code)
	if err != nil {
		panic(err)
	}
}

type StudentDB struct {
	Id        int
	Name      string
	FirstName string
	LastName  string

	DOB  string
	Code string
}
