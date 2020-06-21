package core

import (
	"go-dom-parser/api/sockets"
	"go-dom-parser/domain"
)

// Processor - processor instance
type Processor struct {
	Socket  *sockets.Socket
	Storage *domain.Storage
}
