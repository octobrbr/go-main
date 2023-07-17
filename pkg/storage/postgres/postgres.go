package postgres

import (
	"GoNews/pkg/storage"
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

// Хранилище данных.
type Store struct {
	db *pgxpool.Pool
}

// Конструктор объекта хранилища.
func New(constr string) (*Store, error) {
	db, err := pgxpool.Connect(context.Background(), constr)
	if err != nil {
		return nil, err
	}

	s := Store{
		db: db,
	}

	return &s, nil
}

// Получение списка всех статей из БД,
func (s *Store) Posts() ([]storage.Post, error) {
	rows, err := s.db.Query(context.Background(), `
		SELECT
			posts.id, 
			posts.author_id, 
			posts.title, 
			posts.content, 
			posts.created_at,			
			authors.name
		FROM posts 
		JOIN authors ON posts.author_id = authors.id
		ORDER BY id;
	`)
	if err != nil {
		return nil, err
	}
	var posts []storage.Post

	for rows.Next() {
		var t storage.Post
		err = rows.Scan(
			&t.ID,
			&t.AuthorID,
			&t.Title,
			&t.Content,
			&t.CreatedAt,
			&t.AuthorName,
		)
		if err != nil {
			return nil, err
		}

		posts = append(posts, t)
	}

	return posts, rows.Err()
}

// Добавление статьи в БД
func (s *Store) AddPost(post storage.Post) error {
	var id int
	err := s.db.QueryRow(context.Background(), `
		INSERT INTO posts (author_id, title, content, created_at)
		VALUES ($1, $2, $3, $4) RETURNING id;
		`,
		post.AuthorID,
		post.Title,
		post.Content,
		post.CreatedAt,
	).Scan(&id)
	return err
}

// Обновление статьи в БД
func (s *Store) UpdatePost(post storage.Post) error {

	ct, err := s.db.Exec(context.Background(), `
	UPDATE posts 
	SET author_id = $2,
	title = $3,
	content = $4,
	created_at = $5
	WHERE id = $1;`,
		post.ID, post.AuthorID, post.Title, post.Content, post.CreatedAt)

	if err != nil {
		return err
	}

	if ct.RowsAffected() == 0 {
		return fmt.Errorf("404")
	}

	return nil
}

// Удаление статьи из БД
func (s *Store) DeletePost(post storage.Post) error {

	ct, err := s.db.Exec(context.Background(), `
	DELETE FROM posts WHERE (id = $1);`,
		post.ID)
	if err != nil {
		return err
	}

	if ct.RowsAffected() == 0 {
		return fmt.Errorf("404")
	}

	return nil
}
