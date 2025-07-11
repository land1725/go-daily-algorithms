package day03_is_valid

func isValid(s string) bool {
	stackByte := make([]byte, 0, len(s))
	pair := map[byte]byte{
		')': '(',
		'}': '{',
		']': '[',
	}
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '(', '[', '{':
			stackByte = append(stackByte, s[i])
		case ')', ']', '}':
			if len(stackByte) == 0 || stackByte[len(stackByte)-1] != pair[s[i]] {
				return false
			}
			stackByte = stackByte[:len(stackByte)-1]
		}
	}
	return len(stackByte) == 0
}
