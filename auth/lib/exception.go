package lib

func PanicIfNeed(err interface{}) {
	if err != nil {
		panic(err)
	}
}
