package daysteps

import (
	"errors"
	"fmt"
	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
	"log"
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
		return 0, 0, errors.New("не хватает элементов в слайсе")
	}

	Steps, err := strconv.Atoi(SplittedData[0])
	if err != nil {
		return 0, 0, err
	}
	if Steps <= 0 {
		return 0, 0, errors.New("количество шагов меньше или равно 0")
	}

	Duration, err := time.ParseDuration(SplittedData[1])
	if err != nil {
		return 0, 0, err
	}
	if Duration <= 0 {
		return 0, 0, errors.New("время меньше или равно 0")
	}

	return Steps, Duration, nil
}

func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию
	Steps, Duration, err := parsePackage(data)
	if err != nil {
		log.Println(err)
		return ""
	}
	if Steps <= 0 {
		return ""
	}

	var Distance float64 = (float64(Steps) * stepLength) / float64(mInKm)

	Calories, err := spentcalories.WalkingSpentCalories(Steps, weight, height, Duration)

	if err != nil {
		log.Println(err)
		return ""
	}

	return fmt.Sprintf(`Количество шагов: %d.
Дистанция составила %.2f км.
Вы сожгли %.2f ккал.`+"\n", Steps, Distance, Calories)
}
