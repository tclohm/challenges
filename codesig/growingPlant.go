import . "math"
func growingPlant(upSpeed int, downSpeed int, desiredHeight int) int {
    if desiredHeight <= upSpeed {
        return 1
    }
    res := float64(desiredHeight - upSpeed) / float64(upSpeed - downSpeed) + 1
    return int(Ceil(res))
}
