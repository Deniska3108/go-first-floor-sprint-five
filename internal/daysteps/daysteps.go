package daysteps

import (
	"errors"
	"fmt"
	"go-first-floor-sprint-five/internal/spentcalories"
	"strconv"
	"strings"
	"time"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	// TODO: реализовать функцию
	SplittedData := strings.Split(data, ",")
	if len(SplittedData) != 2 {
		return 0, 0, errors.New("Не хватает элементов в слайсе")
	}

	Steps, err := strconv.Atoi(SplittedData[0])
	if err != nil {
		return 0, 0, err
	}
	if Steps == 0 {
		return 0, 0, errors.New("Количество шагов равно 0")
	}

	WalkTime, err := time.ParseDuration(SplittedData[1])
	if err != nil {
		return 0, 0, err
	}
	if WalkTime == 0 {
		return 0, 0, errors.New("Время равно 0")
	}

	return Steps, WalkTime, nil
}

func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию
	Steps, WalkTime, err := parsePackage(data)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	if Steps == 0 {
		return ""
	}

	var Distance float64 = (float64(Steps) * stepLength) / float64(mInKm)

	Calories := WalkingSpentCalories(weight, height, WalkTime)

	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %2.f км.\nВы сожгли %2.f ккал.\r", Steps, Distance, Calories)
}
