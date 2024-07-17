package cron

import "fmt"

type AfterStartFunc func()

func Start() AfterStartFunc {
	// TODO: use package cron to init cron service
	fmt.Println("hello world")

	return func() {

	}
}
