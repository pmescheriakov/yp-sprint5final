package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	personaldata.Personal
	Steps        int
	TrainingType string
	Duration     time.Duration
}

// Parse
//
//	Метод парсит строку с данными формата "3456,Ходьба,3h00m" и записывает в поля структуры Training.
func (t *Training) Parse(datastring string) (err error) {
	if datastring == "" {
		return errors.New("empty income data")
	}

	items := strings.Split(datastring, ",")
	if len(items) != 3 {
		return errors.New("bad income data-item count")
	}

	steps, activity, dur := items[0], items[1], items[2]

	// steps
	stepsNum, err := strconv.Atoi(steps)
	if err != nil {
		return errors.New("bad income steps count")
	}
	if stepsNum <= 0 {
		return errors.New("steps count <= 0")
	}
	t.Steps = stepsNum

	// activity
	switch activity {
	case "Ходьба":
		t.TrainingType = activity
	case "Бег":
		t.TrainingType = activity
	}

	// duration
	durTime, err := time.ParseDuration(dur)
	if err != nil {
		return errors.New("bad income duration")
	}
	if durTime <= 0 {
		return errors.New("duration <= 0")
	}
	t.Duration = durTime

	return nil
}

func (t Training) ActionInfo() (string, error) {
	var calories float64
	var err error

	distance := spentenergy.Distance(t.Steps, t.Personal.Height)
	speed := spentenergy.MeanSpeed(t.Steps, t.Personal.Height, t.Duration)

	switch t.TrainingType {
	case "Ходьба":
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration)
		if err != nil {
			return "", err
		}
	case "Бег":
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration)
		if err != nil {
			return "", err
		}
	default:
		return "", errors.New("неизвестный тип тренировки")
	}

	return fmt.Sprintf("Тип тренировки: %v\n"+
		"Длительность: %.2f ч.\n"+
		"Дистанция: %.2f км.\n"+
		"Скорость: %.2f км/ч\n"+
		"Сожгли калорий: %.2f\n",
		t.TrainingType, t.Duration.Hours(), distance, speed, calories), nil
}
