package main

import "strings"

func main() {

}

func maskPII(s string) string {
	if strings.Contains(s, "@") {
		return maskEmail(s)
	}
	return maskPhone(s)
}

func maskEmail(s string) string {
	ss := strings.Split(s, "@")
	name := ss[0]
	name = strings.ToLower(name)
	sb := strings.Builder{}
	sb.WriteByte(name[0])
	sb.WriteString("*****")
	sb.WriteByte(name[len(name)-1])
	return sb.String() + "@" + strings.ToLower(ss[1])
}

func maskPhone(s string) string {
	sb := strings.Builder{}
	for _, c := range s {
		if c >= '0' && c <= '9' {
			sb.WriteByte(uint8(c))
		}
	}
	phone := sb.String()
	sb = strings.Builder{}
	ccLen := len(phone) - 10
	if ccLen > 0 {
		sb.WriteString("+")
		for i := 0; i < ccLen; i++ {
			sb.WriteString("*")
		}
		sb.WriteString("-")
	}
	sb.WriteString("***-***-")
	sb.WriteString(phone[len(phone)-4:])
	return sb.String()
}
