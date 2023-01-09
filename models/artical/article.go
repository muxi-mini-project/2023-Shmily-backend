package article

import (
	"2023-Shmily-backend/models"
	"2023-Shmily-backend/models/category"
	"2023-Shmily-backend/models/user"
	"2023-Shmily-backend/pkg/route"
	"strconv"
)

// Article 文章模型
type Article struct {

	// 加载基类
	models.BaseModel
	Title string `gorm:"type:varchar(255);not null;" valid:"title"`
	Body  string `gorm:"type:longtext;not null;" valid:"body"`
	// 关联用户
	UserID uint64 `gorm:"not null;index"`
	User   user.User
	// 关联分类
	CategoryID uint64 `gorm:"not null;default:3;index"`
	// 新增下面这一行
	Category category.Category
}

// Link 方法用来生成文章链接
func (article Article) Link() string {
	return route.Name2URL("articles.show", "id", strconv.FormatUint(article.ID, 10))
}
