package unittest

import (
	"context"
	"fmt"
	"mysql_go/database"
	"mysql_go/entity"
	"mysql_go/repository"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestInsert(t *testing.T) {

	db := repository.NewCommentRepoImpl(database.GetConnection())

	ctx := context.Background()

	comments := entity.Comment{
		Email:   "arifsama02@gmail.com",
		Comment: "Test Comment",
	}

	comment, err := db.Insert(ctx, comments)
	if err != nil {
		panic(err)
	}

	fmt.Println(comment)
}

func TestFindByID(t *testing.T) {

	db := repository.NewCommentRepoImpl(database.GetConnection())

	comment, err := db.FindById(context.Background(), 2)

	if err != nil {
		panic(err)
	}

	fmt.Println(comment)
}

func TestFindAll(t *testing.T) {

	db := repository.NewCommentRepoImpl(database.GetConnection())

	comment, err := db.FindAll(context.Background())

	if err != nil {
		panic(err)
	}

	fmt.Println(comment)
	fmt.Println(comment[0])
	fmt.Println(comment[1])
}
