package main

import (
	"os"
	"strconv"
	"strings"

	"time"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

func main() {

	err := checkEnvVariables()
	if err != nil {
		log.WithError(err).Error("Environment variables were not set")
		return
	}

	err = postReminder()
	if err != nil {
		log.WithError(err).Error("Failed to post a reminder")
		return
	}
}

func checkEnvVariables() error {
	var envVariables = []string{
		"MattermostNotificationsHook",
		"EndOfDayHour",
		"EndOfDayMinute",
		"MinutesBeforeEndToNotify",
		"Message",
		"Description",
		"DaysToNotify",
	}

	for _, envVar := range envVariables {
		if os.Getenv(envVar) == "" {
			return errors.Errorf("Environment variable %s was not set", envVar)
		}
	}

	return nil
}

func postReminder() error {
	locNASA, err := time.LoadLocation("America/New_York")
	if err != nil {
		return err
	}
	nowNASA := time.Now().In(locNASA)

	locAPAC, err := time.LoadLocation("Australia/Sydney")
	if err != nil {
		return err
	}
	nowAPAC := time.Now().In(locAPAC)

	locEMEA, err := time.LoadLocation("Europe/Athens")
	if err != nil {
		return err
	}
	nowEMEA := time.Now().In(locEMEA)

	endOfDayHour, err := strconv.Atoi(os.Getenv("EndOfDayHour"))
	if err != nil {
		return errors.Wrap(err, "failed to parse integer from EndOfDayHour string")
	}

	endOfDayMinute, err := strconv.Atoi(os.Getenv("EndOfDayMinute"))
	if err != nil {
		return errors.Wrap(err, "failed to parse integer from EndOfDayMinute string")
	}

	minutesBeforeEndToNotify, err := strconv.Atoi(os.Getenv("MinutesBeforeEndToNotify"))
	if err != nil {
		return errors.Wrap(err, "failed to parse integer from MinutesBeforeEndToNotify string")
	}

	daysToNotify := strings.Split(os.Getenv("DaysToNotify"), ",")

	log.Info("Checking if any reminder should be sent")

	//It will notify in the last 15minutes before end of day
	if (endOfDayHour-nowNASA.Hour()) < 1 && (endOfDayHour-nowNASA.Hour()) >= 0 && (endOfDayMinute-nowNASA.Minute()) <= minutesBeforeEndToNotify && (endOfDayMinute-nowNASA.Minute()) > 0 {
		for _, day := range daysToNotify {
			if day == nowNASA.Weekday().String() {
				if os.Getenv("NASATimeZoneUsers") != "" {
					log.Info("Its time to remind the NASA users")
					err = sendMattermostNotification(os.Getenv("Message"), os.Getenv("NASATimeZoneUsers"), os.Getenv("Description"))
					if err != nil {
						return errors.Wrap(err, "failed to send Mattermost notification")
					}
				}
			}
		}
	}

	if (endOfDayHour-nowAPAC.Hour()) < 1 && (endOfDayHour-nowAPAC.Hour()) >= 0 && (endOfDayMinute-nowAPAC.Minute()) <= minutesBeforeEndToNotify && (endOfDayMinute-nowAPAC.Minute()) > 0 {
		for _, day := range daysToNotify {
			if day == nowAPAC.Weekday().String() {
				if os.Getenv("APACTimezoneUsers") != "" {
					log.Info("Its time to remind the APAC users")
					err = sendMattermostNotification(os.Getenv("Message"), os.Getenv("APACTimezoneUsers"), os.Getenv("Description"))
					if err != nil {
						return errors.Wrap(err, "failed to send Mattermost notification")
					}
				}
			}
		}
	}

	if (endOfDayHour-nowEMEA.Hour()) < 1 && (endOfDayHour-nowEMEA.Hour()) >= 0 && (endOfDayMinute-nowEMEA.Minute()) <= minutesBeforeEndToNotify && (endOfDayMinute-nowEMEA.Minute()) > 0 {
		for _, day := range daysToNotify {
			if day == nowEMEA.Weekday().String() {
				if os.Getenv("EMEATimezoneUsers") != "" {
					log.Info("Its time to remind the EMEA users")
					err = sendMattermostNotification(os.Getenv("Message"), os.Getenv("EMEATimezoneUsers"), os.Getenv("Description"))
					if err != nil {
						return errors.Wrap(err, "failed to send Mattermost notification")
					}
				}
			}
		}
	}
	return nil
}
