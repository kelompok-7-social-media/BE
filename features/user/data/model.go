package data

import (
	cd "project/features/komentar/data"
	"project/features/posting/data"
	"project/features/user"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Email    string
	Username string
	Password string
	Posting  []data.Posting `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
	Comment  []cd.Comment   `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
}

func ToCore(data User) user.Core {
	return user.Core{
		ID:       data.ID,
		Name:     data.Name,
		Email:    data.Email,
		Username: data.Username,
		Password: data.Password,
	}
}

func (dataModel *User) ModelsToCore() user.Core {
	return user.Core{
		ID:       dataModel.ID,
		Name:     dataModel.Name,
		Email:    dataModel.Email,
		Username: dataModel.Username,
	}
}
func listModelToCore(dataModel []User) []user.Core {
	var dataCore []user.Core
	for _, v := range dataModel {
		dataCore = append(dataCore, v.ModelsToCore())
	}
	return dataCore
}

func CoreToData(data user.Core) User {
	return User{
		Model:    gorm.Model{ID: data.ID},
		Name:     data.Name,
		Email:    data.Email,
		Username: data.Username,
		Password: data.Password,
	}
}
