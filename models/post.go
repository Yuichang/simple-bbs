package models

import (
	"context"
	"database/sql"
	"time"
)

type Post struct {
	ID        int
	Name      string
	Body      string
	CreatedAt time.Time
}

func ListPosts(ctx context.Context, db *sql.DB) ([]Post, error) {

	// 昇順にポストを取り出す
	rows, err := db.QueryContext(ctx, "SELECT id,name,body,created_at FROM post ORDER BY created_at ASC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	posts := []Post{}

	// sqlから取得したデータを一行ずつpostsに追加してく
	for rows.Next() {
		var p Post
		if err := rows.Scan(&p.ID, &p.Name, &p.Body, &p.CreatedAt); err != nil {
			return nil, err
		}
		posts = append(posts, p)
	}
	// ループ中にエラーが発生していないか
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return posts, nil
}
