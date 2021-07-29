import . "regexp"
func isMAC48Address(inputString string) bool {
    re := MustCompile(`^([A-F0-9][A-F0-9]-){5}[A-F0-9][A-F0-9]$`)
    return re.MatchString(inputString)
}
