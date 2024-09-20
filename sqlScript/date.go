package sqlScript

import "time"

type Movement struct {
	Money float32
	Tags  string
	Date  time.Time
}

type MovementRow struct {
	Id  int
	Mov Movement
}

type Money struct {
	Total  float32
	LastUp time.Time
}

type MoneyRow struct {
	Id     int
	RowMon Money
}

var TotalMoney Money

var Movements []MovementRow

func SaveMove(move Movement) {

	var mov MovementRow

	mov.Id = len(Movements) + 1
	mov.Mov = move

	Movements = append(Movements, mov)
	TotalMoney.Total -= mov.Mov.Money
	TotalMoney.LastUp = mov.Mov.Date
}

var TotalMoneys []MoneyRow

func SaveMoney(mon float32, tim time.Time) {
	var mov MoneyRow

	mov.Id = len(TotalMoneys) + 1
	mov.RowMon.Total = mon
	mov.RowMon.LastUp = tim

	TotalMoneys = append(TotalMoneys, mov)
	TotalMoney.Total += mon
	TotalMoney.LastUp = tim
}
