package main

import (
	in "fmt"
	p "pkglvl/pkg_v"
)

func init() {
	in.Println("init() first: executes first")
}
func init() {
	in.Println("init() second: execures second after first")
}

func main() {
	in.Println("main(): executes after all init() functions")

	in.Print(p.AppName)
}

/*
1. В рамках одного файла Go выполняет несколько функций init() строго в том
   порядке, в котором они объявлены в коде (сверху вниз).
2. Если функции init() находятся в разных файлах одного пакета, Go вызывает их
   в лексикографическом порядке (по алфавиту) имён этих файлов.
3. Главная функция main() всегда запускается только после завершения всех init().
*/
