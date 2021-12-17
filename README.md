# Golang JSON Module with `lowerCamelCase`

Base on [official JSON package](https://github.com/golang/go/tree/master/src/encoding/json), auto `lowerCamelCase` struct field if no json tag defined.
No effect on map, only for struct.
More detail could be found at `func lowerCamelCase` in [encode.go](encode.go).

## Usage

```go
package main

import (
	"fmt"

	"github.com/xgomod/json"
)

type User struct {
	ID         int    // -> id
	UUID       int    // -> uuid
	Name       string // -> name
	NickName   string `json:"Nickname"` // -> Nickname
	Birthday   string `json:"-"`
	Age        int    // -> age
	Email      string `json:"e-mail"` // -> e-mail
	QQNumber1  string // -> qqNumber1
	QQNumber2  string `json:"QQNumber2"` // -> QQNumber2
	IPv4       string // -> ipv4
	IPV4       string // -> ipv4
	IPVersion4 string // -> ipVersion4
}

func main() {
	user := User{
		ID:         1,
		UUID:       2,
		Name:       "Machael Jackson",
		NickName:   "MJ",
		Birthday:   "1958-08-29",
		Age:        51,
		Email:      "mj@qq.com",
		QQNumber1:  "19580829",
		QQNumber2:  "20090625",
		IPv4:       "192.168.1.1",
		IPV4:       "192.168.1.2",
		IPVersion4: "192.168.1.3",
	}

	b, _ := json.MarshalIndent(user, "", "  ")
	fmt.Println(string(b))

	jsonStr := `{
		"UUid": 2,
		"Id": 1,
		"Birthday": "1958-08-29",
		"nickname": "MJ",
		"E-mail": "mj@qq.com",
		"age": 51,
		"name": "Machael Jackson",
		"qQNumber2": "20090625",
		"QqNumBer1": "19580829",
		"IPVersion4": "192.168.1.3",
		"IPV4": "192.168.1.2",
		"ipv4": "192.168.1.1"
	}`
	var user2 User
	json.Unmarshal([]byte(jsonStr), &user2)
	fmt.Println(user2)
}
```

[example source code](example/main.go)