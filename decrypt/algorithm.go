package decrypt

func Nimbus(str string) string {
	decryptedString := ""
	for _, c := range str {
		assciCode := int(c)
		character := string(assciCode - 3)
		decryptedString += character
	}
	return decryptedString
}
