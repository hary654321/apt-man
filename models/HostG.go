package models

func UpdateHgCreate(id, uid string) {

	db.Table("hostgroup").Where("id = ?", id).Update("createByID", uid)
}
