package main

import "fmt"

func bruteForceDaily(temps []int) []int {
	length := len(temps)
	answer := make([]int, length, length)
	for day, _ := range temps {
		for futureDay := 1 ; futureDay < length ; futureDay++ {
			if temps[futureDay] > temps[day] {
				answer[day] = futureDay - day
				break
			}
		}
	}
	return answer
}

func daily(temps []int) []int {
    length := len(temps)
    hottest := 0
    answer := make([]int, length, length)

    for currentDay := length - 1 ; currentDay > 0 ; currentDay-- {
    	currentTemp := temps[currentDay]
    	if currentTemp >= hottest {
    		hottest = currentTemp
    		continue
    	}

    	days := 1

    	for temps[currentDay + days] <= currentTemp {
    		days += answer[currentDay + days]
    	}
    	answer[currentDay] = days
    }
    return answer
}

func main() {
	t1 := []int{73,74,75,71,69,72,76,73}
	t2 := []int{30,40,50,60}
	t3 := []int{30,60,90}

	fmt.Println(daily(t1))
	fmt.Println(daily(t2))
	fmt.Println(daily(t3))
}