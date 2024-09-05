package sqlScript

import (
	"database/sql"
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

func CreationTable() *sql.DB {
	db, err := sql.Open("sqlite3", "money.db")
	errorhand.HandlerError(err)

	queryCre := "CREATE TABLE IF NOT EXISTS spendingMoney( id_transition INTEGER PRIMARY KEY, valur_tansaction REAL, tags TEXT, date DATE )"

	_, err = db.Exec(queryCre)
	errorhand.HandlerError(err)

	return db
}
