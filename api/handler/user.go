package handler

import (
	"github.com/Salikhov079/library-auth/api/token"

	pb "github.com/Salikhov079/library-auth/genprotos"

	"github.com/gin-gonic/gin"
)

// @Summary 		Register User
// @Description 	Register page
// @Tags 			User
// @Accept  		json
// @Produce  		json
// @Param   		Register  body     pb.UserReq  true   "Register"
// @Success 		200   {string}   string    "Register Successful"
// @Failure 		401   {string}   string    "Error while Register"
// @Router 			/user/register [post]
func (h *Handler) RegisterUser(ctx *gin.Context) {
	user := &pb.UserReq{}
	err := ctx.BindJSON(user)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	_, err = h.User.RegisterUser(ctx, user)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	ctx.JSON(200, "Success!!!")
}

// @Summary			Update User
// @Description 	Update page
// @Tags 			User
// @Accept  		json
// @Security  		BearerAuth
// @Produce  		json
// @Param     		id 		path   string     true   "User ID"
// @Param   		Update  body   pb.UserRes    true   "Update"
// @Success 		200   {string} string    "Update Successful"
// @Failure 		401   {string} string    "Error while updated"
// @Router 			/user/update/{id} [put]
func (h *Handler) UpdateUser(ctx *gin.Context) {
	user := &pb.UserRes{}
	err := ctx.BindJSON(user)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	_, err = h.User.UpdateUser(ctx, user)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	ctx.JSON(200, "Success!!!")
}

// @Summary			Delete User
// @Description 	Delete page
// @Tags 			User
// @Accept  		json
// @Security  		BearerAuth
// @Produce  		json
// @Param     		id   path     string   true   "User ID"
// @Success 		200 {string}  string   "Delete Successful"
// @Failure 		401 {string}  string   "Error while Deleted"
// @Router 			/user/delete/{id} [delete]
func (h *Handler) DeleteUser(ctx *gin.Context) {
	id := pb.ById{Id: ctx.Param("id")}

	_, err := h.User.DeleteUser(ctx, &id)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	ctx.JSON(200, "Success!!!")
}

// @Summary 		GetAll User
// @Description 	GetAll page
// @Tags 			User
// @Accept  		json
// @Produce  		json
// @Security  		BearerAuth
// @Param 			query  query  pb.UserReq true "Query parameter"
// @Success 		200 {object}  pb.AllUsers     "GetAll Successful"
// @Failure 		401 {string}  string  		  "Error while GetAll"
// @Router 			/user/getall  [get]
func (h *Handler) GetAllUser(ctx *gin.Context) {
	user := &pb.UserReq{}
	user.Email = ctx.Param("email")
	user.UserName = ctx.Param("user_name")

	res, err := h.User.GetAllUsers(ctx, user)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	ctx.JSON(200, res)
}

// @Summary 		GetById User
// @Description 	GetById page
// @Tags 			User
// @Accept  		json
// @Produce  		json
// @Security  		BearerAuth
// @Param     		id   path      string   true    "User ID"
// @Success 		200 {object}   pb.UserRes  "GetById Successful"
// @Failure 		401 {string}   string      "Error while GetByIdd"
// @Router 			/user/get/{id} [get]
func (h *Handler) GetbyIdUser(ctx *gin.Context) {
	id := pb.ById{Id: ctx.Param("id")}

	res, err := h.User.GetByIdUser(ctx, &id)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	ctx.JSON(200, res)
}

// @Summary 		/LoginUser
// @Description 	LoginUser page
// @Tags 			User
// @Accept  		json
// @Produce  		json
// @Param   		Login  body  pb.UserReq     true     "Login"
// @Success 		200 {object}  token.Tokens  "LoginUser Successful"
// @Failure 		401 {string}  string   "Error while LoginUserd"
// @Router 			/user/login [post]
func (h *Handler) LoginUser(ctx *gin.Context) {
	user := &pb.UserReq{}
	err := ctx.ShouldBindJSON(user)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	res, err := h.User.LoginUser(ctx, user)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	t := token.GenereteJWTToken(res)
	ctx.JSON(200, t)
}
