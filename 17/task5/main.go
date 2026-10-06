package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Account struct {
	Login    string  `json:"login"`
	Password string  `json:"-"` // completely ignores it
	Balance  float64 `json:"balance"`
}

func main() {
	accounts := []Account{
		{"Ol1_vka", "djcse!-29", 566.50},
		{"Sergei_Sokolov", "kSdkjOWd_230911", 11223.50},
		{"Mr_Beast", "#zxczsdafxb))jw!&&w112!#", 112345674.50},
	}

	data, err := json.MarshalIndent(accounts, "", "  ")
	if err != nil {
		log.Println("An error occurred: ", err)
	}
	fmt.Println(string(data))
}
