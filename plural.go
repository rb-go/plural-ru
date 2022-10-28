// plural package to make pluralization of russian words
package plural

type Number interface {
	int | int32 | int64 | uint | uint32 | uint64
}

// Noun function to make pluralization for noun
func Noun[n Number](num n, oneWord, twoWord, ThreeWord string) string {
	switch getNounPluralForm(num) {
	case 0:
		return oneWord
	case 1:
		return twoWord
	case 2:
		return ThreeWord
	}
	return ""
}

// Verb function to make pluralization for verb
func Verb[n Number](num n, oneWord, twoWord, ThreeWord string) string {
	switch getVerbPluralForm(num) {
	case 0:
		return oneWord
	case 1:
		return twoWord
	case 2:
		return ThreeWord
	}
	return ""
}

func getNounPluralForm[n Number](a n) int {
	if a%10 == 1 && a%100 != 11 {
		return 0
	} else if a%10 >= 2 && a%10 <= 4 && (a%100 < 10 || a%100 >= 20) {
		return 1
	} else {
		return 2
	}
}

func getVerbPluralForm[n Number](a n) int {
	if a > 1000000 {
		return 2
	}
	if a > 1000 && a < 1000000 && a%1000 == 0 {
		a /= 1000
	}
	if a%10 == 1 && a%100 != 11 || a%10000 == 1000 {
		return 0
	} else if a%10 >= 2 && a%10 <= 4 && (a%100 < 10 || a%100 >= 20) {
		return 1
	}
	return 2
}
