package db

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Name     string
	AuthorID int
	Author   Author
	Tags     []Tag `gorm:"many2many:post_tags"`
}
type Tag struct {
	gorm.Model
	Name string
}

type Author struct {
	ID   int
	Name string
}
