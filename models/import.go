package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func parseJSON(f string) {
	data, err := ioutil.ReadFile(f)
	if err != nil {
		fmt.Fprintf(os.Stderr, "babayaga: %v\n", err)
	}
	byt := []byte(string(data))

	var chapters []map[string][]map[string][]map[string]string
	if err := json.Unmarshal(byt, &chapters); err != nil {
		panic(err)
	}

	for _, chapter := range chapters {
		var c Chapter
		c.ins()

		for _, lesson := range chapter["lessons"] {
			var l Lesson = Lesson{chapterID: c.id}
			l.ins()

			for _, grammar := range lesson["grammars"] {
				var g Grammar = Grammar{
					lessonID: l.id,
					title:    grammar["title"],
					subtitle: grammar["subtitle"],
					text:     grammar["text"]}
				g.ins()
			}
		}
	}
}
