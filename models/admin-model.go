package models

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/nikitamirzani323/gosveltemdb/configs"
	"github.com/nikitamirzani323/gosveltemdb/db"
	"github.com/nikitamirzani323/gosveltemdb/entities"
	"github.com/nikitamirzani323/gosveltemdb/helpers"
)

func Fetch_adminHome() (helpers.ResponseAdmin, error) {
	var obj entities.Model_admin
	var arraobj []entities.Model_admin
	var res helpers.ResponseAdmin
	msg := "Error"
	con := db.CreateCon()
	ctx := context.Background()
	start := time.Now()

	sql_select := `SELECT 
			username , name, idadminlevel,
			statuslogin, lastlogin, joindate, 
			ipaddress, timezone  
			FROM ` + configs.DB_tbl_admin + ` 
			ORDER BY lastlogin DESC 
		`

	row, err := con.QueryContext(ctx, sql_select)

	var no int = 0
	helpers.ErrorCheck(err)
	for row.Next() {
		no += 1
		var (
			idadminlevel_db                                                      int
			username_db, name_db                                                 string
			statuslogin_db, lastlogin_db, joindate_db, ipaddress_db, timezone_db string
		)

		err = row.Scan(
			&username_db, &name_db, &idadminlevel_db,
			&statuslogin_db, &lastlogin_db, &joindate_db,
			&ipaddress_db, &timezone_db)

		helpers.ErrorCheck(err)

		obj.Username = username_db
		obj.Nama = name_db
		obj.Rule = Get_AdminRule("ruleadmingroup", idadminlevel_db)
		obj.Joindate = joindate_db
		obj.Timezone = timezone_db
		obj.Lastlogin = lastlogin_db
		obj.LastIpaddress = ipaddress_db
		obj.Status = statuslogin_db
		arraobj = append(arraobj, obj)
		msg = "Success"
	}
	defer row.Close()

	var objRule entities.Model_adminrule
	var arraobjRule []entities.Model_adminrule
	sql_listrule := `SELECT 
		idadmin 	
		FROM ` + configs.DB_tbl_admingroup + ` 
	`
	row_listrule, err_listrule := con.QueryContext(ctx, sql_listrule)

	helpers.ErrorCheck(err_listrule)
	for row_listrule.Next() {
		var (
			idruleadmin_db string
		)

		err = row_listrule.Scan(&idruleadmin_db)

		helpers.ErrorCheck(err)

		objRule.Idrule = idruleadmin_db
		arraobjRule = append(arraobjRule, objRule)
		msg = "Success"
	}

	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = arraobj
	res.Listrule = arraobjRule
	res.Time = time.Since(start).String()

	return res, nil
}
