package sqlScript

import (
	"context"
	"database/sql"
	"strconv"
	"time"

	"github.com/mattemello/balanceTerminal/errorHand"
)

var db *sql.DB

func CreationTable() *sql.DB {
	var err error
	db, err = sql.Open("sqlite3", "money.db")
	errorhand.HandlerError(err, errorhand.TakeFileLine()+"error in the opening of the db")

	queryCreationSpending := "CREATE TABLE IF NOT EXISTS spendingMoney( id_transition INTEGER PRIMARY KEY, valur_tansaction REAL, tags TEXT, date TEXT, toadd INTEGER);"

	_, err = db.Exec(queryCreationSpending)
	errorhand.HandlerError(err, errorhand.TakeFileLine()+" problem with the query")

	queryCreationMoney := "CREATE TABLE IF NOT EXISTS money( idMoney INTEGER PRIMARY KEY, quantityMoney REAL, lastUpdate DATE)"

	_, err = db.Exec(queryCreationMoney)
	errorhand.HandlerError(err, errorhand.TakeFileLine()+" problem with the query")

	queryCreationTags := "CREATE TABLE IF NOT EXISTS tags(tags string PRIMARY KEY);"

	_, err = db.Exec(queryCreationTags)
	errorhand.HandlerError(err, errorhand.TakeFileLine()+" problem with the query")

	return db
}

func SaveTransaction(mon Movement) error {

	date := strconv.Itoa(mon.Date.Day()) + " " + strconv.Itoa(int(mon.Date.Month())) + " " + strconv.Itoa(mon.Date.Year())

	_, err := db.Exec("INSERT INTO spendingMoney VALUES(" + strconv.Itoa(len(Movements)+1) + ", " + strconv.FormatFloat(float64(mon.Money), 'f', 2, 32) + ", '" + mon.Tags + "', '" + date + "', " + strconv.FormatBool(mon.Add) + ");")

	return err

}

func SaveMoneyDB(num float32, ti time.Time) error {

	_, err := db.Exec("INSERT INTO money VALUES(" + strconv.Itoa(len(TotalMoneys)+1) + ", " + strconv.FormatFloat(float64(num), 'f', 2, 32) + ", '" + ti.String() + "');")

	return err
}

func SaveTags(t string) error {

	_, err := db.Exec("INSERT INTO tags VALUES('" + t + "');")

	return err
}

func TakeTags() {

	rows, err := db.QueryContext(context.Background(), "SELECT * FROM tags;")
	errorhand.HandlerError(err, errorhand.TakeFileLine()+" error in the selection of the row")

	defer rows.Close()

	for rows.Next() {
		var mon string

		err := rows.Scan(&mon)
		errorhand.HandlerError(err, errorhand.TakeFileLine()+" error in the scan of the rows")

		AllTags = append(AllTags, mon)
	}

	if len(AllTags) == 0 {
		SaveTags("Transport")
		SaveTag("Transport")
		SaveTags("Shopping")
		SaveTag("Shopping")
		SaveTags("Food")
		SaveTag("Food")
	}
}

func QuantityMoney() {

	rows, err := db.QueryContext(context.Background(), "SELECT * FROM money;")
	errorhand.HandlerError(err, errorhand.TakeFileLine()+" error in the selection of the row")

	defer rows.Close()

	for rows.Next() {
		var mon MoneyRow

		err := rows.Scan(&mon.Id, &mon.RowMon.Total, &mon.RowMon.LastUp)
		errorhand.HandlerError(err, errorhand.TakeFileLine()+" error in the scan of the rows")

		//fmt.Println(mon.RowMon.LastUp)
		//errorhand.Controll(mon.RowMon.LastUp)

		TotalMoneys = append(TotalMoneys, mon)
	}

	if len(TotalMoneys) == 0 {
		TotalMoney.Total = 0
	} else {
		TotalMoney = TotalMoneys[len(TotalMoneys)-1].RowMon
	}

}

func TakeValue() {

	rows, err := db.QueryContext(context.Background(), "SELECT * FROM spendingMoney;")
	errorhand.HandlerError(err, errorhand.TakeFileLine()+" error in the selection of the row")

	defer rows.Close()

	for rows.Next() {
		var mov MovementRow

		var m string

		err := rows.Scan(&mov.Id, &mov.Mov.Money, &mov.Mov.Tags, &m, &mov.Mov.Add)
		errorhand.HandlerError(err, errorhand.TakeFileLine()+" error in the scan of the rows")

		mov.Mov.Date, err = time.Parse("2 1 2006", m)
		errorhand.HandlerError(err, errorhand.TakeFileLine()+" error in the parse of the date")

		Movements = append(Movements, mov)
	}

}

func DeletPay(toEliminate map[int]bool) {

	for in, el := range toEliminate {
		if el {
			_, err := db.Exec("DELETE FROM spendingMoney WHERE id_transition=" + strconv.Itoa(in) + ";")
			errorhand.HandlerError(err, errorhand.TakeFileLine()+" error in the delet of the row")
			//Movements[].Id = nil
			//TO-DO eliminare dall'array && modify the ammount of total money
		}
	}

}
