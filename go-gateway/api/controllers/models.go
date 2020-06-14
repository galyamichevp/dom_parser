package controllers

import "go-dom-parser/core"

// Controller - container to coomunicate with other infrastructure components
type Controller struct {
	RChan chan string
	Proc  *core.Processor
}
