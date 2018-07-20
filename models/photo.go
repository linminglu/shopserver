package models

import (
	"fmt"

	"github.com/zyx/shop_server/libs"
)

type Photo struct {
	Model
}
type PhotoData struct {
	Name   string `empty:"图片名不能为空"`
	Path   string `empty:"图片地址不能为空"`
	Key    string `empty:"key不能为空"`
	Album  string `empty:"相册不能为空"`
	width  int
	height int
}

func (self *Photo) InitSqlField(sql libs.SqlType) libs.SqlType {
	return self.InitField(self.InitJoinString(sql, true))
}
func (self *Photo) GetModelStruct() interface{} {
	return PhotoData{}
}
func (self *Photo) InitJoinString(sql libs.SqlType, allfield bool) libs.SqlType {
	albumTableName := GetModel(ALBUM).TableName()
	userTableName := GetModel(USER).TableName()

	fieldstr := ""
	if (allfield == true) || (sql.NeedJointable("album") == true) {

		fieldstr += fmt.Sprintf("left join `%s` `album` ON `photo`.`album`=`album`.`id`", albumTableName)
	}
	if (allfield == true) || (sql.NeedJointable("user") == true) {

		fieldstr += fmt.Sprintf("left join `%s` `user` ON `photo`.`upload_user`=`user`.`id`", userTableName)
	}
	return sql.Alias("photo").Join(fieldstr)
}
func (self *Photo) InitField(sql libs.SqlType) libs.SqlType {
	return sql.Field(map[string]string{
		"album.name":        "album_name",
		"user.name":         "upload_user_name",
		"photo.id":          "id",
		"photo.name":        "name",
		"photo.path":        "path",
		"photo.upload_time": "upload_time",
		"photo.upload_user": "upload_user",
		"photo.album":       "album",
		"photo.width":       "width",
		"photo.height":      "height",
		"photo.key":         "key",
		"photo.order_id":    "order_id",
	})
}
