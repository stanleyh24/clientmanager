package service

import (
	"context"
	"fmt"

	"github.com/stanley24/clientmanager/internal/routers/repository"
)

func GetAllRouter() {
	ctx := context.Background()

	routers, err := repository.GetAll(ctx)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Print(routers)

}
