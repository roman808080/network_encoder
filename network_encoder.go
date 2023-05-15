package network_encoder

import (
	"encoding/gob"
	"io"
)

type INetworkWriter interface {
	Write(source any) error
}

type INetworkReader interface {
	Read(source any) error
}

type INetworkEncoder interface {
	GetNetworkWriter() INetworkWriter
	GetNetworkReader() INetworkReader
}

// GobWriter
type GobWriter struct {
	enc *gob.Encoder
}

func NewGobWriter(w io.Writer) *GobWriter {
	return &GobWriter{enc: gob.NewEncoder(w)}
}

func (e *GobWriter) Write(source any) error {
	return e.enc.Encode(source)
}

// GobReader
type GobReader struct {
	dec *gob.Decoder
}

func NewGobReader(r io.Reader) *GobReader {
	return &GobReader{dec: gob.NewDecoder(r)}
}

func (reader *GobReader) Read(destination any) error {
	return reader.dec.Decode(destination)
}

type GobNetworkEncoder struct {
}

func NewGobNetworkEncoder() GobNetworkEncoder {
	return GobNetworkEncoder{}
}

func (*GobNetworkEncoder) GetNetworkWriter(w io.Writer) INetworkWriter {
	return NewGobWriter(w)
}

func (*GobNetworkEncoder) GetNetworkReader(r io.Reader) INetworkReader {
	return NewGobReader(r)
}
