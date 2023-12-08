package models

import "zrDispatch/core/utils/define"

func GetUserInfoBYid(id string) (user *define.User, err error) {

	res := db.Table("user").Where("id", id).Take(&user)
	err = res.Error
	return
}
