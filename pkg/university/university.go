package university

import (
	"fmt"
	"lesson28/pkg/student"
)

type University map[string]*student.Student

func (un University) Put(st student.Student) {
	un[st.Name] = &st
}

func (un University) Get(studentName string) (*student.Student, error) {
	st, ok := un[studentName]
	if !ok {
		return nil, fmt.Errorf("нет такого студента")
	}
	return st, nil
}

func (un University) PrintAll() {
	for _, st := range un {
		stInfo, err := un.Get(st.Name)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(*stInfo)
	}
}
