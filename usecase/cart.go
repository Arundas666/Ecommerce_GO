package usecase

import (
	"errors"
	"firstpro/repository"
	"firstpro/utils/models"
)

func AddToCart(product_id int, user_id int) (models.CartResponse, error) {
	ok, _, err := repository.CheckProduct(product_id)
	//here the second return is category and we will use this later when we need to add the offer details later
	if err != nil {

		return models.CartResponse{}, err
	}

	if !ok {
		return models.CartResponse{}, errors.New("product Does not exist")
	}

	QuantityOfProductInCart, err := repository.QuantityOfProductInCart(user_id, product_id)
	if err != nil {

		return models.CartResponse{}, err
	}
	quantityOfProduct, err := repository.GetQuantityFromProductID(product_id)
	if err != nil {

		return models.CartResponse{}, err
	}
	if quantityOfProduct == 0 {
		return models.CartResponse{}, errors.New("out of stock")
	}
	if quantityOfProduct == QuantityOfProductInCart {
		return models.CartResponse{}, errors.New("stock limit exceeded")
	}
	productPrice, err := repository.GetPriceOfProductFromID(product_id)
	if err != nil {

		return models.CartResponse{}, err
	}
	if QuantityOfProductInCart == 0 {
		err := repository.AddItemIntoCart(user_id, product_id, 1, productPrice)
		if err != nil {

			return models.CartResponse{}, err
		}

	} else {
		currentTotal, err := repository.TotalPriceForProductInCart(user_id, product_id)
		if err != nil {
			return models.CartResponse{}, err
		}
		err = repository.UpdateCart(QuantityOfProductInCart+1, currentTotal+productPrice, user_id, product_id)
		if err != nil {
			return models.CartResponse{}, err
		}

	}
	cartDetails, err := repository.DisplayCart(user_id)
	if err != nil {
		return models.CartResponse{}, err
	}
	cartTotal, err := repository.GetTotalPrice(user_id)
	if err != nil {

		return models.CartResponse{}, err
	}
	return models.CartResponse{
		UserName:   cartTotal.UserName,
		TotalPrice: cartTotal.TotalPrice,
		Cart:       cartDetails,
	}, nil

}

func RemoveFromCart(product_id int, user_id int) (models.CartResponse, error) {
	ok, err := repository.ProductExist(user_id, product_id)
	if err != nil {
		return models.CartResponse{}, err
	}
	if !ok {
		return models.CartResponse{}, errors.New("product does'nt exist in the cart")
	}
	var cartDetails struct {
		Quantity   int
		TotalPrice float64
	}

	cartDetails, err = repository.GetQuantityAndProductDetails(user_id, product_id, cartDetails)
	if err != nil {
		return models.CartResponse{}, err
	}

	cartDetails.Quantity = cartDetails.Quantity - 1

	//remove the product if quantity after deleting is 0
	if cartDetails.Quantity == 0 {
		if err := repository.RemoveProductFromCart(user_id, product_id); err != nil {
			return models.CartResponse{}, err
		}

	}
	if cartDetails.Quantity != 0 {

		product_price, err := repository.GetPriceOfProductFromID(product_id)
		if err != nil {
			return models.CartResponse{}, err
		}
		cartDetails.TotalPrice = cartDetails.TotalPrice - product_price
		err = repository.UpdateCartDetails(cartDetails, user_id, product_id)
		if err != nil {
			return models.CartResponse{}, err
		}
	}
	updatedCart, err := repository.CartAfterRemovalOfProduct(user_id)
	if err != nil {
		return models.CartResponse{}, err
	}
	cartTotal, err := repository.GetTotalPrice(user_id)
	if err != nil {
		return models.CartResponse{}, err
	}
	return models.CartResponse{
		UserName:   cartTotal.UserName,
		TotalPrice: cartTotal.TotalPrice,
		Cart:       updatedCart,
	}, nil

}
func DisplayCart(user_id int) (models.CartResponse, error) {

	cart, err := repository.DisplayCart(user_id)
	if err != nil {
		return models.CartResponse{}, err
	}
	cartTotal, err := repository.GetTotalPrice(user_id)
	if err != nil {
		return models.CartResponse{}, err
	}
	return models.CartResponse{
		UserName:   cartTotal.UserName,
		TotalPrice: cartTotal.TotalPrice,
		Cart:       cart,
	}, nil

}
func EmptyCart(userID int) (models.CartResponse, error) {
	ok, err := repository.CartExist(userID)
	if err != nil {
		return models.CartResponse{}, err
	}
	if !ok {
		return models.CartResponse{}, errors.New("cart already empty")
	}
	if err := repository.EmptyCart(userID); err != nil {
		return models.CartResponse{}, err
	}

	cartTotal, err := repository.GetTotalPrice(userID)

	if err != nil {
		return models.CartResponse{}, err
	}

	cartResponse := models.CartResponse{
		UserName:   cartTotal.UserName,
		TotalPrice: cartTotal.TotalPrice,
		Cart:       []models.Cart{},
	}

	return cartResponse, nil

}

func ApplyCoupon(coupon string, userID int) error {

	cartExist, err := repository.CartExist(userID)
	if err != nil {
		return err
	}

	if !cartExist {
		return errors.New("cart empty, can't apply coupon")
	}

	couponExist, err := repository.CouponExist(coupon)
	if err != nil {
		return err
	}

	if !couponExist {
		return errors.New("coupon does not exist")
	}

	couponValidity, err := repository.CouponValidity(coupon)
	if err != nil {
		return err
	}

	if !couponValidity {
		return errors.New("coupon expired")
	}

	minDiscountPrice, err := repository.GetCouponMinimumAmount(coupon)
	if err != nil {
		return err
	}

	totalPriceFromCarts, err := repository.GetTotalPriceFromCart(userID)
	if err != nil {
		return err
	}

	// if the total Price is less than minDiscount price don't allow coupon to be added
	if totalPriceFromCarts < minDiscountPrice {
		return errors.New("coupon cannot be added as the total amount is less than minimum amount for coupon")
	}

	userAlreadyUsed, err := repository.DidUserAlreadyUsedThisCoupon(coupon, userID)
	if err != nil {
		return err
	}

	if userAlreadyUsed {
		return errors.New("user already used this coupon")
	}

	couponStatus, err := repository.UpdateUsedCoupon(coupon, userID)
	if err != nil {
		return err
	}

	if couponStatus {
		return nil
	}
	return errors.New("could not add the coupon")

}
