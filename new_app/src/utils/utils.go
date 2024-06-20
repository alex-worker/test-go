package utils

func PanicAssert(err error) {
	if err == nil {
		return
	}
	panic(err)
}
