package repositories

// var (
// 	ps []models.Products = []models.Products{
// 		{ID: 1, Name: "Product 1", Price: 1000},
// 		{ID: 2, Name: "Product 2", Price: 2000},
// 		{ID: 3, Name: "Product 3", Price: 3000},
// 	}
// 	pm sync.Mutex
// )

// var db = postgres.NewPostgresDB()

// func TestGetProducts(t *testing.T) {
// 	pr := NewProductsRepo(db)
// 	products, err := pr.GetProducts()
// 	log.Println(&products)
// 	assert.Nil(t, err)
// 	assert.NotNil(t, products)
// }

// func TestGetProduct(t *testing.T) {
// 	pr := NewProductsRepo(nil)
// 	product := pr.GetProduct(1)
// 	log.Println(product)
// 	assert.Equal(t, 1, product.ID)
// }

// func TestAddProduct(t *testing.T) {
// 	pr := NewProductsRepo(nil)
// 	err := pr.AddProduct(&models.Products{
// 		ID:    uint64(len(products) + 1),
// 		Name:  fmt.Sprintf("Product %d", len(products)+1),
// 		Price: 5000,
// 	})
// 	log.Println(err)
// 	assert.Nil(t, err)
// }

// func TestUpdateProduct(t *testing.T) {
// 	var id uint64 = 1
// 	pr := NewProductsRepo(nil)
// 	err := pr.UpdateProduct(id, &models.Products{
// 		Name:  fmt.Sprintf("Product %d updated", id),
// 		Price: 5000,
// 	})
// 	log.Println(pr.GetProduct(id))
// 	log.Println(err)
// 	assert.Equal(t, fmt.Sprintf("Product %d updated", id), pr.GetProduct(id).Name)
// }

// func TestDeleteProduct(t *testing.T) {
// 	var id uint64 = 1
// 	pr := NewProductsRepo(nil)
// 	fmt.Println(pr.GetProduct(id))
// 	err := pr.DeleteProduct(id)
// 	log.Println(err)
// 	assert.Nil(t, pr.GetProduct(id))
// }
