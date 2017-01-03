package helpers

type jsonErr struct {
	Code int    `json:"code"`
	Text string `json:"text"`
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
