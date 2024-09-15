package sqlScript

import (
	"context"
	"database/sql"
	"strconv"

	"github.com/mattemello/balanceTerminal/errorHand"
)

/*_, err = db.Exec("INSERT INTO spendingMoney VALUES(0, 0.5, 'real', 05/12/2021);")
if err != nil {
log.Fatal("29 Error with the database: ", err)
return
}

_, err = db.Exec("DELETE FROM spendingMoney WHERE id_transition=0;")
if err != nil {
log.Fatal("Error with the database: ", err)
return
}*/

/*func CreationRow(value float32, str string, dat time.Month) {

}*/

var db *sql.DB

func CreationTable() *sql.DB {
	var err error
	db, err = sql.Open("sqlite3", "money.db")
	errorhand.HandlerError(err)

	queryCre := "CREATE TABLE IF NOT EXISTS spendingMoney( id_transition INTEGER PRIMARY KEY, valur_tansaction REAL, tags TEXT, date DATETIME )"

	_, err = db.Exec(queryCre)
	errorhand.HandlerError(err)

	return db
}

func SaveValue(mon Movement) error {

	//errorhand.Controll("INSERT INTO spendingMoney VALUES(" + strconv.Itoa(id) + ", " + strconv.FormatFloat(float64(mon.Money), 'f', 2, 32) + ", '" + mon.Tags + "', " + mon.Date.String() + ");")

	_, err := db.Exec("INSERT INTO spendingMoney VALUES(" + strconv.Itoa(len(Movements)+1) + ", " + strconv.FormatFloat(float64(mon.Money), 'f', 2, 32) + ", '" + mon.Tags + "', '" + mon.Date.String() + "');")

	return err

}

func TakeValue() {

	rows, err := db.QueryContext(context.Background(), "SELECT * FROM spendingMoney;")
	errorhand.HandlerError(err)

	defer rows.Close()

	for rows.Next() {
		var mov MovementRow

		err := rows.Scan(&mov.Id, &mov.Mov.Money, &mov.Mov.Tags, &mov.Mov.Date)
		errorhand.HandlerError(err)

		Movements = append(Movements, mov)
	}

}
