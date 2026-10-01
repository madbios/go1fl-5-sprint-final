package spentenergy

import (
	"errors"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if height <= 0 {
		return 0, errors.New("некорректное значение роста")
	}
	if steps <= 0 {
		return 0, errors.New("некорректное значение шагов")
	}
	if weight <= 0 {
		return 0, errors.New("некорректное значение веса")
	}
	if duration <= 0 {
		return 0, errors.New("некорректное значение длительности")
	}
	meanSpd := MeanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	walkSpndClrs := ((weight * meanSpd * durationInMinutes) / 60) * walkingCaloriesCoefficient
	return walkSpndClrs, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if height <= 0 {
		return 0, errors.New("некорректное значение роста")
	}
	if steps <= 0 {
		return 0, errors.New("некорректное значение шагов")
	}
	if weight <= 0 {
		return 0, errors.New("некорректное значение веса")
	}
	if duration <= 0 {
		return 0, errors.New("некорректное значение длительности")
	}
	meanSpd := MeanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	Cal := (weight * meanSpd * durationInMinutes) / minInH
	return Cal, nil

}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if height <= 0 {
		return 0
	}
	if steps <= 0 {
		return 0
	}
	if duration <= 0 {
		return 0
	}
	dist := Distance(steps, height)
	meanSpd := dist / duration.Hours()
	return meanSpd
}

func Distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	if height <= 0 {
		return 0
	}
	if steps <= 0 {
		return 0
	}
	distStep := height * stepLengthCoefficient
	dist := (distStep * float64(steps)) / mInKm
	return dist
}
