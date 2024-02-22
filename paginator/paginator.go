package paginator

import (
	"math"
	"strconv"
	"strings"
	"unicode/utf8"
)

//
func (p *Paginator) Setup() {
    p.PagesNum = 0
    p.Page = 0
    p.FirstPage = 0
    p.PrevPage = 0
    p.NextPage = 0
    p.LastPage = 0
}

//
func (p *Paginator) Calc(rowsNum int64, rowPerPage int64, page int64) {
	if rowsNum == 0 {
	    p.PagesNum = 0
	    p.Page = 0
	    p.FirstPage = 0
	    p.PrevPage = 0
	    p.NextPage = 0
	    p.LastPage = 0
	}
     
    p.PagesNum = int64(math.Floor(float64(rowsNum) / float64(rowPerPage)))
    if rowsNum % rowPerPage > 0 {
        p.PagesNum++
    }

    p.Page = page
    if p.Page < 1 || p.Page > p.PagesNum {
        p.Page = 1
    }

    p.FirstPage = 1
    p.LastPage = p.PagesNum

    p.PrevPage = p.Page - 1
    if p.PrevPage < 1 {
        p.PrevPage = 1
    }

    p.NextPage = p.Page + 1
    if p.NextPage > p.LastPage { 
        p.NextPage = p.LastPage
    }

}

//
func (p *Paginator) PagHTML(link string, pars map[string]string, 
	showPageMode string, classes string) string {
	//
	var html strings.Builder
    var querys string = "?" 

    for k, v := range pars {
    	querys += k + "=" + v + "&"
    }
            
    if querys == "?" {
    	querys = ""
    } else if utf8.RuneCountInString(querys) > 1 { 
    	querys = string([]rune(querys)[0:utf8.RuneCountInString(querys)-1])
    }

    if len(classes) == 0 {
        html.WriteString("<nav aria-label=\"...\"><ul class=\"pagination\">")
    } else {
        html.WriteString("<nav aria-label=\"...\"><ul class=\"pagination ")
        html.WriteString(classes)
        html.WriteString("\">")
    }

    if p.FirstPage != p.PrevPage {
        if showPageMode == "U" {
            html.WriteString("<li class=\"page-item\" id=\"pag-item-")
            html.WriteString(strconv.FormatInt(p.FirstPage, 10))
            html.WriteString("\"><a class=\"page-link\" href=\"")
            html.WriteString(link)
            html.WriteString(strconv.FormatInt(p.FirstPage, 10))
            html.WriteString(querys)
            html.WriteString("\">")
            html.WriteString(strconv.FormatInt(p.FirstPage, 10))
            html.WriteString("</a></li>")
        } else if showPageMode == "Q" {
            var q string = "0"
            if len([]rune(querys)) == 0 {
            	q = "?page=" + strconv.FormatInt(p.FirstPage, 10)
            } else {
            	q = "?page=" + strconv.FormatInt(p.FirstPage, 10) + "&" + string([]rune(querys)[1:])
            }
            html.WriteString("<li class=\"page-item\" id=\"pag-item-")
            html.WriteString(strconv.FormatInt(p.FirstPage, 10))
            html.WriteString("\"><a class=\"page-link\" href=\"")
            html.WriteString(link)
            html.WriteString(q)
            html.WriteString("\">")
            html.WriteString(strconv.FormatInt(p.FirstPage, 10))
            html.WriteString("</a></li>")
        }  else {
            html.WriteString("<li class=\"page-item\" id=\"pag-item-")
            html.WriteString(strconv.FormatInt(p.FirstPage, 10))
            html.WriteString("\"><a class=\"page-link\" href=\"")
            html.WriteString(link)
            html.WriteString(querys)
            html.WriteString("\">")
            html.WriteString(strconv.FormatInt(p.FirstPage, 10))
            html.WriteString("</a></li>")
        }
    }

    if p.PrevPage - p.FirstPage > 1 {
        if showPageMode == "U" {
            html.WriteString("<li class=\"page-item\" id=\"pag-item-")
            html.WriteString(strconv.FormatInt(p.PrevPage - 1, 10))
            html.WriteString("\"><a class=\"page-link\" href=\"")
            html.WriteString(link)
            html.WriteString(strconv.FormatInt(p.PrevPage - 1, 10))
            html.WriteString(querys)
            html.WriteString("\">&#8592;</a></li>")
        } else if showPageMode == "Q" {
            var q string = "0"
            if len([]rune(querys)) == 0 {
            	q = "?page=" + strconv.FormatInt(p.PrevPage - 1, 10)
            } else {
            	q = "?page=" + strconv.FormatInt(p.PrevPage - 1, 10) + "&" + string([]rune(querys)[1:])
            }
            html.WriteString("<li class=\"page-item\" id=\"pag-item-")
            	html.WriteString(strconv.FormatInt(p.PrevPage - 1, 10))
            	html.WriteString("\"><a class=\"page-link\" href=\"")
            	html.WriteString(link)
            	html.WriteString(q)
            	html.WriteString("\">&#8592;</a></li>")
        } else {
            html.WriteString("<li class=\"page-item\" id=\"pag-item-")
            html.WriteString(strconv.FormatInt(p.PrevPage - 1, 10))
            html.WriteString("\"><a class=\"page-link\" href=\"")
            html.WriteString(link)
            html.WriteString(querys)
            html.WriteString("\">&#8592;</a></li>")
        }
    }

    if p.PrevPage != p.Page {
        if showPageMode == "U" {
            html.WriteString("<li class=\"page-item\" id=\"pag-item-")
            html.WriteString(strconv.FormatInt(p.PrevPage, 10))
            html.WriteString("\"><a class=\"page-link\" href=\"")
            html.WriteString(link)
            html.WriteString(strconv.FormatInt(p.PrevPage, 10))
            html.WriteString(querys)
            html.WriteString("\">")
            html.WriteString(strconv.FormatInt(p.PrevPage, 10))
            html.WriteString("</a></li>")
        } else if showPageMode == "Q" {
            var q string = "0"
            if len([]rune(querys)) == 0 {
            	q = "?page=" + strconv.FormatInt(p.PrevPage, 10)
            } else {
            	q = "?page=" + strconv.FormatInt(p.PrevPage, 10) + "&" + string([]rune(querys)[1:])
            }
            html.WriteString("<li class=\"page-item\" id=\"pag-item-")
            html.WriteString(strconv.FormatInt(p.PrevPage, 10))
            html.WriteString("\"><a class=\"page-link\" href=\"")
            html.WriteString(link)
            html.WriteString(q)
            html.WriteString("\">")
            html.WriteString(strconv.FormatInt(p.PrevPage, 10))
            html.WriteString("</a></li>")
        } else {
            html.WriteString("<li class=\"page-item\" id=\"pag-item-")
            html.WriteString(strconv.FormatInt(p.PrevPage, 10))
            html.WriteString("\"><a class=\"page-link\" href=\"")
            html.WriteString(link)
            html.WriteString(querys)
            html.WriteString("\">")
            html.WriteString(strconv.FormatInt(p.PrevPage, 10))
            html.WriteString("</a></li>")
        }
    }

    html.WriteString("<li class=\"page-item active\"><a class=\"page-link\" href=\"javascript:void(0);\">")
    html.WriteString(strconv.FormatInt(p.Page, 10))
    html.WriteString("</a></li>")

    if p.NextPage != p.Page {
        if showPageMode == "U" {
            html.WriteString("<li class=\"page-item\" id=\"pag-item-")
            html.WriteString(strconv.FormatInt(p.NextPage, 10))
            html.WriteString("\"><a class=\"page-link\" href=\"")
            html.WriteString(link)
            html.WriteString(strconv.FormatInt(p.NextPage, 10))
            html.WriteString(querys)
            html.WriteString("\">")
            html.WriteString(strconv.FormatInt(p.NextPage, 10))
            html.WriteString("</a></li>")
        } else if showPageMode == "Q" {
            var q string = "0"
            if len([]rune(querys)) == 0 {
            	q = "?page=" + strconv.FormatInt(p.NextPage, 10)
            } else {
            	q = "?page=" + strconv.FormatInt(p.NextPage, 10) + "&" + string([]rune(querys)[1:])
            }
            html.WriteString("<li class=\"page-item\" id=\"pag-item-")
            html.WriteString(strconv.FormatInt(p.NextPage, 10))
            html.WriteString("\"><a class=\"page-link\" href=\"")
            html.WriteString(link)
            html.WriteString(q)
            html.WriteString("\">")
            html.WriteString(strconv.FormatInt(p.NextPage, 10))
            html.WriteString("</a></li>")
        } else {
            html.WriteString("<li class=\"page-item\" id=\"pag-item-")
            html.WriteString(strconv.FormatInt(p.NextPage, 10))
            html.WriteString("\"><a class=\"page-link\" href=\"")
            html.WriteString(link)
            html.WriteString(querys)
            html.WriteString("\">")
            html.WriteString(strconv.FormatInt(p.NextPage, 10))
            html.WriteString("</a></li>")
        }
    }

    if p.LastPage - p.NextPage > 1 {
        if showPageMode == "U" {
            html.WriteString("<li class=\"page-item\" id=\"pag-item-")
            html.WriteString(strconv.FormatInt(p.NextPage + 1, 10))
            html.WriteString("\"><a class=\"page-link\" href=\"")
            html.WriteString(link)
            html.WriteString(strconv.FormatInt(p.NextPage + 1, 10))
            html.WriteString(querys)
            html.WriteString("\">&#8594;</a></li>")
        } else if showPageMode == "Q" {
            var q string = "0"
            if len([]rune(querys)) == 0 {
            	q = "?page=" + strconv.FormatInt(p.NextPage + 1, 10)
            } else {
            	q = "?page=" + strconv.FormatInt(p.NextPage + 1, 10) + "&" + string([]rune(querys)[1:])
            }
            html.WriteString("<li class=\"page-item\" id=\"pag-item-")
            html.WriteString(strconv.FormatInt(p.NextPage + 1, 10))
            html.WriteString("\"><a class=\"page-link\" href=\"")
            html.WriteString(link)
            html.WriteString(q)
            html.WriteString("\">&#8594;</a></li>")
        } else {
            html.WriteString("<li class=\"page-item\" id=\"pag-item-")
            html.WriteString(strconv.FormatInt(p.NextPage + 1, 10))
            html.WriteString("\"><a class=\"page-link\" href=\"")
            html.WriteString(link)
            html.WriteString(querys)
            html.WriteString("\">&#8594;</a></li>")
        }
    }

    if p.LastPage != p.NextPage {
        if showPageMode == "U" {
            html.WriteString("<li class=\"page-item\" id=\"pag-item-")
            html.WriteString(strconv.FormatInt(p.LastPage, 10))
            html.WriteString("\"><a class=\"page-link\" href=\"")
            html.WriteString(link)
            html.WriteString(strconv.FormatInt(p.LastPage, 10))
            html.WriteString(querys)
            html.WriteString("\">")
            html.WriteString(strconv.FormatInt(p.LastPage, 10))
            html.WriteString("</a></li>")
        } else if showPageMode == "Q" {
            var q string = "0"
            if len([]rune(querys)) == 0 {
            	q = "?page=" + strconv.FormatInt(p.LastPage, 10)
            } else { 
            	q = "?page=" + strconv.FormatInt(p.LastPage, 10) + "&" + string([]rune(querys)[1:])
            }
            html.WriteString("<li class=\"page-item\" id=\"pag-item-")
            html.WriteString(strconv.FormatInt(p.LastPage, 10))
            html.WriteString("\"><a class=\"page-link\" href=\"")
            html.WriteString(link)
            html.WriteString(q)
            html.WriteString("\">")
            html.WriteString(strconv.FormatInt(p.LastPage, 10))
            html.WriteString("</a></li>")
        } else {
            html.WriteString("<li class=\"page-item\" id=\"pag-item-")
            html.WriteString(strconv.FormatInt(p.LastPage, 10))
            html.WriteString("\"><a class=\"page-link\" href=\"")
            html.WriteString(link)
            html.WriteString(querys)
            html.WriteString("\">")
            html.WriteString(strconv.FormatInt(p.LastPage, 10))
            html.WriteString("</a></li>")
        }
    }

    html.WriteString("\n</ul></nav>") 
    return html.String()
}