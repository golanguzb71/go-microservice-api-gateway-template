package helpers

import "reflect"

func PageOfSlice(s interface{}, page, pageSize int) reflect.Value {
	v := reflect.ValueOf(s)
	l := v.Len()
	i := (page - 1) * pageSize
	j := i + pageSize
	if i >= l {
		return reflect.New(reflect.TypeOf(s)).Elem()
	}
	if j > l {
		j = l
	}
	return v.Slice(i, j)
}

func GenerateColumnLetter(n int) string {
	if n <= 0 {
		return ""
	}
	const alphabetSize = 26

	firstLetter := string(rune('A' + (n-1)%alphabetSize))

	remainingPart := GenerateColumnLetter((n - 1) / alphabetSize)

	return remainingPart + firstLetter
}
