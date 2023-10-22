package kreditPlusService

import (
	"net/http"
	models "test-echo/model"
	"test-echo/repository/kreditPlusRepository"
	"test-echo/utils"

	"github.com/labstack/echo/v4"
)

func GetLimit(c echo.Context) (models.ResponseView, error) {
	var res models.ResponseView
	NIK := c.QueryParam("nik")
	if NIK == "" {
		res.ResponseCode = http.StatusBadRequest
		res.ResponseDesc = "Failed"
		res.ResponseTime = utils.DateToStdNow()
		res.Message = "NIK Tidak boleh kosong!!"
		res.Result = nil
		return res, nil
	}
	Limits, err := kreditPlusRepository.CekLimit(NIK)
	if err != nil {
		res.ResponseCode = http.StatusBadRequest
		res.ResponseDesc = "Failed"
		res.ResponseTime = utils.DateToStdNow()
		res.Message = "Gagal Mendapatkan Limits"
		res.Result = nil
		return res, err
	}

	res.ResponseCode = http.StatusOK
	res.ResponseDesc = "Successfully"
	res.ResponseTime = utils.DateToStdNow()
	res.Message = "Berhasil Mendapatkan Limits"
	res.Result = Limits

	return res, nil
}

func GetCustomer(c echo.Context) (models.ResponseView, error) {
	var res models.ResponseView
	NIK := c.QueryParam("nik")
	if NIK == "" {
		res.ResponseCode = http.StatusBadRequest
		res.ResponseDesc = "Failed"
		res.ResponseTime = utils.DateToStdNow()
		res.Message = "NIK Tidak boleh kosong!!"
		res.Result = nil
		return res, nil
	}
	Customer, err := kreditPlusRepository.CekDataUser(NIK)
	if err != nil {
		res.ResponseCode = http.StatusBadRequest
		res.ResponseDesc = "Failed"
		res.ResponseTime = utils.DateToStdNow()
		res.Message = "Gagal Mendapatkan data Customer"
		res.Result = nil
		return res, err
	}
	// format customer
	var obj models.ConsumerFormatter
	Limit, err := kreditPlusRepository.CekLimit(NIK)
	if err != nil {
		res.ResponseCode = http.StatusBadRequest
		res.ResponseDesc = "Failed"
		res.ResponseTime = utils.DateToStdNow()
		res.Message = "Gagal Mendapatkan data Customer"
		res.Result = nil
		return res, err
	}
	obj.Limit.LoanLimit = Limit[0].LoanLimit
	obj.Limit.Tenor = Limit[0].Tenor
	obj.Limit.NIK = Limit[0].NIK
	obj.DOB = Customer.DOB
	obj.NIK = Customer.NIK
	obj.FullName = Customer.FullName
	obj.LegalName = Customer.LegalName
	obj.Salary = Customer.Salary
	res.ResponseCode = http.StatusOK
	res.ResponseDesc = "Successfully"
	res.ResponseTime = utils.DateToStdNow()
	res.Message = "Berhasil Mendapatkan data Customer"
	res.Result = obj

	return res, nil
}
