package advance

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateCart(t *testing.T) {

	body := Cart{
		ProductName: "BIJI KUDA",
		Qty:         10,
		Price:       1200,
		Discount:    0,
	}

	t.Run("Success but discount is 0", func(t *testing.T) {
		body.Discount = 1000000
		actual := CreateNewCart(body.ProductName, body.Qty, body.Price, body.Discount)
		assert.Equal(t, 0, actual.GetDiscount())
	})

	t.Run("Success created cart", func(t *testing.T) {
		body.Discount = 0
		actual := CreateNewCart(body.ProductName, body.Qty, body.Price, body.Discount)
		assert.Equal(t, 12000, actual.GetGrandTotal())
	})

	t.Run("Get Cart Detail", func(t *testing.T) {
		actual := CreateNewCart(body.ProductName, body.Qty, body.Price, body.Discount)
		str := fmt.Sprintf("%s : (%d x %d) - %d = %d", actual.GetProductName(), actual.GetQty(), actual.GetPrice(), actual.GetDiscount(), actual.GetGrandTotal())
		assert.Equal(t, str, actual.GetDetailCart())
	})

}

func TestUpdateCart(t *testing.T) {
	body := Cart{
		ProductName: "BIJI KUDA",
		Qty:         10,
		Price:       1200,
		Discount:    0,
	}
	actual := CreateNewCart(body.ProductName, body.Qty, body.Price, body.Discount)
	t.Run("Success created cart", func(t *testing.T) {

		assert.Equal(t, 12000, actual.GetGrandTotal())
	})

	t.Run("Update cart", func(t *testing.T) {
		actualUpdate := UpdateNewCart(actual, 2)
		assert.Equal(t, 2400, actualUpdate.GetGrandTotal())
	})

	t.Run("Update cart but qty 0", func(t *testing.T) {
		actualUpdate := UpdateNewCart(actual, 0)
		assert.Equal(t, 12000, actualUpdate.GetGrandTotal())
	})
}
