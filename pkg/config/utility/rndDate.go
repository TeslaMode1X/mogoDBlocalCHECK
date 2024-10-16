package utility

import (
	"math/rand"
	"time"
)

func GenerateRandomDate() string {
	rand.Seed(time.Now().UnixNano())

	now := time.Now()

	daysToAdd := rand.Intn(30) + 1
	hoursToAdd := rand.Intn(3) + 1
	minutesToAdd := rand.Intn(59) + 1

	randDate := now.AddDate(0, 0, daysToAdd).
		Add(time.Duration(hoursToAdd) * time.Hour).
		Add(time.Duration(minutesToAdd) * time.Minute)

	return randDate.Format("02.01.2006, 15:04")
}
