package main

import "fmt"

type Food struct {
	rating int
	name string
}


func Init(foods, cuisines, ratings []string) map[string][]Food {
	cuisine_food := make(map[string][]Food)
	for i, name := range foods {
		ratingAndName := Food{ratings[i], name}
		cuisine_food[cuisines[i]] = append(cuisine_food[cuisines[i]], ratingAndFood)
	}

	return cuisine_food
}


//["FoodRatings", "highestRated", "highestRated", "changeRating", "highestRated", "changeRating", "highestRated"]
//[[["kimchi", "miso", "sushi", "moussaka", "ramen", "bulgogi"], ["korean", "japanese", "japanese", "greek", "japanese", "korean"], [9, 12, 8, 15, 14, 7]], ["korean"], ["japanese"], ["sushi", 16], ["japanese"], ["ramen", 16], ["japanese"]]


func main() {

}