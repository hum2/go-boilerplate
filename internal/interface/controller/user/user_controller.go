//go:generate oapi-codegen -include-tags user -generate types -o ./gen/types.gen.go  -package gen $PROJECT_DIR/openapi/oapi-codegen.gen.yaml
//go:generate oapi-codegen -include-tags user -generate gin   -o ./gen/server.gen.go -package gen $PROJECT_DIR/openapi/oapi-codegen.gen.yaml
package user

import (
	"github.com/gin-gonic/gin"
	idVo "github.com/hum2/backend/internal/domain/user/id"
	"github.com/hum2/backend/internal/interface/controller/user/gen"
	usecase "github.com/hum2/backend/internal/usecase/user"
	"net/http"
)

// Controller is a controller of user
type Controller struct {
	usecase *usecase.Usecase
}

// New is a constructor of Controller
func New(u *usecase.Usecase) gen.ServerInterface {
	return &Controller{
		usecase: u,
	}
}

// GetUser is a handler of GET /user
// TODO: fix
func (c *Controller) GetUser(ctx *gin.Context, id gen.UserIdParameter) {
	users, err := c.usecase.FindAll(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	userOutput := make([]gen.UserResponse, len(users))
	for i, user := range users {
		userOutput[i] = gen.UserResponse{
			Id:       user.ID().UUID().String(),
			Name:     user.Name(),
			Birthday: user.BirthdayString(),
		}
	}
	ctx.JSON(http.StatusOK, gin.H{
		"users": userOutput,
	})
}

// GetUsers is a handler of GET /users
func (c *Controller) GetUsers(ctx *gin.Context) {
	users, err := c.usecase.FindAll(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	userOutput := make([]gen.UserResponse, len(users))
	for i, user := range users {
		userOutput[i] = gen.UserResponse{
			Id:       user.ID().UUID().String(),
			Name:     user.Name(),
			Birthday: user.BirthdayString(),
		}
	}
	ctx.JSON(http.StatusOK, gin.H{
		"users": userOutput,
	})
}

// PostUser is a handler of POST /user
func (c *Controller) PostUser(ctx *gin.Context) {
	var input gen.PostUserRequest
	if err := ctx.Bind(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	user, err := c.usecase.Create(ctx.Request.Context(), input.Name, input.Birthday)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"user": gen.UserResponse{
			Id:       user.ID().UUID().String(),
			Birthday: user.BirthdayString(),
			Name:     user.Name(),
		},
	})
}

// PutUser is a handler of PUT /user/:id
func (c *Controller) PutUser(ctx *gin.Context, id gen.UserIdParameter) {
	var input gen.PutUserRequest
	if err := ctx.Bind(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	user, err := c.usecase.FindByID(ctx, idVo.NewFromInput(id.String()))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	// TODO: update
	err = c.usecase.Update(ctx.Request.Context(), user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, nil)
}

// DeleteUser is a handler of DELETE /user/:id
func (c *Controller) DeleteUser(ctx *gin.Context, id gen.UserIdParameter) {
	err := c.usecase.Delete(ctx.Request.Context(), idVo.NewFromInput(id.String()))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, nil)
}
