Это отдельная библиотека
Пример для использования интерфейсов, структур и функции

Для применения скопируйте линк на репозиторий
и добавте на локальный проект по команде go get github.com/zhassulanbaimukhamedov/section1_4

Пример кода как можно использовать данную библиотеку

```
package main

import (
	"fmt"

	c "github.com/zhassulanbaimukhamedov/section1_4"
)

func main() {
	fmt.Println("hello")
	user1 := c.New()
	operation1(user1)
}

func operation1(p c.ICache) {
	p.Set("userId", 34)

	userId := p.Get("userId")
	fmt.Println(userId)

	p.Delete("userId")

	userId = p.Get("userId")
	fmt.Println(userId)
}
```

