package main

import "testing"


func TestTimeConsuming(t *testing.T) {
    m := map[string]int{
        "1122": 3,
        "1111": 4,
        "1234": 0,
        "91212129": 9,
    }

    m2 := map[string]int {
        "1212": 6,
        "1221": 0,
        "123425": 4,
        "123123": 12,
        "12131415": 4,

    }

    for in, expect:=range(m) {
        if findSum(in) != expect {
            panic("Error!")
        }
    }
    for in, expect:=range(m2) {
        if findHalfwaySum(in) != expect {
            panic("Error!")
        }
    }
}