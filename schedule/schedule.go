package main

import (
	"bytes"
	"fmt"
	"math"
	"net/http"
	"os"
	"time"

	"github.com/joaovicdsantos/dayoffreminder/database"
	"github.com/joaovicdsantos/dayoffreminder/model"
)

func main() {
	database.InitDatabase()
	db := database.DBConn

	var dayOffs []model.DayOff
	db.Order("initial_date").Where("(initial_date::date - current_date in (14, 9, 6, 4, 2, 0) or initial_date::date < current_date) and final_date::date > current_date").Find(&dayOffs)

	message := MountSlackMessage(dayOffs)

	var jsonStr = []byte(fmt.Sprintf(`{"text": "%s"}`, message))
	req, _ := http.NewRequest("POST", os.Getenv("SLACK_URL"), bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	client.Do(req)
}

func MountSlackMessage(dayOffs []model.DayOff) string {
	message := ":calendar:  *Calendário DayOffs*  :desert_island:\n\n"

	for _, dayOff := range dayOffs {
		days := math.Round(time.Until(dayOff.InitialDate).Hours()/24) + 1
		if days > 0 {
			message = fmt.Sprintf("%s*%s* estará fora em *%v dia(s)* e voltará dia `%v`.\n",
				message, dayOff.TeamMember, days, dayOff.FinalDate.Format("02-01"))
		} else {
			message = fmt.Sprintf("%s*%s* está fora e voltará dia `%v`.\n",
				message, dayOff.TeamMember, dayOff.FinalDate.Format("02-01"))
		}
	}

	return message
}
