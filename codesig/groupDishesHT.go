import "sort"

func groupingDishes(dishes [][]string) [][]string {
  ht := map[string][]string{}
  for _, dish := range dishes {
    name := dish[0]
    for _, ingredient := range dish[1:] {
      ht[ingredient] = append(ht[ingredient], name)
    }
  }
  
  result := [][]string{}
  
  for ingredient, dishes := range ht {
    if len(dishes) >= 2 {
      sort.Strings(dishes)
      result = append(result, append([]string{ingredient}, dishes...))      
    }
  }
  // nice to know look into our array use an anonymous function to sort
  sort.Slice(result, func(i, j int) bool {
    return result[i][0] < result[j][0]
  })
  return result
}