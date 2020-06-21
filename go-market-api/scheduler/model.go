package scheduler

import (
	"go-dom-parser/domain"
	"go-dom-parser/scheduler/jobs"
)

// Scheduler - job runner to make some periodic processing
type Scheduler struct {
	Jobs    []jobs.Job
	Storage *domain.Storage
}
