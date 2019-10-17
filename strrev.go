package piscine

func StrRev(s string) string {

	var shsh = ""
	for _, a := range s {
		shsh = string(a) + shsh
	}
	return shsh
}
