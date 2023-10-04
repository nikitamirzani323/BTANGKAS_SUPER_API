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

const Fieldpattern_home_redis = "LISTPATTERN_BACKEND"
const Fieldpattern_home_client_redis = "LISTPATTERN_FRONTEND"

func Patternhome(c *fiber.Ctx) error {
	var obj entities.Model_pattern
	var arraobj []entities.Model_pattern
	render_page := time.Now()
	resultredis, flag := helpers.GetRedis(Fieldpattern_home_redis)
	jsonredis := []byte(resultredis)
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		pattern_id, _ := jsonparser.GetString(value, "pattern_id")
		pattern_idcard, _ := jsonparser.GetString(value, "pattern_idcard")
		pattern_idpoin, _ := jsonparser.GetInt(value, "pattern_idpoin")
		pattern_nmpoin, _ := jsonparser.GetString(value, "pattern_nmpoin")
		pattern_status, _ := jsonparser.GetString(value, "pattern_status")
		pattern_status_css, _ := jsonparser.GetString(value, "pattern_status_css")
		pattern_create, _ := jsonparser.GetString(value, "pattern_create")
		pattern_update, _ := jsonparser.GetString(value, "pattern_update")

		obj.Pattern_id = pattern_id
		obj.Pattern_idcard = pattern_idcard
		obj.Pattern_idpoin = int(pattern_idpoin)
		obj.Pattern_nmpoin = pattern_nmpoin
		obj.Pattern_status = pattern_status
		obj.Pattern_status_css = pattern_status_css
		obj.Pattern_create = pattern_create
		obj.Pattern_update = pattern_update
		arraobj = append(arraobj, obj)
	})

	if !flag {
		result, err := models.Fetch_patternHome()
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		helpers.SetRedis(Fieldpattern_home_redis, result, 60*time.Minute)
		fmt.Println("PATTERN MYSQL")
		return c.JSON(result)
	} else {
		fmt.Println("PATTERN CACHE")
		return c.JSON(fiber.Map{
			"status":  fiber.StatusOK,
			"message": "Success",
			"record":  arraobj,
			"time":    time.Since(render_page).String(),
		})
	}
}
func PatternSave(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_patternsave)
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

	// aadmin, listpattern, sData string
	result, err := models.Save_pattern(
		client_admin,
		client.Pattern_List, client.Sdata)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	_deleteredis_pattern()
	return c.JSON(result)
}
func _deleteredis_pattern() {
	val_master := helpers.DeleteRedis(Fieldpattern_home_redis)
	fmt.Printf("Redis Delete BACKEND PATTERN : %d", val_master)

	val_client := helpers.DeleteRedis(Fieldpattern_home_client_redis)
	fmt.Printf("Redis Delete CLIENT PATTERN : %d", val_client)

}
