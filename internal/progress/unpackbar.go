package progress

import (
	"io"
	"math"
	"time"

	"github.com/ActiveState/cli/internal/logging"
	"github.com/vbauerster/mpb/v4"
	"github.com/vbauerster/mpb/v4/decor"
)

// UnpackBar wraps a regular progress bar that is specifically for unpacking
// Note this peculiarities about unpacking:
// - the number of bytes to be read is known.
// - the number of bytes to be written is unknown.
// This struct therefore stores an artificial total (2% about the total bytes to be read)
// Run Complete when you have written all files to disc, and the bar will fill the remaining 2%.
type UnpackBar struct {
	bar           *mpb.Bar
	total         int64 // the total counter for when we are done unpacking
	targetPercent int   // what percentage we should have reached when we are done unpacking
	scaleFactor   int
	barTotal      int64
}

// NewUnpackBar creates a progressbar for unpacking an archive
// The progress bar will stop at targetPercent on unpacking, leaving room to report progress on
// extra tasks
func NewUnpackBar(bytesToRead int64, p *Progress, targetPercent int) *UnpackBar {
	// add a 2% buffer for finishing the last writes
	total := int64(float64(bytesToRead) * 1.02)
	barTotal := int64(float64(total) * 100 / float64(targetPercent))
	return &UnpackBar{bar: p.progress.AddBar(barTotal,
		mpb.BarRemoveOnComplete(),
		mpb.PrependDecorators(
			// synchronize the width with the two total bar decorations
			decor.Merge(
				decor.OnComplete(decor.Spinner(nil, decor.WCSyncSpace), ""),
				decor.WCSyncSpace),
			// decor.Counters(decor.UnitKiB, "%6.1f / %6.1f", decor.WC{W: 20, C: decor.DidentRight}),
		),
		mpb.AppendDecorators(
			decor.Percentage(decor.WC{W: 5}),
		)),
		total:         total,
		targetPercent: targetPercent,
		scaleFactor:   1,
		barTotal:      barTotal,
	}
}

// Complete completes the progress to 100% and should be called after all files are written to disc
func (upb *UnpackBar) Complete() {
	upb.bar.SetCurrent(upb.total)
}

// ReScale sets a scaling factor that the specified number of increments fill the bar up to 100%
func (upb *UnpackBar) ReScale(countsRemaining int) {

	remainingRatio := (100.0 - float64(upb.targetPercent)) / 100.0
	exactScalingFactor := remainingRatio * float64(upb.barTotal) / float64(countsRemaining)

	// round up the scaling factor, so we can use integral steps to get to 100%
	upb.scaleFactor = int(math.Ceil(exactScalingFactor))

	newTotal := int64(float64(countsRemaining) * float64(upb.scaleFactor) / remainingRatio)
	upb.bar.SetTotal(newTotal, false)
	newCurrent := int64(float64(upb.targetPercent) / 100.0 * float64(newTotal))
	upb.bar.SetCurrent(newCurrent)
	logging.Debug("new Total: %d, new current: %d, scaleFactor: %d\n", newTotal, newCurrent, upb.scaleFactor)

	upb.total = newTotal
	upb.barTotal = newTotal
	upb.targetPercent = 100
}

// Increment increments the counter manually this can be used after ReScaleBar to report progress
func (upb *UnpackBar) Increment() {
	upb.bar.IncrBy(upb.scaleFactor)

}

// NewProxyReader wraps a Reader with functionality that automatically updates
// the bar with progress about how many bytes have been read from the underlying
// reader so far.
func (upb *UnpackBar) NewProxyReader(r io.ReadCloser) *ProxyReader {
	return &ProxyReader{
		ReadCloser: r,
		bar:        upb.bar,
		iT:         time.Now(),
		complete:   upb.Complete,
	}
}
