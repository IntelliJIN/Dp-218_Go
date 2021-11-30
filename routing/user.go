package routing

import (
	"net/http"
	"strconv"

<<<<<<< Updated upstream
	model "Dp218Go/domain/entities"
	iface "Dp218Go/domain/interfaces"
	repo "Dp218Go/repositories"

	"github.com/go-chi/render"
=======
	"Dp218Go/services"

>>>>>>> Stashed changes
	"github.com/gorilla/mux"
)

var	userRepo iface.UserRepo
var userIDKey = "userID"

var keyRoutes = []Route{
	{
		Uri:         `/users`,
		Method:    http.MethodGet,
		Handler:	getAllUsers,
	},
	{
		Uri:         `/user/{`+userIDKey+`}`,
		Method:     http.MethodGet,
		Handler:	getUser,
	},
	{
		Uri:         `/user`,
		Method:     http.MethodPost,
		Handler:	createUser,
	},
	{
		Uri:         `/user/{`+userIDKey+`}`,
		Method:     http.MethodPut,
		Handler:	updateUser,
	},
	{
		Uri:         `/user/{`+userIDKey+`}`,
		Method:     http.MethodDelete,
		Handler:	deleteUser,
	},
}

func AddUserHandler(router *mux.Router, repo iface.UserRepo) {
	userRepo = repo
	for _, rt := range keyRoutes{
		router.Path(rt.Uri).HandlerFunc(rt.Handler).Methods(rt.Method)
	}
}

func createUser(w http.ResponseWriter, r *http.Request) {
	user := &model.User{}
	if err := render.Bind(r, user); err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}
	if err := userRepo.AddUser(user); err != nil {
		render.Render(w, r, ErrorRenderer(err))
		return
	}
	if err := render.Render(w, r, user); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
}

func getAllUsers(w http.ResponseWriter, r *http.Request) {
<<<<<<< Updated upstream
	users, err := userRepo.GetAllUsers()
=======
	var users = &models.UserList{}
	var err error
	format := GetFormatFromRequest(r)

	r.ParseForm()
	searchData := r.FormValue("SearchData")
	if len(searchData) == 0 {
		users, err = userService.GetAllUsers()
	} else {
		users, err = userService.FindUsersByLoginNameSurname(searchData)
	}
>>>>>>> Stashed changes
	if err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
	if err := render.Render(w, r, users); err != nil {
		render.Render(w, r, ErrorRenderer(err))
		return
	}
}

func getUser(w http.ResponseWriter, r *http.Request) {
	userId, err := strconv.Atoi(mux.Vars(r)[userIDKey])
	if err != nil {
		render.Render(w, r, ErrorRenderer(err))
	}
	user, err := userRepo.GetUserById(userId)
	if err != nil {
		if err == repo.ErrNoMatch {
			render.Render(w, r, ErrNotFound)
		} else {
			render.Render(w, r, ErrorRenderer(err))
		}
		return
	}
	if err := render.Render(w, r, &user); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	userId, err := strconv.Atoi(mux.Vars(r)[userIDKey])
	if err != nil {
		render.Render(w, r, ErrorRenderer(err))
	}
	err = userRepo.DeleteUser(userId)
	if err != nil {
		if err == repo.ErrNoMatch {
			render.Render(w, r, ErrNotFound)
		} else {
			render.Render(w, r, ServerErrorRenderer(err))
		}
		return
	}
	render.Render(w, r, StatusOK)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	userId, err := strconv.Atoi(mux.Vars(r)[userIDKey])
	if err != nil {
		render.Render(w, r, ErrorRenderer(err))
	}
	userData := model.User{}
	if err := render.Bind(r, &userData); err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}
	userData, err = userRepo.UpdateUser(userId, userData)
	if err != nil {
		if err == repo.ErrNoMatch {
			render.Render(w, r, ErrNotFound)
		} else {
			render.Render(w, r, ServerErrorRenderer(err))
		}
		return
	}
<<<<<<< Updated upstream
	if err := render.Render(w, r, &userData); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
=======

	EncodeAnswer(format, w, &userWithRoleList{userData}, HTMLPath+"user-edit.html")
}

func allUsersOperation(w http.ResponseWriter, r *http.Request) {
	format := GetFormatFromRequest(r)

	r.ParseForm()
	if _, ok := r.Form["ActionType"]; !ok {
		return
	}
	actionType := r.FormValue("ActionType")
	switch actionType {
	case "BlockUser":
		userId, err := strconv.Atoi(r.FormValue("UserID"))
		if err != nil {
			EncodeError(format, w, ErrorRendererDefault(err))
			return
		}
		err = userService.ChangeUsersBlockStatus(userId)
		if err != nil {
			EncodeError(format, w, ErrorRendererDefault(err))
			return
		}
	default:
		EncodeError(format, w, ErrorRendererDefault(fmt.Errorf("unknown users operation")))
	}
	getAllUsers(w, r)
}

func DecodeUserUpdateRequest(r *http.Request, data interface{}) error {

	var err error
	r.ParseForm()
	//userData := models.User{}
	userData := data.(*models.User)

	if _, ok := r.Form["LoginEmail"]; ok {
		userData.LoginEmail = r.FormValue("LoginEmail")

	}
	if _, ok := r.Form["UserName"]; ok {
		userData.UserName = r.FormValue("UserName")
	}
	if _, ok := r.Form["UserSurname"]; ok {
		userData.UserSurname = r.FormValue("UserSurname")
	}
	if _, ok := r.Form["RoleID"]; ok {
		var roleId int
		roleId, err = strconv.Atoi(r.FormValue("RoleID"))
		if err != nil {
			return err
		}
		userData.Role, err = userService.GetRoleById(roleId)
		if err != nil {
			return err
		}
	}
	if _, ok := r.Form["IsBlocked"]; ok {
		userData.IsBlocked, _ = strconv.ParseBool(r.FormValue("IsBlocked"))
	}

	//reflect.ValueOf(data).Elem().Set(reflect.ValueOf(userData))
	return nil
>>>>>>> Stashed changes
}
