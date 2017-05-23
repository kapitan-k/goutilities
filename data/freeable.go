package data

type FnFreeData func(data []byte)

type FreeableData interface {
	Data() []byte
	Free()
}

type FunctionFreeableData struct {
	Dat []byte
	fn  FnFreeData
}

func FunctionFreeableDataCreate(dat []byte, fn FnFreeData) (self FunctionFreeableData) {
	self.Dat = dat
	self.fn = fn
	return
}

func (self *FunctionFreeableData) Data() []byte {
	return self.Dat
}

func (self *FunctionFreeableData) Free() {
	if self.fn != nil {
		self.fn(self.Data())
	}
}
