package symmetric

import "github.com/Quartz-Vision/goslice"

func Encode(data []byte, key []byte) (err error) {
	if len(data) > len(key) {
		return ErrWrongKeyLength
	}

	goslice.Xor(data, key)
	return nil
}

func Decode(data []byte, key []byte) (err error) {
	if len(data) > len(key) {
		return ErrWrongKeyLength
	}

	goslice.Xor(data, key)
	return nil
}
