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
func (u *ItemInfo) Update(oldu *ItemInfo) error {
	oldu.Price = u.Price
	oldu.UpdatedAt = u.UpdatedAt
	return DB.Self.Save(oldu).Error

	// return DB.Self.Model(&oldu).Updates(
	// 	ItemInfo{
	// 		Price:     u.Price,
	// 		UpdatedAt: u.UpdatedAt,
	// 	}).Error
}

//GetItem will return userinfo by username
func GetItem(itemID string) (*ItemInfo, bool, error) {
	u := &ItemInfo{}
	d := DB.Self.Where("item_id=?", itemID).First(&u)
	return u, d.RecordNotFound(), d.Error
}

//Create is to Create item from db
func (u *TypeInfo) Create() error {
	if GetType(u.Typename, u.URL) {
		return DB.Self.Create(&u).Error
	}
	return nil
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

//Create is to Create CityInfo item from db
func (u *CityInfo) Create() error {
	if GetCity(u.Cityname, u.Shortcut) {
		return DB.Self.Create(&u).Error
	}
	return nil
}

//GetCity will return TypeInfo by username
func GetCity(cityname, shortcut string) bool {
	u := &TypeInfo{}
	return DB.Self.Where("cityname=?", cityname).Where("shortcut=?", shortcut).First(&u).RecordNotFound()
}

//GetType will return TypeInfo by username
func GetType(typename, url string) bool {
	u := &TypeInfo{}
	return DB.Self.Where("typename=?", typename).Where("url=?", url).First(&u).RecordNotFound()
}

//GetAllType will return TypeInfo by username
func GetAllType() (*[]TypeInfo, error) {
	u := &[]TypeInfo{}
	d := DB.Self.Find(u, "is_coped=?", false)
	return u, d.Error
}

//GetAllCity will return TypeInfo by username
func GetAllCity() (*[]CityInfo, error) {
	u := &[]CityInfo{}
	d := DB.Self.Find(u)
	return u, d.Error
}
