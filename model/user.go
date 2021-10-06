package model

import (
	"database/sql"
	"errors"

	"github.com/WahyuSiddarta/unimportant/conf"
	"github.com/WahyuSiddarta/unimportant/helper"
	"github.com/WahyuSiddarta/unimportant/structure"
)

func (db SQLDB) MGetUserProfileByUserID(user_id string) (structure.UserLogin, error) {
	var user structure.UserLogin = structure.UserLogin{}
	stmt, err := db.DB[conf.DBPSQL].Preparex("SELECT email, password, created_at, user_id, status from d_user WHERE user_id=$1")
	if err != nil {
		Logger.Error().Msgf("MGetUserProfileByUserID-1:%s", err.Error())
		return structure.UserLogin{}, err
	}
	defer stmt.Close()

	err = stmt.QueryRowx(user_id).StructScan(&user)
	if err != nil {
		if err != sql.ErrNoRows {
			Logger.Error().Msgf("MGetUserProfileByUserID-x:%s", err.Error())
		}
		return user, err
	}
	return user, nil
}
func (db SQLDB) MGetUserByEmail(email string) (structure.UserLogin, error) {
	var user structure.UserLogin = structure.UserLogin{}
	stmt, err := db.DB[conf.DBPSQL].Preparex("SELECT email, password, created_at, user_id, status from d_user WHERE email=$1")
	if err != nil {
		Logger.Error().Msgf("MGetUserByEmail-1:%s", err.Error())
		return structure.UserLogin{}, err
	}
	defer stmt.Close()

	err = stmt.QueryRowx(email).StructScan(&user)
	if err != nil {
		if err != sql.ErrNoRows {
			Logger.Error().Msgf("MGetUserByEmail-x:%s", err.Error())
		}
		return user, err
	}
	return user, nil
}

func (db SQLDB) MCreateNewUser(email string, password string) (string, error) {
	user_id := helper.NewUUID()

	tx, err := db.DB[conf.DBPSQL].Beginx()
	if err != nil {
		return user_id, err
	}

	var test string = ""
	err = tx.QueryRowx(`SELECT user_id FROM d_user WHERE email = $1`, email).Scan(&test)
	if err != nil {
		if err != sql.ErrNoRows {
			tx.Rollback()
			Logger.Error().Msgf("MRegisterNew-0:%s", err.Error())
			return user_id, err
		} else if test != "" {
			tx.Rollback()
			Logger.Error().Msgf("MRegisterNew-username already exist:%s", err.Error())
			err = errors.New("user already exist")
			return user_id, err
		}
	}

	stmt1, err := tx.Prepare(`INSERT into d_user_login (email, password, created_at, user_id, status) values ($1, $2, now(), $3, 0)`)
	if err != nil {
		tx.Rollback()
		Logger.Error().Msgf("MRegisterNew-1:%s", err.Error())
		return user_id, err
	}
	defer stmt1.Close()

	_, err = stmt1.Exec(email, password, user_id)
	if err != nil {
		tx.Rollback()
		Logger.Error().Msgf("MRegisterNew-2:%s", err.Error())
		return user_id, err
	}

	stmt2, err := tx.Preparex(`INSERT into d_user_profile (email, created_at, user_id) values ($1, now, $2)`)
	if err != nil {
		tx.Rollback()
		Logger.Error().Msgf("MRegisterNew-3:%s", err.Error())
		return user_id, err
	}
	defer stmt2.Close()

	_, err = stmt2.Exec(email, user_id)
	if err != nil {
		tx.Rollback()
		Logger.Error().Msgf("MRegisterNew-4:%s", err.Error())
		return user_id, err
	}

	tx.Commit()
	return user_id, nil
}
