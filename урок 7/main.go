package main

import (
	"fmt"
	"math"
)

func main() {

	// Задание 1
	bannerWidth := 12
	bannerHeight := 8

	bannerArea := bannerWidth * bannerHeight
	fmt.Println(bannerArea)

	halfBannerArea := bannerArea / 2
	fmt.Println(halfBannerArea)

	bannerBorderLength := (bannerWidth + bannerHeight) * 2
	fmt.Println(bannerBorderLength)

	// Задание 2
	boxCount := 29
	leftoverBoxes := boxCount % 5
	fmt.Println(leftoverBoxes)

	// Задание 3
	tempMorning := 18
	tempAfternoon := 25
	tempEvening := 20

	totalTemp := tempMorning + tempAfternoon + tempEvening
	averageTemp := totalTemp / 3
	fmt.Println(averageTemp)

	// Задание 4
	knownWords := 47
	wordsGoal := 120

	progressPercent := float64(knownWords) / float64(wordsGoal) * 100
	fmt.Println(progressPercent, "%")

	// Задание 5
	coins := 0

	coins += 500
	fmt.Println(coins)

	coins += 1200
	fmt.Println(coins)

	coins /= 2
	fmt.Println(coins)

	coins *= 2
	fmt.Println(coins)

	coins -= 300
	fmt.Println(coins)

	// Задание 6
	participants := 42
	groupCount := 8

	participantsPerGroup := participants / groupCount
	fmt.Println(participantsPerGroup)

	// Задание 7
	fmt.Println(20 - 4*3)
	fmt.Println((20 - 4) * 3)

	/*
		Результаты отличаются из-за порядка выполнения математических операций.
		Сначала выполняется умножение, поэтому 20 - 4 * 3 = 8.
		В скобках сначала выполняется вычитание,
		поэтому (20 - 4) * 3 = 48.
	*/

	// Задание 8
	squareValue := 81
	squareRoot := math.Sqrt(float64(squareValue))

	multiplier := 5
	exponent := 2
	powerResult := math.Pow(float64(multiplier), float64(exponent))

	fmt.Println(squareRoot)
	fmt.Println(powerResult)
}