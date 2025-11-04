package main

import (
	"fmt"

	"gopkg.in/telebot.v4"
)

// РАЗДЕЛЫ ПРИЛОЖЕНИЯ

func firstMeeting(context telebot.Context, dataAboutUser *User) {

	context.Send(&telebot.Sticker{File: telebot.File{FileID: "CAACAgIAAxkBAAIBT2jO4pDTCO3VsPAGdU8lStRIetWPAAJ7AAPBnGAM2xofth1UNog2BA"}})

	context.Send(
		fmt.Sprintf("<b>Приветик, %s! 👋\nЭто МАИнчик</b>\n\nПару сек, и ты окунешься с головой в новое общение!\n\n<blockquote>Продолжая, ты принимаешь <b><a href='https://telegra.ph/MAInchik--politika-polzovaniya-servera-09-28-2'>наши условия</a></b></blockquote>\n\n<b>ВЛЕТАЙ СКОРЕЕ! Да, ты) 👇</b>", dataAboutUser.Personal.Login),
		&telebot.SendOptions{
			ParseMode:             telebot.ModeHTML,
			DisableWebPagePreview: true,
			ReplyMarkup: &telebot.ReplyMarkup{
				InlineKeyboard: [][]telebot.InlineButton{
					{{Text: "В меню 🏠", Data: "menu"}, {Text: "Моя анкета 🤠", Data: "editMyPage"}},
				},
			},
		},
	)
}

func menu(context telebot.Context, dataAboutUser *User) {

	dataAboutUser.Action = "menu"

	fakeActsCount := AllStatisticsData.ActsCount + AllStatisticsData.FakeActsCount
	fakeUsersCount := len(UsersData) - 200

	maleCount, femaleCount := 0, 0
	{
		for _, obj := range UsersData {
			switch obj.Personal.Gender {
			case "Муж":
				maleCount++
			}
		}
		femaleCount = fakeUsersCount - maleCount
	}

	context.Bot().Edit(telebot.StoredMessage{
		ChatID:    dataAboutUser.ChatID,
		MessageID: dataAboutUser.MessageID,
	},
		fmt.Sprintf("%s, <b>%s!</b>\n\n<blockquote>Анкета - <b>%v</b>\n<b><a href='https://t.me/MAInchik_bot/?start=settings'>Фильтр и видимость</a></b>\n\nЛайков всего: <b>%v шт</b>\nНажатий в боте: <b>%v шт</b></blockquote>\n\n<b>Делай, че по кайфу 😉</b>",

			greetingText(), dataAboutUser.Personal.Login, map[bool]string{true: map[bool]string{true: "АКТИВНА ☑️", false: "СКРЫТА 👻"}[dataAboutUser.PageIsShowing], false: "НЕ ЗАКОНЧЕНА ⚠️"}[dataAboutUser.Personal.IsVerified], dotFormatNumber(AllStatisticsData.LikesCount), dotFormatNumber(fakeActsCount)),
		&telebot.SendOptions{
			ParseMode:             telebot.ModeHTML,
			DisableWebPagePreview: true,
			ReplyMarkup: &telebot.ReplyMarkup{
				InlineKeyboard: [][]telebot.InlineButton{
					{{Text: "Смотреть " + map[bool]string{true: fmt.Sprintf("всех (%d+) 🔥", fakeUsersCount), false: map[bool]string{true: fmt.Sprintf("мальчиков (%d+) 😎", maleCount), false: fmt.Sprintf("девчонок (%d+) 🥰", femaleCount)}[dataAboutUser.PagesGender == "Муж"]}[dataAboutUser.PagesGender == "Все"], Data: "showRecomendations"}},

					{{Text: "Лайкнули ✨", Data: "showLikedMe"}, {Text: "Моя анкета 🤠", Data: "editMyPage"}},

					{{Text: "Создатели 🤔", Data: "aboutUs"}, {Text: "Настройки ⚙️", Data: "settings"}},

					{{Text: fmt.Sprintf("Мой ранг: %s", dataAboutUser.Rank), Data: "aboutRank"}},
				},
			},
		},
	)
}

func help(context telebot.Context, dataAboutUser *User) {
	dataAboutUser.Action = "help"

	context.Send(
		&telebot.Photo{
			File:    telebot.File{FileID: "AgACAgIAAxkBAAIJtGjcPqKVr_ePtUeAdVA_PRNReHvcAAIH_jEbKYfpSu4qNUb_3CmlAQADAgADeQADNgQ"},
			Caption: "<b>Солнце, ты в правильном месте, пиши поддержке в ЛС 🤗\n\nТы можешь:</b><blockquote>- пожаловаться на любую анкету\n- решить тех. вопрос\n- закинуть крутую идею</blockquote>\n\n<b>Мы УЖЕ ждем твое сообщение 👇</b>",
		},
		&telebot.SendOptions{
			ParseMode:             telebot.ModeHTML,
			DisableWebPagePreview: true,
			ReplyMarkup: &telebot.ReplyMarkup{
				InlineKeyboard: [][]telebot.InlineButton{
					{{Text: "Пиши скорее 💭", URL: "https://t.me/te1ron"}},
				},
			},
		},
	)
}

func settings(context telebot.Context, dataAboutUser *User) {

	dataAboutUser.Action = "settings"

	context.Bot().Edit(telebot.StoredMessage{
		ChatID:    dataAboutUser.ChatID,
		MessageID: dataAboutUser.MessageID,
	},
		fmt.Sprintf("<b>⚙️ Настройки • <code>%v</code>\n\nТвоя анкета</b> - <a href='https://t.me/MAInchik_bot/?start=editMyPage'>изменить</a><blockquote><b>%s, %v</b>\n<i>«%s»</i>\n\n<b>%s</b> - <a href='https://t.me/MAInchik_bot/?start=aboutRank'>что это?</a></blockquote>\n\n<b>За все время:</b><blockquote>Поставлено <b>%v</b>\nПолучено <b>%v</b>\n\nТы с нами с <b>%v</b></blockquote>\n\n<b><a href='https://telegra.ph/MAInchik--politika-polzovaniya-servera-09-28-2'>Условия и правила</a></b>", dataAboutUser.ChatID, dataAboutUser.Personal.Login, declension(dataAboutUser.Personal.Age, "год", "года", "лет", true), trimWithDots(dataAboutUser.Personal.About, 15), dataAboutUser.Rank, declension(dataAboutUser.Statistics.LikesCount, "лайк", "лайка", "лайков", true), declension(dataAboutUser.Statistics.PageLikesCount, "лайк", "лайка", "лайков", true), dataAboutUser.Date.Format("02.01.06")),
		&telebot.SendOptions{
			ParseMode:             telebot.ModeHTML,
			DisableWebPagePreview: true,
			ReplyMarkup: &telebot.ReplyMarkup{
				InlineKeyboard: [][]telebot.InlineButton{
					{{Text: fmt.Sprintf("Фильтр по полу: %s 🔄", dataAboutUser.PagesGender), Data: "toggleEditPagesGender"}},

					{{Text: "Просмотры ❌", Data: "resetViewedPages"}, {Text: fmt.Sprintf("Видимость: %s", map[bool]string{true: "☑️", false: "👻"}[dataAboutUser.PageIsShowing]), Data: "togglePageIsShowing"}},

					{{Text: "⬅️В меню", Data: "menu"}, {Text: "Помощь 💭", Data: "help"}},
				},
			},
		},
	)

}

func aboutUs(context telebot.Context, dataAboutUser *User) {

	_, err := context.Bot().Edit(telebot.StoredMessage{
		ChatID:    dataAboutUser.ChatID,
		MessageID: dataAboutUser.MessageID,
	},
		&telebot.Photo{
			File:    telebot.File{FileID: "AgACAgIAAxkBAAIIlmjZT5ZimMgjQmkbOIScd58xwb2bAAIeBDIb9RHISjBBopaf1w4gAQADAgADdwADNgQ"},
			Caption: fmt.Sprintf("<b>❝ а кто мы по правде?</b>\n- 2 первокурсника\n\n<b><a href='https://t.me/qu1z3x'>Давид</a></b> - разраб, дизайнер, CEO\n<blockquote><i>Сделал так, чтобы вы все пользовались МАИнчиком с улыбкой на лице (2338 строк кода). И весь визуал — каждая картинка здесь и в постах выточена им мышкой в Figma. Держит глобальную дорогу и стиль проекта.</i> - <b><a href='https://github.com/qu1z3x'>GitHub</a></b> </blockquote>\n\n<b><a href='https://t.me/te1ron'>Артем</a></b> - пиар, поддержка, SMM\n<blockquote><i>Создатель идей и текстов в канале, промоутер, тестировщик, инвестор проекта и голос коммьюнити МАИнчика. Именно он отвечал тебе в поддержке, именно его сотка по русскому слышна в грамотности каждого поста.</i></blockquote>\n\n<b>Спасибо за прочтение, %s</b>", dataAboutUser.Personal.Login),
		},
		&telebot.SendOptions{
			ParseMode:             telebot.ModeHTML,
			DisableWebPagePreview: true,
			ReplyMarkup: &telebot.ReplyMarkup{
				InlineKeyboard: [][]telebot.InlineButton{
					{{Text: "⬅️В меню", Data: "menuWithDelete"}, {Text: fmt.Sprintf("❤️‍🔥 %v", AllStatisticsData.AboutUsLikes), Data: "likeAboutUs"}},
				},
			},
		},
	)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func aboutRank(context telebot.Context, dataAboutUser *User) {

	context.Send(
		fmt.Sprintf("<b>ЭТО РАНГ АНКЕТЫ 🤔</b>\n\n<blockquote><b>👉 Получить можно за красивые глазки и за активность</blockquote>\n\nТвой ранг: <b>«%s»</b>\n%s", dataAboutUser.Rank, map[bool]string{true: "- он особенный и виден всем", false: "- он обычный и не отображается"}[dataAboutUser.Rank != "Юзер 😀"]),
		&telebot.SendOptions{
			ParseMode:             telebot.ModeHTML,
			DisableWebPagePreview: true,
			ReplyMarkup: &telebot.ReplyMarkup{
				InlineKeyboard: [][]telebot.InlineButton{
					{{Text: genderDeclension(dataAboutUser.Personal.Gender, "Понял", "Поняла") + " 👍", Data: "deleteMessage"}},
				},
			},
		},
	)
}

func aboutChannel(context telebot.Context, dataAboutUser *User) {

	context.Bot().Edit(telebot.StoredMessage{
		ChatID:    dataAboutUser.ChatID,
		MessageID: dataAboutUser.MessageID,
	},
		fmt.Sprintf("<b>Давай так</b>\n👉 Мы делаем <i>бесплатно и круто,</i> а ты <b>подписываешься на канал</b>\n\n<a href='https://t.me/mainchik'><b>МАИнчик | знакомства 💜</b></a>\n\n<b>Исправляйся, %s) 👆</b>", genderDeclension(dataAboutUser.Personal.Gender, "дружище", "подруга")),
		&telebot.SendOptions{
			ParseMode:             telebot.ModeHTML,
			DisableWebPagePreview: true,
			ReplyMarkup: &telebot.ReplyMarkup{
				InlineKeyboard: [][]telebot.InlineButton{
					{{Text: "⬅️В меню", Data: "menu"}},
				},
			},
		},
	)
}

func showPrivateUserData(context telebot.Context, dataAboutUser *User, userID int64) {

	var dataAboutCertainUser *User = nil
	for i := range UsersData {
		if UsersData[i].ChatID == userID {
			dataAboutCertainUser = &UsersData[i]
			break
		}
	}

	if dataAboutCertainUser != nil {

		dataAboutUser.CurrentPageID = dataAboutCertainUser.ChatID

		_, err := context.Bot().Edit(telebot.StoredMessage{
			ChatID:    dataAboutUser.ChatID,
			MessageID: dataAboutUser.MessageID,
		},
			fmt.Sprintf("<b>%v • <code>%v</code></b>\n\n<b>Подробно:</b><blockquote>Ранг: <b>«%v»</b>\n\nВсего <b>%v</b>\nС нами <b>с %v</b></blockquote>", dataAboutCertainUser.Personal.Login, dataAboutCertainUser.ChatID, dataAboutCertainUser.Rank, declension(dataAboutCertainUser.Statistics.ActsCount, "нажатие", "нажатия", "нажатий", true), dataAboutCertainUser.Date.Format("02.01.06")),
			&telebot.SendOptions{
				ParseMode:             telebot.ModeHTML,
				DisableWebPagePreview: true,
				ReplyMarkup: &telebot.ReplyMarkup{
					InlineKeyboard: [][]telebot.InlineButton{
						{{Text: "Контакт 🙈", URL: fmt.Sprintf("tg://user?id=%v", dataAboutCertainUser.ChatID)}, {Text: "Анкета 🤠", Data: fmt.Sprintf("showPage%v", dataAboutCertainUser.ChatID)}},

						{{Text: map[bool]string{true: "", false: fmt.Sprintf("✏️ Ранг: %s", map[bool]string{true: "... ❌", false: dataAboutCertainUser.Rank}[dataAboutUser.Action == "EditRank"])}[dataAboutCertainUser.ChatID == qu1z3xID], Data: "toggleEditRank"}},

						{{Text: map[bool]string{true: "", false: fmt.Sprintf("Заблок: %v", boolIcon(dataAboutCertainUser.InBlackList))}[dataAboutCertainUser.ChatID == qu1z3xID], Data: fmt.Sprintf("toggleBlockUser%v", dataAboutCertainUser.ChatID)}, {Text: map[bool]string{true: "", false: fmt.Sprintf("Модерация: %v", boolIcon(dataAboutCertainUser.Personal.IsVerified))}[dataAboutCertainUser.ChatID == qu1z3xID], Data: fmt.Sprintf("toggleIsVerified%v", dataAboutCertainUser.ChatID)}},

						{{Text: "Закрыть панель ✖️", Data: "deleteMessage"}},
					},
				},
			},
		)

		if err != nil {
			sendDataAboutError(dataAboutUser, err)
			return
		}
	} else {
		context.Bot().Edit(telebot.StoredMessage{
			ChatID:    dataAboutUser.ChatID,
			MessageID: dataAboutUser.MessageID,
		},
			"<b>Не нашел пользователя 😔</b>",
			&telebot.SendOptions{
				ParseMode:             telebot.ModeHTML,
				DisableWebPagePreview: true,
				ReplyMarkup: &telebot.ReplyMarkup{
					InlineKeyboard: [][]telebot.InlineButton{
						{{Text: "Удалить это ❌", Data: "deleteMessage"}},
					},
				},
			},
		)
	}
}

func blackListMessage(context telebot.Context, dataAboutUser *User) {
	context.Send(
		fmt.Sprintf("<b>%s <a href='https://telegra.ph/MAInchik--politika-polzovaniya-servera-09-28-2'>наши простые условия</a> 😔</b>\n\nПока ты не можешь пользоваться <b>МАИнчиком</b>\n\n<b>Ты всегда можешь написать нам в ЛС 👇</b>", genderDeclension(dataAboutUser.Personal.Gender, "Дружище, ты нарушил", "Подруга, ты нарушила")),
		&telebot.SendOptions{
			ParseMode:             telebot.ModeHTML,
			DisableWebPagePreview: true,
			ReplyMarkup: &telebot.ReplyMarkup{
				InlineKeyboard: [][]telebot.InlineButton{
					{{Text: "Пиши скорее 💭", URL: "https://t.me/te1ron"}},
				},
			},
		},
	)
}

// func reminderAlerts(context telebot.Context, dataAboutUser *User) {

// }
