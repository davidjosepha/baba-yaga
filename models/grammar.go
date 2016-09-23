package models

import (
	"fmt"
	"os"
)

type Grammar struct {
	id       int
	lessonID int
	title    string
	subtitle string
	text     string
}

const gramtab = "grammars"

func (g *Grammar) Lesson() *Lesson {
	// TODO
	return nil
}

func (g *Grammar) SetLesson(l *Lesson) {
	// TODO
	// sets g.lessonID to l.id
	// if l not in db, insert (?)
}

func (g *Grammar) String() string {
	return fmt.Sprintf("{ Grammar ID: %d, Lesson ID: %d }", g.id, g.Lesson().id)
}

func (g *Grammar) Select(id int) {
	q := fmt.Sprintf("SELECT id,lesson_id,title,subtitle,text FROM %s WHERE id = $1", gramtab)
	err := db.QueryRow(q, id).Scan(&g.id, &g.lessonID, &g.title, &g.subtitle, &g.text)
	if err != nil {
		fmt.Fprintf(os.Stderr, "babayaga: %v\n", err)
	}
}

func (g *Grammar) Insert() {
	q := fmt.Sprintf(
		"INSERT INTO %s (lesson_id,title,subtitle,text) VALUES ($1,$2,$3,$4)",
		gramtab)
	res, err := db.Exec(q, g.lessonID, g.title, g.subtitle, g.text)
	if err != nil {
		fmt.Fprintf(os.Stderr, "babayaga: %v\n", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		fmt.Fprintf(os.Stderr, "babayaga: %v\n", err)
	}
	g.id = int(id)
}

func (g *Grammar) Update() {
	// TODO
}

func (g *Grammar) Delete() {
	// TODO
}

func mkGramTab() {
	q := fmt.Sprintf(
		"CREATE TABLE IF NOT EXISTS %s (id integer primary key autoincrement, lesson_id integer references %s(id), title varchar(256), subtitle varchar(256), text text);",
		gramtab, lsntab)
	if _, err := db.Exec(q); err != nil {
		fmt.Fprintf(os.Stderr, "babayaga: %v\n", err)
	}
}
