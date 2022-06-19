package model

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/joaovicdsantos/dayoffreminder/utils"
	"gorm.io/gorm"
)

type DayOff struct {
	Id           string    `json:"id" gorm:"not null"`
	InitialDate  time.Time `json:"initial_date" gorm:"not null"`
	FinalDate    time.Time `json:"final_date" gorm:"not null"`
	TeamMember   string    `json:"team_member" gorm:"not null"`
	TeamMemberId string    `json:"team_member_id" gorm:"not null"`
}

func (do *DayOff) BeforeCreate(tx *gorm.DB) (err error) {
	do.Id = uuid.New().String()
	return
}

func (do *DayOff) SlackRequestToDayOff(slackRequest SlackRequest) error {
	text := slackRequest.Text
	teamMember := slackRequest.UserName
	teamMemberId := slackRequest.UserId

	match, _ := regexp.MatchString("^[0-9]{2}-[0-9]{2}-[0-9]{4} [0-9]{2}-[0-9]{2}-[0-9]{4}$", text)
	if !match {
		return fmt.Errorf("invalid format! DD-MM-YYYY DD-MM-YYYY")
	}

	var dates []time.Time
	stringDates := strings.Split(text, " ")
	for _, element := range stringDates {
		isDate, date := utils.ValidateDateFormat(element)
		if !isDate {
			return fmt.Errorf("invalid format! DD-MM-YYYY DD-MM-YYYY")
		}
		dates = append(dates, date)
	}

	if dates[0].Before(time.Now()) || dates[0].After(dates[1]) {
		return fmt.Errorf("start date greater than end date or current date greater than start date")
	}

	*do = DayOff{
		InitialDate:  dates[0],
		FinalDate:    dates[1],
		TeamMember:   teamMember,
		TeamMemberId: teamMemberId,
	}

	return nil
}
