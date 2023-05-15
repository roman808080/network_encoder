package network_encoder

type INetworkWriter interface {
	Write(source any) error
}

type INetworkReader interface {
	Write(source any) error
}

type INetworkEncoder interface {
	GetNetworkWriter() INetworkWriter
	GetNetworkReader() INetworkReader
}
