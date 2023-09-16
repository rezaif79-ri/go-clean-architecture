package data

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/rezaif79-ri/go-clean-architecture/intf"
)

type UserDataLayer struct {
	DBConn *sql.DB
}

func NewUserDataLayer(conn *sql.DB) intf.UserDataLayer {
	return &UserDataLayer{
		DBConn: conn,
	}
}

// GetUserProfileByID implements intf.UserDataLayer.
func (*UserDataLayer) GetUserProfileByID(ctx context.Context, userID int) {
	fmt.Println("get user profile data by id: ", userID)

}

// TestUserProfileData implements intf.UserDataLayer.
func (*UserDataLayer) TestUserProfileData(context.Context) {
	fmt.Println("test user profile data")
}
