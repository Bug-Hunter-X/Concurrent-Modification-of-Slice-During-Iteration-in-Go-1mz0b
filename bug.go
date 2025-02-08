func MyFunc(s []string) {for i := range s {fmt.Println(s[i])}}

This function iterates over a slice of strings and prints each element.  It might seem correct, but it has a subtle issue that can lead to unexpected behavior, especially with concurrent access or modification of the slice.