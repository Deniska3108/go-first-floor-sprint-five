package spentcalories

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

func parseTraining(data string) (int, string, time.Duration, error) {
	// TODO: реализовать функцию
	SplittedData := strings.Split(data, ",")
	if len(SplittedData) != 3 {
		return 0, "", 0, errors.New("не хватает элементов в слайсе" + "\n")
	}

	Steps, err := strconv.Atoi(SplittedData[0])
	if err != nil {
		return 0, "", 0, err
	}
	if Steps == 0 {
		return 0, "", 0, errors.New("количество шагов равно 0" + "\n")
	}

	ActivityType := SplittedData[1]

	Duration, err := time.ParseDuration(SplittedData[2])
	if err != nil {
		return 0, "", 0, err
	}
	if Duration == 0 {
		return 0, "", 0, errors.New("время равно 0" + "\n")
	}

	return Steps, ActivityType, Duration, nil
}

func distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	return (height * stepLengthCoefficient * float64(steps)) / mInKm
}

func meanSpeed(steps int, height float64, Duration time.Duration) float64 {
	// TODO: реализовать функцию
	if Duration <= 0 {
		return 0
	}

	return distance(steps, height) / Duration.Hours()
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	// TODO: реализовать функцию
	Steps, ActivityType, Duration, err := parseTraining(data)

	if err != nil {
		log.Println(err)
		return "", err
	}

	switch ActivityType {
	case "Ходьба":
		Distance := distance(Steps, height)

		AverageSpeed := meanSpeed(Steps, height, Duration)

		Calories, err := WalkingSpentCalories(Steps, weight, height, Duration)
		if err != nil {
			return "", err
		}

		resultWalk := fmt.Sprintf(`Тип тренировки: %s.
Длительность: %v ч.
Дистанция: %.2f км.
Скорость: %.2f км/ч.
Сожгли калорий: %2.f`+"\n", ActivityType, Duration.Hours(), Distance, AverageSpeed, Calories)

		return resultWalk, nil

	case "Бег":

		Distance := distance(Steps, height)

		AverageSpeed := meanSpeed(Steps, height, Duration)

		Calories, err := RunningSpentCalories(Steps, weight, height, Duration)
		if err != nil {
			return "", err
		}

		resultRun := fmt.Sprintf(`Тип тренировки: %s.
Длительность: %.2f ч.
Дистанция: %.2f км.
Скорость: %.2f км/ч.
Сожгли калорий: %.2f`+"\n", ActivityType, Duration.Hours(), Distance, AverageSpeed, Calories)

		return resultRun, nil

	default:
		return "", errors.New("неизвестный тип тренировки" + "\n")
	}
}

func RunningSpentCalories(steps int, weight, height float64, Duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps == 0 {
		return 0, errors.New("количество шагов равно 0" + "\n")
	}

	if Duration <= 0 {
		return 0, errors.New("время равно 0" + "\n")
	}

	return (weight * meanSpeed(steps, height, Duration) * Duration.Minutes()) / minInH, nil

}

func WalkingSpentCalories(steps int, weight, height float64, Duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps == 0 {
		return 0, errors.New("количество шагов равно 0" + "\n")
	}

	if Duration <= 0 {
		return 0, errors.New("время равно 0" + "\n")
	}

	calories := (weight * meanSpeed(steps, height, Duration) * Duration.Minutes()) / minInH

	return calories * walkingCaloriesCoefficient, nil
}
