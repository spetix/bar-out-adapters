package protocols

import "os"

type baseOutput struct {
	Device *os.File
}
