package leetcodego
func GCD(a, b int) int {
    //credit to vishinoi29
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func gcdOfStrings(str1 string, str2 string) string {
    if str1 + str2 == str2 + str1{
        return str1[:GCD(len(str1), len(str2))]
    }else{
        return ""
    }
}