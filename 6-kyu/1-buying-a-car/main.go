package main

import (
	"fmt"
	"math"
)

func main() {
	// Best Practices
	var priceOldCar float64 = 2000
	var priceNewCar float64 = 8000
	var percentage = 1.5
	saverate := 1000
	months := 0

	for ; (priceOldCar + float64(months*saverate)) < priceNewCar; months++ {
		if months%2 == 1 {
			percentage += 0.5
		}
		priceOldCar -= priceOldCar * (percentage / 100)
		priceNewCar -= priceNewCar * (percentage / 100)
	}
	fmt.Printf("priceNewCar:%v saveing:%v priceOldCar:%v result:%v\n",
		priceNewCar, priceOldCar, months*saverate,
		int(priceOldCar+float64(months*saverate)-priceNewCar))
}

// NbMonths :
func NbMonths(startPriceOld, startPriceNew, savingperMonth int, percentLossByMonth float64) [2]int {
	result := [2]int{}
	var priceOldCar = float64(startPriceOld)
	var priceNewCar = float64(startPriceNew)
	var saverate = float64(savingperMonth)

	var months int
	var saving float64
	percentLossByMonth = percentLossByMonth / 100

	for (priceOldCar+saving)-priceNewCar < 0 {
		months++
		if months%2 == 0 {
			percentLossByMonth += 0.005
		}
		saving += saverate
		priceOldCar -= priceOldCar * percentLossByMonth
		priceNewCar -= priceNewCar * percentLossByMonth
	}

	result[0] = months
	result[1] = int(math.Round((priceOldCar + saving) - priceNewCar))

	return result
}

// NbMonthsBestPractices :
func NbMonthsBestPractices(startPriceOld, startPriceNew, savingperMonth int, percentLossByMonth float64) [2]int {
	months := 0
	priceOld := float64(startPriceOld)
	priceNew := float64(startPriceNew)

	for ; priceOld+float64(months*savingperMonth) < priceNew; months++ {
		if months%2 == 1 {
			percentLossByMonth += 0.5
		}
		priceOld -= priceOld * percentLossByMonth / 100.0
		priceNew -= priceNew * percentLossByMonth / 100.0
	}

	return [2]int{months, int(priceOld + float64(months*savingperMonth) - priceNew + 0.5)}
}
