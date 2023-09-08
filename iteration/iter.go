package iteration

import "strings"

func Repeat(s string, times int) string {
    //var repeated string
    //for i := 0; i < times; i++ {
    //    repeated += char
    //}
    //return repeates
    return strings.Repeat(s, times)
}

