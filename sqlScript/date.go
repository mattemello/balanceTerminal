package sqlScript

import "time"

type Movement struct {
	Money float32
	Tags  string
	Date  time.Time
}

/*func (move Movement) addMoney(num float32) {
	move.money = num
}
func (move Movement) addTags(tag string) {
	move.tags = tag
}
func (move Movement) addDay(day int) {
	move.day = day
}
func (move Movement) addMonth(month int) {
	move.month = month
}
func (move Movement) addYear(year int) {
	move.year = year
}*/
