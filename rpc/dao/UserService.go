package dao

import "rpc/init"
import "rpc/dao/domain"

// GetByID 根据ID获取到用户
func GetByID(Id int64) domain.User {
	findUser := &domain.User{ID: Id}
	init.GormDB.Find(findUser)
	// 查询后直接返回
	// TODO 语法有待考究
	return *findUser
}

// Insert 插入一个新用户
func Insert(user domain.User) {
	init.GormDB.Create(user)
	// TODO 调研下有没有更简介的写法
}

// Update 更新User
func Update(user domain.User) {
	update := &domain.User{
		ID:         user.ID,
		Code:       user.Code,
		Name:       user.Name,
		NickName:   user.NickName,
		IdCard:     user.IdCard,
		Mobile:     user.Mobile,
		CreateTime: user.CreateTime,
		ModifyTime: user.ModifyTime,
	}
	init.GormDB.Updates(update)
}

