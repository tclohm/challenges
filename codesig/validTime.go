func validTime(time string) bool {
    minhour, maxhour := "00", "23"
    minmin, maxmin := "00", "60"
    
    hour := time[0:2]
    minute := time[3:]
    
    if (minmin <= minute && minute < maxmin) && (minhour <= hour && hour <= maxhour) {
        return true
    }
    
    return false
}