package paginator

import (
	"strconv"
	"strings"
	"unicode/utf8"

    "github.com/ANDferraresso/beagleTK/orm"
)

//
func (p *Paginator) ManageFilter(dbEngine string, model *orm.Table, GET_ map[string]string,
	extRefs bool, quickList []string, listAllConds [][7]string, 
    options map[string][]map[string]string, listURL string, btnName string, 
    colSpan int) ([][7]string, map[string]string, string, string) {
    //
    conds := listAllConds
	var filterHTML strings.Builder
    var filterJs strings.Builder

	columns := []string{}
	if len(quickList) > 0 {
		columns = append(columns, quickList...)
	} else {
		columns = append(columns, model.ColumnsOrder...)
    }

    filterList := map[string]string{}
    filterArr := map[string]string{}

    if extRefs == false {
        // columns
		for _, v := range columns {   
	    	if _, ok := GET_["filter_" + v]; ok {
		    	if GET_["filter_" + v] != "" {
	                t := string([]rune(GET_["filter_" + v])[0:30])
	                filterList[v] = t
	                filterArr["filter_" + v] = t	    	
	            }
	        }
		}
        // conds
        for k, v := range filterList {   
            if v != "" {
            	conds = append(conds, [7]string{"W", "AND", model.Name, k, "LIKE", "%" + v + "%", ""})
            } 
        }
        // HTML
        for _, v := range columns {
            if _, ok := options[v]; ok {
                filterHTML.WriteString("\n      <th scope=\"col\"><select class=\"form-select form-select-sm\" name=\"filter_" + v + "\" id=\"filter_" + v + "_id\">")
                filterHTML.WriteString("<option value=\"\"></option>")
                for _, elem := range options[v] {
                	// NB: C'è in solo elemento in ciascuna mappa
                	for optK, optV := range elem {
                		if _, ok = filterArr["filter_" + v]; ok {
                			if filterArr["filter_" + v] == optK {
                				filterHTML.WriteString("<option value=\"" + optK + "\" selected=\"selected\">" + optV + "</option>")
                			} else {
                				filterHTML.WriteString("<option value=\"" + optK + "\">" + optV + "</option>")
                			}
                		} else {
                			filterHTML.WriteString("<option value=\"" + optK + "\">" + optV + "</option>")
                		}
                	}
                }
                filterHTML.WriteString("</select></th>\n")
            } else {
            	if _, ok = filterArr["filter_" + v]; ok {
            		filterHTML.WriteString("      <th scope=\"col\"><input type=\"text\" maxlength=\"30\" class=\"form-control form-control-sm\" name=\"filter_" +  v + "\" id=\"filter_" +  v + "_id\" value=\"" + filterArr["filter_" + v] + "\"/></th>\n")
            	} else {
            		filterHTML.WriteString("      <th scope=\"col\"><input type=\"text\" maxlength=\"30\" class=\"form-control form-control-sm\" name=\"filter_"  + v + "\" id=\"filter_" +  v + "_id\" /></th>\n")
            	}
            }
        }
        filterHTML.WriteString("\n      <th scope=\"col\" colspan=\"" + strconv.Itoa(colSpan) + "\"><button type=\"button\" id=\"" + btnName + "\" class=\"btn btn-secondary btn-sm\">Filter</button></th>\n")
        // Js
        filterJs.WriteString("\nlet btn = document.getElementById(\"" + btnName + "\");")
        filterJs.WriteString("\nif btn != null) {")
        filterJs.WriteString("\n  btn.addEventListener(\"click\", function (event) {")
        filterJs.WriteString("\n    event.preventDefault();")
        filterJs.WriteString("\n    let filters = {};")
        filterJs.WriteString("\n    let location = \"" + listURL + "\";")
        for _, v := range columns {
        	filterJs.WriteString("\n    filters[\"" + v + "\"] = String(document.getElementById(\"filter_" + v + "_id\").value).substring(0, 30);")	
        }
        filterJs.WriteString("\n    for (let key in filters) {")
        filterJs.WriteString("\n      if (filters[key] !== \"\") location += \"filter_\" + key + \"=\" + filters[key] + \"&\";") 
        filterJs.WriteString("\n    }")
        filterJs.WriteString("\n    if (location.charAt(location.length - 1) == \"?\") location = location.substring(0, location.length - 1);")
        filterJs.WriteString("\n    else if (location.charAt(location.length - 1) == \"&\") location = location.substring(0, location.length - 1);")
        filterJs.WriteString("\n    document.location.href = location;")
        filterJs.WriteString("\n  });")
        filterJs.WriteString("\n}")
    } else {
        // extRefs == TRUE
        // columns
        for _, v := range columns {   
        	if _, ok := model.FKeys[v]; ok {
            	if _, ok := GET_["filter_" + v + "_FK"]; ok {
            		if GET_["filter_" + v + "_FK"] != "" {
            			t := string([]rune(GET_["filter_" + v + "_FK"])[0:30])
                    	filterList[v + "_FK"] = t
                    	filterArr["filter_" + v + "_FK"] = t
            		}
                }
            } else {
            	if _, ok = GET_["filter_" + v]; ok {
            		if GET_["filter_" + v] != "" {
            			t := string([]rune(GET_["filter_" + v])[0:30])
                    	filterList[v] = t
                    	filterArr["filter_" + v] = t
            		}
            	}
            }  
        }

        // conds
        for k, v := range filterList {   
            if utf8.RuneCountInString(k) >= 3 && string([]rune(k)[utf8.RuneCountInString(k)-3 : utf8.RuneCountInString(k)]) == "_FK" {
		        if v != ""  {
                    if dbEngine == "sqlite3" {
                        conds = append(conds, [7]string{"W", "AND", "", k, "LIKE", "%" + v + "%", ""})
       				} else if dbEngine == "mysql" {
                        conds = append(conds, [7]string{"W", "AND", "", k, "LIKE", "%" + v + "%", ""})
                    }
                }  
	        } else {
	        	if v != "" {
	        		conds = append(conds, [7]string{"W", "AND", model.Name, k, "LIKE", "%" + v + "%", ""})
	        	}
            }
        }

        // HTML
        for _, v := range columns {   
        	if _, ok := options[v]; ok {
        		if _, ok := model.FKeys[v]; ok {
        			filterHTML.WriteString("\n      <th scope=\"col\"><select class=\"form-select form-select-sm\" name=\"filter_" + v + "_FK\" id=\"filter_" + v + "_FK_id\">")
        		} else {
        			filterHTML.WriteString("\n      <th scope=\"col\"><select class=\"form-select form-select-sm\" name=\"filter_" + v + "\" id=\"filter_" + v + "_id\">")
        		}
        		filterHTML.WriteString("<option value=\"\"></option>")
                for _, elem := range options[v] {
                	// NB: C'è in solo elemento in ciascuna mappa
                	for  optK, optV := range elem {
                		if _, ok = model.FKeys[v]; ok {
                			if _, ok = filterArr["filter_" + v + "_FK"]; ok {
                				if filterArr["filter_" + v + "_FK"] == optK {
                					filterHTML.WriteString("<option value=\"" + optK + "\" selected=\"selected\">" + optV + "</option>")
                				} else {
                					filterHTML.WriteString("<option value=\"" + optK + "\">" + optV + "</option>")
                				} 
                			} else {
                				filterHTML.WriteString("<option value=\"" + optK + "\">" + optV + "</option>")
                			}
                		} else {
                			if _, ok = filterArr["filter_" + v]; ok {
                				if filterArr["filter_" + v] == optK {
                					filterHTML.WriteString("<option value=\"" + optK + "\" selected=\"selected\">" + optV + "</option>")
                				} else {
                					filterHTML.WriteString("<option value=\"" + optK + "\">" + optV + "</option>")
                				}
                			}
                		}
                	}
                }
                filterHTML.WriteString("</select></th>\n")
            } else {
        		if _, ok := model.FKeys[v]; ok {
        			if _, ok := filterArr["filter_" + v + "_FK"]; ok {
                        filterHTML.WriteString("      <th scope=\"col\"><input type=\"text\" maxlength=\"30\" class=\"form-control form-control-sm\" name=\"filter_" + v + "_FK\" id=\"filter_" + v + "_FK_id\" value=\"" + filterArr["filter_" +  v + "_FK"] + "\" /></th>\n")
                    } else {
                        filterHTML.WriteString("      <th scope=\"col\"><input type=\"text\" maxlength=\"30\" class=\"form-control form-control-sm\" name=\"filter_" + v + "_FK\" id=\"filter_" + v + "_FK_id\" value=\"\" /></th>\n")
                    }
                } else {
                    if _, ok := filterArr["filter_" + v]; ok {
                        filterHTML.WriteString("      <th scope=\"col\"><input type=\"text\" maxlength=\"30\" class=\"form-control form-control-sm\" name=\"filter_" + v + "\" id=\"filter_" + v + "_id\" value=\"" + filterArr["filter_" + v] + "\" /></th>\n")
                    } else {
                        filterHTML.WriteString("      <th scope=\"col\"><input type=\"text\" maxlength=\"30\" class=\"form-control form-control-sm\" name=\"filter_" + v + "\" id=\"filter_" + v + "_id\" value=\"\" /></th>\n")
                    }
                }
            }
  		}

        filterHTML.WriteString("      <th scope=\"col\" colspan=\"" + strconv.Itoa(colSpan) + "\"><button type=\"button\" id=\"" + btnName + "\" class=\"btn btn-secondary btn-sm\">Filter</button></th>\n")
        
        // Js
        filterJs.WriteString("\nlet btn = document.getElementById(\"" + btnName + "\");")
        filterJs.WriteString("\nif btn != null) {")
        filterJs.WriteString("\n  btn.addEventListener(\"click\", function (event) {")
        filterJs.WriteString("\n    event.preventDefault();")
        filterJs.WriteString("\n    let filters = {};")
        filterJs.WriteString("\n    let location = \"" + listURL + "\"")
        for _, v := range columns {   
        	if _, ok := model.FKeys[v]; ok {
        		if _, ok := options[v]; ok {
                    filterJs.WriteString("\n    filters[\"" + v + "_FK\"] = String(document.getElementById(\"filter_" + v + "_FK_id\").value).substring(0, 30);")
                } else {
                    filterJs.WriteString("\n    filters[\"" + v + "_FK\"] = String(document.getElementById(\"filter_" + v + "_FK_id\").value).substring(0, 30);")
                }
            } else {
                filterJs.WriteString("\n    filters[\"" + v + "\"] = String(document.getElementById(\"filter_" + v + "_id\").value).substring(0, 30);")
            }
        }
        filterJs.WriteString("\n    for (let key in filters) {")
        filterJs.WriteString("\n      if filters[key] !== \"\") location += \"filter_\" + key + \"=\" + filters[key] + \"&\";")
        filterJs.WriteString("\n    }")
        filterJs.WriteString("\n    if location.charAt(location.length - 1) == \"?\") location = location.substring(0, location.length - 1);")
        filterJs.WriteString("\n    else if location.charAt(location.length - 1) == \"&\") location = location.substring(0, location.length - 1);")
        filterJs.WriteString("\n    document.location.href = location;")
        filterJs.WriteString("\n  });")
        filterJs.WriteString("\n});")
    }

    return conds, filterArr, filterJs.String(), filterHTML.String()
}