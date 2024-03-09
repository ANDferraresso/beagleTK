package form

import (
	"github.com/ANDferraresso/beagleTK/orm"
)

func (form *Form) ImportTable(table *orm.Table, tableT *orm.Dictio, extRefs bool, quickList []string, notRequired []string, prefix string) {
    columns := []string{}
    if len(quickList) > 0 {
        columns = append(columns, quickList...)
    } else {
        for _, v := range table.ColumnsOrder {
            columns = append(columns, v)
        }
    }

    form.Name = table.Name
    form.Prefix = prefix
    for _, v := range columns {
    	form.FieldsOrder = append(form.FieldsOrder, v)
    	if tableT != nil {
	        form.Fields[v] = &Field{
	            Name:      v,
	            Title:     tableT.Title[v],
	            MinLength: table.Columns[v].MinLength,
	            MaxLength: table.Columns[v].MaxLength,
	            Checks:    table.Columns[v].Checks,
	        }
	        form.UIs[v] = &UI{
	            Attrs:      make(map[string]string),
	            Default:    "",
	            Widget:     "",
	            WsUrl:      "",
	            WsCallback: "",
	            Opts:       tableT.Opts[v],
	        }
	    } else {
	        form.Fields[v] = &Field{
	            Name:      v,
	            Title:     "",
	            MinLength: table.Columns[v].MinLength,
	            MaxLength: table.Columns[v].MaxLength,
	            Checks:    table.Columns[v].Checks,
	        }
	        form.UIs[v] = &UI{
	            Attrs:      make(map[string]string),
	            Default:    "",
	            Widget:     "",
	            WsUrl:      "",
	            WsCallback: "",
	            Opts:       map[string]string{},
	        }
	    }

        // default
        /*
           "0": "Non definito",
           "1": "NULL",
           "2": "Stringa vuota",
           "3": "Come definito",
           "4": "CURRENT_TIMESTAMP"
        */
        if table.Columns[v].UcDefault == "3" {
            form.UIs[v].Default = table.Columns[v].Default
        }

        if !stringInSlice(v, notRequired) {
            form.Required = append(form.Required, v)
            form.UIs[v].Attrs["required"] = "required"
        }

        if table.Columns[v].UI_Widget != "" {
            form.UIs[v].Widget = table.Columns[v].UI_Widget
        }
        if table.Columns[v].UI_WsUrl != "" {
            form.UIs[v].WsUrl = table.Columns[v].UI_WsUrl
        }
        if table.Columns[v].UI_WsCallback != "" {
            form.UIs[v].WsCallback = table.Columns[v].UI_WsCallback
        }
    }

    form.AddCsrfField("")
}