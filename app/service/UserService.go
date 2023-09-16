package service

import (
	"context"
	"fmt"

	"github.com/rezaif79-ri/go-clean-architecture/app/model"
	"github.com/rezaif79-ri/go-clean-architecture/intf"
)

type UserService struct {
	userDL intf.UserDataLayer
}

func NewUserService(userDataLayer intf.UserDataLayer) intf.UserService {
	return &UserService{
		userDL: userDataLayer,
	}
}

// TestUserProfile implements intf.UserService.
func (us *UserService) TestUserProfile(ctx context.Context) {
	fmt.Println("test user profile service")
	us.userDL.TestUserProfileData(ctx)
}

// PrintUserProfile implements intf.UserService.
func (us *UserService) PrintUserProfile(ctx context.Context, user *model.User) {
	fmt.Println("print user profile service")
	us.userDL.GetUserProfileByID(ctx, user.UserId)

}
