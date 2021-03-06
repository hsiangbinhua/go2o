package partner

import (
	"com/ording"
	"com/ording/dao"
	"encoding/json"
	"net/http"
	"ops/cf/app"
	"ops/cf/web/ui/tree"
)

type commC struct {
	*ording.Webcgi
	app.Context
}

func (this *mainC) GeoLocation(w http.ResponseWriter, r *http.Request) {
	this.Webcgi.GeoLocation(w, r)
}

//地区Json
func (this *mainC) ChinaJson(w http.ResponseWriter, r *http.Request) {
	var node *tree.TreeNode = dao.Common().GetChinaTree()
	json, _ := json.Marshal(node)
	w.Write(json)
}
