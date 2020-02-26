package handler

import (
	"net/http"
	"strconv"

	"github.com/nektro/mantle/pkg/db"

	"github.com/nektro/go-util/util"
	etc "github.com/nektro/go.etc"
)

// SaveOAuth2InfoCb saves info from go.oauth to user session cookie
func SaveOAuth2InfoCb(w http.ResponseWriter, r *http.Request, provider string, id string, name string, oa2resp map[string]interface{}) {
	ru := db.QueryUserBySnowflake(provider, id, name)
	util.Log("[user-login]", provider, id, ru.UUID, name)
	sess := etc.GetSession(r)
	sess.Values["user"] = ru.UUID
	sess.Save(r, w)
	ru.SetName(name)
}

// Invite is handler for /invite
func Invite(w http.ResponseWriter, r *http.Request) {
	_, user, _ := apiBootstrapRequireLogin(r, w, http.MethodGet, false)

	if o, _ := strconv.ParseBool(db.Props.Get("public")); o {
		if !user.IsMember {
			user.SetAsMember(true)
			util.Log("[user-join]", "User", user.UUID, "just became a member and joined the server")
		}
		w.Header().Add("Location", "./chat/")
		w.WriteHeader(http.StatusFound)
	}
}

// ApiAbout is handler for /api/about
func ApiAbout(w http.ResponseWriter, r *http.Request) {
	writeAPIResponse(r, w, true, http.StatusOK, db.Props.GetAll())
}
