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
	render_page := time.Now()
	resultredis, flag := helpers.GetRedis(Fieldcompanyadminrule_home_redis)
	jsonredis := []byte(resultredis)
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		companyadminrule_id, _ := jsonparser.GetString(value, "companyadminrule_id")
		companyadminrule_idcompany, _ := jsonparser.GetString(value, "companyadminrule_idcompany")
		companyadminrule_nmrule, _ := jsonparser.GetString(value, "companyadminrule_nmrule")
		companyadminrule_rule, _ := jsonparser.GetString(value, "companyadminrule_rule")
		companyadminrule_create, _ := jsonparser.GetString(value, "companyadminrule_create")
		companyadminrule_update, _ := jsonparser.GetString(value, "companyadminrule_update")

		obj.Companyadminrule_id = companyadminrule_id
		obj.Companyadminrule_idcompany = companyadminrule_idcompany
		obj.Companyadminrule_nmrule = companyadminrule_nmrule
		obj.Companyadminrule_rule = companyadminrule_rule
		obj.Companyadminrule_create = companyadminrule_create
		obj.Companyadminrule_update = companyadminrule_update
		arraobj = append(arraobj, obj)
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
			"status":  fiber.StatusOK,
			"message": "Success",
			"record":  arraobj,
			"time":    time.Since(render_page).String(),
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
		client.Companyadminrule_id, client.Companyadminrule_idcompany,
		client.Companyadminrule_nmrule, client.Companyadminrule_rule, client.Sdata)
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

	val_master_adminrule := helpers.DeleteRedis(Fieldcompanyadminrule_home_redis)
	fmt.Printf("Redis Delete BACKEND COMPANY ADMIN RULE : %d", val_master_adminrule)
}
