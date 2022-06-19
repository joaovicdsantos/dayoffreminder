package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/joaovicdsantos/dayoffreminder/database"
	"github.com/joaovicdsantos/dayoffreminder/model"
	"github.com/joaovicdsantos/dayoffreminder/utils"
	"gorm.io/gorm"
)

func main() {
	database.InitDatabase()
	db := database.DBConn

	var dayOffs []model.DayOff
	db.Order("initial_date").Where("initial_date::date - current_date in (15, 10, 7, 5, 3, 1) " +
		" or (initial_date::date <= current_date and final_date::date >= current_date)").Find(&dayOffs)

	if len(dayOffs) > 0 {
		today := todayFromDb(db)
		message := mountSlackMessage(dayOffs, today)

		var jsonStr = []byte(fmt.Sprintf(`{"text": "%s"}`, message))
		req, _ := http.NewRequest("POST", os.Getenv("SLACK_URL"), bytes.NewBuffer(jsonStr))
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		client.Do(req)
	}
}

func mountSlackMessage(dayOffs []model.DayOff, today time.Time) string {
	message := ":calendar:  *Calendário DayOffs*  :desert_island:\n\n"

	for _, dayOff := range dayOffs {
		daysSinceLeaving := dayOff.InitialDate.Sub(today).Hours() / 24
		daysToArrival := dayOff.FinalDate.Sub(today).Hours() / 24
		message = fmt.Sprintf("%s%s\n", message, generateMessage(dayOff, int(daysSinceLeaving), int(daysToArrival)))
	}

	return message
}

func generateMessage(dayOff model.DayOff, daysSinceLeaving, daysToArrival int) string {
	currentDirectory, _ := os.Getwd()
	jsonFile, err := os.Open(fmt.Sprintf("%s/schedule/messages.json", currentDirectory))
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()

	byteValueJSON, _ := ioutil.ReadAll(jsonFile)

	var message model.Messages
	json.Unmarshal(byteValueJSON, &message)

	var strMessage string
	if daysSinceLeaving > 7 {
		strMessage = message.GoingOut.Normal[random(0, len(message.GoingOut.Normal)-1)]
	} else if daysSinceLeaving <= 7 && daysSinceLeaving > 0 {
		strMessage = message.GoingOut.Week[random(0, len(message.GoingOut.Week)-1)]
	} else if daysSinceLeaving == 0 {
		strMessage = message.GoingOut.Today[random(0, len(message.GoingOut.Today)-1)]
	} else if daysToArrival > 7 {
		strMessage = message.GetIn.Normal[random(0, len(message.GoingOut.Normal)-1)]
	} else if daysToArrival <= 7 && daysToArrival > 0 {
		strMessage = message.GetIn.Week[random(0, len(message.GoingOut.Week)-1)]
	} else if daysToArrival == 0 {
		strMessage = message.GetIn.Today[random(0, len(message.GoingOut.Today)-1)]
	}

	strMessage = strings.Replace(strMessage, "<teamMember>", fmt.Sprintf("<@%s>", dayOff.TeamMemberId), -1)
	strMessage = strings.Replace(strMessage, "<initialDate>", dayOff.InitialDate.Format("02-01"), -1)
	strMessage = strings.Replace(strMessage, "<finalDate>", dayOff.FinalDate.Format("02-01"), -1)
	strMessage = strings.Replace(strMessage, "<daysSinceLeaving>", fmt.Sprint(daysSinceLeaving), -1)
	strMessage = strings.Replace(strMessage, "<daysToArrival>", fmt.Sprint(daysToArrival), -1)

	return strMessage
}

func random(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}

func todayFromDb(db *gorm.DB) time.Time {
	var result string
	db.Raw("SELECT current_date").Scan(&result)
	_, today := utils.ValidateDateFormatWithFormat(result[0:10], "2006-01-02")
	return today
}
