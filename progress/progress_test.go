package progress_test

import (
	"bytes"
	"io"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/dGilli/terminal-ui/progress"
)

func TestProgressUpdate(t *testing.T) {
	testCases := []struct {
		name     string
        progress int
        expects  int // progress bar length relative to terminal wdith
	}{
        {
            name:     "should update progress correctly",
            progress: 50,
            expects:  50,
        },
        {
            name:     "should complete correctly",
            progress: 100,
            expects:  100,
        },
        {
            name:     "should clamp progress to total",
            progress: 150,
            expects:  100,
        },
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			buf := &bytes.Buffer{}

			p := progress.New(progress.Config{
				Writer: buf,
                Width: 100,
			})

            p.Update(tc.progress)

			time.Sleep(time.Second * 1)

			data, err := io.ReadAll(buf)
			assert.NoError(t, err)

			assert.Equal(t, strings.Repeat("#", tc.expects), string(data))
		})
	}
}

func TestProgressWorksAsync(t *testing.T) {
}

func TestAdjustsToTerminalSize(t *testing.T) {
}
