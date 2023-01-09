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
	"log"
	"os"
	"strconv"
)

var DEBUG bool = false
var TOKEN string

var ANNOUNCEMENT_ID string = os.Getenv("ANNOUNCEMENT_ID")
var LOG_CHANNEL string = os.Getenv("LOG_CHANNEL_ID")

func init() {
	debug := os.Getenv("DEBUG")
	if debug == "true" {
		DEBUG = true
	} else if debug == "false" {
		DEBUG = false
	} else {
		panic("Tο debug variable έχει λάθος τιμή, πρέπει να είναι είτε true ή false")
	}
	TOKEN = os.Getenv("TOKEN")
	log.Println(TOKEN)
	if TOKEN == "" {
		panic("Δεν έχει δωθεί TOKEN, σου σπάω το πρόγραμμα τώρα πριν γίνει αργότερα.")
	}
}

func main() {
	a_id, err := strconv.ParseInt(ANNOUNCEMENT_ID, 10, 64)
	if err == nil {
		panic(err)
	}
	bot := Bot{Token: TOKEN, Announcement_ID: a_id, Debug: DEBUG}
	bot.Setup()
	bot.Update()
}
