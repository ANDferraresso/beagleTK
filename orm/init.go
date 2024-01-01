package orm

import (
	"database/sql"
	"fmt"
	"reflect"

	_ "github.com/go-sql-driver/mysql"

	"github.com/ANDferraresso/beagleTK/validator"
)

type RowValues map[string]interface{}

type FKeys struct {
	ToTable  string
	ToColumn string
	ToRefs   []string
}

type CondSql struct {
	Wquery string
	Hquery string
	Binds  []interface{}
}

type Column struct {
	Type          string
	Length        string
	NotNull       bool
	UcDefault     string
	Default       string
	MinLength     string
	MaxLength     string
	Checks        []validator.Check
	UI_Widget     string
	UI_WsUrl      string
	UI_WsCallback string
}

type Table struct {
	Name             string
	Primary          []string
	Uniques          [][]string
	ColumnsInUniques []string
	Indexes          [][]string
	FKeys            map[string]FKeys
	ColumnsOrder     []string
	Columns          map[string]Column
}

type Dictio struct {
	EntTitle string                         `json:"EntTitle"`
	Title    map[string]string              `json:"Title"`
	Msg      map[string]string              `json:"Msg"`
	Info     map[string]string              `json:"Info"`
	Opts     map[string][]map[string]string `json:"Opts"`
}

type TableTab struct {
	T             *Dictio
	Data          []RowValues
	HasData       bool
	CsrfToken     string
	FormHasErrors bool
}

type DbRes struct {
	Err     bool
	Msg     string
	Data    []RowValues
}

type OptRes struct {
	Err     bool
	Msg     string
	Data    []map[string]string
}

func elabConds(conds [][7]string) *CondSql {
	cs := CondSql{
		Wquery: "",
		Hquery: "",
		Binds:  []interface{}{},
	}

	for _, cond := range conds {
		// W = WHERE
		// H = HAVING
		// Esempio: cond = "W", "AND", "user", "rid", "=', 2, ""
		// Esempio: cond = "W", "AND (", "user", "gender", "=", "M", ")""
		// Esempio: cond = "W", "AND", "", "rid", "=', 2, ""
		// Esempio: cond = "W", "AND", "", "rid", "IS NULL", "", ""
		// Esempio: cond = "W", "AND", "", "rid", "IS NOT NULL", "", ""
		if cond[0] == "W" {
			if cond[2] != "" {
				cs.Wquery = fmt.Sprintf("%s %s `%s`.`%s` %s", cs.Wquery, cond[1], cond[2], cond[3], cond[4])
			} else {
				cs.Wquery = fmt.Sprintf("%s %s `%s` %s", cs.Wquery, cond[1], cond[3], cond[4])
			}
			if cond[4] == "IS NULL" || cond[4] == "IS NOT NULL" {
				cs.Wquery = fmt.Sprintf("%s %s", cs.Wquery, cond[6])
			} else {
				cs.Wquery = fmt.Sprintf("%s ? %s", cs.Wquery, cond[6])
				cs.Binds = append(cs.Binds, cond[5])
			}
		} else if cond[0] == "H" {
			if cond[2] != "" {
				cs.Hquery = fmt.Sprintf("%s %s `%s`.`%s` %s", cs.Hquery, cond[1], cond[2], cond[3], cond[4])
			} else {
				cs.Hquery = fmt.Sprintf("%s %s `%s` %s", cs.Hquery, cond[1], cond[3], cond[4])
			}
			if cond[4] == "IS NULL" || cond[4] == "IS NOT NULL" {
				cs.Hquery = fmt.Sprintf("%s %s", cs.Hquery, cond[6])
			} else {
				cs.Hquery = fmt.Sprintf("%s ? %s", cs.Hquery, cond[6])
				cs.Binds = append(cs.Binds, cond[5])
			}
		}
	}

	if cs.Hquery != "" {
		cs.Hquery = fmt.Sprintf("HAVING 1 %s", cs.Hquery)
	}

	return &cs
}

func ManageDbmsError(dbRes *DbRes, debug string, err error, query string) *DbRes {
	dbRes.Err = true
	dbRes.Msg = "DBMS ERROR"
	dbRes.Data = []RowValues{}

	switch debug {
	case "0":

	case "1":
		dbRes.Msg += ": " + err.Error() + " " + query
	default:

	}

	return dbRes
}

func ManageNullString(s string) sql.NullString {
	if len(s) == 0 {
		return sql.NullString{}
	}
	return sql.NullString{
		String: s,
		Valid:  true,
	}
}

func ManageDbResValue(rb sql.RawBytes) interface{} {
	if rb == nil {
		return nil
	} else {
		return string(rb)
	}
}

func ManageValues(values *[]interface{}, v interface{}) {
	if (reflect.TypeOf(v)) == nil {
		*values = append(*values, ManageNullString(""))
	} else {
		t := v.(interface{})
		switch t.(type) {
		case nil:
			*values = append(*values, ManageNullString(""))
		case bool:
			if v.(bool) == true {
				*values = append(*values, 1)
			} else {
				*values = append(*values, 0)
			}
		case int:
			*values = append(*values, v.(int))
		case int32:
			*values = append(*values, v.(int32))
		case int64:
			*values = append(*values, v.(int64))
		case string:
			*values = append(*values, v.(string))
		default:
			*values = append(*values, "")
		}
	}
}
