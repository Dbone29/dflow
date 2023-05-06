package events

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

var ApiInistalized = &Event{}

type ApiInistalizedPayload struct {
	Gin  *gin.Engine
	Time time.Time
}

type ApiInistalizedHandler struct{}

func (h ApiInistalizedHandler) Handle(payload EventPayload) {
	data, ok := payload.(ApiInistalizedPayload)
	if !ok {
		fmt.Println("Failed to cast payload to ApiInistalizedPayload")
		return
	}

	fmt.Printf("User created: %s, Time: %s\n", data.Gin.BasePath(), data.Time)
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
