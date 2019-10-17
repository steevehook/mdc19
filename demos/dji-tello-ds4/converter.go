package main

func int16To100(i int16) int16 {
	maxInt16 := int16(^uint16(0)>>1)
	unit := maxInt16 / 100
	return i / unit * -1
}
