package db

import (
	"database/sql"
	"github.com/jinzhu/gorm"
)

//Student model
type Student struct {
	Name string `gorm:"column:name;type:varchar(20);not null"`
	Age  int    `gorm:"column:age;type:int;not null"`
	ID   int    `gorm:"column:id_student;primary_key"`
}

//TableName return table name for gorm
func (*Student) TableName() string {
	return "student"
}

//GormConnection returns gorm connection
func GormConnection() (*gorm.DB, error) {
	db, err := gorm.Open("postgres", "host=localhost user=postgres password=password port=5432 sslmode=disable")
	return db, err

}

//SQLConnection returns sql connection
func SQLConnection() (*sql.DB, error) {
	db, err := sql.Open("postgres", "host=localhost user=postgres password=password port=5432 sslmode=disable")
	return db, err
}

//RowInsert insert data using database/sql
func RowInsert(db *sql.DB, student *Student) error {
	_, err := db.Exec("INSERT INTO student(name, age) values($1, $2)", student.Name, student.Age)
	return err
}

//GormInsert insert data using gorm
func GormInsert(db *gorm.DB, student *Student) error {
	return db.Exec("INSERT INTO student(name, age) values(?, ?)", student.Name, student.Age).Error
}
func gormInsertUsingCreate(db *gorm.DB, student *Student) error {
	return db.Create(&student).Error
}
func rowsSelect(db *sql.DB, limit int) ([]Student, error) {
	rows, err := db.Query("SELECT * FROM student limit $1", limit)
	if err != nil {
		return nil, err
	}
	students := make([]Student, 0, limit)
	defer rows.Close()
	var student Student
	for rows.Next() {
		err = rows.Scan(&student.Name, &student.Age, &student.ID)
		if err != nil {
			return nil, err
		}
		students = append(students, student)
	}
	return students, nil
}

func gormSelectRows(db *gorm.DB, limit int) ([]Student, error) {
	students := make([]Student, 0, limit)
	err := db.Limit(limit).Find(&students).Error
	return students, err
}
func singleRowQuery(db *sql.DB, id int) (*Student, error) {
	var student Student
	err := db.QueryRow("SELECT * FROM student where id_student=$1", id).
		Scan(&student.Name, &student.Age, &student.ID)
	if err != nil {
		return nil, err
	}
	return &student, nil
}
func gormSingleRow(db *gorm.DB, id int) (*Student, error) {
	var student Student
	err := db.Where("id_student=?", id).Find(&student).Error
	if err != nil {
		return nil, err
	}
	return &student, nil
}
