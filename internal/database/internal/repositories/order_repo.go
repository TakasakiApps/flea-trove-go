package repositories

import "github.com/TakasakiApps/flea-trove-go/internal/models"

type OrderRepository BaseRepository

func (or *OrderRepository) CreateOrder(order *models.Order) int64 {
	tx := or.db.Table(models.OrderTable).Create(order)

	return tx.RowsAffected
}

func (or *OrderRepository) GetOrderById(orderId uint) (order *models.Order) {
	var orderRes models.Order
	tx := or.db.Table(models.OrderTable).Where("id = ?", orderId).First(&orderRes)
	if tx.RowsAffected != 1 {
		return nil
	}
	order = &orderRes
	return
}

func (or *OrderRepository) GetOrderList() (orderList []models.Order) {
	or.db.Table(models.OrderTable).Find(&orderList)
	return
}

func (or *OrderRepository) GetOrderByUser(user string) (orderList []models.Order) {
	or.db.Table(models.OrderTable).Where("user = ?", user).Find(&orderList)
	return
}

func (or *OrderRepository) UpdateOrder(order *models.Order) int64 {
	tx := or.db.Table(models.OrderTable).
		Where("id = ?", order).
		Select("*").
		Omit("id", "created_at", "update_at", "delete_at", "user", "product_id").
		Updates(order)
	return tx.RowsAffected
}

func (or *OrderRepository) DeleteOrderById(orderId uint) int64 {
	tx := or.db.Table(models.OrderTable).Delete(&models.Order{ID: orderId})

	return tx.RowsAffected
}
