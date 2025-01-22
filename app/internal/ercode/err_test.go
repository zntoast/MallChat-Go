package ercode

import (
	"errors"
	"fmt"
	"net/http"
	"testing"

	"go.uber.org/zap"
)

func TestFieldsMarshal(t *testing.T) {
	// fields := Fields{}
	// fields = append(fields, zap.Error(errors.New("test error")))
	// fields = append(fields, zap.String("key", "value"))
	// fields = append(fields, zap.Int("key2", 123))
	// fields = append(fields, zap.Bool("key3", true))
	// fields = append(fields, zap.Float64("key4", 3.14))
	// fields = append(fields, zap.String("key5", "中文"))
	// fields = append(fields, zap.String("key6", "😀"))
	// str := fields.marshalJSON()
	// t.Log()
	// t.Logf("%s", string(str))
}

func TestErrPrint(t *testing.T) {
	err := New(http.StatusForbidden, "账号错误", zap.String("name", "错误"), zap.Error(errors.New("test error")))
	fmt.Println(err)
}
