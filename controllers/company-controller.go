package controllers

import (
	"fmt"
	"time"

	"github.com/buger/jsonparser"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/nikitamirzani323/BTANGKAS_SUPER_API/entities"
	"github.com/nikitamirzani323/BTANGKAS_SUPER_API/helpers"
	"github.com/nikitamirzani323/BTANGKAS_SUPER_API/models"
)

const Fieldcompany_home_redis = "COMPANY_BACKEND"
const Fieldcompanyadminrule_home_redis = "COMPANYADMINRULE_BACKEND"
const Fieldcompanyadmin_home_redis = "COMPANYADMIN_BACKEND"
const Fieldcompany_home_client_redis = "COMPANY_FRONTEND"

func Companyhome(c *fiber.Ctx) error {
	var obj entities.Model_company
	var arraobj []entities.Model_company
	var objcurr entities.Model_currshare
	var arraobjcurr []entities.Model_currshare
	render_page := time.Now()
	resultredis, flag := helpers.GetRedis(Fieldcompany_home_redis)
	jsonredis := []byte(resultredis)
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	listcurr_RD, _, _, _ := jsonparser.Get(jsonredis, "listcurr")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		company_id, _ := jsonparser.GetString(value, "company_id")
		company_startjoin, _ := jsonparser.GetString(value, "company_startjoin")
		company_endjoin, _ := jsonparser.GetString(value, "company_endjoin")
		company_name, _ := jsonparser.GetString(value, "company_name")
		company_idcurr, _ := jsonparser.GetString(value, "company_idcurr")
		company_nmowner, _ := jsonparser.GetString(value, "company_nmowner")
		company_phoneowner, _ := jsonparser.GetString(value, "company_phoneowner")
		company_emailowner, _ := jsonparser.GetString(value, "company_emailowner")
		company_url, _ := jsonparser.GetString(value, "company_url")
		company_status, _ := jsonparser.GetString(value, "company_status")
		company_status_css, _ := jsonparser.GetString(value, "company_status_css")
		company_create, _ := jsonparser.GetString(value, "company_create")
		company_update, _ := jsonparser.GetString(value, "company_update")

		obj.Company_id = company_id
		obj.Company_startjoin = company_startjoin
		obj.Company_endjoin = company_endjoin
		obj.Company_name = company_name
		obj.Company_idcurr = company_idcurr
		obj.Company_nmowner = company_nmowner
		obj.Company_phoneowner = company_phoneowner
		obj.Company_emailowner = company_emailowner
		obj.Company_url = company_url
		obj.Company_status = company_status
		obj.Company_status_css = company_status_css
		obj.Company_create = company_create
		obj.Company_update = company_update
		arraobj = append(arraobj, obj)
	})
	jsonparser.ArrayEach(listcurr_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		curr_id, _ := jsonparser.GetString(value, "curr_id")

		objcurr.Curr_id = curr_id
		arraobjcurr = append(arraobjcurr, objcurr)
	})
	if !flag {
		result, err := models.Fetch_companyHome()
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		helpers.SetRedis(Fieldcompany_home_redis, result, 60*time.Minute)
		fmt.Println("COMPANY MYSQL")
		return c.JSON(result)
	} else {
		fmt.Println("COMPANY CACHE")
		return c.JSON(fiber.Map{
			"status":   fiber.StatusOK,
			"message":  "Success",
			"record":   arraobj,
			"listcurr": arraobjcurr,
			"time":     time.Since(render_page).String(),
		})
	}
}
func Companyadminrulehome(c *fiber.Ctx) error {
	var obj entities.Model_companyadminrule
	var arraobj []entities.Model_companyadminrule
	var objcompany entities.Model_companyshare
	var arraobjcompany []entities.Model_companyshare
	render_page := time.Now()
	resultredis, flag := helpers.GetRedis(Fieldcompanyadminrule_home_redis)
	jsonredis := []byte(resultredis)
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	listcompany_RD, _, _, _ := jsonparser.Get(jsonredis, "listcompany")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		companyadminrule_id, _ := jsonparser.GetInt(value, "companyadminrule_id")
		companyadminrule_idcompany, _ := jsonparser.GetString(value, "companyadminrule_idcompany")
		companyadminrule_nmrule, _ := jsonparser.GetString(value, "companyadminrule_nmrule")
		companyadminrule_rule, _ := jsonparser.GetString(value, "companyadminrule_rule")
		companyadminrule_create, _ := jsonparser.GetString(value, "companyadminrule_create")
		companyadminrule_update, _ := jsonparser.GetString(value, "companyadminrule_update")

		obj.Companyadminrule_id = int(companyadminrule_id)
		obj.Companyadminrule_idcompany = companyadminrule_idcompany
		obj.Companyadminrule_nmrule = companyadminrule_nmrule
		obj.Companyadminrule_rule = companyadminrule_rule
		obj.Companyadminrule_create = companyadminrule_create
		obj.Companyadminrule_update = companyadminrule_update
		arraobj = append(arraobj, obj)
	})
	jsonparser.ArrayEach(listcompany_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		company_id, _ := jsonparser.GetString(value, "company_id")
		company_name, _ := jsonparser.GetString(value, "company_name")

		objcompany.Company_id = company_id
		objcompany.Company_name = company_name
		arraobjcompany = append(arraobjcompany, objcompany)
	})
	if !flag {
		result, err := models.Fetch_companyadminruleHome()
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		helpers.SetRedis(Fieldcompanyadminrule_home_redis, result, 60*time.Minute)
		fmt.Println("COMPANY ADMIN GROUP MYSQL")
		return c.JSON(result)
	} else {
		fmt.Println("COMPANY ADMIN GROUP CACHE")
		return c.JSON(fiber.Map{
			"status":      fiber.StatusOK,
			"message":     "Success",
			"record":      arraobj,
			"listcompany": arraobjcompany,
			"time":        time.Since(render_page).String(),
		})
	}
}
func Companyadminhome(c *fiber.Ctx) error {
	var obj entities.Model_companyadmin
	var arraobj []entities.Model_companyadmin
	var objcompany entities.Model_companyshare
	var arraobjcompany []entities.Model_companyshare
	var objrule entities.Model_companyadminrule_share
	var arraobjrule []entities.Model_companyadminrule_share
	render_page := time.Now()
	resultredis, flag := helpers.GetRedis(Fieldcompanyadmin_home_redis)
	jsonredis := []byte(resultredis)
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	listcompany_RD, _, _, _ := jsonparser.Get(jsonredis, "listcompany")
	listrule_RD, _, _, _ := jsonparser.Get(jsonredis, "listrule")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		companyadmin_id, _ := jsonparser.GetString(value, "companyadmin_id")
		companyadmin_idrule, _ := jsonparser.GetInt(value, "companyadmin_idrule")
		companyadmin_idcompany, _ := jsonparser.GetString(value, "companyadmin_idcompany")
		companyadmin_tipe, _ := jsonparser.GetString(value, "companyadmin_tipe")
		companyadmin_username, _ := jsonparser.GetString(value, "companyadmin_username")
		companyadmin_ipaddress, _ := jsonparser.GetString(value, "companyadmin_ipaddress")
		companyadmin_lastlogin, _ := jsonparser.GetString(value, "companyadmin_lastlogin")
		companyadmin_name, _ := jsonparser.GetString(value, "companyadmin_name")
		companyadmin_phone1, _ := jsonparser.GetString(value, "companyadmin_phone1")
		companyadmin_phone2, _ := jsonparser.GetString(value, "companyadmin_phone2")
		companyadmin_status, _ := jsonparser.GetString(value, "companyadmin_status")
		companyadmin_status_css, _ := jsonparser.GetString(value, "companyadmin_status_css")
		companyadmin_create, _ := jsonparser.GetString(value, "companyadmin_create")
		companyadmin_update, _ := jsonparser.GetString(value, "companyadmin_update")

		obj.Companyadmin_id = companyadmin_id
		obj.Companyadmin_idcompany = companyadmin_idcompany
		obj.Companyadmin_idrule = int(companyadmin_idrule)
		obj.Companyadmin_tipe = companyadmin_tipe
		obj.Companyadmin_username = companyadmin_username
		obj.Companyadmin_ipaddress = companyadmin_ipaddress
		obj.Companyadmin_lastlogin = companyadmin_lastlogin
		obj.Companyadmin_name = companyadmin_name
		obj.Companyadmin_phone1 = companyadmin_phone1
		obj.Companyadmin_phone2 = companyadmin_phone2
		obj.Companyadmin_status = companyadmin_status
		obj.Companyadmin_status_css = companyadmin_status_css
		obj.Companyadmin_create = companyadmin_create
		obj.Companyadmin_update = companyadmin_update
		arraobj = append(arraobj, obj)
	})
	jsonparser.ArrayEach(listcompany_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		company_id, _ := jsonparser.GetString(value, "company_id")
		company_name, _ := jsonparser.GetString(value, "company_name")

		objcompany.Company_id = company_id
		objcompany.Company_name = company_name
		arraobjcompany = append(arraobjcompany, objcompany)
	})
	jsonparser.ArrayEach(listrule_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		companyadminrule_id, _ := jsonparser.GetInt(value, "companyadminrule_id")
		companyadminrule_idcompany, _ := jsonparser.GetString(value, "companyadminrule_idcompany")
		companyadminrule_nmrule, _ := jsonparser.GetString(value, "companyadminrule_nmrule")

		objrule.Companyadminrule_id = int(companyadminrule_id)
		objrule.Companyadminrule_idcompany = companyadminrule_idcompany
		objrule.Companyadminrule_nmrule = companyadminrule_nmrule
		arraobjrule = append(arraobjrule, objrule)
	})
	if !flag {
		result, err := models.Fetch_companyadminHome()
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		helpers.SetRedis(Fieldcompanyadmin_home_redis, result, 60*time.Minute)
		fmt.Println("COMPANY ADMIN  MYSQL")
		return c.JSON(result)
	} else {
		fmt.Println("COMPANY ADMIN  CACHE")
		return c.JSON(fiber.Map{
			"status":      fiber.StatusOK,
			"message":     "Success",
			"record":      arraobj,
			"listcompany": arraobjcompany,
			"listrule":    arraobjrule,
			"time":        time.Since(render_page).String(),
		})
	}
}
func CompanySave(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_companysave)
	validate := validator.New()
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	err := validate.Struct(client)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element helpers.ErrorResponse
			element.Field = err.StructField()
			element.Tag = err.Tag()
			errors = append(errors, &element)
		}
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "validation",
			"record":  errors,
		})
	}
	user := c.Locals("jwt").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	temp_decp := helpers.Decryption(name)
	client_admin, _ := helpers.Parsing_Decry(temp_decp, "==")

	// aadmin, idrecord, code, idcurr, nmcompany, nmowner, phoneowner, emailowner, url, status, sData string
	result, err := models.Save_company(
		client_admin,
		client.Company_id, client.Company_idcurr,
		client.Company_name, client.Company_nmowner, client.Company_phoneowner, client.Company_emailowner,
		client.Company_url, client.Company_status,
		client.Sdata)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	_deleteredis_company()
	return c.JSON(result)
}
func CompanyadminruleSave(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_companyadminrulesave)
	validate := validator.New()
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	err := validate.Struct(client)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element helpers.ErrorResponse
			element.Field = err.StructField()
			element.Tag = err.Tag()
			errors = append(errors, &element)
		}
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "validation",
			"record":  errors,
		})
	}
	user := c.Locals("jwt").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	temp_decp := helpers.Decryption(name)
	client_admin, _ := helpers.Parsing_Decry(temp_decp, "==")

	// admin, idrecord, idcompany, name, rule, sData string
	result, err := models.Save_companyadminrule(
		client_admin,
		client.Companyadminrule_idcompany,
		client.Companyadminrule_nmrule, client.Companyadminrule_rule, client.Sdata, client.Companyadminrule_id)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	_deleteredis_company()
	return c.JSON(result)
}
func CompanyadminSave(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_companyadminsave)
	validate := validator.New()
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	err := validate.Struct(client)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element helpers.ErrorResponse
			element.Field = err.StructField()
			element.Tag = err.Tag()
			errors = append(errors, &element)
		}
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "validation",
			"record":  errors,
		})
	}
	user := c.Locals("jwt").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	temp_decp := helpers.Decryption(name)
	client_admin, _ := helpers.Parsing_Decry(temp_decp, "==")

	// admin, idrecord, idcompany, idrule, username, password, name, phone1, phone2, status, sData string
	result, err := models.Save_companyadmin(
		client_admin,
		client.Companyadmin_id, client.Companyadmin_idcompany,
		client.Companyadmin_username, client.Companyadmin_password, client.Companyadmin_name,
		client.Companyadmin_phone1, client.Companyadmin_phone2, client.Companyadmin_status, client.Sdata, client.Companyadmin_idrule)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	_deleteredis_company()
	return c.JSON(result)
}
func _deleteredis_company() {
	val_master := helpers.DeleteRedis(Fieldcompany_home_redis)
	fmt.Printf("Redis Delete BACKEND COMPANY : %d", val_master)

	val_master_admin := helpers.DeleteRedis(Fieldcompanyadmin_home_redis)
	fmt.Printf("Redis Delete BACKEND COMPANY ADMIN  : %d", val_master_admin)

	val_master_adminrule := helpers.DeleteRedis(Fieldcompanyadminrule_home_redis)
	fmt.Printf("Redis Delete BACKEND COMPANY ADMIN RULE : %d", val_master_adminrule)
}
