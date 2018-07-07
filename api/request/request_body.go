package request

import "io"

type Body interface {
	IoReaderMap() (*map[string]io.Reader, error)
}
