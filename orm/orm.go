package orm

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

/*
https://stackoverflow.com/questions/17845619/how-to-call-the-scan-variadic-function-using-reflection
https://golang.hotexamples.com/examples/database.sql/-/RawBytes/golang-rawbytes-function-examples.html
https://ddcode.net/2019/05/11/how-do-golang-sql-rawbytes-convert-to-actual-types-of-variables/
http://go-database-sql.org/varcols.html
https://github.com/go-sql-driver/mysql/issues/682
https://github.com/go-sql-driver/mysql/issues/585
*/

func (table *Table) Select(dbConn *sql.DB, debug string, countRows bool, extRefs bool, quickList []string,
	conds [][7]string, order [][2]string, limit int, offset int) *DbRes {

	dbRes := DbRes{Err: false, Msg: "", Data: []RowValues{}}

	var query strings.Builder
	var columnsSql strings.Builder
	var values []interface{}
	var colsN int
	var rv []RowValues

	columns := []string{}
	if len(quickList) > 0 {
		columns = append(columns, quickList...)
	} else {
		columns = append(columns, ColumnsOrder...)
		/*
		for _, v := range table.ColumnsOrder {
			columns = append(columns, v)
		}
		*/
	}

	if countRows {
		colsN = 1
	} else {
		colsN = len(columns)
	}
	record := make([]interface{}, colsN)
	for i, _ := range record {
		record[i] = new(sql.RawBytes)
	}

	// Elabora condizioni
	cs := elabConds(conds)

	//
	if !extRefs {
		if countRows {
			// SELECT
			if stringInSlice(table.ColumnsOrder, "rid") {
				query.WriteString("SELECT COUNT(*) FROM (SELECT `rid` FROM `")
				query.WriteString(table.Name)
				query.WriteString("` WHERE 1 ")
				query.WriteString(cs.Wquery)
				query.WriteString(" ")
				query.WriteString(cs.Hquery)
				query.WriteString(") AS `CNT`")
			} else {
				query.WriteString("SELECT COUNT(*) FROM (SELECT * FROM `")
				query.WriteString(table.Name)
				query.WriteString("` WHERE 1 ")
				query.WriteString(cs.Wquery)
				query.WriteString(" ")
				query.WriteString(cs.Hquery)
				query.WriteString(") AS `CNT`")
			}
			// WHERE (condizioni)
			cs = elabConds(conds)
		} else {
			// SELECT
			for _, v := range columns {
				// Columns
				columnsSql.WriteString("`")
				columnsSql.WriteString(v)
				columnsSql.WriteString("`, ")
			}
			// WHERE (condizioni)
			cs = elabConds(conds)
			// ORDER
			orderSql := ""
			if len(order) == 0 {
				orderSql = "1"
			} else {
				for _, ord := range order {
					orderSql = fmt.Sprintf("`%s`.`%s` %s, ", table.Name, ord[0], ord[1])
				}
				orderSql = strings.TrimRight(orderSql, ", ")
			}

			query.WriteString("SELECT ")
			query.WriteString(strings.TrimRight(columnsSql.String(), ", "))
			query.WriteString(" FROM `")
			query.WriteString(table.Name)
			query.WriteString("` WHERE 1")
			query.WriteString(cs.Wquery)
			query.WriteString(" ")
			query.WriteString(cs.Hquery)
			query.WriteString(" ORDER BY ")
			query.WriteString(orderSql)
			if limit > -1 {
				query.WriteString(" LIMIT ")
				query.WriteString(strconv.Itoa(offset))
				query.WriteString(", ")
				query.WriteString(strconv.Itoa(limit))
			}

		}

		values = append(values, cs.Binds...)
	} else {
		// extRefs = true
		//

		/*

		   // Trova ripetizioni e copia array
		   $fKeys = array();
		   $tablesArr = array($this->tDef['name'] => 1);
		   foreach ($this->tDef['fKeys'] as $k => $v) {
		       if ((!empty($quickList) && in_array($k, $quickList)) || (empty($quickList) && !in_array($k, $excluding))) {
		           if (isset($tablesArr[ $v['toTable'] ]))
		               $tablesArr[ $v['toTable'] ] += 1;
		           else
		               $tablesArr[ $v['toTable'] ] = 1;
		           $fKeys[$k] = $v;
		       }
		   }

		   //
		   foreach ($tablesArr as $k => $v) {
		       if ($v > 1) {
		           $c = 1;
		           foreach ($fKeys as $k1 => $v1) {
		               if ($k === $v1['toTable']) {
		                   $fKeys[$k1]['toTable'] .= "_" . (string)$c;
		                   $c++;
		               }

		           }
		       }
		   }

		   $fromSQL = "";
		   foreach ($tablesArr as $k => $v) {
		       if ($v === 1) $fromSQL .= " `". $k ."`, ";
		       else {
		           $c = 1;
		           while ($c <= $v) {
		               $fromSQL .= " `" . $k . "` AS `". $k ."_".(string)$c ."`, ";
		               $c++;
		           }
		       }
		   }
		   $fromSQL = rtrim($fromSQL, ", ");

		   // SELECT
		   $selectSQL = "";
		   $whereSQL = "";
		   $havingSQL = "";
		   foreach($columns as $v) {
		       if (!array_key_exists($v, $fKeys)) {
		           $selectSQL .= "`".$this->tDef['name']."`.`$v`, ";
		       }
		       else {
		           $fKey = $fKeys[$v];
		           if (DBMS_ENGINE === 'mysql') {
		               $selectSQL .= "`".$this->tDef['name']."`.`$v`, CONCAT(";
		               foreach($fKey['toRefs'] as $x) {
		                   $selectSQL .= "`".$fKey['toTable']."`.`$x`, \" \", ";
		               }
		               $selectSQL = rtrim($selectSQL, ", \" \", ");
		               $selectSQL .= ') AS `'.$v."_FK`, ";
		           }
		           else if (DBMS_ENGINE === 'sqlite3') {
		               $selectSQL .= "`".$this->tDef['name']."`.`$v`, ";
		               foreach($fKey['toRefs'] as $x) {
		                   $selectSQL .= "`".$fKey['toTable']."`.`$x` || \" \" || ";
		               }
		               $selectSQL = rtrim($selectSQL, " || \" \" || ");
		               $selectSQL .= ' AS `'.$v."_FK`, ";
		           }
		           // WHERE (foreign keys)
		           // $whereSQL .= " AND `".$this->tDef['name']."`.`$v` = `".$fKey['toTable']."`.`".$fKey['toColumn']."`";
		           $whereSQL .= " AND `".$this->tDef['name']."`.`$v` = `".$fKey['toTable']."`.`".$fKey['toColumn']."`";
		       }
		   }
		   $selectSQL = rtrim($selectSQL, ", ");

		   // Condizioni
		   list($w_condSQL, $h_condSQL, $binds) = $this->elabConds($conds);

		   $sql = "";
		   if ($countRows) {
		       if ($selectSQL === '') {
		           if (in_array("rid", $tableCls))
		               $sql .= "SELECT COUNT(*) FROM (SELECT `rid` FROM $fromSQL WHERE 1 $whereSQL $w_condSQL $h_condSQL) AS `CNT`";
		           else
		               $sql .= "SELECT COUNT(*) FROM (SELECT * FROM $fromSQL WHERE 1 $whereSQL $w_condSQL $h_condSQL) AS `CNT`";
		       }
		       else {
		           $sql .= "SELECT COUNT(*) FROM (SELECT $selectSQL FROM $fromSQL WHERE 1 $whereSQL $w_condSQL $h_condSQL) AS `CNT`";
		       }
		   }
		   else {
		       // ORDER
		       $orderSQL = "";
		       if (empty($order)) {
		           $orderSQL = "1";
		       }
		       else {
		           foreach ($order as $ord) {
		               if (mb_substr($ord[0], 0, 1) === "`")
		                   $orderSQL .= "`".$ord[0]."` ".$ord[1].", ";
		               else
		                   $orderSQL .= "`".$this->tDef['name']."`.`".$ord[0]."` ".$ord[1].", ";
		           }
		           $orderSQL = rtrim($orderSQL, ", ");
		       }

		       if ($limit === NULL)
		           $sql = "SELECT $selectSQL FROM $fromSQL WHERE 1 $whereSQL $w_condSQL $h_condSQL ORDER BY $orderSQL";
		       else
		           $sql = "SELECT $selectSQL FROM $fromSQL WHERE 1 $whereSQL $w_condSQL $h_condSQL ORDER BY $orderSQL LIMIT $offset, $limit";
		   }
		*/

	}

	// Lancia query
	if countRows {
		var cnt int64 = 0
		err := dbConn.QueryRow(query.String(), values...).Scan(&cnt)
		if err != nil {
			return ManageDbmsError(&dbRes, debug, err, query.String())
		} else {
			vals := RowValues{}
			vals["CNT"] = cnt //string(raw)
			rv = append(rv, vals)
		}
	} else {  
		rows, err := dbConn.Query(query.String(), values...)
		if err != nil {
			return ManageDbmsError(&dbRes, debug, err, query.String())
		} else {
			// Ottieni tipo di dati delle colonne (columnTypes)
			// columnType.Name() nome colonna
			// columnType.DatabaseTypeName() tipo di data (database)
			// columnType.ScanType() tipo di dato (go)
			// for _, columnType := range columnTypes {
			//     fmt.Println(columnType.Name(), " ", columnType.DatabaseTypeName(), " ", columnType.ScanType())
			// }
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
	}

	dbRes = DbRes{Err: false, Msg: "", Data: rv}
	return &dbRes
}

func (table *Table) InsertRow(dbConn *sql.DB, debug string, ps *map[string]interface{}) *DbRes {
	dbRes := DbRes{Err: false, Msg: "", Data: []RowValues{}}

	var query strings.Builder
	var columnsSql strings.Builder
	var valuesSql strings.Builder
	var values []interface{}

	for k, v := range *ps {
		// Columns
		columnsSql.WriteString("`")
		columnsSql.WriteString(k)
		columnsSql.WriteString("`, ")

		// Values
		valuesSql.WriteString("?, ")
		ManageValues(&values, v)
	}

	query.WriteString("INSERT INTO `")
	query.WriteString(table.Name)
	query.WriteString("`(")
	query.WriteString(strings.TrimRight(columnsSql.String(), ", "))
	query.WriteString(" ) VALUES(")
	query.WriteString(strings.TrimRight(valuesSql.String(), ", "))
	query.WriteString(")")

	// Lancia query
	res, err := dbConn.Exec(query.String(), values...)
	if err != nil {
		return ManageDbmsError(&dbRes, debug, err, query.String())
	} else {
		lastInsertId, err := res.LastInsertId()
		if err != nil {
			lastInsertId = int64(0)
		} else {
			lastInsertId = int64(lastInsertId)
		}
		/*
		   []RowValues{ map[string]interface{"lastInsertId": lastInsertId} }
		       m = make(map[string]Vertex)
		       m["Bell Labs"] = Vertex{
		           40.68433, -74.39967,
		       }
		*/
		dbRes = DbRes{Err: false, Msg: "", Data: []RowValues{map[string]interface{}{"lastInsertId": lastInsertId}}}
	}

	return &dbRes
}

func (table *Table) Update(dbConn *sql.DB, debug string, ps *map[string]interface{}, conds [][7]string) *DbRes {
	dbRes := DbRes{Err: false, Msg: "", Data: []RowValues{}}

	var query strings.Builder
	var columnsSql strings.Builder
	var values []interface{}

	for k, v := range *ps {
		// Columns
		columnsSql.WriteString("`")
		columnsSql.WriteString(k)
		columnsSql.WriteString("` = ?,")

		// Values
		ManageValues(&values, v)
	}

	// WHERE (condizioni)
	cs := elabConds(conds)

	query.WriteString("UPDATE `")
	query.WriteString(table.Name)
	query.WriteString("` SET ")
	query.WriteString(strings.TrimRight(columnsSql.String(), ", "))
	query.WriteString(" WHERE 1 ")
	query.WriteString(cs.Wquery)
	query.WriteString(" ")
	query.WriteString(cs.Hquery)

	values = append(values, cs.Binds...)

	// Lancia query
	res, err := dbConn.Exec(query.String(), values...)
	if err != nil {
		return ManageDbmsError(&dbRes, debug, err, query.String())
	} else {
		rowsAffected, err := res.RowsAffected()
		if err != nil {
			rowsAffected = int64(0)
		} else {
			rowsAffected = int64(rowsAffected)
		}

		dbRes = DbRes{Err: false, Msg: "", Data: []RowValues{map[string]interface{}{"rowsAffected": rowsAffected}}}
	}

	return &dbRes
}

func (table *Table) DeleteRows(dbConn *sql.DB, debug string, conds [][7]string) *DbRes {
	dbRes := DbRes{Err: false, Msg: "", Data: []RowValues{}}

	var query strings.Builder
	var values []interface{}

	// WHERE (condizioni)
	cs := elabConds(conds)

	query.WriteString("DELETE FROM `")
	query.WriteString(table.Name)
	query.WriteString("` WHERE 1 ")
	query.WriteString(cs.Wquery)
	query.WriteString(" ")
	query.WriteString(cs.Hquery)

	values = append(values, cs.Binds...)

	// Lancia query
	res, err := dbConn.Exec(query.String(), values...)
	if err != nil {
		return ManageDbmsError(&dbRes, debug, err, query.String())
	} else {
		rowsAffected, err := res.RowsAffected()
		if err != nil {
			rowsAffected = int64(0)
		} else {
			rowsAffected = int64(rowsAffected)
		}

		dbRes = DbRes{Err: false, Msg: "", Data: []RowValues{map[string]interface{}{"rowsAffected": rowsAffected}}}
	}

	return &dbRes
}
