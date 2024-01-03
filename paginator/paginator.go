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
func (p *Paginator) Calc(rowsNum int, rowPerPage int, page int) {
	if rowsNum == 0 {
	    p.PagesNum = 0
	    p.Page = 0
	    p.FirstPage = 0
	    p.PrevPage = 0
	    p.NextPage = 0
	    p.LastPage = 0
	}
     
    p.PagesNum = math.Floor(rowsNum / rowPerPage)
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
func (p *Paginator) PagHTML(rowsNum int, rowPerPage int, Page int, link string, pars map[string]string, 
	showPageMode string, classes string) string {
	//
	p.Calc(rowsNum, rowPerPage, page)
	var html strings.Builder
    var querys string = "?" 

    for k, v := range pars {
    	querys .= k + "=" + v + "&"
    }
            
    if querys == "?" {
    	querys = ""
    } else if utf8.RuneCountInString(querys) > 1 { 
    	querys = string([]rune(querys)[0:utf8.RuneCountInString(querys)-1])
    }

    if len(classes) {
        html.WriteString("<nav aria-label=\"...\"><ul class=\"pagination\">")
    } else {
        html.WriteString("<nav aria-label=\"...\"><ul class=\"pagination ")
        html.WriteString(classes)
        html.WriteString("\">")
    }

    if p.FirstPage != p.PrevPage {
        if showPageMode == "U" {
            html.WriteString("<li class=\"page-item\" id=\"pag-item-")
            html.WriteString(strconv.Itoa(p.FirstPage))
            html.WriteString("\"><a class=\"page-link\" href=\"")
            html.WriteString(link)
            html.WriteString(strconv.Itoa(p.FirstPage))
            html.WriteString(querys)
            html.WriteString("\">")
            html.WriteString(strconv.Itoa(p.FirstPage))
            html.WriteString("</a></li>")
        } else if showPageMode == "Q" {
            var q string = 0
            if mb_strlen(querys) == 0 {
            	q = "?Page=" + strconv.Itoa(p.FirstPage)
            } else {
            	q = "?Page=" + strconv.Itoa(p.FirstPage) + "&" + string([]rune(querys)[1:])
            }
            html.WriteString("<li class=\"page-item\" id=\"pag-item-")
            html.WriteString(strconv.Itoa(p.FirstPage))
            html.WriteString("\"><a class=\"page-link\" href=\"")
            html.WriteString(link)
            html.WriteString(q)
            html.WriteString("\">")
            html.WriteString(strconv.Itoa(p.FirstPage))
            html.WriteString("</a></li>")
        }  else {
            html.WriteString("<li class=\"page-item\" id=\"pag-item-")
            html.WriteString(strconv.Itoa(p.FirstPage))
            html.WriteString("\"><a class=\"page-link\" href=\"")
            html.WriteString(link.querys)
            html.WriteString("\">")
            html.WriteString(strconv.Itoa(p.FirstPage))
            html.WriteString("</a></li>")
        }
    }

    if p.PrevPage - p.FirstPage > 1 {
        if showPageMode == "U" {
            html.WriteString("<li class=\"page-item\" id=\"pag-item-")
            html.WriteString(strconv.Itoa(p.PrevPage - 1))
            html.WriteString("\"><a class=\"page-link\" href=\"")
            html.WriteString(link)
            html.WriteString(strconv.Itoa(p.PrevPage - 1))
            html.WriteString(querys)
            html.WriteString("\">&#8592;</a></li>")
        } else if showPageMode == "Q" {
            var q string = 0
            if mb_strlen(querys) == 0 {
            	q = "?Page=" + strconv.Itoa(p.PrevPage - 1)
            } else {
            	q = "?Page=" + strconv.Itoa(p.PrevPage - 1) + "&" + string([]rune(querys)[1:])
            }
            html.WriteString("<li class=\"page-item\" id=\"pag-item-")
            	html.WriteString(strconv.Itoa(p.PrevPage - 1))
            	html.WriteString("\"><a class=\"page-link\" href=\"")
            	html.WriteString(link)
            	html.WriteString(q)
            	html.WriteString("\">&#8592;</a></li>")
        } else {
            html.WriteString("<li class=\"page-item\" id=\"pag-item-")
            html.WriteString(strconv.Itoa(p.PrevPage - 1))
            html.WriteString("\"><a class=\"page-link\" href=\"")
            html.WriteString(link)
            html.WriteString(querys)
            html.WriteString("\">&#8592;</a></li>")
        }
    }

    if p.PrevPage != p.Page {
        if showPageMode == "U" {
            html.WriteString("<li class=\"page-item\" id=\"pag-item-")
            html.WriteString(strconv.Itoa(p.PrevPage))
            html.WriteString("\"><a class=\"page-link\" href=\"")
            html.WriteString(link)
            html.WriteString(strconv.Itoa(p.PrevPage))
            html.WriteString(querys)
            html.WriteString("\">")
            html.WriteString(strconv.Itoa(p.PrevPage))
            html.WriteString("</a></li>")
        } else if showPageMode == "Q" {
            var q string = 0
            if mb_strlen(querys) == 0 {
            	q = "?Page=" + strconv.Itoa(p.PrevPage)
            } else {
            	q = "?Page=" + strconv.Itoa(p.PrevPage) + "&" + string([]rune(querys)[1:])
            }
            html.WriteString("<li class=\"page-item\" id=\"pag-item-")
            html.WriteString(strconv.Itoa(p.PrevPage))
            html.WriteString("\"><a class=\"page-link\" href=\"")
            html.WriteString(link)
            html.WriteString(q)
            html.WriteString("\">")
            html.WriteString(strconv.Itoa(p.PrevPage))
            html.WriteString("</a></li>")
        } else {
            html.WriteString("<li class=\"page-item\" id=\"pag-item-")
            html.WriteString(strconv.Itoa(p.PrevPage))
            html.WriteString("\"><a class=\"page-link\" href=\"")
            html.WriteString(link)
            html.WriteString(querys)
            html.WriteString("\">")
            html.WriteString(strconv.Itoa(p.PrevPage))
            html.WriteString("</a></li>")
        }
    }

    html.WriteString("<li class=\"page-item active\"><a class=\"page-link\" href=\"javascript:void(0);\">")
    html.WriteString(p.Page)
    html.WriteString("</a></li>")

    if p.NextPage != p.Page {
        if showPageMode == "U" {
            html.WriteString("<li class=\"page-item\" id=\"pag-item-")
            html.WriteString(strconv.Itoa(p.NextPage))
            html.WriteString("\"><a class=\"page-link\" href=\"")
            html.WriteString(link)
            html.WriteString(strconv.Itoa(p.NextPage))
            html.WriteString(querys)
            html.WriteString("\">")
            html.WriteString(strconv.Itoa(p.NextPage))
            html.WriteString("</a></li>")
        } else if showPageMode == "Q" {
            var q string = 0
            if mb_strlen(querys) == 0 {
            	q = "?Page=".p.NextPage
            } else {
            	q = "?Page=".p.NextPage."&".mb_substr(querys, 1, mb_strlen(querys))
            }
            html.WriteString("<li class=\"page-item\" id=\"pag-item-")
            html.WriteString(strconv.Itoa(p.NextPage))
            html.WriteString("\"><a class=\"page-link\" href=\"")
            html.WriteString(link)
            html.WriteString(q)
            html.WriteString("\">")
            html.WriteString(strconv.Itoa(p.NextPage))
            html.WriteString("</a></li>")
        } else {
            html.WriteString("<li class=\"page-item\" id=\"pag-item-")
            html.WriteString(strconv.Itoa(p.NextPage))
            html.WriteString("\"><a class=\"page-link\" href=\"")
            html.WriteString(link)
            html.WriteString(querys)
            html.WriteString("\">")
            html.WriteString(strconv.Itoa(p.NextPage))
            html.WriteString("</a></li>")
        }
    }

    if p.LastPage - p.NextPage > 1 {
        if showPageMode == "U" {
            html.WriteString("<li class=\"page-item\" id=\"pag-item-")
            html.WriteString(strconv.Itoa(p.NextPage + 1))
            html.WriteString("\"><a class=\"page-link\" href=\"")
            html.WriteString(link)
            html.WriteString(strconv.Itoa(p.NextPage + 1))
            html.WriteString(querys)
            html.WriteString("\">&#8594;</a></li>")
        } else if showPageMode == "Q" {
            var q string = 0
            if mb_strlen(querys) == 0 {
            	q = "?Page=" + strconv.Itoa(p.NextPage + 1)
            } else {
            	q = "?Page=" + strconv.Itoa( p.NextPage + 1) + "&" + string([]rune(querys)[1:])
            }
            html.WriteString("<li class=\"page-item\" id=\"pag-item-")
            html.WriteString(strconv.Itoa(p.NextPage + 1))
            html.WriteString("\"><a class=\"page-link\" href="')
            html.WriteString(link)
            html.WriteString(q)
            html.WriteString("\">&#8594;</a></li>")
        } else {
            html.WriteString("<li class=\"page-item\" id=\"pag-item-")
            html.WriteString(strconv.Itoa(p.NextPage + 1))
            html.WriteString("\"><a class=\"page-link\" href=\"")
            html.WriteString(link)
            html.WriteString(querys)
            html.WriteString("\">&#8594;</a></li>")
        }
    }

    if p.LastPage != p.NextPage {
        if showPageMode == "U" {
            html.WriteString("<li class=\"page-item\" id=\"pag-item-")
            html.WriteString(strconv.Itoa(p.LastPage))
            html.WriteString("\"><a class=\"page-link\" href=\"")
            html.WriteString(link)
            html.WriteString(strconv.Itoa(p.LastPage))
            html.WriteString(querys)
            html.WriteString("\">")
            html.WriteString(strconv.Itoa(p.LastPage))
            html.WriteString("</a></li>")
        } else if showPageMode == "Q" {
            var q string = 0
            if mb_strlen(querys) == 0 {
            	q = "?Page=" + strconv.Itoa(p.LastPage)
            } else { 
            	q = "?Page=" + strconv.Itoa(p.LastPage) + "&" + string([]rune(querys)[1:])
            }
            html.WriteString("<li class=\"page-item\" id=\"pag-item-")
            html.WriteString(strconv.Itoa(p.LastPage))
            html.WriteString("\"><a class=\"page-link\" href=\"")
            html.WriteString(link)
            html.WriteString(q)
            html.WriteString("\">")
            html.WriteString(strconv.Itoa(p.LastPage))
            html.WriteString("</a></li>")
        } else {
            html.WriteString("<li class=\"page-item\" id=\"pag-item-")
            html.WriteString(strconv.Itoa(p.LastPage))
            html.WriteString("\"><a class=\"page-link\" href=\"")
            html.WriteString(link)
            html.WriteString(querys)
            html.WriteString("\">")
            html.WriteString(strconv.Itoa(p.LastPage))
            html.WriteString("</a></li>")
        }
    }

    html.WriteString("\n</ul></nav>") 
    return html.String()
}