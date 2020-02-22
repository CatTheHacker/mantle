package main

import (
	"github.com/nektro/mantle/pkg/db"
	"github.com/nektro/mantle/pkg/iconst"

	. "github.com/nektro/go-util/alias"
)

func queryAllChannels() []db.Channel {
	result := []db.Channel{}
	rows := db.DB.Build().Se("*").Fr(iconst.TableChannels).Exe()
	for rows.Next() {
		result = append(result, *db.ScanChannel(rows))
	}
	rows.Close()
	return result
}

func queryUserByUUID(uid string) (*db.User, bool) {
	rows := db.DB.Build().Se("*").Fr(iconst.TableUsers).Wh("uuid", uid).Exe()
	if !rows.Next() {
		return &db.User{}, false
	}
	ru := db.ScanUser(rows)
	rows.Close()
	return ru, true
}

func queryUserBySnowflake(provider string, flake string, name string) *db.User {
	rows := db.DB.Build().Se("*").Fr(iconst.TableUsers).Wh("provider", provider).Wh("snowflake", flake).Exe()
	if rows.Next() {
		ru := db.ScanUser(rows)
		rows.Close()
		return ru
	}
	// else
	id := db.DB.QueryNextID(iconst.TableUsers)
	uid := newUUID()
	now := T()
	roles := ""
	if id == 1 {
		roles += "o"
		db.Props.Set("owner", uid)
	}
	db.DB.QueryPrepared(true, F("insert into %s values ('%d', '%s', '%s', '%s', '0', '0', ?, '', '%s', '%s', '%s')", iconst.TableUsers, id, provider, flake, uid, now, now, roles), name)
	return queryUserBySnowflake(provider, flake, name)
}

func queryAssertUserName(uid string, name string) {
	db.DB.Build().Up(iconst.TableUsers, "name", name).Wh("uuid", uid).Exe()
}

func queryAllRoles() []db.Role {
	result := []db.Role{}
	rows := db.DB.Build().Se("*").Fr(iconst.TableRoles).Or("position", "asc").Exe()
	for rows.Next() {
		result = append(result, *db.ScanRole(rows))
	}
	rows.Close()
	return result
}
