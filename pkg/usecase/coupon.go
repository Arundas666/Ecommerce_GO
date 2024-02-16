package usecase

import (
	"errors"
	"firstpro/pkg/repository"
	"firstpro/pkg/utils/models"
)

func AddCoupon(coupon models.AddCoupon) (string, error) {

	// if coupon already exist and if it is expired revalidate it. else give back an error message saying the coupon already exist

	couponExist, err := repository.CouponExist(coupon.Coupon)
	if err != nil {
		return "", err
	}

	if couponExist {

		alreadyValid, err := repository.CouponRevalidateIfExpired(coupon.Coupon)

		if err != nil {
			return "", err
		}

		if alreadyValid {
			return "The coupon which is valid already exists", nil
		}

		return "Made the coupon valid", nil

	}

	err = repository.AddCoupon(coupon)
	if err != nil {
		return "", err
	}

	return "successfully added the coupon", nil
}

func AddProductOffer(productOffer models.ProductOfferReceiver) error {

	return repository.AddProductOffer(productOffer)

}
func AddCategoryOffer(categoryOffer models.CategoryOfferReceiver) error {

	return repository.AddCategoryOffer(categoryOffer)

}
func GetCoupon() ([]models.Coupon, error) {

	return repository.GetCoupon()

}

func ExpireCoupon(couponID int) error {

	// check whether coupon exist
	couponExist, err := repository.ExistCoupon(couponID)
	if err != nil {
		return err
	}

	// if it exists expire it, if already expired send back relevant message
	if couponExist {
		err = repository.CouponAlreadyExpired(couponID)
		if err != nil {
			return err
		}

		return nil
	}

	return errors.New("coupon does not exist")

}
