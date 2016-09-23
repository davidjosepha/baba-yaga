package models

import (
	"fmt"
	"os"
)

// TODO: typedefs for ChapterID, LessonID, GrammarID

type Chapter struct {
	id    int
	title string
}

const chtab = "chapters"

func (c *Chapter) Lessons() []*Lesson {
	// TODO
	return nil
}

func (c *Chapter) String() string {
	return fmt.Sprintf("{ Chapter ID: %d, # Lessons: %d }", c.id, len(c.Lessons()))
}

func (c *Chapter) Select(id int) {
	q := fmt.Sprintf("SELECT id,title FROM %s WHERE id = $1;", chtab)
	err := db.QueryRow(q, id).Scan(&c.id, &c.title)
	if err != nil {
		fmt.Fprintf(os.Stderr, "babayaga: %v\n", err)
	}
}

func (c *Chapter) Insert() {
	q := fmt.Sprintf(
		"INSERT INTO %s (title) VALUES ($1)",
		chtab)
	res, err := db.Exec(q, c.title)
	if err != nil {
		fmt.Fprintf(os.Stderr, "babayaga: %v\n", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		fmt.Fprintf(os.Stderr, "babayaga: %v\n", err)
	}
	c.id = int(id)
}

func (c *Chapter) Update() {
	// TODO
}

func (c *Chapter) Delete() {
	// TODO
}

func mkChTab() {
	q := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (id integer primary key autoincrement, title varchar(256));", chtab)
	if _, err := db.Exec(q); err != nil {
		fmt.Fprintf(os.Stderr, "babayaga: %v\n", err)
	}
}
