package extractor

import "io"

type Extractor interface {
	GetFilesBody() map[string]io.ReadCloser
}
