package main

import (
	"testing"
)
type Test_for_sum struct{
	a int
	b int
	out int
}
var tests_sum = []Test_for_sum{
	{1,2,3},
	{4,5,9},
}

func TestSum(t *testing.T){
	for i ,test:= range tests_sum{
		result := Sum(test.a,test.b)
		if result != test.out{
			t.Errorf("%d: Result(%d+%d) = %d; want %d",i,test.a,test.b,result,test.out)
		}
	}
}