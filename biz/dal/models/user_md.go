package models

import (
	"github.com/bachhieu/test/pkg/constant"
	"time"
)

type UserMD struct {
	ID           string                `db:"id" json:"id"`
	Username     string                `db:"username" json:"username"`
	Email        string                `db:"email" json:"email"`
	PasswordHash string                `db:"password_hash" json:"-"`
	FullName     string                `db:"full_name" json:"full_name"`
	Role         string                `db:"role" json:"role"`
	Provider     constant.ProviderUser `db:"provider" json:"provider"` // Auth by login/fb/google
	CreatedTime  int32                 `db:"created_time" json:"created_time"`
	UpdatedTime  int32                 `db:"updated_time" json:"updated_time"`
}

func (current *UserMD) AllowedUpdateFields(req *UserMD) *UserMD {
	return &UserMD{
		ID:           current.ID,
		Username:     req.Username,
		Email:        req.Email,
		PasswordHash: req.PasswordHash,
		FullName:     req.FullName,
		Role:         req.Role,
		Provider:     req.Provider,
		CreatedTime:  current.CreatedTime,
		UpdatedTime:  int32(time.Now().Unix()),
	}
}
