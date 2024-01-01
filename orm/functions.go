package orm

import (
	"database/sql"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

// 
func ResultsFromSelect(dbConn *sql.DB, debug string, 
	columns []string, record []interface{}, query strings.Builder) *DbRes  {
	//
	dbRes := DbRes{Err: false, Msg: "", Data: []RowValues{}}
	var values []interface{}
	var rv []RowValues

	rows, err := dbConn.Query(query.String(), values...)
	if err != nil {
		return ManageDbmsError(&dbRes, debug, err, query.String())
	} else {
		_, err := rows.ColumnTypes()
		if err != nil {
			return ManageDbmsError(&dbRes, debug, err, query.String())
		} else {
			// Ottieni un record alla volta.
			defer rows.Close()
			for rows.Next() {
				err := rows.Scan(record...)
				if err != nil {
					return ManageDbmsError(&dbRes, debug, err, query.String())
				} else {
					vals := RowValues{}
					for i, k := range columns {
						raw := *record[i].(*sql.RawBytes)
						vals[k] = ManageDbResValue(raw)
						i++
					}
					rv = append(rv, vals)
				}
			}
		}
	}

	dbRes = DbRes{Err: false, Msg:"", Data: rv}
	return &dbRes
}