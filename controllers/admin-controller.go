package controllers

import (
	"log"
	"time"

	"github.com/buger/jsonparser"
	"github.com/gofiber/fiber/v2"
	"github.com/nikitamirzani323/gosveltemdb/helpers"
	"github.com/nikitamirzani323/gosveltemdb/models"
)

func Adminhome(c *fiber.Ctx) error {
	field_redis := "LISTADMIN_AGENT"

	var obj entities.responseredis_adminhome
	var arraobj []entities.responseredis_adminhome
	var obj_listruleadmin entities.responseredis_adminhome_listruleadmin
	var arraobj_listruleadmin []entities.responseredis_adminhome_listruleadmin
	render_page := time.Now()
	resultredis, flag := helpers.GetRedis(field_redis)
	jsonredis := []byte(resultredis)
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	listruleadmin_RD, _, _, _ := jsonparser.Get(jsonredis, "listruleadmin")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		admin_no, _ := jsonparser.GetInt(value, "admin_no")
		admin_username, _ := jsonparser.GetString(value, "admin_username")
		admin_nama, _ := jsonparser.GetString(value, "admin_nama")
		admin_rule, _ := jsonparser.GetString(value, "admin_rule")
		admin_joindate, _ := jsonparser.GetString(value, "admin_joindate")
		admin_timezone, _ := jsonparser.GetString(value, "admin_timezone")
		admin_lastlogin, _ := jsonparser.GetString(value, "admin_lastlogin")
		admin_lastipaddres, _ := jsonparser.GetString(value, "admin_lastipaddres")
		admin_status, _ := jsonparser.GetString(value, "admin_status")

		obj.Admin_no = int(admin_no)
		obj.Admin_username = admin_username
		obj.Admin_nama = admin_nama
		obj.Admin_rule = admin_rule
		obj.Admin_joindate = admin_joindate
		obj.Admin_timezone = admin_timezone
		obj.Admin_lastlogin = admin_lastlogin
		obj.Admin_lastipaddres = admin_lastipaddres
		obj.Admin_status = admin_status
		arraobj = append(arraobj, obj)
	})
	jsonparser.ArrayEach(listruleadmin_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		adminrule_idruleadmin, _ := jsonparser.GetInt(value, "adminrule_idruleadmin")
		adminrule_name, _ := jsonparser.GetString(value, "adminrule_name")

		obj_listruleadmin.Adminrule_idruleadmin = int(adminrule_idruleadmin)
		obj_listruleadmin.Adminrule_name = adminrule_name
		arraobj_listruleadmin = append(arraobj_listruleadmin, obj_listruleadmin)
	})
	if !flag {
		result, err := models.Fetch_adminHome()
		helpers.SetRedis(field_redis, result, 5*time.Minute)
		log.Println("MYSQL")
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		return c.JSON(result)
	} else {
		log.Println("cache")
		return c.JSON(fiber.Map{
			"status":        fiber.StatusOK,
			"message":       "Success",
			"record":        arraobj,
			"listruleadmin": arraobj_listruleadmin,
			"time":          time.Since(render_page).String(),
		})
	}
}
