// Salt provides common logger interface.
package ajika

type Salt interface {
	Debugf(template string, args ...interface{})
	Debug(args ...interface{})
	Error(args ...interface{})
}

type MockSalt struct {
}

func (salt MockSalt) Debugf(template string, args ...interface{}) {
}

func (salt MockSalt) Debug(args ...interface{}) {
}

func (salt MockSalt) Error(args ...interface{}) {
}
