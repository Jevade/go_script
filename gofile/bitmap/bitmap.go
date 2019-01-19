package bitmap

import "errors"

//Bitmap is to sort with bitmap
type Bitmap struct {
	data []uint64
	size uint
}

const (
	byteSize     = 64
	byteSizeStep = 6
	mask         = 0x3f
)

//NewBitmap is to return a Bitmap
func NewBitmap(size uint) *Bitmap {
	length := size >> byteSizeStep
	p := Bitmap{data: make([]uint64, length, length), size: size}
	return &p
}

//Set is to set 1 in indexth bit
func (bitmap *Bitmap) Set(index uint) error {
	if index > bitmap.size {
		return errors.New("Length out of range")
	}
	bitmap.data[index>>byteSizeStep] |= (1 << (index & mask))
	return nil
}

//Cle is to set 0 in indexth bit
func (bitmap *Bitmap) Cle(index uint) error {
	if index > bitmap.size {
		return errors.New("Length out of range")
	}
	bitmap.data[index>>byteSizeStep] |= ^(1 << (index & mask))
	return nil
}

// IsSet to determine if indexth bit is set
func (bitmap *Bitmap) IsSet(index uint) bool {
	return (bitmap.data[index>>byteSizeStep] & (1 << (index & mask))) > 0
}
