package piscine


func HasCode(dec string) string{
	result :=""
	size := len(dec)

	for i := 0; i < len(dec); i++ {
		h :=(int(dec[i]) + size) %127

		if h < 33 {
			h += 33
		}
		result +=string(rune(h))
	}
	return result
}
//
//using for range 
func HasCode(dec string) string{
	result :=""
	size := len(dec)

	for _, r := dec {
		h :=(int(r) + size) %127

		if h < 33 {
			h += 33
		}
		result +=string(rune(h))
	}
	return result
}
