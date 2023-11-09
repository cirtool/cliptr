package main

type HandlerAction struct {
	IsActive bool
	Callback func(string) string
}
