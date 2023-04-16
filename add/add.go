package add

import (
	"About_me_Bot/config"
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	start = "/start - start the game with your bot"
	links = "/links - https://habr.com/ru/post/351060/, https://go-telegram-bot-api.dev/"
	about = "EMINEM Born\tMarshall Bruce Mathers III\nOctober 17, 1972 (age 50)\nSt. Joseph, Missouri, U.S.\nOther names\t\nSlim ShadyEvil (as part of Bad Meets Evil)M&MMC Double M\nOccupation\t\nRappersingersongwriterrecord producerrecord executiveactor\nYears active\t1988–present[1]\nWorks\t\nAlbumssinglesproductionvideography"
)

func Add(cfg *config.Config) error {
	bot, err := tgbotapi.NewBotAPI(cfg.Btoken)
	if err != nil {
		return err
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
		}
		switch update.Message.Command() {
		case "start":
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, start)
			_, err := bot.Send(msg)
			if err != nil {
				log.Printf("Error: '%v'\n", err)
			}
			continue
		case "links":
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, links)
			_, err := bot.Send(msg)
			if err != nil {
				log.Printf("Error: '%v'\n", err)
			}
			continue
		case "about":
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("%s%s", "/about - ", about))
			_, err := bot.Send(msg)
			if err != nil {
				log.Printf("Error: '%v'\n", err)
			}
			continue
		case "help":
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("%s%s%s%s", "/help - ", "/start,", "/links,", "/about"))
			_, err := bot.Send(msg)
			if err != nil {
				log.Printf("Error: '%v'\n", err)
			}
			continue
		default:
			fmt.Println("default")
			if err := Add(cfg); err != nil {
				log.Printf("func Add() failed. Error: '%v'\n", err)
			}
		}
	}
	return err
}
