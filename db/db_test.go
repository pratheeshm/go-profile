package db

import (
	_ "github.com/lib/pq"
	"testing"
)

func BenchmarkRowInsert(b *testing.B) {
	db, err := SQLConnection()
	if err != nil {
		b.Error(err)
	}
	student := &Student{Name: "John", Age: 21}
	for i := 0; i < b.N; i++ {
		err := RowInsert(db, student)
		if err != nil {
			b.Fatal(err)
		}
	}
}
func BenchmarkGormInsert1(b *testing.B) {
	db, err := GormConnection()
	if err != nil {
		b.Error(err)
	}
	student := &Student{Name: "John", Age: 21}
	for i := 0; i < b.N; i++ {
		err := GormInsert(db, student)
		if err != nil {
			b.Fatal(err)
		}
	}
}
func BenchmarkGormInsertUsingCreate(b *testing.B) {
	db, err := GormConnection()
	if err != nil {
		b.Error(err)
	}
	student := &Student{Name: "John", Age: 21}
	for i := 0; i < b.N; i++ {
		err := gormInsertUsingCreate(db, student)
		if err != nil {
			b.Fatal(err)
		}
	}
}
func BenchmarkRowsSelect(b *testing.B) {
	db, err := SQLConnection()
	if err != nil {
		b.Error(err)
	}
	for i := 0; i < b.N; i++ {
		result, err := rowsSelect(db, 100)
		if err != nil {
			b.Fatal(err)
		}
		if len(result) != 100 {
			b.Fatalf("lenght of result is not 100, but got %v", len(result))
		}
	}
}
func BenchmarkGormSelectRows(b *testing.B) {
	db, err := GormConnection()
	if err != nil {
		b.Error(err)
	}
	for i := 0; i < b.N; i++ {
		result, err := gormSelectRows(db, 100)
		if err != nil {
			b.Fatal(err)
		}
		if len(result) != 100 {
			b.Fatalf("lenght of result is not 100, but got %v", len(result))
		}
	}
}
func BenchmarkSingleRowQuery(b *testing.B) {
	db, err := SQLConnection()
	if err != nil {
		b.Error(err)
	}
	for i := 0; i < b.N; i++ {
		stud, err := singleRowQuery(db, 100)
		if err != nil {
			b.Fatal(err)
		}
		if stud.ID != 100 {
			b.Fatal("Id is not 100")
		}
	}
}
func BenchmarkGormSingleRow(b *testing.B) {
	db, err := GormConnection()
	if err != nil {
		b.Error(err)
	}
	for i := 0; i < b.N; i++ {
		stud, err := gormSingleRow(db, 100)
		if err != nil {
			b.Fatal(err)
		}
		if stud.ID != 100 {
			b.Fatal("Id is not 100")
		}
	}
}
