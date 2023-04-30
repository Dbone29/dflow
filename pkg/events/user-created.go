package events

import (
	"fmt"
)

type UserCreatedHandler struct{}

func (h UserCreatedHandler) Handle(payload EventPayload) {
	data, ok := payload.(UserCreatedPayload)
	if !ok {
		fmt.Println("Failed to cast payload to UserCreatedPayload")
		return
	}
	fmt.Printf("User created: %s, Time: %s\n", data.Email, data.Time)
}

/*func main() {
	handler := UserCreatedHandler{}

	UserCreated.Register(handler)

	payload := UserCreatedPayload{
		Email: "user@example.com",
		Time:  time.Now(),
	}

	UserCreated.Trigger(payload)
}*/
