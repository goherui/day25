package service

import (
	"context"
	__ "day25/proto"
	"day25/service/basic/config"
	"day25/service/model"
	"net/http"
)

// server is used to implement helloworld.GreeterServer.
type Server struct {
	__.UnimplementedStreamGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *Server) CreateUser(_ context.Context, in *__.CreateUserReq) (*__.CreateUserResp, error) {
	var user model.User
	err := user.FindUser(config.DB, in.Username)
	if err == nil {
		return &__.CreateUserResp{
			Code: http.StatusBadRequest,
			Msg:  "用户已存在",
		}, nil
	}
	user = model.User{
		Username: in.Username,
		Email:    in.Email,
		Age:      in.Age,
	}
	err = user.CreateUser(config.DB)
	if err != nil {
		return &__.CreateUserResp{
			Code: http.StatusBadRequest,
			Msg:  "添加失败",
		}, nil
	}
	return &__.CreateUserResp{
		Code: http.StatusOK,
		Msg:  "添加成功",
	}, nil
}
func (s *Server) DelUser(_ context.Context, in *__.DelUserReq) (*__.DelUserResp, error) {
	var user model.User
	err := user.FindUserid(config.DB, in.Id)
	if err != nil {
		return &__.DelUserResp{
			Code: http.StatusBadRequest,
			Msg:  "用户不存在",
		}, nil
	}
	err = user.DelUser(config.DB, in.Id)
	if err != nil {
		return &__.DelUserResp{
			Code: http.StatusBadRequest,
			Msg:  "删除失败",
		}, nil
	}
	return &__.DelUserResp{
		Code: http.StatusOK,
		Msg:  "删除成功",
	}, nil
}
func (s *Server) UpdateUser(_ context.Context, in *__.UpdateUserReq) (*__.UpdateUserResp, error) {
	var user model.User
	err := user.FindUserid(config.DB, in.Id)
	if err != nil {
		return &__.UpdateUserResp{
			Code: http.StatusBadRequest,
			Msg:  "用户不存在",
		}, nil
	}
	user = model.User{
		Username: in.Username,
		Email:    in.Email,
		Age:      in.Age,
	}
	err = user.UpdateUser(config.DB, in.Id)
	if err != nil {
		return &__.UpdateUserResp{
			Code: http.StatusBadRequest,
			Msg:  "修改失败",
		}, nil
	}
	return &__.UpdateUserResp{
		Code: http.StatusOK,
		Msg:  "修改成功",
	}, nil
}
func (s *Server) UserList(_ context.Context, in *__.UserListReq) (*__.UserListResp, error) {
	var user model.User
	users, err := user.GetUserList(config.DB, in.Page, in.Size)
	if err != nil {
		return &__.UserListResp{
			Msg: "获取用户列表失败",
		}, nil
	}
	var userList []*__.User
	for _, u := range users {
		userList = append(userList, &__.User{
			Id:       int64(u.ID),
			Username: u.Username,
			Email:    u.Email,
			Age:      u.Age,
		})
	}
	return &__.UserListResp{
		UserList: userList,
		Msg:      "获取用户列表成功",
	}, nil
}
