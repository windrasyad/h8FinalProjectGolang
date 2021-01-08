package home

import (
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

var (
	// table is the table name.
	table = "article"
)

// Article defines the model.
type Article struct {
	ArticleID uint32         `db:"article_id"`
	Tittle    string         `db:"tittle"`
	Issian    string         `db:"issian"`
	UserID    uint32         `db:"user_id"`
	CreatedAt mysql.NullTime `db:"created_at"`
	UpdatedAt mysql.NullTime `db:"updated_at"`
	DeletedAt mysql.NullTime `db:"deleted_at"`
	Publish   uint32         `db:"publish"`
}

// Connection is an interface for making queries.
type Connection interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Get(dest interface{}, query string, args ...interface{}) error
	Select(dest interface{}, query string, args ...interface{}) error
}

// ByID gets an Article by ArticleID.
func ByID(db Connection, ArticleID string, userID string) (Article, bool, error) {
	result := Article{}
	err := db.Get(&result, fmt.Sprintf(`
		SELECT article_id, tittle, issian, user_id, created_at, updated_at, deleted_at, publish
		FROM %v
		WHERE article_id = ?
			AND deleted_at IS NULL
		LIMIT 1
		`, table),
		ArticleID)
	return result, err == sql.ErrNoRows, err
}

// ByUserID gets all Articles for a user.
func ByUserID(db Connection, userID string) ([]Article, bool, error) {
	var result []Article
	err := db.Select(&result, fmt.Sprintf(`
		SELECT article_id, tittle, issian, user_id, created_at, updated_at, deleted_at, publish
		FROM %v
		WHERE user_id = ?
			AND deleted_at IS NULL
		`, table),
		userID)
	return result, err == sql.ErrNoRows, err
}

// ByUserIDPaginate gets Articles for a user based on page and max variables.
func ByUserIDPaginate(db Connection, userID string, max int, page int) ([]Article, bool, error) {
	var result []Article
	err := db.Select(&result, fmt.Sprintf(`
		SELECT article_id, tittle, issian, user_id, created_at, updated_at, deleted_at, publish
		FROM %v
		WHERE publish = ?
			AND deleted_at IS NULL
		LIMIT %v OFFSET %v
		`, table, max, page),
		1)
	return result, err == sql.ErrNoRows, err
}

// ByUserIDCount counts the number of Articles for a user.
func ByUserIDCount(db Connection, userID string) (int, error) {
	var result int
	err := db.Get(&result, fmt.Sprintf(`
		SELECT count(*)
		FROM %v
		WHERE publish = ?
			AND deleted_at IS NULL
		`, table),
		1)
	return result, err
}
