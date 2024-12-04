// Дана структура Human (с произвольным набором полей и методов).
// Реализовать встраивание методов в структуре Action от родительской
// структуры Human (аналог наследования).

package main

import "fmt"

type Human struct {
	Name string
	Age  int
}

func (h *Human) SayName() string {
	return fmt.Sprintf("My name is %s\n", h.Name)
}

type Action struct {
	Human
}

func (a *Action) IntroduceHuman() string {
	return fmt.Sprintf("Hello! %s\n", a.SayName())
}

type Namer interface {
	SayName() string
}

func main() {
	human := &Human{Name: "Ann", Age: 29}
	action := &Action{Human: *human}

	// использование метода Human происходит снаружи
	fmt.Println(action.SayName())

	// использование метода Human происходит внутри реализации
	fmt.Println(action.IntroduceHuman())

	//при использовании встраивания(композиции) этот код компилируется, так как набор методов родительского класса полностью
	//доступен извне напрямую из дочернего, соответственно интерфейс реализуется. Если пользоваться не встраиванием, а просто положить структуру
	//Human в структуру Action, то action не будет реализовывать данный интерфейс, необходимо будет явно обращаться к
	//структуре Human и ее методам, код не скомпилируется.
	var namer Namer
	namer = action
	fmt.Println(namer.SayName())

}
