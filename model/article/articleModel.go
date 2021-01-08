package article

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
			AND user_id = ?
			AND deleted_at IS NULL
		LIMIT 1
		`, table),
		ArticleID, userID)
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
		WHERE user_id = ?
			AND deleted_at IS NULL
		LIMIT %v OFFSET %v
		`, table, max, page),
		userID)
	return result, err == sql.ErrNoRows, err
}

// ByUserIDCount counts the number of Articles for a user.
func ByUserIDCount(db Connection, userID string) (int, error) {
	var result int
	err := db.Get(&result, fmt.Sprintf(`
		SELECT count(*)
		FROM %v
		WHERE user_id = ?
			AND deleted_at IS NULL
		`, table),
		userID)
	return result, err
}

// Create adds an Article.
func Create(db Connection, tittle string, issian string, userID string) (sql.Result, error) {
	result, err := db.Exec(fmt.Sprintf(`
		INSERT INTO %v
		(tittle, issian, user_id, publish)
		VALUES
		(?,?,?)
		`, table),
		tittle, issian, userID, 1)
	return result, err
}

// Update makes changes to an existing Article.
func Update(db Connection, tittle string, issian string, articleID string, userID string) (sql.Result, error) {
	result, err := db.Exec(fmt.Sprintf(`
		UPDATE %v
		SET tittle = ?, issian = ?
		WHERE article_id = ?
			AND user_id = ?
			AND deleted_at IS NULL
		LIMIT 1
		`, table),
		tittle, issian, articleID, userID)
	return result, err
}

// DeleteHard removes an Article.
func DeleteHard(db Connection, articleID string, userID string) (sql.Result, error) {
	result, err := db.Exec(fmt.Sprintf(`
		DELETE FROM %v
		WHERE article_id = ?
			AND user_id = ?
			AND deleted_at IS NULL
		`, table),
		articleID, userID)
	return result, err
}

// DeleteSoft marks an Article as removed.
func DeleteSoft(db Connection, articleID string, userID string) (sql.Result, error) {
	result, err := db.Exec(fmt.Sprintf(`
		UPDATE %v
		SET deleted_at = NOW()
		WHERE article_id = ?
			AND user_id = ?
			AND deleted_at IS NULL
		LIMIT 1
		`, table),
		articleID, userID)
	return result, err
}
