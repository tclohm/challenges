import ("strings")

func exist(char string, in []string) (int, bool) {
    for i, r := range in {
        if string(r) == char {
            return i, true
        }
    }
    return -1, false
}

func createAnagram(s string, t string) int {
    // min # of replacement ops needed to get an anagram of string t from the string s
    sarr := strings.Split(s, "")
    tarr := strings.Split(t, "")
    
    for _, r := range sarr {
        c := string(r)
        nth, valid := exist(c, tarr)
        if valid {
            first := tarr[:nth]
            second := tarr[nth+1:]
            tarr = append(first, second...)
        }
    }

    return len(tarr)
}