/*
   The GPLv3 License (GPLv3)

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

import "fmt"

// SubCommand of UniWa bot
// @Plato /info etc
var Commands map[string]string = map[string]string{
	"/info":    "Γειά σου καλώς όρισες στην ομάδα, με λένε Πλάτωνα.\n Για περισσότερες πληροφορίες μπες στα announcements και διάλεξε τα κατάλληλα κανάλια που επιθυμείς.",
	"/contact": "Επικοινωνία με γραμματεία του τμήματος.\n Ραντεβού πλέον δεν απαιτείται. \n http://www.ice.uniwa.gr/contact/",
}

// a way to spread 56 courses in a easy way
// @Plato /find_server Φυσική
var CoursesLinks map[string]string = map[string]string{
	"Μαθηματική-Ανάλυση-1":        "https://t.me/+dQRIteGCUJ0yYThk",
	"Γραμμική-Άλγεβρα":            "https://t.me/+i7GU8-cAZbQ5YTRk",
	"Εισαγωγή-Υπολογιστών":        "https://t.me/+HoaZwRX0lqxmZGJk",
	"Προγραμματισμός-Υπολογιστών": "https://t.me/+DO3TZ8xsTX9kZWE8",
	"Διακριτά-Μαθηματικά":         "https://t.me/+mGJ_2vFO61sxMGU0",
	"Φυσική":                      "https://t.me/+ItCW2lMbrkplYTY0",
	// δίνω την ευκαρία σε κάποιον να κάνει contribute να βάλει περισσότερα μαθήματα
	// contact @bill88t
}

// Turns a whole hash table into a string
func CreateList(content map[string]string) string {
	txt := ""
	for key, value := range content {
		txt += fmt.Sprintf("%s: %s \n", key, value)
	}
	txt += "\n"
	return txt
}

// Retuns an output of all courses to the end user
func ListLessons() string {
	return CreateList(CoursesLinks)
}

// Retuns an output of all available subcommands to the end user
func Help() string {
	return CreateList(Commands)
}
