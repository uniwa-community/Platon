/*
   BELOW SCRIPT IS SIGNED UNDER The GPLv3 License (GPLv3)

   Copyright (c) 2022 Author

   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU General Public License for more details.

   You should have received a copy of the GNU General Public License
   along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

type Tab struct {
	Date    string
	Content string
	Type    string
	Link    string
}

// Wrapper is a session that saves up 126 announcements in cache.
// Every N hour a new Wrapper is executed.
// It is assumed if a new ID is being spotted, then a new announcement has been created.
type Wrapper struct {
	Collector *colly.Collector
	Cache     []Tab
}

// Setup creates a new Collector for web scrapping while also
// it sets a timeout of 120 seconds limit to avoid wasting time in case the site is too damn slow
func (w *Wrapper) Setup() {
	w.Collector = colly.NewCollector()
	w.Collector.SetRequestTimeout(120 * time.Second)
	w.Collector.OnError(func(_ *colly.Response, err error) {
		log.Println("Πρόβλημα σύνδεσης με την ιστοσελίδα της σχολής", err)
	})
}

// It will print out the cache that is
// a slice of Tab with their contents
func (w *Wrapper) PrintAnnouncements() {
	fmt.Printf("%+v", w.Cache)
}

// Visits the website for announcements and saves them in Cache
func (w *Wrapper) GrabAnnouncements() {
	w.Collector.OnHTML("div[data-url]", func(e *colly.HTMLElement) {

		// The person who wrote the announcements I diagnose him with seizure
		txt := strings.ReplaceAll(strings.ReplaceAll(e.Text, "\t", ""), "\n", "")

		split := strings.Split(txt, " ")

		var content []string
		for _, txt := range split {
			if txt != "" {
				content = append(content, txt)
			}
		}

		var tab Tab = Tab{
			Date:    content[0],
			Content: strings.Join(content[1:len(content)-2], " "),
			Type:    content[len(content)-1],
			Link:    e.Attr("data-url"),
		}

		w.Cache = append(w.Cache, tab)
	})

	w.Collector.Visit("http://www.ice.uniwa.gr/announcements-all/")
}
