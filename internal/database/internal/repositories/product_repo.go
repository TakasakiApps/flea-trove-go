package repositories

import (
	"github.com/TakasakiApps/flea-trove-go/internal/models"
)

type ProductRepository BaseRepository

func (pr *ProductRepository) GetProductList() (productList []models.Product) {
	pr.db.Table(models.ProductTable).Find(&productList)
	return
}

func (pr *ProductRepository) GetProductListByUser(account string) (productList []models.Product) {
	pr.db.Table(models.ProductTable).Where("user = ?", account).Find(&productList)
	return
}

func (pr *ProductRepository) GetProductListByKeyword(keyword string) (productList []models.Product) {
	pr.db.Table(models.ProductTable).Where("name like ?", "%"+keyword+"%").Find(&productList)
	return
}

func (pr *ProductRepository) CreateProduct(product *models.Product) int64 {
	tx := pr.db.Table(models.ProductTable).Create(product)
	return tx.RowsAffected
}

func (pr *ProductRepository) UpdateProduct(product *models.Product) int64 {
	tx := pr.db.Table(models.ProductTable).
		Where("id = ?", product.ID).
		// 选择所有字段更新
		Select("*").
		// 忽略的字段
		Omit("id", "created_at", "update_at", "delete_at", "image_id", "user").
		Updates(product)
	return tx.RowsAffected
}

func (pr *ProductRepository) GetProductById(id uint) (product *models.Product) {
	var productRes models.Product
	tx := pr.db.Table(models.ProductTable).Where("id = ?", id).First(&productRes)
	if tx.RowsAffected != 1 {
		return nil
	}
	product = &productRes
	return
}
