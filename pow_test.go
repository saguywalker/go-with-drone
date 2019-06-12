package main

import (
    "testing"
)

type testpair struct{
    values []int
    result int
}

var tests = []testpair{
    {[]int{2,4}, 16},
    {[]int{3,5}, 243},
    {[]int{2,12}, 4096},
    {[]int{10,0}, 1},
    {[]int{25,2}, 625},
    {[]int{10,1}, 10},
}

func TestPow(t *testing.T){
    for _, pair := range tests{
        v := Pow(pair.values[0], pair.values[1])
        if v != pair.result{
            t.Error("For", pair.values,"expected",pair.result,"got",v)
        }
    }
}
