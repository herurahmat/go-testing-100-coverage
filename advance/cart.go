package advance

import "fmt"

type Cart struct {
	ProductName string
	Qty         int
	Price       int
	Discount    int
}

func (s Cart) GetProductName() string {
	return s.ProductName
}

func (s Cart) GetQty() int {
	return s.Qty
}

func (s Cart) GetPrice() int {
	return s.Price
}

func (s Cart) GetDiscount() int {
	subTotal := (s.Price * s.Qty)
	if s.Discount > subTotal {
		return 0
	}
	return s.Discount
}

func (s Cart) GetGrandTotal() int {
	subTotal := (s.Price * s.Qty)
	grandTotal := subTotal - s.Discount
	return grandTotal
}

func CreateNewCart(productName string, qty int, price int, discount int) *Cart {
	return &Cart{
		ProductName: productName,
		Qty:         qty,
		Price:       price,
		Discount:    discount,
	}
}

func UpdateNewCart(oldCart *Cart, qty int) *Cart {
	if qty <= 0 {
		return oldCart
	}

	return &Cart{
		ProductName: oldCart.ProductName,
		Qty:         qty,
		Price:       oldCart.Price,
		Discount:    oldCart.Discount,
	}
}

func (s Cart) GetDetailCart() string {

	str := fmt.Sprintf("%s : (%d x %d) - %d = %d", s.GetProductName(), s.GetQty(), s.GetPrice(), s.GetDiscount(), s.GetGrandTotal())
	return str
}
