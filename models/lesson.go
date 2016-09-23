package models

import (
	"fmt"
	"os"
)

type Lesson struct {
	id        int
	chapterID int
}

const lsntab = "lessons"

func (l *Lesson) Chapter() *Chapter {
	// TODO
	return nil
}

func (l *Lesson) SetChapter(c *Chapter) {
	// TODO
	// sets l.chapterID to c.id
	// if c not in db, insert (?)
}

func (l *Lesson) Grammars() []*Grammar {
	// TODO
	return nil
}

func (l *Lesson) String() string {
	return fmt.Sprintf("{ Lesson ID: %d, Chapter ID: %d, # Grammars: %d }",
		l.id, l.chapterID, len(l.Grammars()))
}

func (l *Lesson) Select(id int) {
	q := fmt.Sprintf("SELECT id,chapter_id FROM %s WHERE id = $1", lsntab)
	err := db.QueryRow(q, id).Scan(&l.id, &l.chapterID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "babayaga: %v\n", err)
	}
}

func (l *Lesson) Insert() {
	q := fmt.Sprintf(
		"INSERT INTO %s (chapter_id) VALUES ($1)",
		lsntab)
	res, err := db.Exec(q, l.chapterID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "babayaga: %v\n", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		fmt.Fprintf(os.Stderr, "babayaga: %v\n", err)
	}
	l.id = int(id)
}

func (l *Lesson) Update() {
	// TODO
}

func (l *Lesson) Delete() {
	// TODO
}

func mkLsnTab() {
	q := fmt.Sprintf(
		"CREATE TABLE IF NOT EXISTS %s (id integer primary key autoincrement, chapter_id integer references %s(id));",
		lsntab, chtab)
	if _, err := db.Exec(q); err != nil {
		fmt.Fprintf(os.Stderr, "babayaga: %v\n", err)
	}
}
