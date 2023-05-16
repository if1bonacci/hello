package middleware

type LogInterface interface {
	Init(title string)
	AddMessage(message string)
	Post()
}
