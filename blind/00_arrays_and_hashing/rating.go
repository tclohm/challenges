package main

import "fmt"

type Food struct {
	rating int
	name string
}

type FoodRatings struct {
	system map[string][]Food
}


func Init(foods, cuisines []string, ratings []int) FoodRatings {
	cuisine_food := FoodRatings{make(map[string][]Food)}
	for i, name := range foods {
		ratingAndName := Food{ratings[i], name}
		cuisine_food.system[cuisines[i]] = append(cuisine_food.system[cuisines[i]], ratingAndName)
	}

	return cuisine_food
}

func (this *FoodRatings) ChangeRating(food string, rating int) {
	var index int
	var genre string
	for key, values := range this.system {
		for i, item := range values {
			if item.name == food {
				index = i
				genre = key
			}
		}
	}
	
	this.system[genre][index].rating = rating
}

func (this *FoodRatings) HighestRating(cuisine string) string {
	var highest Food
	if foods, ok := this.system[cuisine]; ok {
		highest = foods[0]
		for _, food := range foods {
			if food.rating > highest.rating {
				highest = food
			}
		}
	}
	return highest.name
}


//["FoodRatings", "highestRated", "highestRated", "changeRating", "highestRated", "changeRating", "highestRated"]
//[[["kimchi", "miso", "sushi", "moussaka", "ramen", "bulgogi"], ["korean", "japanese", "japanese", "greek", "japanese", "korean"], [9, 12, 8, 15, 14, 7]], ["korean"], ["japanese"], ["sushi", 16], ["japanese"], ["ramen", 16], ["japanese"]]
func main() {
	foods := []string{"kimchi", "miso", "sushi", "moussaka", "ramen", "bulgogi"}
	genres := []string{"korean", "japanese", "japanese", "greek", "japanese", "korean"}
	ratings := []int{9, 12, 8, 15, 14, 7}

	Rating := Init(foods, genres, ratings)
	//fmt.Println(Rating.system["japanese"])
	fmt.Println(Rating.HighestRating("korean"))
	fmt.Println(Rating.HighestRating("japanese"))
	Rating.ChangeRating("sushi", 16)
	fmt.Println(Rating.HighestRating("japanese"))
	fmt.Println(Rating.HighestRating("greek"))

}