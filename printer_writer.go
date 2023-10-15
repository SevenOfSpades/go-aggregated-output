package output

import (
	"fmt"
	"io"
	"time"
)

type writerPrinter struct {
	writer io.Writer
}

func NewWriterPrinter(writer io.Writer) Printer {
	return &writerPrinter{writer: writer}
}

func (h *writerPrinter) Debug(record Record, _ Output) {
	_, _ = h.writer.Write([]byte(fmt.Sprintf("DEBUG | %s | %s\n", time.Now().Format(time.DateTime), record.Message())))
}

func (h *writerPrinter) Info(record Record, _ Output) {
	_, _ = h.writer.Write([]byte(fmt.Sprintf("INFO | %s | %s\n", time.Now().Format(time.DateTime), record.Message())))
}

func (h *writerPrinter) Warning(record Record, _ Output) {
	_, _ = h.writer.Write([]byte(fmt.Sprintf("WARNING | %s | %s\n", time.Now().Format(time.DateTime), record.Message())))
}

func (h *writerPrinter) Error(record Record, _ Output) {
	_, _ = h.writer.Write([]byte(fmt.Sprintf("ERROR | %s | %s\n", time.Now().Format(time.DateTime), record.Message())))
}
