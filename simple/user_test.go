package simple

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserCreate(t *testing.T) {

	actual := CreateNewUser("budi", "1234")

	t.Run("Get Name", func(t *testing.T) {
		assert.Equal(t, "budi", actual.GetName())
	})

	t.Run("Get Phone", func(t *testing.T) {
		assert.Equal(t, "1234", actual.GetPhone())
	})
}
