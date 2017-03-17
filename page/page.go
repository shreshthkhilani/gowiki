package page

import (
    "database/sql"
    _ "github.com/lib/pq"
    "log"
)

var db *sql.DB = nil

type Page struct {
    ID int
    Title string
    Body string
}

func (p *Page) isArticle() (bool, error) {
    var id int
    query := `SELECT id FROM articles WHERE title=$1`
    err := db.QueryRow(query, p.Title).Scan(&id)
    switch {
        case err == sql.ErrNoRows:
            return false, nil
        case err != nil:
            return false, err
        default:
            return true, nil
    }
}

func (p *Page) insertArticle() error {
    query := `INSERT INTO articles(title, body) VALUES($1, $2)`
    stmt, err := db.Prepare(query)
    if err != nil {
        return err
    }
    defer stmt.Close()
    _, err = stmt.Exec(p.Title, p.Body)
    if err != nil {
        return err
    }
    return nil
}

func (p *Page) updateArticle() error {
    query := `UPDATE articles SET body=$1 WHERE title=$2`
    stmt, err := db.Prepare(query)
    if err != nil {
        return err
    }
    defer stmt.Close()
    _, err = stmt.Exec(p.Body, p.Title)
    if err != nil {
        return err
    }
    return nil
}

func (p *Page) Save() error {
    exists, err := p.isArticle()
    if err != nil {
        return err
    }
    if exists {
        return p.updateArticle()
    } else {
        return p.insertArticle()
    }
}

func LoadPage(title string) (*Page, error) {
    var id int
    var body string
    query := `SELECT id, body FROM articles WHERE title=$1`
    err := db.QueryRow(query, title).Scan(&id, &body)
    switch {
        case err != nil:
            return nil, err
        default:
            return &Page{ID: id, Title: title, Body: body}, nil
    }
}

func init() {
    var err error
    db, err = sql.Open("postgres", "postgres://shreshth:abc123@localhost/shreshth?sslmode=disable")
    if err != nil {
        log.Fatal("Postgres: Unable to open DB.", err)
    }
}