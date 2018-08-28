package main

//go:generate goplater model --type User  --suffix _q --tmpl ./q.tmpl

type User struct {
	Name     string `db:"name" json:"name"`
	Age      int    `db:"age" json:"age"`
	FullName string `db:"full_name" json:"full_name"`
}
