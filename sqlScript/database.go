package sqlScript

import (
	"context"
	"database/sql"
	"strconv"

	"github.com/mattemello/balanceTerminal/errorHand"
)

var db *sql.DB

func CreationTable() *sql.DB {
	var err error
	db, err = sql.Open("sqlite3", "money.db")
	errorhand.HandlerError(err, errorhand.TakeFileLine()+"error in the opening of the db")

	queryCreationSpending := "CREATE TABLE IF NOT EXISTS spendingMoney( id_transition INTEGER PRIMARY KEY, valur_tansaction REAL, tags TEXT, date DATETIME )"

	_, err = db.Exec(queryCreationSpending)
	errorhand.HandlerError(err, errorhand.TakeFileLine()+" problem with the query")

	queryCreationMoney := "CREATE TABLE IF NOT EXISTS money( idMoney INTEGER PRIMARY KEY, quantityMoney REAL, lastUpdate DATE)"

	_, err = db.Exec(queryCreationMoney)
	errorhand.HandlerError(err, errorhand.TakeFileLine()+" problem with the query")

	return db
}

func SaveTransaction(mon Movement) error {

	_, err := db.Exec("INSERT INTO spendingMoney VALUES(" + strconv.Itoa(len(Movements)+1) + ", " + strconv.FormatFloat(float64(mon.Money), 'f', 2, 32) + ", '" + mon.Tags + "', '" + mon.Date.String() + "');")

	return err

}

func SaveMoney(num float32) error {

	//TO-DO: save the money in the db
	_, err := db.Exec("INSERT INTO spendingMoney VALUES(" + strconv.Itoa(len(Movements)+1) + ",")

	return err
}

func QuantityMoney() float32 {
	var val float32

	return val
}

func TakeValue() {

	rows, err := db.QueryContext(context.Background(), "SELECT * FROM spendingMoney;")
	errorhand.HandlerError(err, errorhand.TakeFileLine()+" error in the selection of the row")

	defer rows.Close()

	for rows.Next() {
		var mov MovementRow

		err := rows.Scan(&mov.Id, &mov.Mov.Money, &mov.Mov.Tags, &mov.Mov.Date)
		errorhand.HandlerError(err, errorhand.TakeFileLine()+" error in the scan of the rows")

		Movements = append(Movements, mov)
	}

}
