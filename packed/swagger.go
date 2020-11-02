package packed

import "github.com/gogf/gf/os/gres"

func init() {
	if err := gres.Add("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"); err != nil {
		panic("add binary content to resource manager failed: " + err.Error())
	}
}
