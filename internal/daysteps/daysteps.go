package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	personaldata.Personal
	Steps    int
	Duration time.Duration
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	if datastring == "" {
		return errors.New("empty income data")
	}

	items := strings.Split(datastring, ",")
	if len(items) != 2 {
		return errors.New("bad income data-item count")
	}

	steps, dur := items[0], items[1]

	// steps
	stepsNum, err := strconv.Atoi(steps)
	if err != nil {
		return errors.New("bad income steps count")
	}
	if stepsNum <= 0 {
		return errors.New("steps count <= 0")
	}
	ds.Steps = stepsNum

	// duration
	durTime, err := time.ParseDuration(dur)
	if err != nil {
		return errors.New("bad income duration")
	}
	if durTime <= 0 {
		return errors.New("duration <= 0")
	}
	ds.Duration = durTime

	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	var calories float64
	var err error

	distance := spentenergy.Distance(ds.Steps, ds.Personal.Height)

	calories, err = spentenergy.WalkingSpentCalories(ds.Steps, ds.Personal.Weight, ds.Personal.Height, ds.Duration)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Количество шагов: %v.\n"+
		"Дистанция составила %.2f км.\n"+
		"Вы сожгли %.2f ккал.\n",
		ds.Steps, distance, calories), nil
}
