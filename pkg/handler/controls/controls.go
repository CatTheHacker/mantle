package controls

import (
	"net/http"

	"github.com/nektro/mantle/pkg/db"

	"github.com/gorilla/mux"
	"github.com/nektro/go-util/arrays/stringsu"
	etc "github.com/nektro/go.etc"
	"github.com/nektro/go.etc/dbt"
	"github.com/nektro/go.etc/htp"
)

var formMethods = []string{http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete}

// GetUser asserts a user is logged in
func GetUser(c *htp.Controller, r *http.Request, w http.ResponseWriter) *db.User {
	l := etc.JWTGetClaims(c, r)
	//
	userID := l["sub"].(string)
	user, ok := db.QueryUserByUUID(dbt.UUID(userID))
	c.Assert(ok, "500: unable to find user: "+userID)

	method := r.Method
	if stringsu.Contains(formMethods, method) {
		r.Method = http.MethodPost
		r.ParseMultipartForm(0)
		r.Method = method
	}

	w.Header().Add("x-m-jwt-iss", l["iss"].(string))
	w.Header().Add("x-m-jwt-sub", l["sub"].(string))

	return user
}

// GetMemberUser asserts the user is a member and not banned
func GetMemberUser(c *htp.Controller, r *http.Request, w http.ResponseWriter) *db.User {
	u := GetUser(c, r, w)
	c.Assert(u.IsMember, "403: you are not a member of this server")
	c.Assert(!u.IsBanned, "403: you are banned")
	return u
}

// GetUIDFromPath retrieves the uuid from the url path ans asserts its validity as a ulid
func GetUIDFromPath(c *htp.Controller, r *http.Request) dbt.UUID {
	up := mux.Vars(r)["uuid"]
	uo := dbt.UUID(up)
	c.Assert(dbt.IsUUID(uo), "400: 'uid' must be a valid ULID")
	return uo
}
