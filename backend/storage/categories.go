package storage

// Category descrubes a row from gumsite_categories that is a table in postgesql
type Category struct {
	ID               int    `json:"id"`
	CategoryName     string `json:"category_name"`
	CategoryQuantity int    `json:"category_quantity"`
}

// TestCategory returns a Category that is used
// for testing categories.go
func TestCategory() Category {
	return Category{
		ID:               1,
		CategoryName:     "Milk",
		CategoryQuantity: 1,
	}
}

// InsertCategory inserts a new row into gumsite_categories
func InsertCategory(store Storage, c Category) error {
	if _, err := store.db.Query("INSERT INTO gumsite_categories(category_name, category_quantity)VALUES($1,$2)",
		c.CategoryName,
		c.CategoryQuantity,
	); err != nil {
		return err
	}
	return nil
}

// SelectAllCategories returns an array of Category
// that is equvalent of gumsite_categories
func SelectAllCategories(store Storage) (*[]Category, error) {
	rows, err := store.db.Query("SELECT id, category_name, category_quantity FROM gumsite_categories")
	if err != nil {
		return nil, err
	}
	var id int
	var name string
	var quantity int

	c := []Category{}

	for rows.Next() {
		rows.Scan(&id, &name, &quantity)
		c = append(c, Category{
			ID:               id,
			CategoryName:     name,
			CategoryQuantity: quantity,
		})
	}
	return &c, nil
}

func SelectCategoryByName(store Storage, name string) (*Category, error) {
	rows, err := store.db.Query("SELECT id, category_name, category_quantity FROM gumsite_categories WHERE category_name=$1", name)
	if err != nil {
		return nil, err
	}
	c := Category{}
	for rows.Next() {
		rows.Scan(&c.ID, &c.CategoryName, &c.CategoryQuantity)
	}
	return &c, nil
}

// DeleteCategory deletes a row from gumsite_categories
func DeleteCategory(store Storage, name string) error {
	if _, err := store.db.Query("DELETE FROM gumsite_categories WHERE category_name=$1",
		name,
	); err != nil {
		return err
	}
	return nil
}
