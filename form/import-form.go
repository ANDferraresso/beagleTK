package form

import (
	"github.com/ANDferraresso/beagleTK/orm"
)

func (form *Form) ImportForm(form *orm.Table, formT *orm.Dictio, extRefs bool, quickList []string, notRequired []string, prefix string) {
    $fieldCls = array_keys($fDef['fields']);
    fields := []string{}
    if len(quickList) > 0 {
        fields = append(fields, quickList...)
    } else {
        for _, v := range form.ColumnsOrder { // foreach($fieldCls as $v)  
            columns = append(columns, v)
        }
    }

    form.Name = form.Name
    form.Prefix = prefix

    for _, v := range fields {
        form.FieldsOrder = append(form.FieldsOrder, v)
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



    //
    public function importForm($fDef, &$formT, $extRefs, &$quickList, &$notRequired, $prefix) {   
        $fieldCls = array_keys($fDef['fields']);
        $fields = array();
        if (count($quickList) > 0) 
            $fields = $quickList;
        else 
            foreach($fieldCls as $v) 
                array_push($fields, $v);

        $this->fDef['name'] = $fDef['name'];
        $this->fDef['prefix'] = $prefix;
        $this->fDef['required'] = array();
        $this->fDef['dontValidate'] = array();
        $this->fDef['fields'] = array();
        $this->fDef['ui'] = array();

        foreach($fields as $v) {
            $this->importField($fDef, $formT, $v, $extRefs, $notRequired);
        }

        // Add _csrf
        $this->addCsrfField();
    }

    //
    public function importField($fDef, &$formT, $cKey, $extRefs, &$notRequired) {   
        $field = $fDef['fields'][$cKey];
        $this->fDef['fields'][$cKey] = array(
            "name" => $cKey,
            "title" => $formT['title'][$cKey],
            "minLength" => $field['minLength'],
            "maxLength" => $field['maxLength'],
            "checks" => $field['checks']
        );

        $this->fDef['ui'][$cKey] = array(
            "attrs" => array(), 
            "default" => '', 
            "widget" => '', 
            "wsUrl" => NULL, 
            "wsCallback" => NULL,
            "options" => array()
        );

        if ($field['default'] == NULL) 
            $this->fDef['ui'][$cKey]['default'] = "";
        else
            $this->fDef['ui'][$cKey]['default'] = $field['default'];

        if (isset($formT['opts'][$cKey]))
            $this->fDef['ui'][$cKey]['options'] = $formT['opts'][$cKey];
              
        if ($field['ui_widget'] === "input-checkbox") {
            if (in_array($cKey, $notRequired))
                ;
            else {
                array_push($this->fDef['required'], $cKey);
                $this->fDef['ui'][$cKey]['attrs']['required'] = "required";
            }
        }
        else {  
            if (in_array($cKey, $notRequired))
                ;
            else {
                array_push($this->fDef['required'], $cKey);
                $this->fDef['ui'][$cKey]['attrs']['required'] = "required";
            }
        }

        $this->fDef['ui'][$cKey]['widget'] = $field['ui_widget'];
        $this->fDef['ui'][$cKey]['wsUrl'] = $field['ui_wsUrl'];
        $this->fDef['ui'][$cKey]['wsCallback'] = $field['ui_wsCallback']; 

        return;
    }