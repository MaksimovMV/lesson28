package main

import (
	"bufio"
	"fmt"
	"lesson28/pkg/student"
	"lesson28/pkg/university"
	"os"
)

func main() {
	un := make(university.University)
	scStudentInfo := bufio.NewScanner(os.Stdin)
	fmt.Println("Введите данные студентов в формате \"Имя Возраст Курс\":")
	for scStudentInfo.Scan() {
		infoStr := scStudentInfo.Text()
		st, err := student.NewStudent(infoStr)
		if err != nil {
			fmt.Println(err)
			continue
		}
		un.Put(st)
	}
	fmt.Println("Студенты из хранилища:")
	un.PrintAll()
}
