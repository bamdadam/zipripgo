package compression

type Encorder interface {
	Encode([]byte) []byte
}

type Decoder interface {
	Decode([]byte) []byte
}
