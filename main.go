package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"acuvim-ii/acuvim"
)

func main() {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	now := time.Now().In(loc)
	timestamp := now.Format(time.RFC3339)

	pm, err := acuvim.Polling(6, "/dev/ttyUSB0")
	if err != nil {
		log.Fatalf("[ERROR] acuvim:%s", err.Error())
	}

	pm.Device_ID = "gon_001"
	pm.Timestamp = timestamp
	jw, errN := json.MarshalIndent(pm, " ", " ")
	if errN == nil {
		fmt.Println(string(jw))
	}

}
