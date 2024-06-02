package main

//type Mapper struct{ slice []int }
//func (mapper *Mapper) test() { mapper.slice = append(mapper.slice, 2, 3, 4) } // OUTPUT: 3
//func (mapper Mapper)test() { mapper.slice = append(mapper.slice, 2, 3, 4) } // OUTPUT: 0

func main() {
	// ----------------------------------------------------------
	//size := 1000000
	//values1 := make([]int32, size)
	//startTime1 := time.Now().UnixMilli()
	//for index := 0; index < size; index++ {
	//	values1[index] = int32(index)
	//}
	//endTime1 := time.Now().UnixMilli()
	//time1 := endTime1 - startTime1
	//var values2 []int32
	//startTime2 := time.Now().UnixMilli()
	//for index := 0; index < size; index++ {
	//	values2 = append(values2, int32(index))
	//}
	//endTime2 := time.Now().UnixMilli()
	//time2 := endTime2 - startTime2
	//fmt.Println(fmt.Sprintf("time1: %d, time2: %d", time1, time2))
	// OUTPUT: time1 < time2

	// ----------------------------------------------------------
	//defer func() { fmt.Println("function 1") }()
	//defer func() { fmt.Println("function 2") }()
	//
	//defer func() {
	//	fmt.Println("function 3")
	//	panic("function 4")
	//}()
	//defer func() { fmt.Println("function 5") }()
	//defer func() {
	//	if r := recover(); r != nil {
	//		fmt.Println("function 6")
	//	}
	//}()
	// OUTPUT: function 5 function 3 function 2 function 1 panic: function 4

	// ----------------------------------------------------------
	//obj := Mapper{}
	//obj.test()

	//fmt.Println(fmt.Sprintf("%v", len(obj.slice)))
	//OUTPUT: 0

	// ----------------------------------------------------------
	//function()
	//fmt.Println("finished")
	// OUTPUT: some work-1 panic: panic

	// ----------------------------------------------------------
	//defer func() {
	//	if err := recover(); err != nil {
	//		fmt.Println("not finished")
	//	}
	//}()
	//mapper := make(map[int]struct{})
	//mapper[0] = struct{}{}
	//mapper[1] = struct{}{}
	//mapper[2] = struct{}{}
	//mapper[3] = struct{}{}
	//mapper[4] = struct{}{}
	//mapper[5] = struct{}{}
	//
	//for key := range mapper {
	//	mapper[key] = struct{}{}
	//	delete(mapper, key)
	//}
	//fmt.Println("finished")
	// OUTPUT: finished

	// ----------------------------------------------------------
	//values := make([]interface{}, 10)
	//
	//for index := 0; index < len(values); index++ {
	//	values[index] += values[index]
	//}
	//fmt.Println(values)
	// OUTPUT: ERROR (compile-time)

	// ----------------------------------------------------------
	//arr1 := [2]int{2, 3}
	//arr2 := [...]int{2, 3}
	//
	//fmt.Println(arr1 == arr2)
	// OUTPUT: true

	// ----------------------------------------------------------
	//defer func() { fmt.Println("function 1") }()
	//defer func() { fmt.Println("function 2") }()
	//
	//defer func() {
	//	if r := recover(); r != nil {
	//		fmt.Println("function 6")
	//	}
	//}()
	//defer func() {
	//	fmt.Println("function 3")
	//	panic("function 4")
	//}()
	//defer func() { fmt.Println("function 5") }()
	// OUTPUT: function 5 function 3 function 6 function 2 function 1

	// ----------------------------------------------------------
	//mapper := make(map[int]struct{})
	//prepare(mapper)
	//fmt.Println(fmt.Sprintf("values are: %v", mapper))
	// OUTPUT: values are: map[6:{}]

	// ----------------------------------------------------------
	//var sampleMap = map[string]interface{}{"key1": 1, "key2": []map[string]interface{}}
	//fmt.Println(sampleMap["key3"])
	// OUTPUT: ERROR

	// ----------------------------------------------------------
	//channel1 := make(chan bool, 2)
	//fmt.Println("task begin")
	//channel1 <- true
	//fmt.Println("task complete")

	// OUTPUT (if buffered channel): task begin task complete
	// OUTPUT (if buffered channel): task begin fatal error: deadlock

	// ----------------------------------------------------------
	//var x int
	//arr := [3]int{3, 2, 5}
	//x, arr = arr[0] , []arr[1:]
	//fmt.Println(x, arr)

	// OUTPUT: ERROR

	// ----------------------------------------------------------
	//var err = function()
	//if err != nil {
	//	fmt.Println("not finished")
	//} else {
	//	fmt.Println("finished")
	//}

	// OUTPUT: ERROR

	// ----------------------------------------------------------
	//mapper := make(map[string][]string)
	//
	//fruits := []string{"orange", "banana", "apple", "kiwi", "watermelon"}
	//mapper["fruit"] = fruits
	//
	//for index := range fruits {
	//	if index%2 == 0 {
	//		fruits = fruits[:index]
	//	}
	//}
	//fmt.Println(fmt.Sprintf("%v", fruits))
	//fmt.Println(fmt.Sprintf("%v", mapper["fruit"]))

	// OUTPUT: fruits: [orange banana apple kiwi] mapper["fruit"]=[orange banana apple kiwi watermelon]

	// ----------------------------------------------------------
	//var sampleMap = map[string]interface{}{"key1": 1, "key2": []map[string]interface{}{}}
	//sampleMap["key2"] = append(sampleMap["key2"].([]map[string]interface{}), map[string]interface{}{"key3": 1})
	//fmt.Println(sampleMap["key2"])

	// OUTPUT: [map[key3:1]]
	// ----------------------------------------------------------
}

//func function() (err error) {
//	var err error // if this line is not there: OUTPUT => not finished
//	defer func() {
//		if r := recover(); r != nil {
//			err = errors.New("some error occured")
//		}
//	}()
//	panic("panic")
//	return err
//}
// ----------------------------------------------------------
//func function() {
//	if r := recover(); r != nil {
//		fmt.Println("panic recovered")
//	}
//	fmt.Println("some work-1")
//	panic("panic")
//	fmt.Println("some work-2")
//}
// ----------------------------------------------------------
//func prepare(mapper map[int]struct{}) {
//	for i := 0; i <= 5; i++ {
//		defer func() { mapper[i] = struct{}{} }()
//	}
//}
