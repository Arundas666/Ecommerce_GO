package usecase

import (
	"errors"
	"firstpro/config"
	"firstpro/helper"
	"firstpro/repository"
	"firstpro/utils/models"
)

func SendOTP(phone string) error {
	cfg, err := config.LoadConfig()
	if err!=nil{
		return err
	}


	ok := repository.FindUserByMobileNumber(phone)
	if !ok {
		return errors.New("the user does not exist")
	}
	helper.TwilioSetup(cfg.ACCOUNTSID,cfg.AUTHTOKEN)
	_,err=helper.TwilioSendOTP(phone,cfg.SERVICESSID)
	if err!=nil{
		return errors.New("error occured while generating otp")
	}
	

	return nil
}
func VerifyOTP(code models.VerifyData)(models.TokenUser,error){
	cfg, err := config.LoadConfig()
	if err!=nil{
		return models.TokenUser{},err
	}
	helper.TwilioSetup(cfg.ACCOUNTSID, cfg.AUTHTOKEN)
	err = helper.TwilioVerifyOTP(cfg.SERVICESSID, code.Code, code.User.PhoneNumber)
	if err != nil {
		return models.TokenUser{}, errors.New("error while verifying")
	}

}
