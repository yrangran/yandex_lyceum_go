package main


func Sum(a, b int) int {
	return a + b
}

func Length(a int) string {
	switch {
	case a < 0:
		return "negative"
	case a == 0:
		return "zero"
	case a < 10:
		return "short"
	case a < 100:
		return "long"
	}
	return "very long"
}






// func main() {

// }





// func TimeAgo(pastTime time.Time) string {
// 	nowTime := time.Now()
// 	dif := nowTime.Sub(pastTime)
// 	second := int(dif.Seconds())
// 	minut := second/60
// 	hours := minut/60
// 	days := hours/24
// 	months := days/30
// 	years := months/365
// 	switch {
// 	case years > 0:
// 		return fmt.Sprintf("%d year%s ago",years, pluralize(years))
// 	case months > 0:
// 		return fmt.Sprintf("%d month%s ago",months, pluralize(months))
// 	case days > 0:
// 		return fmt.Sprintf("%d day%s ago",days, pluralize(days))
// 	case hours > 0:
// 		return fmt.Sprintf("%d hour%s ago",hours, pluralize(hours))
// 	case minut > 0:
// 		return fmt.Sprintf("%d minute%s ago",minut, pluralize(minut))
// 	default:
// 		return fmt.Sprintf("%d second%s ago",second, pluralize(second))
// 	}
	
// }

// func pluralize(n int) string{
// 	if n == 1{
// 		return ""
// 	}
// 	return "s"
// }

// func ParseStringToTime(dateString, format string) (time.Time, error) {
// 	tmp := time.Now().Format(format)
// 	t, err := time.Parse(tmp, dateString)
// 	if err != nil {
// 		return time.Time{}, fmt.Errorf("ошибка при парсинге времени: %w", err)
// 	}
// 	return t, nil
// }

// func FormatTimeToString(timestamp time.Time, format string) string {
// 	formattedTime := timestamp.Format(format)
// 	return formattedTime
// }

// func TimeDifference(start, end time.Time) time.Duration {
// 	diff := end.Sub(start)
// 	return diff
// }

// func NextWorkday(start time.Time) time.Time {
//     switch start.Weekday() {
//     case time.Sunday:
//         return start.AddDate(0, 0, 1)
//     case time.Monday:
//         return start.AddDate(0, 0, 2)
//     case time.Tuesday:
//         return start.AddDate(0, 0, 3)
//     case time.Wednesday:
//         return start.AddDate(0, 0, 4)
//     case time.Thursday:
//         return start.AddDate(0, 0, 5)
//     case time.Friday:
//         return start.AddDate(0, 0, 3)
//     case time.Saturday:
//         return start.AddDate(0, 0, 2)
//     }
//     return time.Time{}
// }

// func AreAnagrams(str1, str2 string) bool{
//     str1 = strings.ToLower(str1)
//     str2 = strings.ToLower(str2)
//     arr1 := []rune(str1)
//     sort.Slice(arr1, func(i,j int)bool{
//         return arr1[i]<arr1[j]
//     })
//     arr2 := []rune(str2)
//     sort.Slice(arr2, func(i,j int)bool{
//         return arr2[i]<arr2[j]
//     })
//     return reflect.DeepEqual(arr1, arr2)
// }

//var InvalidInputError = errors.New("invalid input, please provide two integers")

// func SumTwoIntegers(a, b string) (int, error){
//     //проверка первого числа
//     numA, errA := parseInt(a)
//     if errA != nil{
//         return 0, fmt.Errorf("first number: %w", errA)
//     }
//     numB, errB := parseInt(b)
//     if errB != nil{
//         return 0, fmt.Errorf("first number: %w", errA)
//     }
//     result := numA + numB
//     return result, nil
// }

// func parseInt(s string)(int,error){
//     s = strings.TrimSpace(s)
//     //определяем знак
//     var sign int
//     if s[0] == '-'{
//         sign = -1
//         s = s[1:]
//     }else if s[0] == '+'{
//         sign = 1
//         s = s[1:]
//     }else{
//         sign = 1
//     }
//     result := 0
//     for _, char := range s{
//         if !unicode.IsDigit(char){
//             return 0, InvalidInputError
//         }
//         digit := int(char - '0')
//         result = result*10 + digit
//     }
//     return sign * result, nil
// }

// func IntToBinary(num int) (string, error){
//     if num<0{
//         return "",InvalidInputError
//     }
//     var result string
//     result = fmt.Sprintf("%b",num)
//     return result,nil
// }

// func Factorial(n int) (int, error) {

//     if n<0{
//         return 0, InvalidInputError
//     }
//     if n == 1 || n==0{
//         return 1, nil
//     }
//     result,err := Factorial(n-1)
//     if err != nil{
//         return 0, fmt.Errorf("failed calculation: %w", err)
//     }
//     return result * n,nil
// }

// func GetCharacterAtPosition(str string, position int) (rune, error){
//     runes := []rune(str)
//     if len(runes)<position{
//         return 0, fmt.Errorf("position out of range")
//     }
//     return runes[position],nil
// }

// func DivideIntegers (a,b int) (float64,error){
//     if b == 0{
//         return 0, fmt.Errorf("division by zero is not allowed")
//     }
//     return float64(a)/float64(b), nil
// }
// func ConcatStringsAndInt(str1, str2 string, num int)string{
//     return fmt.Sprintf(str1+str2+"%d",num)
// }
// type Logger interface{
//     Log(message string)
// }

// type LogLevel string

// const(
//     Error LogLevel = "ERROR"
//     Info LogLevel = "INFO"
// )

// type Log struct{
//     LogLevel LogLevel
// }

// func (l Log) Log(message string){
//     switch l.LogLevel{
//     case Error:
//         fmt.Println(Error+":",message)
//     case Info:
//         fmt.Println(Info+":",message)
//     default:
//         fmt.Println("UNKNOWN LEVEL:", message)
//     }
// }

// type Animal interface{
//     MakeSound()
// }

// type Dog struct{
//     name string
// }

// type Cat struct{
//     name string
// }

// func (d Dog) MakeSound(){
//     fmt.Println("Гав!")
// }

// func (c Cat) MakeSound(){
//     fmt.Println("Мяу!")
// }

// type Shape interface{
//     Area() float64
// }

// type Rectangle struct{ //прямоугольник
//     Width float64
//     Length float64
// }

// type Circle struct{ //круг
//     radius float64
// }

// func CalculateArea(s Shape)float64{
//     return s.Area()
// }

// func (r Rectangle) Area() float64{
//     return r.Width*r.Length
// }

// func (c Circle) Area() float64{
//     return math.Pi * math.Pow(c.radius,2)
// }

// type Task struct{
//     summary string
//     description string
//     deadline time.Time
//     priority int
// }
// func (t Task) IsOverdue()bool{
//     return  time.Now().After(t.deadline)
// }
// func (t Task) IsTopPriority()bool{
//     return t.priority>=4
// }

// type Note struct{
//     title string
//     text string
// }

// type ToDoList struct{
//     name string
//     tasks []Task
//     notes []Note
// }

// func (t ToDoList) TasksCount()int{
//     return len(t.tasks)
// }
// func (t ToDoList) NotesCount()int{
//     return len(t.notes)
// }
// func (t ToDoList) CountTopPrioritiesTasks()int{
//     count :=0
//     for _, task:= range t.tasks{
//         if task.IsOverdue(){
//             count++
//         }
//     }
//     return count
// }

// func (t ToDoList) CountOverdueTasks()int{
//     count:=0
//     for _, task := range t.tasks{
//         if task.IsOverdue(){
//             count++
//         }
//     }
//     return count
// }

// func main() {

// }

//type Student struct{
//     name string
//     solvedProblems int
//     scoreForOneTask float64
//     passingScore float64
// }

// func (s Student) IsExcellentStudent()bool{
//     alltask := float64(s.solvedProblems) * s.scoreForOneTask
//     return alltask >= s.passingScore
// }

// type Employee struct{
// 	name string
// 	position string
// 	salary float64
// 	bonus float64
// }

// func (e Employee) CalculateTotalSalary(){
// 	fmt.Printf("Employee: %s\nPosition: %s\nTotal Salary: %.2f",e.name,e.position,(e.salary+e.bonus))
// }

// func DeleteLongKeys(m map[string]int) map[string]int{
//     res := make(map[string]int)
//     for key , value := range m{
//         if len(key) > 6 {
//             res[key] = value
//         }else{
//             continue
//         }
//     }
//     return res
// }

// func CountingSort(contacts []string) map[string]int{
//     res := make (map[string]int)
//     for _, num := range contacts {
//         res[num]++
//     }
//     return res
// }

// func SwapKeysAndValues (m map[string]string) map[string]string{
//     var res map[string]string
//     res = make(map[string]string)
//     for key, value := range m{
//         res[value] = key
//     }
//     return res
// }

// func SumOfValuesInMap(m map[int]int) int{
//     var sum int
//     for _, val:= range m{
//         sum+=val
//     }
//     return sum
// }

// func FindMaxKey(m map[int]int) int{
//     var max int
//     first := true
//     for key, _ := range m{
//         if first{
//             max = key
//             first = false
//             continue
//         }
//         if key > max{
//             max = key
//         }
//     }
//     return max
// }
