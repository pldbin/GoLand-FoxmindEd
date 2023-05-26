package bot

import (
	"About_me_Bot/config"
	"About_me_Bot/holiday"
	"About_me_Bot/logger"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strings"
	"time"
)

const (
	start = "/start - start the game with your bot and Please select a country to get information about its holidays"
	links = "/links - https://habr.com/ru/post/351060/, https://go-telegram-bot-api.dev/"
	about = "EMINEM Born\tMarshall Bruce Mathers III\nOctober 17, 1972 (age 50)\nSt. Joseph, Missouri, U.S.\nOther names\t\nSlim ShadyEvil (as part of Bad Meets Evil)M&MMC Double M\nOccupation\t\nRappersingersongwriterrecord producerrecord executiveactor\nYears active\t1988–present[1]\nWorks\t\nAlbumssinglesproductionvideography"
)

var countries = map[string]string{
	"🇺🇸": "US",
	"🇬🇧": "GB",
	"🇨🇦": "CA",
	"🇦🇺": "AU",
}

func Add(cfg *config.Config, logger logger.LogrusLogger) error {
	bot, err := tgbotapi.NewBotAPI(cfg.Btoken)
	if err != nil {
		return err
	}

	bot.Debug = true
	logger.Infof("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	countriesFlags := []string{}
	for k := range countries {
		countriesFlags = append(countriesFlags, k)
	}
	holidayincountry := false

	for update := range updates {
		switch update.Message.Command() {

		case "start":
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, start)

			var newKeyboardButtonRows [][]tgbotapi.KeyboardButton
			for i := 0; i < len(countriesFlags); i += 2 {
				newKeyboardButton1 := tgbotapi.NewKeyboardButton(countriesFlags[i])
				newKeyboardButton2 := tgbotapi.NewKeyboardButton(" ")
				if i+1 < len(countries) {
					newKeyboardButton2 = tgbotapi.NewKeyboardButton(countriesFlags[i+1])
				}
				newKeyboardButtonRow := tgbotapi.NewKeyboardButtonRow(newKeyboardButton1, newKeyboardButton2)
				newKeyboardButtonRows = append(newKeyboardButtonRows, newKeyboardButtonRow)
			}

			msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
				newKeyboardButtonRows...,
			)
			_, err := bot.Send(msg)

			if err != nil {
				logger.Errorf("Error: '%v'\n", err)
			}
			holidayincountry = true
			continue
		case "links":
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, links)
			_, err := bot.Send(msg)
			if err != nil {
				logger.Errorf("Error: '%v'\n", err)
			}
			continue
		case "about":
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("%s%s", "/about - ", about))
			_, err := bot.Send(msg)
			if err != nil {
				logger.Errorf("Error: '%v'\n", err)
			}
			continue
		case "help":
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("%s%s%s%s", "/help - ", "/start,", "/links,", "/about"))
			_, err := bot.Send(msg)
			if err != nil {
				logger.Errorf("Error: '%v'\n", err)
			}
			continue
		case "exit":
			holidayincountry = false
		default:
			logger.Info("We don't have any other routes, please set another value. ")
			if update.Message == nil { // If we got a non-command message
				continue
			}
			if update.Message != nil { // If we got a message
				logger.Infof("[%s] %s", update.Message.From.UserName, update.Message.Text)
			}

			if holidayincountry {
				flag := update.Message.Text
				if v, e := countries[flag]; e {
					h, err := holiday.GetHolidays(v, time.Now(), cfg.HolidaysAPI)
					if err != nil {
						logger.Errorf("Error: '%v'\n", err)
					}

					var holidaysMsg string
					if len(h) == 0 {
						holidaysMsg = "No new holidays found for " + flag + " for today.\n"
						logger.Info(holidaysMsg)
					} else {
						var holidaysName []string
						for i := 0; i < len(h); i++ {
							holidaysName = append(holidaysName, h[i].Name)
						}

						holidaysMsg = strings.Join(holidaysName, ", ")
					}

					msg := tgbotapi.NewMessage(update.Message.Chat.ID, holidaysMsg)

					_, err = bot.Send(msg)
					if err != nil {
						logger.Errorf("Error: '%v'\n", err)
					}
				}
			}
			continue
		}
	}
	return err
}
