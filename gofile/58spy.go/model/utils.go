package model

// Create is to create ItemInfo in db
func (u *ItemInfo) Create() error {
	return DB.Self.Create(&u).Error
}

//DeleteItem is to delete item from db
func DeleteItem(itemID string) error {
	item := ItemInfo{}
	item.ItemID = itemID
	return DB.Self.Delete(&item).Error
}

//Update is to update userinfo in db
func (u *ItemInfo) Update() error {
	return DB.Self.Model(&u).Updates(
		ItemInfo{
			Price: u.Price,
		}).Error
	// return DB.Self.Save(u).Error
}

//GetItem will return userinfo by username
func GetItem(itemID string) (*ItemInfo, error) {
	u := &ItemInfo{}
	d := DB.Self.Where("item_id=?", itemID).First(&u)
	return u, d.Error
}

//Create is to Create item from db
func (u *TypeInfo) Create() error {
	if _, err := GetType(u.Typename); err == nil {
		return nil
	}
	return DB.Self.Create(&u).Error
}

//DeleteType is to delete item from db
func DeleteType(id uint64) error {
	typeInfo := TypeInfo{}
	typeInfo.ID = id
	return DB.Self.Delete(&typeInfo).Error
}

//Update is to update userinfo in db
func (u *TypeInfo) Update() error {
	return DB.Self.Model(&u).Updates(
		TypeInfo{
			URL: u.URL,
		}).Error
	// return DB.Self.Save(u).Error
}

//GetType will return TypeInfo by username
func GetType(typename string) (*TypeInfo, error) {
	u := &TypeInfo{}
	d := DB.Self.Where("typename=?", typename).First(&u)
	return u, d.Error
}

//GetAllType will return TypeInfo by username
func GetAllType() (*[]TypeInfo, error) {
	u := &[]TypeInfo{}
	d := DB.Self.Find(u, "is_coped=?", false)
	return u, d.Error
}
