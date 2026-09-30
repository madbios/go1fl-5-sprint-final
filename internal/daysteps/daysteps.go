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
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	parts := strings.Split(datastring, ",")
	if len(parts) != 2 {
		return errors.New("Ошибка разделения")
	}
	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return errors.New("Ошибка преобразования числа")
	}
	if steps <= 0 {
		return errors.New("Ошибка количества шагов")
	}
	ds.Steps = steps
	duration, err := time.ParseDuration(parts[1])
	if err != nil {
		return errors.New("ошибка преобразования времени")
	}
	if duration >= 24*time.Hour || duration <= 0 {
		return errors.New("Ошибка времени")
	}
	ds.Duration = duration
	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	Distance := spentenergy.Distance(ds.Steps, ds.Height)

	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", errors.New("Ошибка подсчета каллорий")
	}
	pr := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", ds.Steps, Distance, calories)
	return pr, nil
}
