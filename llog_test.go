package llog_test

import (
	"testing"

	"github.com/andrewslotin/llog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWriter_Write(t *testing.T) {
	messages := []string{
		"fatal: something went wrong",
		"missing something",
		"error doing something",
		"failed to do something",
		"warn: something went wrong",
		"warning: something went wrong",
		"debug is set to true",
		"random stuff",
		"debug: here is what's happened",
	}

	examples := map[llog.Level][]string{
		llog.FatalLevel: messages[:2],
		llog.ErrorLevel: messages[:4],
		llog.WarnLevel:  messages[:6],
		llog.InfoLevel:  messages[:8],
		llog.DebugLevel: messages,
	}

	for lvl, expected := range examples {
		t.Run(lvl.String(), func(t *testing.T) {
			m := &writerMock{}
			w := llog.NewWriter(m, lvl)

			for _, msg := range messages {
				_, err := w.Write([]byte(msg))
				require.NoError(t, err)
			}

			assert.Equal(t, expected, m.Messages)
		})
	}
}

type writerMock struct {
	Messages []string
}

func (w *writerMock) Write(p []byte) (n int, err error) {
	w.Messages = append(w.Messages, string(p))
	return len(p), nil
}
