package controllers

import (
	"fmt"
	"strconv"
	"time"

	"github.com/buger/jsonparser"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/nikitamirzani323/BTANGKAS_SUPER_API/entities"
	"github.com/nikitamirzani323/BTANGKAS_SUPER_API/helpers"
	"github.com/nikitamirzani323/BTANGKAS_SUPER_API/models"
)

const Fieldlistpattern_home_redis = "LISTPATTERNGAME_BACKEND"
const Fieldlistpatterndetail_home_redis = "LISTPATTERNGAMEDETAIL_BACKEND"
const Fieldlistpattern_home_client_redis = "LISTPATTERN_FRONTEND"

func Listpatternhome(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_listpattern)
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
	fmt.Println(client.Listpattern_page)
	if client.Listpattern_search != "" {
		val_listpattern := helpers.DeleteRedis(Fieldlistpattern_home_redis + "_" + strconv.Itoa(client.Listpattern_page) + "_" + client.Listpattern_search + "_" + client.Listpattern_search_status)
		fmt.Printf("Redis Delete BACKEND LISTPATTERN : %d", val_listpattern)
	}
	var obj entities.Model_listpattern
	var arraobj []entities.Model_listpattern
	render_page := time.Now()
	resultredis, flag := helpers.GetRedis(Fieldlistpattern_home_redis + "_" + strconv.Itoa(client.Listpattern_page) + "_" + client.Listpattern_search + "_" + client.Listpattern_search_status)
	jsonredis := []byte(resultredis)
	perpage_RD, _ := jsonparser.GetInt(jsonredis, "perpage")
	totalrecord_RD, _ := jsonparser.GetInt(jsonredis, "totalrecord")
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		listpattern_id, _ := jsonparser.GetString(value, "listpattern_id")
		listpattern_nmlistpattern, _ := jsonparser.GetString(value, "listpattern_nmlistpattern")
		listpattern_status, _ := jsonparser.GetString(value, "listpattern_status")
		listpattern_status_css, _ := jsonparser.GetString(value, "listpattern_status_css")
		listpattern_create, _ := jsonparser.GetString(value, "listpattern_create")
		listpattern_update, _ := jsonparser.GetString(value, "listpattern_update")

		obj.Listpattern_id = listpattern_id
		obj.Listpattern_nmlistpattern = listpattern_nmlistpattern
		obj.Listpattern_status = listpattern_status
		obj.Listpattern_status_css = listpattern_status_css
		obj.Listpattern_create = listpattern_create
		obj.Listpattern_update = listpattern_update
		arraobj = append(arraobj, obj)
	})

	if !flag {
		//search, status string, page int
		result, err := models.Fetch_listpatternHome(client.Listpattern_search, client.Listpattern_search_status, client.Listpattern_page)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		helpers.SetRedis(Fieldlistpattern_home_redis+"_"+strconv.Itoa(client.Listpattern_page)+"_"+client.Listpattern_search+"_"+client.Listpattern_search_status, result, 60*time.Minute)
		fmt.Println("LISTPATTERN MYSQL")
		return c.JSON(result)
	} else {
		fmt.Println("LISTPATTERN CACHE")
		return c.JSON(fiber.Map{
			"status":      fiber.StatusOK,
			"message":     "Success",
			"perpage":     perpage_RD,
			"totalrecord": totalrecord_RD,
			"record":      arraobj,
			"time":        time.Since(render_page).String(),
		})
	}
}
func Listpatterndetailhome(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_listpatterndetail)
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

	var obj entities.Model_listpatterndetail
	var arraobj []entities.Model_listpatterndetail
	render_page := time.Now()
	resultredis, flag := helpers.GetRedis(Fieldlistpatterndetail_home_redis + "_" + client.Listpatterndetail_idlistpattern)
	jsonredis := []byte(resultredis)
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		listpatterndetail_id, _ := jsonparser.GetInt(value, "listpatterndetail_id")
		listpatterndetail_idpattern, _ := jsonparser.GetString(value, "listpatterndetail_idpattern")
		listpatterndetail_nmpoin, _ := jsonparser.GetString(value, "listpatterndetail_nmpoin")
		listpatterndetail_status, _ := jsonparser.GetString(value, "listpatterndetail_status")
		listpatterndetail_status_css, _ := jsonparser.GetString(value, "listpatterndetail_status_css")
		listpatterndetail_create, _ := jsonparser.GetString(value, "listpatterndetail_create")
		listpatterndetail_update, _ := jsonparser.GetString(value, "listpatterndetail_update")

		obj.Listpatterndetail_id = int(listpatterndetail_id)
		obj.Listpatterndetail_idpattern = listpatterndetail_idpattern
		obj.Listpatterndetail_nmpoin = listpatterndetail_nmpoin
		obj.Listpatterndetail_status = listpatterndetail_status
		obj.Listpatterndetail_status_css = listpatterndetail_status_css
		obj.Listpatterndetail_create = listpatterndetail_create
		obj.Listpatterndetail_update = listpatterndetail_update
		arraobj = append(arraobj, obj)
	})

	if !flag {
		//search, status string, page int
		result, err := models.Fetch_listpatterndetailHome(client.Listpatterndetail_idlistpattern)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		helpers.SetRedis(Fieldlistpatterndetail_home_redis+"_"+client.Listpatterndetail_idlistpattern, result, 60*time.Minute)
		fmt.Println("LISTPATTERNDETAIL MYSQL")
		return c.JSON(result)
	} else {
		fmt.Println("LISTPATTERNDETAIL CACHE")
		return c.JSON(fiber.Map{
			"status":  fiber.StatusOK,
			"message": "Success",
			"record":  arraobj,
			"time":    time.Since(render_page).String(),
		})
	}
}
func ListpatternSave(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_listpatternsave)
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

	// admin, name, status, idrecord, sData string
	result, err := models.Save_listpattern(
		client_admin,
		client.Listpattern_nmlistpattern, client.Listpattern_status, client.Listpattern_id, client.Sdata)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	_deleteredis_listpattern("", client.Listpattern_search, client.Listpattern_search_status, client.Listpattern_page)
	return c.JSON(result)
}
func ListpatterndetailSave(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_listpatterndetailsave)
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

	// admin, idlistpattern, idpattern, sData string
	result, err := models.Save_listpatterndetail(
		client_admin,
		client.Listpatterndetail_idlistpattern, client.Listpatterndetail_idpattern, client.Sdata)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	_deleteredis_listpattern(client.Listpatterndetail_idlistpattern, "", "", 0)
	return c.JSON(result)
}
func _deleteredis_listpattern(idlistpattern, search, status string, page int) {
	val_master := helpers.DeleteRedis(Fieldlistpattern_home_redis + "_" + strconv.Itoa(page) + "_" + search + "_" + status)
	fmt.Printf("Redis Delete BACKEND LISTPATTERN : %d\n", val_master)

	val_masterdetail := helpers.DeleteRedis(Fieldlistpatterndetail_home_redis + "_" + idlistpattern)
	fmt.Printf("Redis Delete BACKEND LISTPATTERNDETAIL : %d\n", val_masterdetail)
}
