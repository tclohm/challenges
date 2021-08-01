func exist(file string, in []string) bool {
    for _, f := range in {
        if file == f {
            return true
        }
    }
    return false
}

func extend(file string, number int, in []string) bool {
    for _, f := range in {
        if file + "(" + strconv.Itoa(number) + ")" == f {
            return false
        }
    }
    return true
}

func fileNaming(names []string) []string {
    files := []string{}
    for _, file := range names {
        if exist(file, files) {
            i := 1
            for {
                if extend(file, i, files) {
                    newfile := file + "(" + strconv.Itoa(i) + ")"
                    files = append(files, newfile)
                    break
                }
                i++
            }
        } else {
            files = append(files, file)
        }
    }
    
    return files
}