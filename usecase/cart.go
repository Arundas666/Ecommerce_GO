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
