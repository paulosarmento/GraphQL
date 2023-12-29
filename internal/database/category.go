package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Category struct {
	db          *sql.DB
	ID          string
	Name        string
	Description string
}

func NewCategory(db *sql.DB) *Category {
	return &Category{db: db}
}

func (c *Category) Create(name string, description string) (Category, error) {
	id := uuid.New().String()
	_, err := c.db.Exec("INSERT INTO categories (id, name, description) VALUES ($1, $2, $3)", id, name, description)

	if err != nil {
		return Category{}, err
	}
	return Category{ID: id, Name: name, Description: description}, nil

}
func (c *Category) FindAll() ([]Category, error) {
	rows, err := c.db.Query("SELECT id, name, description FROM categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	categories := []Category{}
	for rows.Next() {
		var id, name, description string
		if err := rows.Scan(&id, &name, &description); err != nil {
			return nil, err
		}
		categories = append(categories, Category{ID: id, Name: name, Description: description})
	}
	return categories, nil
}

// graphql
// mutation createCategory {
//   createCategory(input: {name: "Tecnologioa", description: "Cursos de Tecnologia"}){
//     id
//     name
//     description
//   }
// }

// mutation createCourse {
//   createCourse(input: {name: "FullCycle", description: "THE BEST!", categoryId:"5939f8fc-1ed0-4030-924e-513f525d6a30"}){
//     id
//     name
//   }
// }

// query queryCategories {
//   categories {
//     id
//     name
//     description
//   }
// }
// query querycourses {
//   courses {
//     id
//     name
//   }
// }
//
// tabelas
// sqlite3 data.db
// create table categories (id string, name string, description string);
// create table courses (id string, string, name string, description string, category_id string);
