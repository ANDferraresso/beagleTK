package mux

import (
	"database/sql"
	"regexp"
	"slices"
	"strings"

	"github.com/fasthttp/session/v2"

	_ "github.com/go-sql-driver/mysql"
	"github.com/valyala/fasthttp"
)

type BaseMux struct {
	DbConn  *sql.DB
	Session *session.Session
}

func (bm *BaseMux) Mux(ctx *fasthttp.RequestCtx, setting *map[string]string, routes *map[string]Route, urlDictio *map[string]map[string]string) {
	var ctrl Controller
	var foundRoute []string
	var idx int
	var path string
	var pathArray []string
	var params map[string]string
	var tree map[string]Route

	// Store Database connection.
	ctx.SetUserValue("dbConn", bm.DbConn)
	//	// Store Session connection.
	//	ctx.SetUserValue("session", bm.Session)

	ctrl = Controller{Method: "", Querys: nil, Handler: nil}
	foundRoute = nil
	idx = 0
	path = strings.Trim(string(ctx.Path()), "/")
	pathArray = strings.Split(path, "/")
	params = make(map[string]string)
	tree = (*routes)["ROOT"].Sons
	actLangsArr := strings.Split((*setting)["actived_langs"], "|")

Loop2:
	for ok := true; ok; ok = (idx < len(pathArray) && tree != nil) {
		var flag = false
	Loop1:
		for k, _ := range tree {
			var t = pathArray[idx]
			// Verifica se ci sono parametri dentro a < >
			if len(k) >= 2 && (k[0:1] == "<" && k[len(k)-1:] == ">") {
				// Espressione regolare (se c'è)
				if strings.Index(k, ":") > 0 {
					// C'è espressione regolare
					parts := strings.Split(k[1:len(k)-1], ":")
					if parts[0] == "lang" && (*setting)["lang_in_url"] == "1" {
						re := regexp.MustCompile("^" + parts[1] + "$")
						if re.Match([]byte(t)) {
							if slices.Contains(actLangsArr, t) {
								params["lang"] = t
								foundRoute = append(foundRoute, k)
								flag = true
							} else {
								flag = false
							}
						} else {
							flag = false
						}
					} else {
						re := regexp.MustCompile("^" + parts[1] + "$")
						if re.Match([]byte(t)) {
							params[parts[0]] = t
							foundRoute = append(foundRoute, k)
							flag = true
						} else {
							flag = false
						}
					}
				} else {
					// Non c'è espressione regolare
					if k[1:len(k)-1] == "lang" && (*setting)["lang_in_url"] == "1" {
						if slices.Contains(actLangsArr, t) {
							params["lang"] = t
							foundRoute = append(foundRoute, k)
							flag = true
						} else {
							flag = false
						}
					} else {
						params[k[1:len(k)-1]] = t
						foundRoute = append(foundRoute, k)
						flag = true
					}
				}
			} else if len(k) >= 2 && k[0:1] == "{" && k[len(k)-1:] == "}" && (*setting)["lang_in_url"] == "1" {
				// Verifica se ci sono parti da tradurre dentro a { }
				if (*urlDictio)[params["lang"]][k[1:len(k)-1]] == t {
					foundRoute = append(foundRoute, k)
					flag = true
				} else {
					flag = false
				}
			} else if k == t {
				foundRoute = append(foundRoute, k)
				flag = true
			} else {
				flag = false
			}

			//
			if flag {
				if idx == len(pathArray)-1 {
					if tree[k].Ctrl.Handler == nil {

					} else {
						methodSlice := strings.Split(tree[k].Ctrl.Method, "|")
						f := false
						for _, method := range methodSlice {
							if method == string(ctx.Method()) {
								f = true
								break
							}
						}
						if f {
							ctrl = tree[k].Ctrl
						}
					}
					break Loop2
				} else {
					tree = tree[k].Sons
					break Loop1
				}
			} else {
				continue
			}
		}

		idx++
	}

	//
	if ctrl.Handler == nil {
		// ctx.SetStatusCode(fasthttp.StatusNotFound)
		ctx.SetUserValue("foundRoute", nil)
		ctx.SetUserValue("pathArray", nil)
		ctx.SetUserValue("params", nil)
		ctx.SetContentType("text/html")
		ctx.Error("404 Not found", fasthttp.StatusNotFound)
	} else {
		ctx.SetUserValue("foundRoute", foundRoute)
		ctx.SetUserValue("pathArray", pathArray)
		ctx.SetUserValue("params", params)
		ctx.SetStatusCode(fasthttp.StatusOK)
		ctrl.Handler(ctx, bm.Session)
	}
	defer bm.DbConn.Close()
}
