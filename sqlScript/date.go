package sqlScript

import "time"

type Movement struct {
	Money float32
	Tags  string
	Date  time.Time
	Add   bool
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

var AllTags []string

var TotalMoney Money

var Movements []MovementRow

//TO-DO make it a map?

func SaveMove(move Movement) {

	var mov MovementRow

	mov.Id = len(Movements) + 1
	mov.Mov = move

	Movements = append(Movements, mov)
	if mov.Mov.Add {
		TotalMoney.Total += mov.Mov.Money
	} else {
		TotalMoney.Total -= mov.Mov.Money
	}
	TotalMoney.LastUp = time.Now()

	SaveMoneyDB(TotalMoney.Total, TotalMoney.LastUp)
	TotalMoneys = append(TotalMoneys, MoneyRow{len(TotalMoneys) + 1, TotalMoney})
}

func SaveTag(s string) {
	AllTags = append(AllTags, s)
}

var TotalMoneys []MoneyRow
