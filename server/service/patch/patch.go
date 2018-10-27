package patch

import (
	"path/filepath"
	"plugin"
)

var defDir = "patch"

//运行补丁 补丁必须有一个RunPatch()的方法
func RunPatch(path *string) (err error) {
	funcName := "RunPatch"
	fullPath := filepath.Join(defDir, *path)
	p, err := plugin.Open(fullPath)
	if err != nil {
		return
	}

	newF, err := p.Lookup(funcName)
	if err != nil {
		return
	}
	newF.(func())()
	return
}
