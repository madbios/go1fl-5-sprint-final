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
	// TODO: добавить поля
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	parts := strings.Split(datastring, ",")
	if len(parts) != 3 {
		return errors.New("Ошибка разбиения на слайс")
	}
	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return errors.New("Ошибка первого числа(некорректное)")
	}
	t.Steps = steps
	if parts[1] != "Ходьба" && parts[1] != "Бег" {
		return errors.New("Неизвестный тип тренировки")
	}
	t.TrainingType = parts[1]
	duration, err := time.ParseDuration(parts[2])
	if err != nil {
		return errors.New("Ошибка преобразования времени")
	}
	if duration >= 24*time.Hour || duration <= 0 {
		return errors.New("Ошибка преобразования времени")
	}
	t.Duration = duration
	return nil
}

func (t Training) ActionInfo() (string, error) {
	var err error
	var calories float64
	Distance := spentenergy.Distance(t.Steps, t.Height)
	meansSpeed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)
	switch t.TrainingType {
	case "Бег":
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	case "Ходьба":
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	default:
		return "", errors.New("неизвестный тип тренировки")
	}
	if err != nil {
		return "", fmt.Errorf("ошибка расчёта калорий: %w", err)
	}
	result := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nCожгли калорий: %.2f", t.TrainingType, t.Duration.Hours(), Distance, meansSpeed, calories)
	return result, nil
}
