package main

import "fmt"

func addStrings(num1 string, num2 string) string {
	n := len(num1)
	if n < len(num2) {
		n = len(num2)
	}
	n1 := len(num1)
	n2 := len(num2)

	sb := make([]byte, n+1)

	i, j := n1-1, n2-1
	var s1, s2, index, tmp, sum int
	for ; i >= 0 && j >= 0; {
		s1 = int(num1[i] - '0')
		s2 = int(num2[j] - '0')
		sum = s1 + s2 + tmp
		sb[index] = byte(sum%10 + '0')
		tmp = sum / 10
		index += 1
		i -= 1
		j -= 1
	}

	if i >= 0 {
		for i >= 0 || tmp > 0 {
			s1 = 0
			if i >= 0 {
				s1 = int(num1[i] - '0')
			}
			sum = s1 + tmp
			sb[index] = byte(sum%10 + '0')
			tmp = sum / 10
			index += 1
			i -= 1
		}
	}

	if j >= 0  {
		for j >= 0 || tmp > 0 {
			s2 = 0
			if j >= 0 {
				s2 = int(num2[j]-'0')
			}
			sum = s2 + tmp
			sb[index] = byte(sum%10 + '0')
			tmp = sum / 10
			index += 1
			j -= 1
		}
	}

	if tmp != 0 {
		sb[index] = byte(tmp + '0')
		index += 1
	}

	sb = sb[:index]
	if len(sb) == 1 {
		return string(sb)
	}
	l, r := 0, len(sb)-1
	for l < r {
		sb[l], sb[r] = sb[r], sb[l]
		l += 1
		r -= 1
	}
	return string(sb)
}

func main() {
	fmt.Println(addStrings("0", "0"))
	fmt.Println(addStrings("923", "83"))
	fmt.Println(addStrings("123", "87"))
	fmt.Println(addStrings("1", "9"))
	fmt.Println(addStrings("9", "99"))
	fmt.Println(addStrings("6994", "36"))
	fmt.Println(addStrings("71", "168899993"))
}
