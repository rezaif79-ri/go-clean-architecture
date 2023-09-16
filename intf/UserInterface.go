package intf

import (
	"context"

	"github.com/rezaif79-ri/go-clean-architecture/app/model"
)

type UserService interface {
	PrintUserProfile(ctx context.Context, userProfile *model.User)
	TestUserProfile(ctx context.Context)
}

type UserDataLayer interface {
	GetUserProfileByID(ctx context.Context, userID int)
	TestUserProfileData(context.Context)
}
