package helper

func PanicIfErr(err error, text string) {
	if err != nil {
		panic(text)
	}
}
