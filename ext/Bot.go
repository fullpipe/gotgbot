package ext

import (
	"log"
	"encoding/json"
	"strconv"
	"gotgbot/types"
	"net/url"
)

type Bot struct {
	Token string
}

func (b Bot) GetMe() types.User {
	v := url.Values{}

	r := Get(b, "getChat", v)

	var u types.User
	json.Unmarshal(r.Result, &u)

	if !r.Ok {
		log.Fatal("You done goofed, API Res for getMe was not OK")
	}

	return u

}

func (b Bot) GetUserProfilePhotos(user_id int) types.UserProfilePhotos {
	v := url.Values{}
	v.Add("user_id", strconv.Itoa(user_id))


	r := Get(b, "getUserProfilePhotos", v)
	if !r.Ok {
		log.Println("You done goofed")
		log.Println(r)
	}

	var userProfilePhotos types.UserProfilePhotos
	json.Unmarshal(r.Result, &userProfilePhotos)

	return userProfilePhotos
}


func (b Bot) GetFile(file_id string) types.File {
	v := url.Values{}
	v.Add("file_id", file_id)

	r := Get(b, "getFile", v)
	if !r.Ok {
		log.Fatal("You done goofed, API Res for getFile was not OK")
	}

	var f types.File
	json.Unmarshal(r.Result, &f)

	return f
}

// TODO: options here
// TODO: r.OK or unmarshal??
func (b Bot) AnswerCallbackQuery(callback_query_id string) bool {
	v := url.Values{}
	v.Add("callback_query_id", callback_query_id)

	r := Get(b, "answerCallbackQuery", v)
	if !r.Ok {
		log.Fatal("You done goofed, API Res for answerCallbackQuery was not OK")
	}

	var bb bool
	json.Unmarshal(r.Result, &bb)

	return bb
}

func (b Bot) Send(msg Sendable) (*types.Message, error) {
	return msg.send()
}