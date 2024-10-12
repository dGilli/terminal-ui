package progress

import (
	"context"
	"io"
	"math"
	"os"
	"sync"
	"time"

	"golang.org/x/term"
)

type Progress struct {
	writer     io.Writer
	glyph      rune
    cols       int
    current    int
	cancelFunc context.CancelFunc
	doneCh     chan struct{}
	lock       sync.RWMutex
}

type Config struct {
	Writer io.Writer
	Glyph  *rune
    Cols   *int
}

func New(cfg Config) *Progress {
	p := &Progress{
		writer: os.Stderr,
		glyph:  '#',
        cols:   80,
	}

	if cfg.Writer != nil {
		p.writer = cfg.Writer
	}

	if cfg.Glyph != nil {
		p.glyph = *cfg.Glyph
	}

	if cfg.Cols != nil {
		p.cols = *cfg.Cols
	} else if term.IsTerminal(0) {
        width, _, err := term.GetSize(0)
        if err == nil {
            p.cols = width
        }
    }

	return p
}

type Bar struct {
    cols int32
}

func (b *Bar) Update(current float64) {
    numGlyphs := int(math.Floor(current * float64(b.cols / 100)))

    ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		time.Sleep(5 * time.Second)
		done <- true
	}()

	for {
		select {
		case t := <-ticker.C:
            p.writer.Write([]byte("#"))
		case <-done:
			ticker.Stop()
			return
		}
	}

    for i := 0; i < numGlyphs; i++ {
        b.writer.Write([]byte("#"))
	    time.Sleep(time.Millisecond * 100)
        if i == 10 {
            p.writer.Write([]byte("\n"))
        }
    }
}

func (p *Progress) Update(current int) {

    /*
	if p.isComplete() {
		return
	}

    complete := 80
    if current > complete {
		current = complete
	}

	p.current = current
    */

    //repeatedRune := strings.Repeat(string(p.glyph), current / 100 * 80)

    /*
    for i := 0; i < 10; i++ {
        p.writer.Write([]byte("#"))
	    time.Sleep(time.Millisecond * 100)
        if i == 10 {
            p.writer.Write([]byte("\n"))
        }
    }
    p.writer.Write([]byte("\n"))
*/

    /*
######################################################### / 100%
    */

    ticker := time.NewTicker(1 * time.Second)
	done := make(chan bool)

	go func() {
		time.Sleep(5 * time.Second)
		done <- true
	}()

	for {
		select {
		case t := <-ticker.C:
            p.writer.Write([]byte("#"))
		case <-done:
			ticker.Stop()
			return
		}
	}

    //repeatedRune := strings.Repeat(string(p.glyph), 10)
    //p.writer.Write([]byte(repeatedRune))
}

func (p *Progress) isComplete() bool {
	p.lock.RLock()
	defer p.lock.RUnlock()

	return p.doneCh != nil
}
