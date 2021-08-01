func prod(n int) int {
    product := 1
    for n != 0 {
        product *= n % 10
        n /= 10
    }
    return product
}
func digitsProduct(product int) int {
    // find the smallest positive product whose digit equal to product
    for i := 1 ; i <= 5000 ; i++ {
        if prod(i) == product {
            return i
        }
    }
    return -1
}