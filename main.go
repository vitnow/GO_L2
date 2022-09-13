package main

import "log"

type Config struct {
	Address string
	Port    int
}

func main() {
	log.Panicln(getConfig())
}

//func div(a, b float64) (ret float64, err error) {
//	defer func() {
//		if e := recover(); e != nil {
//			fmt.Println("My error:", err)
//			err = e.(error)
//		}
//	}()
//	defer func() {
//		fmt.Println("div defer", ret, err)
//		if err := recover(); err != nil {
//			fmt.Println("err in dev:", err)
//			panic(err)
//		}
//	}()
//	if b == 0 {
//		panic("divide by zero")
//		return 0, errors.New("divide by zero")
//	}
//	return a / b, nil
//
//}
//
//func main() {
//	defer func() {
//		if err := recover(); err != nil {
//			log.Println("We've got an error:", err)
//		}
//	}()
//
//	// 1. Задание
//	res, err := div(1, 0)
//	fmt.Println("div:", res, err)
//
//	// 2. Задание
//
//	fmt.Println("next exercise 2")
//	time.Sleep(3 * time.Second)
//
//	numbers := []string{
//		"one",
//		"two",
//		"three",
//	}
//	fmt.Println("My favorite number:", numbers[len(numbers)])
//
//	fmt.Println("next exercise 3")
//	time.Sleep(3 * time.Second)
//
//	// 3. Задание
//	for i := 0; i < 1000000; i++ {
//		newFile, err := os.Create(fmt.Sprintf("%v.txt", i))
//		if err != nil {
//			log.Println("Not enough resources:", err)
//			defer recover()
//		}
//		defer newFile.Close()
//	}
//}
