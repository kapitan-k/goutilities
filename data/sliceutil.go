package data

func CopyBuf(in []byte) (out []byte) {
	out = make([]byte, len(in))
	copy(out, in)
	return
}
