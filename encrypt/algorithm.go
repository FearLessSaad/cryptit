package encrypt

func Nimbus(str string) string {
	encryptedString := ""
	for _, c := range str {
		assciCode := int(c)
		character := string(assciCode + 3)
		encryptedString += character
	}
	return encryptedString
}
