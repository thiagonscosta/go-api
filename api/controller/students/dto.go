import (
	"github.com/google/uuid"
)

package students

type StudentInput struct {
	ID string
	UUID uuid.UUID
	Name string `json:"name"`
	Age int `json:"age"`
}