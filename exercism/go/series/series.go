package series

func All(n int, str string) (out []string) {
	for i := 0; i <= len(str)-n; i += 1 {
		out = append(out, str[i:i+n])
	}
	return out
}

func UnsafeFirst(n int, s string) string {
	lst := All(n, s)
	if len(lst) > 0 {
		return lst[0]
	}
	return ""
}

func First(n int, s string) (first string, ok bool) {
	lst := All(n, s)
	if len(lst) > 0 {
		return lst[0], true
	}
	return "", false

}
