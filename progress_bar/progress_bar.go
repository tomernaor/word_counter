/*
 * Copyright (c) 2025 Twelve
 * All rights reserved.
 *
 * This software is the confidential and proprietary information
 * of Twelve. ("Confidential Information"). You shall not
 * disclose such Confidential Information and shall use it only in
 * accordance with the terms of the license agreement you entered
 * into with Twelve.
 */

package progress_bar

import (
	"fmt"
	"sync/atomic"

	"github.com/vbauerster/mpb/v8"
	"github.com/vbauerster/mpb/v8/decor"
)

const (
	Green = "\033[32m"
	Reset = "\033[0m"
)

type ProgressBars struct {
	p          *mpb.Progress
	total      *mpb.Bar
	errorCount int64
}

// NewProgressBars creates new progress bars
func NewProgressBars(total int64) *ProgressBars {
	progressBars := &ProgressBars{p: mpb.New(mpb.WithWidth(60))}

	// create progress bar
	progressBars.total = progressBars.p.AddBar(int64(total),
		mpb.PrependDecorators(
			decor.Name(Green+"Fetching: "+Reset),
			decor.CountersNoUnit("%d/%d"),
			decor.Any(func(statistics decor.Statistics) string {
				// this will be called each render
				return fmt.Sprintf(" Errors:%d", atomic.LoadInt64(&progressBars.errorCount))
			}),
		),
	)

	return progressBars
}

// IncProgress increment progress
func (p *ProgressBars) IncProgress() {
	p.total.Increment()
}

// IncError increments error count
func (b *ProgressBars) IncError() {
	atomic.AddInt64(&b.errorCount, 1)
}

// Wait waits for bars to complete
func (p *ProgressBars) Wait() {
	p.p.Wait()
}
