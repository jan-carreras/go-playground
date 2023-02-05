package variablelengthquantity

import (
	"errors"
)

func EncodeVarint(input []uint32) (encoded []byte) {
	for _, i := range input {
		e := []byte{byte(i % 128)}
		for i >>= 7; i != 0; i >>= 7 {
			e = append([]byte{128 + byte(i%128)}, e...)
		}
		encoded = append(encoded, e...)
	}

	return encoded
}

func DecodeVarint(in []byte) (rsp []uint32, err error) {
	complete := false // Keeps track of the start/end of the sequence
	d := uint32(0)    // Accumulator to parse each individual digit

	for _, b := range in {
		d += uint32(b &^ 128) // Ignore 8th bit, add the rest

		complete = b&128 == 0 // Is a complete sequence?
		if complete {
			rsp, d = append(rsp, d), 0 // Store result, and restart accumulator
			continue
		}
		d <<= 7

	}
	if !complete {
		return nil, errors.New("high bits are messed up")
	}

	return rsp, nil
}
