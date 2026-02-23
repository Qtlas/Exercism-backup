package birdwatcher


// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
    var somme int = 0
    for _, val := range birdsPerDay {
        somme += val
    }
    return somme
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
    return TotalBirdCount(birdsPerDay[week*7-7:week*7])
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
    for i, _ := range birdsPerDay {
        birdsPerDay[i] = birdsPerDay[i] + (i+1)&1
    }
    return birdsPerDay
}
