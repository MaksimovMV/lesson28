package student

import (
	"fmt"
	"strconv"
	"strings"
)

type Student struct {
	Name  string
	Age   int
	Grade int
}

func NewStudent(infoStr string) (st Student, err error) {
	data := strings.Fields(infoStr)

	if len(data) != 3 {
		err = fmt.Errorf("некорректный ввод данных")
		return st, err
	}

	st.Name = data[0]

	st.Age, err = strconv.Atoi(data[1])
	if err != nil {
		fmt.Println("Не удалось считать возраст:")
		return st, err
	}

	st.Grade, err = strconv.Atoi(data[2])
	if err != nil {
		fmt.Println("Не удалось считать курс:")
		return st, err
	}

	return st, err
}
