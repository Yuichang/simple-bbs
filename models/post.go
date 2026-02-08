package models

import (
	"context"
	"database/sql"
	"time"
)

type Post struct {
	ID               int
	Name             string
	Body             string
	CreatedAt        time.Time
	CreatedAtDisplay string
}

// dbからポスト一覧を取得する
func ListPosts(ctx context.Context, db *sql.DB) ([]Post, error) {

	// 昇順にポストを取り出す
	rows, err := db.QueryContext(ctx, "SELECT id,name,body,created_at FROM posts ORDER BY created_at ASC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	posts := []Post{}

	// dbから取得したデータを一行ずつpostsに追加してく
	for rows.Next() {
		var p Post
		if err := rows.Scan(&p.ID, &p.Name, &p.Body, &p.CreatedAt); err != nil {
			return nil, err
		}
		p.CreatedAtDisplay = p.CreatedAt.Format("2006-01-02 15:04:05")
		posts = append(posts, p)
	}
	// ループ中にエラーが発生していないか
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return posts, nil
}

func CreatePost(ctx context.Context, db *sql.DB, name, body string) error {
	_, err := db.ExecContext(ctx, "INSERT INTO posts (name, body, created_at) VALUES (?, ?, ?)", name, body, time.Now())
	return err
}

func DeletePost(ctx context.Context, db *sql.DB, id string) error {
	_, err := db.ExecContext(ctx, "DELETE FROM posts WHERE id = ?", id)
	return err
}

func CreateAccount(ctx context.Context, db *sql.DB, username, gender,hashedPassword string) error {

	_, err := db.ExecContext(ctx, "INSERT INTO users (name, gender,hashed_password,created_at) VALUES (?, ?, ?,?)", username, gender, hashedPassword,time.Now())
	return err
}