package utils

import "io"

// GeneralReader ..
type GeneralReader struct {
	r io.Reader
}

// NewGeneralReader ..
func NewGeneralReader(r io.Reader) *GeneralReader {
	return &GeneralReader{r: r}
}

func (g *GeneralReader) Read(p []byte) (n int, err error) {
	return g.r.Read(p)
}

func (g *GeneralReader) ReadByte() (byte, error) {
	p := make([]byte, 1)
	_, err := g.r.Read(p)
	return p[0], err
}

// GeneralWriter ..
type GeneralWriter struct {
	w io.Writer
}

// NewGeneralWriter ..
func NewGeneralWriter(w io.Writer) *GeneralWriter {
	return &GeneralWriter{w: w}
}

func (g *GeneralWriter) Write(p []byte) (n int, err error) {
	return g.w.Write(p)
}

func (g *GeneralWriter) WriteByte(c byte) error {
	_, err := g.w.Write([]byte{c})
	return err
}
