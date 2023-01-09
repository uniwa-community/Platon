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
	"strconv"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// Bot on constructor needs Token, Announcement_ID
// Debug (optinal) if enabled it will send stdout to terminal
// DO NOT PUT ANYTHING ON API, Announcements!  MAKE SURE TO RUN .Setup() AFTER THAT!
type Bot struct {
	Token           string
	Announcement_ID int64
	Announcements   []string
	API             *tgbotapi.BotAPI
	Cache           []Tab
	Debug           bool

	// TODO
	// Αν πεθάνει το πρόγραμα προφανώς θα
	// πρέπει να ξανασκανάρει σε N ώρες όχι επιτόπου
}

// Setups contents of Bot and gains Telegram Session
func (b *Bot) Setup() {
	bot, err := tgbotapi.NewBotAPI(b.Token)
	if err != nil {
		panic("Εισαγωγή Token είναι λάθος, το πρόγραμμα δεν μπορεί να συνεχιστεί χωρίς αυτό")
	}
	b.API = bot
	log.Printf("Authorized on account %s", bot.Self.UserName)
	b.Cache = FillCacheWithAnnouncements(false)
}

// SubCommands are a bot interface for users
// such as finding information about a Teacher, Classroom etc.
func (b *Bot) EnableSubCommands() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := b.API.GetUpdatesChan(u)
	if err != nil {
		log.Println(err)
	}
	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Printf("In channel %d [%s] %s", update.Message.Chat.ID, update.Message.From.UserName, update.Message.Text)

			// Log message before going to handler
			name := update.Message.From.FirstName + update.Message.From.LastName
			log_msg := fmt.Sprintf("User: %s send %s in %s", name, update.Message.Text, update.Message.Chat.Title)
			l_id, err := strconv.ParseInt(LOG_CHANNEL, 10, 64)
			if err == nil {
				log.Println("Failed to parse log_id")
				continue
			}
			_, err2 := b.API.Send(tgbotapi.NewMessage(l_id, log_msg))
			if err2 != nil {
				log.Println("Το μποτ δεν μπορεί να στείλει αυτό το περιεχόμενο, υπάρχει rate limit.")
			}
			b.CommandHandler(update)
		}
	}
}

// Double Checks Cache to avoid duplicate spam before posting
func (b *Bot) CompareAnnouncements() {
	log.Println("To Bot συλλέγει ενημερώσεις")
	for i, new_content := range FillCacheWithAnnouncements(false) {
		// This kind of check may seem a little unnecessary, but we better be safe first
		if b.Cache[i].Link != new_content.Link &&
			b.Cache[i].Content != new_content.Content &&
			b.Cache[i].Date != new_content.Date {
			txt := new_content.Type + " " + new_content.Content + "\n" + new_content.Link
			b.Announcements = append(b.Announcements, txt)
		}
	}
}

// Retuns every possible announcement from the page
func FillCacheWithAnnouncements(debug bool) []Tab {
	wrap := Wrapper{}
	wrap.Setup()
	wrap.GrabAnnouncements()
	if debug {
		wrap.PrintAnnouncements()
	}
	return wrap.Cache
}

// Yells and pings to everyone
func (b *Bot) MakeAnAnnouncement() {
	for _, announcement := range b.Announcements {
		_, err := b.API.Send(tgbotapi.NewMessage(b.Announcement_ID, announcement))
		if err != nil {
			log.Println("Το μποτ δεν μπορεί να στείλει αυτό το περιεχόμενο, υπάρχει rate limit.")
		}
	}
}

// Executes Bot's functionalities
func (b *Bot) Update() {

	go func() {
		for {
			b.CompareAnnouncements()
			b.MakeAnAnnouncement()
			time.Sleep(5 * time.Hour)
		}
	}()

	b.EnableSubCommands()
}

// A wrapper to handle commands in a more organized way
func (b *Bot) CommandHandler(update tgbotapi.Update) {

	length, cmd := HandleXargs(strings.Split(update.Message.Text, " "))

	// update.Message.Text
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

	// reply to whoever run the command
	msg.ReplyToMessageID = update.Message.MessageID

	if value, ok := Commands[cmd[0]]; ok {
		msg.Text = value
		b.API.Send(msg)
		return
	}

	switch cmd[0] {
	case "/help":
		msg.Text = Help()
	case "/find_server":
		switch length {
		case 1:
			msg.Text = ListLessons()
		case 2:
			msg.Text = FindLesson(cmd)
		default:
			msg.Text = FindLessons(cmd)
		}
	}

	b.API.Send(msg)
}

func HandleXargs(xargs []string) (int, []string) {

	length := len(xargs)
	if length == 1 {
		return 1, xargs
	}

	if xargs[0] == "@UniWa_bot" {
		xargs = xargs[1:]
		length -= 1
	}

	return length, xargs
}

func FindLesson(xargs []string) string {
	lesson := xargs[1]
	log.Println(lesson)
	value, ok := CoursesLinks[lesson]
	if !ok {
		return "Δεν υπάρχει αυτό το μάθημα."
	}

	return value
}

func FindLessons(lessons []string) string {
	txt := ""
	for _, course := range lessons[1:] {
		value, ok := CoursesLinks[course]
		if !ok {
			txt += course + ": Δεν υπάρχει αυτό το μάθημα.\n"
			continue
		}
		txt += fmt.Sprintf("%s: %s\n", course, value)
	}
	return txt
}
