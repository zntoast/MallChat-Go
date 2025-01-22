package ercode

import (
	"encoding/json"
	"math"
	"time"
	"unicode/utf8"

	"go.uber.org/zap"
	"go.uber.org/zap/buffer"
	"go.uber.org/zap/zapcore"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

var (
	errBuffer        = buffer.NewPool()
	CodeError        = 5000000000
	CodeUnknownError = 9999999999
)

const _hex = "0123456789abcdef"

type Fields []zap.Field

func (fs Fields) marshalJSON(buf *buffer.Buffer) {
	for i, field := range fs {
		if i != 0 {
			buf.AppendByte(',')
		}
		buf.AppendByte('"')
		buf.AppendString(field.Key)
		switch field.Type {
		case zapcore.BoolType:
			buf.AppendString(`":`)
			buf.AppendBool(field.Integer == 1)
		case zapcore.Int8Type,
			zapcore.Int16Type,
			zapcore.Int32Type,
			zapcore.Int64Type:
			buf.AppendString(`":`)
			buf.AppendInt(field.Integer)
		case zapcore.Uint8Type,
			zapcore.Uint16Type,
			zapcore.Uint32Type,
			zapcore.Uint64Type:
			buf.AppendString(`":`)
			buf.AppendUint(uint64(field.Integer))
		case zapcore.StringType:
			buf.AppendString(`":"`)
			safeAddString(buf, field.String)
			buf.AppendByte('"')
		case zapcore.Float32Type:
			buf.AppendString(`":`)
			val := math.Float32frombits(uint32(field.Integer))
			switch {
			case math.IsNaN(float64(val)):
				buf.AppendString(`"NaN"`)
			case math.IsInf(float64(val), 1):
				buf.AppendString(`"+Inf"`)
			case math.IsInf(float64(val), -1):
				buf.AppendString(`"-Inf"`)
			default:
				buf.AppendFloat(float64(val), 32)
			}
		case zapcore.Float64Type:
			buf.AppendString(`":`)
			val := math.Float64frombits(uint64(field.Integer))
			switch {
			case math.IsNaN(val):
				buf.AppendString(`"NaN"`)
			case math.IsInf(val, 1):
				buf.AppendString(`"+Inf"`)
			case math.IsInf(val, -1):
				buf.AppendString(`"-Inf"`)
			default:
				buf.AppendFloat(val, 64)
			}
		case zapcore.ArrayMarshalerType:
			buf.AppendString(`":[`)
			arrayMarshaler := field.Interface.(zapcore.ArrayMarshaler)
			_ = arrayMarshaler.MarshalLogArray(&arrayEncoder{buf: buf})
			buf.AppendByte(']')
		case zapcore.ErrorType:
			err := field.Interface.(error)
			buf.AppendString(`":"`)
			safeAddString(buf, err.Error())
			buf.AppendByte('"')
		default:
			data, err := json.Marshal(field.Interface)
			if err != nil {
				buf.AppendString(`":null`)
				return
			}
			buf.AppendString(`":"`)
			safeAddString(buf, string(data))
			buf.AppendByte('"')
		}
	}

}

type errorEntry struct {
	code    int
	message string
	detail  Fields
}

func (err *errorEntry) Error() string {
	buf := errBuffer.Get()
	defer buf.Free()
	p := message.NewPrinter(language.English)
	err.marshalJSON(p, buf)
	return buf.String()
}

func New(code int, message string, detail ...zap.Field) error {
	return &errorEntry{
		code:    code,
		message: message,
		detail:  detail,
	}
}

func (e *errorEntry) marshalJSON(printer *message.Printer, buf *buffer.Buffer) {
	buf.AppendString(`{"code":`)
	buf.AppendInt(int64(e.code))
	buf.AppendString(`,"msg":"`)
	if printer != nil {
		safeAddString(buf, printer.Sprintf(e.message))
	} else {
		safeAddString(buf, e.message)
	}
	buf.AppendByte('"')
	if len(e.detail) != 0 {
		buf.AppendString(`,"debug":{`)
		e.detail.marshalJSON(buf)
		buf.AppendByte('}')
	}
	buf.AppendByte('}')
}

func safeAddString(buf *buffer.Buffer, s string) {
	for i := 0; i < len(s); {
		if tryAddRuneSelf(buf, s[i]) {
			i++
			continue
		}
		r, size := utf8.DecodeRuneInString(s[i:])
		if tryAddRuneError(buf, r, size) {
			i++
			continue
		}
		buf.AppendString(s[i : i+size])
		i += size
	}
}

func tryAddRuneSelf(buf *buffer.Buffer, b byte) bool {
	if b >= utf8.RuneSelf {
		return false
	}
	if 0x20 <= b && b != '\\' && b != '"' {
		buf.AppendByte(b)
		return true
	}
	switch b {
	case '\\', '"':
		buf.AppendByte('\\')
		buf.AppendByte(b)
	case '\n':
		buf.AppendByte('\\')
		buf.AppendByte('n')
	case '\r':
		buf.AppendByte('\\')
		buf.AppendByte('r')
	case '\t':
		buf.AppendByte('\\')
		buf.AppendByte('t')
	default:
		// Encode bytes < 0x20, except for the escape sequences above.
		buf.AppendString(`\u00`)
		buf.AppendByte(_hex[b>>4])
		buf.AppendByte(_hex[b&0xF])
	}
	return true
}

func tryAddRuneError(buf *buffer.Buffer, r rune, size int) bool {
	if r == utf8.RuneError && size == 1 {
		buf.AppendString(`\ufffd`)
		return true
	}
	return false
}

type arrayEncoder struct {
	idx int
	buf *buffer.Buffer
}

func (ae *arrayEncoder) AppendBool(v bool) {
	if ae.idx != 0 {
		ae.buf.AppendByte(',')
	}
	ae.buf.AppendBool(v)
	ae.idx++
}

func (ae *arrayEncoder) AppendByteString([]byte) {

}

func (ae *arrayEncoder) AppendComplex128(complex128) {

}

func (ae *arrayEncoder) AppendComplex64(complex64) {

}

func (ae *arrayEncoder) appendFloat(val float64, bitSize int) {
	if ae.idx != 0 {
		ae.buf.AppendByte(',')
	}
	switch {
	case math.IsNaN(val):
		ae.buf.AppendString(`"NaN"`)
	case math.IsInf(val, 1):
		ae.buf.AppendString(`"+Inf"`)
	case math.IsInf(val, -1):
		ae.buf.AppendString(`"-Inf"`)
	default:
		ae.buf.AppendFloat(val, bitSize)
	}
	ae.idx++
}

func (ae *arrayEncoder) AppendFloat64(v float64) {
	ae.appendFloat(v, 64)
}

func (ae *arrayEncoder) AppendFloat32(v float32) {
	ae.appendFloat(float64(v), 32)
}

func (ae *arrayEncoder) AppendInt(v int) {
	ae.AppendInt64(int64(v))
}

func (ae *arrayEncoder) AppendInt64(v int64) {
	if ae.idx != 0 {
		ae.buf.AppendByte(',')
	}
	ae.buf.AppendInt(v)
	ae.idx++
}

func (ae *arrayEncoder) AppendInt32(v int32) {
	ae.AppendInt64(int64(v))
}

func (ae *arrayEncoder) AppendInt16(v int16) {
	ae.AppendInt64(int64(v))
}

func (ae *arrayEncoder) AppendInt8(v int8) {
	ae.AppendInt64(int64(v))
}

func (ae *arrayEncoder) AppendString(v string) {
	if ae.idx != 0 {
		ae.buf.AppendByte(',')
	}
	ae.buf.AppendByte('"')
	safeAddString(ae.buf, v)
	ae.buf.AppendByte('"')
	ae.idx++
}

func (ae *arrayEncoder) AppendUint(v uint) {
	ae.AppendUint64(uint64(v))
}

func (ae *arrayEncoder) AppendUint64(v uint64) {
	if ae.idx != 0 {
		ae.buf.AppendByte(',')
	}
	ae.buf.AppendUint(v)
	ae.idx++
}

func (ae *arrayEncoder) AppendUint32(v uint32) {
	ae.AppendUint64(uint64(v))
}

func (ae *arrayEncoder) AppendUint16(v uint16) {
	ae.AppendUint64(uint64(v))
}

func (ae *arrayEncoder) AppendUint8(v uint8) {
	ae.AppendUint64(uint64(v))
}

func (ae *arrayEncoder) AppendUintptr(uintptr) {

}

func (ae *arrayEncoder) AppendDuration(time.Duration) {

}

func (ae *arrayEncoder) AppendTime(time.Time) {

}

func (ae *arrayEncoder) AppendArray(zapcore.ArrayMarshaler) error {
	return nil
}

func (ae *arrayEncoder) AppendObject(zapcore.ObjectMarshaler) error {
	return nil
}

func (ae *arrayEncoder) AppendReflected(value interface{}) error {
	return nil
}

func GetCode(e error) int {
	if x, ok := e.(*errorEntry); ok {
		return x.code
	}
	return CodeUnknownError
}

func GetMessage(e error) string {
	if x, ok := e.(*errorEntry); ok {
		return x.message
	}
	return ""
}
