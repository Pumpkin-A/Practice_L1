package main

import (
	"context"
	"fmt"
)

// Конкретная реализация клиента, которого нужно адартировать
type RealClient struct{}

// Его метод, который нужно адаптировать
func (rc *RealClient) DoRealRequest(ctx context.Context, param1, param2 string, param3 int, isFlag bool) error {
	fmt.Println("Do Real request")
	return nil
}

// Интерфейс для пользователя - несовестим с интерфейсом RealClient
type UserInterface interface {
	DoRequest1(ctx context.Context, param string) error
}

// Интерфейс для админа - несовестим с интерфейсом RealClient
type AdminInterface interface {
	DoRequestA(ctx context.Context, param1, param2 string) error
	DoRequestB(ctx context.Context, param int) error
}

// ---> Адаптер, который реализует данный интерфейс для пользователя
type UserClientAdapter struct {
	client *RealClient
}

func (ca UserClientAdapter) DoRequest1(ctx context.Context, param string) error {
	return ca.client.DoRealRequest(ctx, param, "param_1", 12, false)
}

// <---

// ---> Адаптер, который реализует данный интерфейс для админа
type AdminClientAdapter struct {
	client *RealClient
}

func (ca UserClientAdapter) DoRequestA(ctx context.Context, param1, param2 string) error {
	return ca.client.DoRealRequest(ctx, param1, param2, 0, true)
}

func (ca UserClientAdapter) DoRequestB(ctx context.Context, param int) error {
	return ca.client.DoRealRequest(ctx, "", "", param, true)
}

// <---

func main() {
	// Создаем реального клиента
	realClient := RealClient{}

	// И далее передаем его в объекты с несовестимым интерфейсом, используя для этого реализации паттерна "адаптер"

	doUserRequests(UserClientAdapter{client: &realClient})

	doAdminRequests(UserClientAdapter{client: &realClient})
}

// Работаем с адаптированным клиентом через интерфейс, который реализует адаптер
func doUserRequests(client UserInterface) {
	if err := client.DoRequest1(context.Background(), "param_1"); err != nil {
		fmt.Println(err.Error())
		return
	}
}

func doAdminRequests(client AdminInterface) {
	if err := client.DoRequestA(context.Background(), "12", "10"); err != nil {
		fmt.Println(err.Error())
		return
	}

	if err := client.DoRequestB(context.Background(), 100); err != nil {
		fmt.Println(err.Error())
		return
	}
}
