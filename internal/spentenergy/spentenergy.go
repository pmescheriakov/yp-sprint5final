package spentenergy

import (
	"errors"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // Количество метров в километре.
	minInH                     = 60   // Количество минут в часе.
	stepLengthCoefficient      = 0.45 // Коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // Коэффициент для расчета калорий при ходьбе.
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, errors.New("steps count <= 0")
	}
	if duration <= 0. {
		return 0, errors.New("duration <= 0")
	}
	if weight <= 0. {
		return 0, errors.New("weight <= 0")
	}
	if height <= 0. {
		return 0, errors.New("height <= 0")
	}

	return walkingCaloriesCoefficient * (duration.Minutes() * weight * MeanSpeed(steps, height, duration) / float64(minInH)), nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, errors.New("steps count <= 0")
	}
	if duration <= 0 {
		return 0, errors.New("duration <= 0")
	}
	if weight <= 0. {
		return 0, errors.New("weight <= 0")
	}
	if height <= 0. {
		return 0, errors.New("height <= 0")
	}

	return duration.Minutes() * weight * MeanSpeed(steps, height, duration) / float64(minInH), nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if steps <= 0 || duration <= 0 {
		return 0.
	}

	return Distance(steps, height) / duration.Hours()
}

func Distance(steps int, height float64) float64 {
	return float64(steps) * height * stepLengthCoefficient / float64(mInKm)
}
