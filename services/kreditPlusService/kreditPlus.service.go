package kreditPlusService

import (
	"net/http"
	"strconv"
	models "test-echo/model"
	"test-echo/repository/kreditPlusRepository"
	"test-echo/utils"

	"github.com/google/uuid"

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

func CreateTransaction(c echo.Context) (models.ResponseView, error) {
	var obj models.Transaction
	var res models.ResponseView
	obj.ContractNo = uuid.New()
	obj.AdminFee = 10000
	otr := c.FormValue("otr")
	f, err := strconv.ParseFloat(otr, 64)
	if err != nil {
		return res, err
	}
	obj.OTR = f
	obj.AssetType = c.FormValue("asset_type")
	obj.AssetName = c.FormValue("asset_name")
	num, err := strconv.Atoi(c.FormValue("tenor"))
	if err != nil {
		return res, err
	}
	obj.Installments = num
	obj.InterestRate = 11.5
	obj.NIK = c.FormValue("nik")
	// check limit berdasarkan NIK dan tenor angsuran
	limits, err := kreditPlusRepository.CekLimit(obj.NIK)
	if err != nil {
		return res, err
	}
	// cek tenor
	limitAcc := 0.0
	for _, v := range limits {
		if v.Tenor == num {
			limitAcc = v.LoanLimit
			break
		}
	}
	if limitAcc == 0.0 {
		res.ResponseCode = http.StatusBadRequest
		res.ResponseDesc = "Failed"
		res.ResponseTime = utils.DateToStdNow()
		res.Message = "tenor tidak tersedia"
		res.Result = nil
		return res, err
	}
	// cek limit by tenor
	if f <= limitAcc {
		createTranscation, err := kreditPlusRepository.CreateTransaction(obj)
		if err != nil {
			res.ResponseCode = http.StatusBadRequest
			res.ResponseDesc = "Failed"
			res.ResponseTime = utils.DateToStdNow()
			res.Message = "Gagal membuat transaction"
			res.Result = nil
			return res, err
		}
		res.ResponseCode = http.StatusOK
		res.ResponseDesc = "Successfully"
		res.ResponseTime = utils.DateToStdNow()
		res.Message = "Berhasil Mendapatkan data Customer"
		res.Result = createTranscation
		return res, nil

	} else {
		res.ResponseCode = http.StatusBadRequest
		res.ResponseDesc = "Failed"
		res.ResponseTime = utils.DateToStdNow()
		res.Message = "Harga Barang tidak sesuai Limit anda"
		res.Result = nil
		return res, err
	}
}
