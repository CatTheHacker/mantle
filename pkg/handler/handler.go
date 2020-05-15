package handler

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/nektro/mantle/pkg/db"
	"github.com/nektro/mantle/pkg/ws"

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
	ru.SetName(strings.ReplaceAll(name, " ", ""))
}

// InviteGet is handler for GET /invite
func InviteGet(w http.ResponseWriter, r *http.Request) {
	etc.WriteHandlebarsFile(r, w, "/invite.hbs", map[string]interface{}{
		"data": db.Props.GetAll(),
		"code": r.URL.Query().Get("code"),
	})
}

// InvitePost is handler for POST /invite
func InvitePost(w http.ResponseWriter, r *http.Request) {
	if ok, _ := strconv.ParseBool(db.Props.Get("public")); !ok {
		s := etc.GetSession(r)
		s.Values["code"] = r.URL.Query().Get("code")
		s.Save(r, w)
	}
	w.Header().Add("Location", "./login")
	w.WriteHeader(http.StatusFound)
}

// Verify is handler for /verify
func Verify(w http.ResponseWriter, r *http.Request) {
	sess, user, err := apiBootstrapRequireLogin(r, w, http.MethodGet, false)
	if err != nil {
		return
	}
	if !db.IsUID(user.UUID) {
		user.ResetUID()
		sess := etc.GetSession(r)
		sess.Values["user"] = user.UUID
		sess.Save(r, w)
	}
	if user.IsMember {
		w.Header().Add("Location", "./chat/")
		w.WriteHeader(http.StatusFound)
		return
	}
	if o, _ := strconv.ParseBool(db.Props.Get("public")); o {
		if !user.IsMember {
			user.SetAsMember(true)
		}
		w.Header().Add("Location", "./chat/")
		w.WriteHeader(http.StatusFound)
		return
	}
	code, ok := sess.Values["code"].(string)
	if !ok {
		writeAPIResponse(r, w, false, http.StatusBadRequest, "invite code required to enter")
		return
	}
	inv, ok := db.QueryInviteByCode(code)
	if !ok {
		writeAPIResponse(r, w, false, http.StatusBadRequest, "invalid invite code")
		return
	}
	if inv.IsFrozen {
		writeAPIResponse(r, w, false, http.StatusBadRequest, "invite is frozen and can not be used")
		return
	}
	if inv.MaxUses > 0 && inv.Uses >= inv.MaxUses {
		writeAPIResponse(r, w, false, http.StatusBadRequest, "invite use count has been exceeded")
		return
	}
	switch inv.Mode {
	case 0:
		// permanent
	case 1:
		//
	case 2:
		s := time.Since(inv.ExpiresOn.T())
		if s > 0 {
			writeAPIResponse(r, w, false, http.StatusBadRequest, "invite is expired")
			return
		}
	}
	//
	inv.Use()
	user.SetAsMember(true)
	for _, item := range inv.GivenRoles {
		user.AddRole(item)
	}
	w.Header().Add("Location", "./chat/")
	w.WriteHeader(http.StatusFound)
}

func Chat(w http.ResponseWriter, r *http.Request) {
	etc.WriteHandlebarsFile(r, w, "/chat/index.hbs", map[string]interface{}{
		"data": db.Props.GetAll(),
	})
}

// ApiAbout is handler for /api/about
func ApiAbout(w http.ResponseWriter, r *http.Request) {
	writeAPIResponse(r, w, true, http.StatusOK, db.Props.GetSome("name", "owner", "public", "description", "cover_photo", "profile_photo", "version"))
}

func ApiPropertyUpdate(w http.ResponseWriter, r *http.Request) {
	_, user, err := apiBootstrapRequireLogin(r, w, http.MethodPut, true)
	if err != nil {
		return
	}
	if hGrabFormStrings(r, w, "p_name", "p_value") != nil {
		return
	}
	n := r.Form.Get("p_name")
	v := r.Form.Get("p_value")
	usp := ws.UserPerms{}.From(user)
	if !usp.ManageServer {
		writeAPIResponse(r, w, false, http.StatusForbidden, "users require the manage_server permission to update properties.")
		return
	}
	if !db.Props.Has(n) {
		writeAPIResponse(r, w, false, http.StatusBadRequest, "specified property does not exist.")
		return
	}
	db.Props.Set(n, v)
	db.CreateAudit(db.ActionSettingUpdate, user, "", n, v)
	writeAPIResponse(r, w, true, http.StatusOK, []string{n, v})
}
