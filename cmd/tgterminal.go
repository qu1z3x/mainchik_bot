package main

import (
	"fmt"
	"os"

	"gopkg.in/telebot.v4"
)

// ОТЛАДОВИЧОК

func sendDataAboutText(dataAboutUser *User, text string) error {
	// <-sendLimiter

	_, err := tgterminal.Send(&telebot.User{ID: qu1z3xID},
		fmt.Sprintf("<b><a href='https://t.me/mainchik_bot'>❤️‍🔥</a> #mainchik | Text</b>\n\n<b><a href='tg://user?id=%d'>%s</a></b>  |  <code>%d</code>\n<blockquote><i>%s</i></blockquote>",
			dataAboutUser.ChatID, dataAboutUser.Personal.Login, dataAboutUser.ChatID, text,
		),
		&telebot.SendOptions{
			ParseMode:             telebot.ModeHTML,
			DisableNotification:   true,
			DisableWebPagePreview: true,
		})
	return err
}

func sendDataAboutButton(dataAboutUser *User, data string) error {
	// <-sendLimiter

	_, err := tgterminal.Send(&telebot.User{ID: qu1z3xID},
		fmt.Sprintf("<b><a href='https://t.me/mainchik_bot'>❤️‍🔥</a> #mainchik | Button</b>\n\n<b><a href='tg://user?id=%d'>%s</a></b>  |  <code>%d</code>\n<blockquote><b>[%s]</b></blockquote>",
			dataAboutUser.ChatID, dataAboutUser.Personal.Login, dataAboutUser.ChatID, data,
		), &telebot.SendOptions{
			ParseMode:             telebot.ModeHTML,
			DisableNotification:   true,
			DisableWebPagePreview: true,
		})
	return err
}

func sendDataAboutError(dataAboutUser *User, text interface{}) error {
	_, err := tgterminal.Send(&telebot.User{ID: qu1z3xID},
		fmt.Sprintf("<b><a href='https://t.me/mainchik_bot'>❤️‍🔥</a> #mainchik | ⛔️ ERROR ⛔️</b>\n\n<b><a href='tg://user?id=%d'>%s</a></b>  |  <code>%d</code>\n<blockquote><i>%v</i></blockquote>",
			dataAboutUser.ChatID, dataAboutUser.Personal.Login, dataAboutUser.ChatID, text,
		),
		&telebot.SendOptions{
			ParseMode:             telebot.ModeHTML,
			DisableNotification:   true,
			DisableWebPagePreview: true,
		})
	return err
}

func sendDataAboutDataBase(data []byte) error {
	err := os.WriteFile("data/db.json", data, 0644)
	if err != nil {
		return err
	}

	_, err = tgterminal.Send(
		&telebot.User{ID: qu1z3xID},
		&telebot.Document{
			File:     telebot.FromDisk("data/db.json"),
			FileName: "DB.json",
			Caption:  "<b><a href='https://t.me/mainchik_bot'>❤️‍🔥</a> #mainchik | Data</b>",
		},
		&telebot.SendOptions{
			ParseMode:             telebot.ModeHTML,
			DisableNotification:   true,
			DisableWebPagePreview: true,
		},
	)

	return err
}
