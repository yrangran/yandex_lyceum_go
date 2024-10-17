package main 
import (
	"testing"
	"errors"
)


var ErrInvalidUTF8 = errors.New("invalid utf8")

func GetUTFLength(input []byte) (int, error) {
if !utf8.Valid(input) {
	return 0, ErrInvalidUTF8
}

return utf8.RuneCount(input), nil
}



func TestGetUTFLength(t *testing.T) {
    // Тест на валидный UTF-8
    validInput := []byte("Hello, 世界!")
    length, err := GetUTFLength(validInput)
    if err != nil {
        t.Errorf("Expected no error, but got: %v", err)
    }
    if length != 10 {
        t.Errorf("Expected length 11, but got: %d", length)
    }

    // Тест на невалидный UTF-8
    invalidInput := []byte{0xFF} // Байт, который не является началом UTF-8 последовательности
    _, err = GetUTFLength(invalidInput)
    if err == nil || err.Error() != ErrInvalidUTF8.Error() {
        t.Errorf("Expected error %q, but got: %v", ErrInvalidUTF8, err)
    }

    // Тест на пустой байтовый слайс
    emptyInput := []byte{}
    length, err = GetUTFLength(emptyInput)
    if err != nil {
        t.Errorf("Expected no error, but got: %v", err)
    }
    if length != 0 {
        t.Errorf("Expected length 0, but got: %d", length)
    }
}




















// type Test_length struct {
// 	in  int
// 	out string
// }

// var tests = []Test_length{
// 	{-1, "negative"},
// 	{5, "short"},
// 	{0, "zero"},
// 	{24, "long"},
// 	{888, "very long"},
// }

// func TestLength(t *testing.T) {
// 	for i, test := range tests {
// 		size := Length(test.in)
// 		if size != test.out {
// 			t.Errorf("#%d: Size(%d)=%s; want %s", i, test.in, size, test.out)
// 		}
// 	}
// }

// type Test struct{
// 	a,b int
// 	out int
// }


// var tests = []Test{
// 	{1,2,2},
// 	{5,4,20},
// 	{3,2,6},
// 	{2,0,0},
// }

// func TestMultiply(t *testing.T){
// 	for i, test := range tests{
// 		result := Multiply(test.a,test.b)
// 		if result != test.out{
// 			t.Errorf("%d: Result(%d*%d) = %d; want %d",i,test.a,test.b,result,test.out)
// 		}
// 	}
// }

// type Test struct{
// 	name string
// 	s string
// 	out string
// }

// var tests =[]Test{
// 	{"first","apple","ppl"},
// 	{"second with y","appyyleyy","ppyylyy"},
// 	{"only consonants","bbbtttrrr","bbbtttrrr"},
// 	{"for fun","big muskules","bg mskls"},
// }

// func TestDeleteVowels(t *testing.T){
// 	cases := []struct{
// 		//имя теста
// 		name string
// 		//входные данные
// 		s string
// 		//то что должно получится
// 		out string
// 	}{
// 		//тестовые данные 1
// 		{
// 			name: "first",
// 			s: "apple",
// 			out: "ppl",
// 		},
// 		//тестовые данные 2
// 		{
// 			name: "second with y",
// 			s: "appyyleyy",
// 			out: "ppyylyy",
// 		},

// 	}
// 	for _, tc:= range cases{
// 		tc := tc
// 		//запуск отдельного теста
// 		t.Run(tc.name, func(t *testing.T){
// 			got :=  DeleteVowels(tc.s)
// 			if got != tc.out{
// 				t.Errorf("DeleteVowels(%s) = %s; want %s",tc.s,got,tc.out)
// 			}
// 		})
// 	}
// }