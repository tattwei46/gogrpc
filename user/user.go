package user

import proto "go-grpc/user/pb"

type User struct {
	ID       string `json:"id"`
	UserID   string `json:"user_id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u User) ToProto() (user *proto.User) {
	return &proto.User{
		Id:       u.ID,
		UserId:   u.UserID,
		Email:    u.Email,
		Password: u.Password,
	}
}

func ToUser(user *proto.User) User {
	return User{
		ID:       user.GetId(),
		UserID:   user.GetUserId(),
		Email:    user.GetEmail(),
		Password: user.GetPassword(),
	}
}
