func depositProfit(deposit int, rate int, threshold int) int {
    // how long will it take for your balance to pass the threshold
    var years int = 0
    var futuredeposit float64 = float64(deposit)
    for futuredeposit < float64(threshold) {
        interest := futuredeposit * (float64(rate) / 100.0)
        futuredeposit += interest
        years++
    }
    return years
