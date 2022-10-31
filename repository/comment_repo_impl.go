package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"mysql_go/entity"
	"strconv"
)

type CommentRepoImpl struct {
	DB *sql.DB
}

func NewCommentRepoImpl(db *sql.DB) CommentRepo {
	return &CommentRepoImpl{DB: db}
}

func (repo *CommentRepoImpl) Insert(ctx context.Context, comment entity.Comment) (entity.Comment, error) {

	query := "Insert Into comments(email,comment)Values(?,?)"

	result, err := repo.DB.ExecContext(ctx, query, comment.Email, comment.Comment)

	if err != nil {
		return comment, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return comment, err
	}

	comment.Id = int32(id)

	fmt.Println("Sukses Insert Comment Id :", id)

	return comment, nil

}

func (repo *CommentRepoImpl) FindById(ctx context.Context, id int32) (entity.Comment, error) {

	query := "select id,email,comment from comments where id =?"

	rows, err := repo.DB.QueryContext(ctx, query, id)

	defer rows.Close()

	comment := entity.Comment{}

	if err != nil {
		return comment, err
	}

	if rows.Next() {
		//ada
		rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		return comment, nil
	}
	//tidak ada
	return comment, errors.New("Id" + strconv.Itoa(int(id)) + "Not Found")

}

func (repo *CommentRepoImpl) FindAll(ctx context.Context) ([]entity.Comment, error) {

	query := "select id,email,comment from comments"

	rows, err := repo.DB.QueryContext(ctx, query)

	defer rows.Close()
	if err != nil {
		return nil, err
	}

	comments := []entity.Comment{}
	for rows.Next() {
		comment := entity.Comment{}
		rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		comments = append(comments, comment)
	}
	return comments, nil
}
