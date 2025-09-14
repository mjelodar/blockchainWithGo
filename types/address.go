package types

import "encoding/hex"

type Address [20]uint8

func (a Address) toSlice() []byte {
	slice := make([]byte, 20)
	for i := 0; i < 20; i++ {
		slice[i] = a[i]
	}
	return slice
}

func (a Address) String() string {
	return hex.EncodeToString(a.toSlice())
}

func AddressFromBytes(b []byte) Address {
	if len(b) != 20 {
		panic("NewAddressFromBytes: input byte slice must be exactly 20 bytes long")
	}

	var value Address
	for i := 0; i < 20; i++ {
		value[i] = b[i]
	}
	return value
}
