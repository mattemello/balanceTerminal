package sqlScript

import (
	"database/sql"
	"strconv"
	//"time"

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

	_, err := db.Exec("INSERT INTO spendingMoney VALUES(1, " + strconv.FormatFloat(float64(mon.Money), 'E', 2, 32) + ", '" + mon.Tags + "', " + mon.Date + ");")

	return err

}
