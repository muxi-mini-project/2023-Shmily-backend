package article

import (
	"2023-Shmily-backend/pkg/logger"
	"2023-Shmily-backend/pkg/model"
	"2023-Shmily-backend/pkg/pagination"
	"2023-Shmily-backend/pkg/route"
	"2023-Shmily-backend/pkg/types"
	"net/http"
)

// Get 通过ID获取文章
func Get(idstr string) (Article, error) {
	var article Article
	id := types.StringToUint64(idstr)
	// 增加 Preload("Category")
	if err := model.DB.Debug().Preload("User").Preload("Category").First(&article, id).Error; err != nil {
		return article, err
	}
	return article, nil

}

// GetAll 获取全部文章
func GetAll(r *http.Request, perPage int) ([]Article, pagination.ViewData, error) {

	// 1. 初始化分页实例
	db := model.DB.Model(Article{}).Order("created_at desc")
	_pager := pagination.New(r, db, route.Name2URL("home"), perPage)

	// 2. 获取视图数据
	viewData := _pager.Paging()

	// 3. 获取数据
	var articles []Article
	_pager.Results(&articles)

	return articles, viewData, nil
}

// Create 创建文章，通过 article.ID 来判断是否创建成功
func (article *Article) Create() (err error) {
	if err = model.DB.Create(&article).Error; err != nil {
		logger.LogError(err)
		return err
	}

	return nil
}

// Update 更新文章
func (article *Article) Update() (rowsAffected int64, err error) {
	result := model.DB.Save(&article)
	if err = result.Error; err != nil {
		logger.LogError(err)
		return 0, err
	}

	return result.RowsAffected, nil
}

// Delete 删除文章
func (article *Article) Delete() (rowsAffected int64, err error) {
	result := model.DB.Delete(&article)
	if err = result.Error; err != nil {
		logger.LogError(err)
		return 0, err
	}

	return result.RowsAffected, nil
}

// GetByCategoryID 获取分类相关的文章
func GetByCategoryID(cid string, r *http.Request, perPage int) ([]Article, pagination.ViewData, error) {

	// 1. 初始化分页实例
	db := model.DB.Model(Article{}).Where("category_id = ?", cid).Order("created_at desc")
	_pager := pagination.New(r, db, route.Name2URL("categories.show", "id", cid), perPage)

	// 2. 获取视图数据
	viewData := _pager.Paging()

	// 3. 获取数据
	var articles []Article
	_pager.Results(&articles)

	return articles, viewData, nil
}
