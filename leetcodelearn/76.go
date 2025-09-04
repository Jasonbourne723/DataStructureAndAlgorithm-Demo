package leetcodelearn

func MinWindow(s string, t string) string {

	var sub string
	if len(s) < len(t) {
		return sub
	}
	left := 0

	count_t := [58]int{}
	count_s := [58]int{}
	count_s_e := [58]int{}

	for _, item := range t {
		count_t[item-'A']++
	}

	for i, _ := range t {
		val := count_t[s[i]-'A']
		if val == 0 {
			continue
		} else {
			if val == count_s[s[i]-'A'] {
				count_s_e[s[i]-'A']++
			} else {
				count_s[s[i]-'A']++
			}
		}
	}
	if count_s == count_t {
		return s[0:len(t)]
	}

	for i := len(t); i < len(s); i++ {
		val := count_t[s[i]-'A']
		if val == 0 {
			continue
		} else {
			if val == count_s[s[i]-'A'] {
				count_s_e[s[i]-'A']++
			} else {
				count_s[s[i]-'A']++
			}
		}

		if count_s == count_t {
			for {
				if count_t[s[left]-'A'] == 0 {
					left++
					continue
				}
				if count_s_e[s[left]-'A'] > 0 {
					count_s_e[s[left]-'A']--
					left++
					continue
				}
				break
			}
			length := i - left + 1
			if len(sub) == 0 || length < len(sub) {
				sub = s[left : i+1]
			}
		}
	}
	return sub
}
