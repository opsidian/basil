// Copyright (c) 2017 Opsidian Ltd.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package job

import (
	"errors"
	"sync/atomic"

	"github.com/opsidian/basil/basil"
)

// Scheduler handles workers and schedules jobs
type Scheduler struct {
	logger       basil.Logger
	maxWorkers   int
	maxQueueSize int
	workers      []Worker
	jobQueue     chan basil.Job
	quit         chan bool
	lastID       int64
}

// NewScheduler creates a new scheduler instance
func NewScheduler(logger basil.Logger, maxWorkers int, maxQueueSize int) *Scheduler {
	return &Scheduler{
		logger:       logger,
		workers:      make([]Worker, maxWorkers),
		maxWorkers:   maxWorkers,
		maxQueueSize: maxQueueSize,
		jobQueue:     make(chan basil.Job, maxQueueSize),
		quit:         make(chan bool),
	}
}

// Start creates and starts the workers
func (s *Scheduler) Start() {
	s.logger.Debug().
		Int("workers", s.maxWorkers).
		Int("queueSize", s.maxQueueSize).
		Msg("job scheduler starting")

	for i := 0; i < s.maxWorkers; i++ {
		s.workers[i] = NewWorker(s.logger, s.jobQueue)
		s.workers[i].Start()
	}
}

// Stop stops all the workers and the dispatcher process
func (s *Scheduler) Stop() {
	go func() {
		s.quit <- true
	}()

	for i := 0; i < s.maxWorkers; i++ {
		s.workers[i].Stop()
	}

	s.logger.Debug().Msg("job scheduler stopped")
}

// Schedule schedules a new job
func (s *Scheduler) ScheduleJob(job basil.Job) error {
	job.SetJobID(int(atomic.AddInt64(&s.lastID, 1)))

	if job.Lightweight() {
		go func() {
			job.Run()
		}()
		return nil
	} else {
		select {
		case s.jobQueue <- job:
			return nil
		case <-s.quit:
			return errors.New("job scheduler is shutting down")
		}
	}
}
