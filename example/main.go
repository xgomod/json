package main

import (
	"fmt"
	"json"
)

type User struct {
	ID        int `json:"-"`
	Name      string
	NickName  string `json:"Nickname"`
	Age       int
	Email     string `json:"e-mail"`
	QQNumber1 string
	QQNumber2 string `json:"qqNumber2"`
}

func main() {
	user := User{
		ID:        1,
		Name:      "Machael Jackson",
		NickName:  "MJ",
		Age:       51,
		Email:     "mj@qq.com",
		QQNumber1: "19580829",
		QQNumber2: "20090625",
	}

	b, _ := json.Marshal(user)
	fmt.Println(string(b)) // {"name":"Machael Jackson","Nickname":"MJ","age":51,"e-mail":"mj@qq.com","QQNumber1":"19580829","qqNumber2":"20090625"}

	jsonStr := `{"Name":"Machael Jackson","nickname":"MJ","age":51,"E-mAil":"mj@qq.com","QqNumBer1":"19580829","qqNumbeR2":"20090625"}`
	var user2 User
	json.Unmarshal([]byte(jsonStr), &user2)
	fmt.Println(user2) // {0 Machael Jackson MJ 51 mj@qq.com 19580829 20090625}
}
