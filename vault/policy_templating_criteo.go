package vault

import (
	"context"
	"strings"
)

func criteoFuncFactory(ctx context.Context) interface{} {
	return func() interface{} {
		return map[string]interface{}{
			"string": map[string]interface{}{
				"replace": strings.Replace,
			},
		}
	}
}

func init() {
	templateFuncFactories["criteo"] = criteoFuncFactory
}
