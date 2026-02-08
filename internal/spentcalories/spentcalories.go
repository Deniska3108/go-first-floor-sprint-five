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
		return 0, "", 0, errors.New("Не хватает элементов в слайсе")
	}

	Steps, err := strconv.Atoi(SplittedData[0])
	if err != nil {
		return 0, "", 0, err
	}
	if Steps == 0 {
		return 0, "", 0, errors.New("Количество шагов равно 0")
	}

	ActivityType := SplittedData[1]

	SpentTime, err := time.ParseDuration(SplittedData[2])
	if err != nil {
		return 0, "", 0, err
	}
	if SpentTime == 0 {
		return 0, "", 0, errors.New("Время равно 0")
	}

	return Steps, ActivityType, SpentTime, nil
}

func distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	return (height * stepLengthCoefficient * float64(steps)) / mInKm
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if duration <= 0 {
		return 0
	}

	return distance(steps, height) / duration.Hours()
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	// TODO: реализовать функцию
	Steps, ActivityType, SpentTime, err := parseTraining(data)

	if err != nil {
		log.Println(err)
		return "", err
	}

	switch ActivityType {
	case "Ходьба":
		Distance := distance(Steps, height)

		AverageSpeed := meanSpeed(Steps, height, SpentTime)

		Calories, err := WalkingSpentCalories(Steps, weight, height, SpentTime)
		if err != nil {
			return "", err
		}

		resultWalk := fmt.Sprintf(`Тип тренировки: %s.\nДлительность: %v ч.\nДистанция: %2.f км.\n
		Скорость: %2.f км/ч.\nСожгли калорий: %2.f\r`, ActivityType, SpentTime, Distance, AverageSpeed, Calories)

		return resultWalk, nil

	case "Бег":

		Distance := distance(Steps, height)

		AverageSpeed := meanSpeed(Steps, height, SpentTime)

		Calories, err := RunningSpentCalories(Steps, weight, height, SpentTime)
		if err != nil {
			return "", err
		}

		resultRun := fmt.Sprintf(`Тип тренировки: %s.\nДлительность: %v ч.\nДистанция: %2.f км.\n
		Скорость: %2.f км/ч.\nСожгли калорий: %2.f\r`, ActivityType, SpentTime, Distance, AverageSpeed, Calories)

		return resultRun, nil

	default:
		return "", errors.New("неизвестный тип тренировки")
	}
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps == 0 {
		return 0, errors.New("Количество шагов равно 0")
	}

	if duration <= 0 {
		return 0, errors.New("Время равно 0")
	}

	return (weight * meanSpeed(steps, height, duration) * duration.Minutes()) / minInH, nil

}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps == 0 {
		return 0, errors.New("Количество шагов равно 0")
	}

	if duration <= 0 {
		return 0, errors.New("Время равно 0")
	}

	calories := (weight * meanSpeed(steps, height, duration) * duration.Minutes()) / minInH

	return calories * walkingCaloriesCoefficient, nil
}
