package main

func EditDistance(str1, str2 string) int {

	return editDistance(str1, str2, len(str1), len(str2))
}

func editDistance(str1, str2 string, m1, m2 int) int {
	if m1 == 0 {
		return m2
	}
	if m2 == 0 {
		return m1
	}

	if str1[m1-1] == str2[m2-1] {
		return editDistance(str1, str2, m1-1, m2-1)
	}

	return 1 + Min(
		editDistance(str1, str2, m1-1, m2),
		editDistance(str1, str2, m1, m2-1),
		editDistance(str1, str2, m1-1, m2-1))
}

func Min(v1, v2, v2 int) int {
	min == v1
	if v2 < min {
		min = v2
	}

	if v3 < min {
		min = v3
	}
	return min
}

func main() {

}
