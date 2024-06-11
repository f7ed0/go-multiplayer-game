package objects

type IntBuffer struct {
	buf  []int
	size int
}

func NewIntBuffer(size int) IntBuffer {
	return IntBuffer{
		buf:  []int{},
		size: size,
	}
}

func (buf *IntBuffer) Stack(item int) {
	if len(buf.buf) < buf.size {
		buf.buf = append(buf.buf, item)
		return
	}
	buf.buf = buf.buf[1:]
	buf.buf = append(buf.buf, item)
}

func (buf *IntBuffer) GetValues() []int {
	var ret []int
	copy(ret, buf.buf)
	return ret
}

func (buf *IntBuffer) GetMean() float64 {
	if len(buf.buf) <= 0 {
		return 0
	}
	sum := 0
	for _, it := range buf.buf {
		sum += it
	}
	return float64(sum) / float64(len(buf.buf))
}
