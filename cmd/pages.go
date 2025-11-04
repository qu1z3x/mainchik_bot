package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"

	"gopkg.in/telebot.v4"
)

// ФУНКЦИИ ПОКАЗА АНКЕТ

func showPage(context telebot.Context, dataAboutUser *User, pageID int64, showLikeButtons bool) {

	dataAboutUser.CurrentPageID = pageID

	var dataAboutCertainUser *User = nil
	for i := range UsersData {
		if UsersData[i].ChatID == pageID {
			dataAboutCertainUser = &UsersData[i]
			break
		}
	}

	if dataAboutCertainUser != nil {
		pageText := fmt.Sprintf("<b>%v%s, %v</b>\n<blockquote><i>%s</i></blockquote>", map[bool]string{true: fmt.Sprintf("<a href='https://t.me/MAInchik_bot/?start=aboutRank'>%v</a>\n\n", dataAboutCertainUser.Rank), false: ""}[dataAboutCertainUser.Rank != "Юзер 😀"], dataAboutCertainUser.Personal.Login, declension(dataAboutCertainUser.Personal.Age, "год", "года", "лет", true), dataAboutCertainUser.Personal.About)

		pageButtons := [][]telebot.InlineButton{
			{{Text: "Закрыть анкету ✖️", Data: "deleteMessage"}},
		}

		if showLikeButtons {
			pageButtons = [][]telebot.InlineButton{
				{{Text: map[bool]string{true: "Админка этой анкеты 🥶", false: ""}[dataAboutUser.ChatID == qu1z3xID || dataAboutUser.ChatID == artemID], Data: fmt.Sprintf("showPrivateUserData%v", pageID)}},

				{{Text: "👎", Data: fmt.Sprintf("dislikePage%v", pageID)}, {Text: "❤️‍🔥", Data: fmt.Sprintf("likePage%v", pageID)}},
			}

			//? //////////////////  ЕСЛИ ПОПАЛСЯ САМ СЕБЕ  //////////////////

			if pageID == dataAboutUser.ChatID {
				pageButtons = [][]telebot.InlineButton{
					{{Text: "Это же ты! нравится? 😍", Data: "-"}}, {{Text: "Изменить 👎", Data: "editMyPage"}, {Text: "Да ❤️‍🔥", Data: "showRecomendations"}},
				}
			}
		}

		_, err := context.Bot().Send(
			context.Chat(),
			&telebot.Photo{
				File:    telebot.File{FileID: dataAboutCertainUser.Personal.MediaID},
				Caption: pageText,
			},
			&telebot.SendOptions{
				ParseMode:             telebot.ModeHTML,
				DisableWebPagePreview: true,
				ReplyMarkup: &telebot.ReplyMarkup{
					InlineKeyboard: pageButtons,
				},
			},
		)

		//! ЕСЛИ ПРОБЛЕМА С ФОТОГРАФИЕЙ - ОТСЫЛАЕТ ПРОСТО АНКЕТУ БЕЗ НЕЕ

		if err != nil {
			context.Send(pageText,
				&telebot.SendOptions{
					ParseMode:             telebot.ModeHTML,
					DisableWebPagePreview: true,
					ReplyMarkup: &telebot.ReplyMarkup{
						InlineKeyboard: pageButtons,
					},
				},
			)

			return
		}

	}
}

func showingAlgorithm(context telebot.Context, dataAboutUser *User, mode string) {

	dataAboutUser.Action = mode

	switch mode {
	case "recomendations":

		if len(dataAboutUser.ViewedPages) > 0 && len(dataAboutUser.ViewedPages)%5 == 0 && rand.Intn(100) < 50 {
			showAd(context, dataAboutUser)
			return
		}

		for i := 0; i < 4000; i++ {
			randUser := &UsersData[rand.Intn(len(UsersData))]

			if randUser.PageIsShowing &&
				randUser.Personal.IsVerified &&
				!randUser.InBlackList &&
				randUser.ChatID != dataAboutUser.ChatID &&
				!contains(randUser.LikedPages, dataAboutUser.ChatID) &&
				!contains(dataAboutUser.LikedPages, randUser.ChatID) &&
				!contains(dataAboutUser.ViewedPages, randUser.ChatID) &&

				(dataAboutUser.PagesGender == "Все" || dataAboutUser.PagesGender == randUser.Personal.Gender) {

				showPage(context, dataAboutUser, randUser.ChatID, true)
				return
			}
		}

		context.Send(
			fmt.Sprintf("<b>Новые анкеты %sзакончились, возвращайся позже 😉</b>", map[bool]string{true: "", false: map[bool]string{true: "мальчиков ", false: "девчонок "}[dataAboutUser.PagesGender == "Муж"]}[dataAboutUser.PagesGender == "Все"]),
			&telebot.SendOptions{
				ParseMode:             telebot.ModeHTML,
				DisableWebPagePreview: true,
				ReplyMarkup: &telebot.ReplyMarkup{
					InlineKeyboard: [][]telebot.InlineButton{
						{{Text: "Сбросить просмотры 🔄", Data: "resetViewedPages"}},
					},
				},
			},
		)

		if msg, err := context.Bot().Send(context.Chat(), "ㅤ"); err == nil {
			dataAboutUser.MessageID = strconv.Itoa(msg.ID)
		}

		menu(context, dataAboutUser)
	case "likedMe":

		// ПОКАЗ АНКЕТ КОТОРЫЕ ТЕБЯ ЛАЙКНУЛИ (+ ДОБАВЛЕНИЕ ИХ В СПИСОК "ПРОСМОТРЕННЫХ")

		for i := range UsersData {
			user := &UsersData[i]
			if contains(user.LikedPages, dataAboutUser.ChatID) && !user.InBlackList {
				showPage(context, dataAboutUser, user.ChatID, true)
				return
			}
		}

		context.Send(
			"Пока это все твои <b>«фанатики» 😊</b>",
			&telebot.SendOptions{
				ParseMode:             telebot.ModeHTML,
				DisableWebPagePreview: true,
			},
		)

		if msg, err := context.Bot().Send(context.Chat(), "ㅤ"); err == nil {
			dataAboutUser.MessageID = strconv.Itoa(msg.ID)
		}

		menu(context, dataAboutUser)
	}
}

func editMyPage(context telebot.Context, dataAboutUser *User) {

	if dataAboutUser != nil {

		// ФУНКЦИЯ ПОКАЗА ВИДА ПОЛЯ ВО ВРЕМЯ РЕДАКТИРОВАНИЯ

		editingView := func(value interface{}, buttonName string) string {
			if strings.Contains(dataAboutUser.Action, buttonName) {
				switch buttonName {
				case "Media":
					return "... ❌"
				case "Login":
					return "... ❌"
				case "Age":
					return "... 16-30 ❌"
				case "About":
					return "... ❌"
				}
			} else {
				switch buttonName {
				case "Age":
					if dataAboutUser.Personal.Age == 0 {
						return "➕"
					} else {
						return fmt.Sprintf("%v ✔️", value)
					}
				case "About":
					if dataAboutUser.Personal.About == "" {
						return "➕"
					} else {
						return "✔️"
					}
				case "Media":
					if dataAboutUser.Personal.MediaID == "AgACAgIAAxkBAAIJsmjcPf09HQ-MwghHpi58OQACMepPAAIE_jEbKYfpSkjOa9LBjv7eAQADAgADeQADNgQ" {
						return "➕"
					} else {
						return "✔️"
					}
				default:
					return fmt.Sprintf("%v ✔️", value)
				}
			}
			return fmt.Sprintf("%v ✔️", value)
		}

		//

		context.Bot().Edit(telebot.StoredMessage{
			ChatID:    dataAboutUser.ChatID,
			MessageID: dataAboutUser.MessageID,
		},
			&telebot.Photo{
				File:    telebot.File{FileID: dataAboutUser.Personal.MediaID},
				Caption: fmt.Sprintf("<b>%v%s, %v</b>\n<blockquote><i>%s</i></blockquote>\n\n<b>Нажми и вписывай данные 👇</b>", map[bool]string{true: fmt.Sprintf("<a href='https://t.me/MAInchik_bot/?start=aboutRank'>%v</a>\n\n", dataAboutUser.Rank), false: ""}[dataAboutUser.Rank != "Юзер 😀"], dataAboutUser.Personal.Login, declension(dataAboutUser.Personal.Age, "год", "года", "лет", true), map[bool]string{true: dataAboutUser.Personal.About, false: "дизайнер из Санкт-Петербурга..\n\n<b>(меньше 300 символов)</b>"}[dataAboutUser.Personal.About != ""]),
			}, &telebot.SendOptions{
				ParseMode:             telebot.ModeHTML,
				DisableWebPagePreview: true,
				ReplyMarkup: &telebot.ReplyMarkup{
					InlineKeyboard: [][]telebot.InlineButton{
						{{Text: fmt.Sprintf("Имя: %s", editingView(trimWithDots(dataAboutUser.Personal.Login, 25), "Login")), Data: "toggleEditLogin"}},

						{{Text: fmt.Sprintf("Лет: %s", editingView(dataAboutUser.Personal.Age, "Age")), Data: "toggleEditAge"}, {Text: fmt.Sprintf("Пол: %s 🔄", dataAboutUser.Personal.Gender), Data: "toggleEditGender"}},

						{{Text: fmt.Sprintf("О себе: %s", editingView(dataAboutUser.Personal.About, "About")), Data: "toggleEditAbout"}, {Text: fmt.Sprintf("Фотка: %s", editingView(dataAboutUser.Personal.MediaID, "Media")), Data: "toggleEditMedia"}},

						{{Text: "Продолжить 🚀", Data: "applyPageChanges"}},
					},
				},
			},
		)
	}
}

func showAd(context telebot.Context, dataAboutUser *User) {

	//? ///////////////////////  РЕКЛАМНЫЕ МЕСТА  ///////////////////////////

	var adsList []func() = []func(){

		// ПОКАЗ АНКЕТЫ САМОГО ЮЗЕРА

		func() {
			showPage(context, dataAboutUser, dataAboutUser.ChatID, true)
		},

		// РЕКЛАМА СОЗДАТЕЛЕЙ - ФОТО

		func() {
			AllStatisticsData.firstAdViewsCount++

			_, err := context.Bot().Send(
				context.Chat(),
				&telebot.Photo{
					File:    telebot.File{FileID: "AgACAgIAAxkBAAIr6mkIuJ_wFRGega5xvC7qAbe82frIAAJ2C2sba4BISLv8y3Bth2hWAQADAgADeQADNgQ"},
					Caption: "<b><a href='https://t.me/MAInchik_bot/?start=aboutRank'>СОЗДАТЕЛИ 🥶</a>\n\nДавид и Артем</b>\n- два перво...\n\n",
				},
				&telebot.SendOptions{
					ParseMode: telebot.ModeHTML,
					ReplyMarkup: &telebot.ReplyMarkup{
						InlineKeyboard: [][]telebot.InlineButton{
							{
								{Text: "Дальше 👎", Data: "dislikeAd"},
								{Text: "Глянуть ❤️‍🔥", Data: "aboutUs"},
							},
						},
					},
				},
			)

			//! ЕСЛИ ПРОБЛЕМА С МЕДИА - ПРОСТО ПРОПУСК РЕКЛАМЫ

			if err != nil {
				showingAlgorithm(context, dataAboutUser, "recomendations")
				return
			}
		},

		// РЕКЛАМА МОИХ РИЛСОВ - ВИДЕО

		// func() {
		// 	AllStatisticsData.firstAdViewsCount++

		// 	var reelsList []string = []string{"BAACAgIAAxkBAAIr4GkIqo-TDu0AAWQPMRqUs0IHtbajAQACfYoAAiwGSEheQ4BDzvElfDYE"}

		// 	_, err := context.Bot().Send(
		// 		context.Chat(),
		// 		&telebot.Video{
		// 			File:    telebot.File{FileID: reelsList[rand.Intn(len(reelsList))]},
		// 			Caption: "<b>СТОП, ЧЕ? 🤯</b>\n- это единственное видео в этом чате\n\n<b><a href='https://t.me/qu1z3x'>заходит? больше такого</a> ←</b>\n\n* вкл vpn",
		// 		},
		// 		&telebot.SendOptions{
		// 			ParseMode: telebot.ModeHTML,
		// 			ReplyMarkup: &telebot.ReplyMarkup{
		// 				InlineKeyboard: [][]telebot.InlineButton{
		// 					{
		// 						{Text: "Дальше 👎", Data: "dislikeAd"},
		// 						{Text: "Инста ❤️‍🔥", URL: "https://t.me/qu1z3x"},
		// 					},
		// 				},
		// 			},
		// 		},
		// 	)

		// 	//! ЕСЛИ ПРОБЛЕМА С МЕДИА - ПРОСТО ПРОПУСК РЕКЛАМЫ

		// 	if err != nil {
		// 		showingAlgorithm(context, dataAboutUser, "recomendations")
		// 		return
		// 	}
		// },
	}

	var adNumber int = rand.Intn(len(adsList))
	for i := 0; i < 70; i++ {

		if contains(dataAboutUser.ViewedPages, int64(adNumber)) {
			adNumber = rand.Intn(len(adsList))
		} else {
			break
		}
	}

	dataAboutUser.Action = "showAd"
	dataAboutUser.CurrentPageID = int64(adNumber)

	dataAboutUser.ViewedPages = append(dataAboutUser.ViewedPages, int64(adNumber))

	adsList[adNumber]()
}
