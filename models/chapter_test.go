package models

import (
	"testing"
)

func init() {
	Open(false)
}

func TestChapterSel(t *testing.T) {
	chapter := Chapter{id: 3}
	var c Chapter
	c.Select(chapter.id)
	if c != chapter {
		t.Errorf("Loaded incorrect chapter (expected %s, loaded %s)", chapter, c)
	}
	Close()
}

func TestChapterIns(t *testing.T) {
}

func TestChapterUpd(t *testing.T) {
}

func TestChapterDel(t *testing.T) {
}
