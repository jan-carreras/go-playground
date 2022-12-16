package paasio

import (
	"io"
	"sync"
)

type ReaderWriterCounter struct {
	WriterCounter
	ReaderCounter
}

type WriterCounter struct {
	byteCount int
	callCount int
	w         io.Writer
	mux       sync.Mutex
}

type ReaderCounter struct {
	byteCount int
	callCount int
	r         io.Reader
	mux       sync.Mutex
}

func (c *ReaderCounter) Read(p []byte) (n int, err error) {
	bytes, err := c.r.Read(p)
	c.mux.Lock()
	defer c.mux.Unlock()
	c.byteCount += bytes
	c.callCount += 1
	return bytes, err
}

func (c *ReaderCounter) ReadCount() (n int64, nops int) {
	c.mux.Lock()
	defer c.mux.Unlock()
	return int64(c.byteCount), c.callCount
}

func (c *WriterCounter) Write(p []byte) (n int, err error) {
	bytes, err := c.w.Write(p)
	c.mux.Lock()
	defer c.mux.Unlock()
	c.byteCount += bytes
	c.callCount += 1
	return bytes, err
}

func (c *WriterCounter) WriteCount() (n int64, nops int) {
	c.mux.Lock()
	defer c.mux.Unlock()
	return int64(c.byteCount), c.callCount
}

func NewWriteCounter(w io.Writer) WriteCounter {
	return &WriterCounter{w: w}
}

func NewReadWriteCounter(rw io.ReadWriter) ReadWriteCounter {
	return &ReaderWriterCounter{
		WriterCounter: WriterCounter{
			w: rw.(io.Writer),
		},
		ReaderCounter: ReaderCounter{
			r: rw.(io.Reader),
		},
	}
}

func NewReadCounter(r io.Reader) ReadCounter {
	return &ReaderCounter{
		r: r,
	}
}
