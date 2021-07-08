func phoneCall(min1 int, min2_10 int, min11 int, s int) int {
    time := 0
    for s > 0 {
        if time >= 10 && s - min11 >= 0 {
            time += 1
            s -= min11
        } else if time >= 1 && time < 10 && s - min2_10 >= 0 {
            time += 1
            s -= min2_10
        } else if time == 0 && s - min1 >= 0 {
            time += 1
            s -= min1
        } else {
            return time
        }
    }
    return time