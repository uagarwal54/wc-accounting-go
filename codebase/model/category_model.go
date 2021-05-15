package model

import (
	"fmt"

	"github.com/beego/beego/client/orm"
)

type (
	ItemCategory struct {
		SrNum        int    `orm:"column(srNum);pk"`
		CategoryId   string `orm:"column(categoryId)"`
		CategoryName string `orm:"column(categoryName); unique"`
	}
	Categories struct {
		CategoryList []ItemCategory
	}
)

func init() {
	orm.RegisterModel(new(ItemCategory))
}

func (categories *Categories) ReadAllCategoryData() (err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable("itemCategory").All(&categories.CategoryList)
	return
}

// category name is a unique value
func (categoryInst *ItemCategory) ReadCategoryByName() (err error) {
	// orm.Debug = true
	o := orm.NewOrm()
	err = o.QueryTable("itemCategory").Filter("categoryName", categoryInst.CategoryName).One(categoryInst)
	return
}

// Fetch category by category id
func (categoryInst *ItemCategory) ReadCategoryByCategoryId() (err error) {
	o := orm.NewOrm()
	err = o.QueryTable("itemCategory").Filter("categoryId", categoryInst.CategoryId).One(&categoryInst)
	return
}

// Fetch the number of items in the `itemCategory` table
func CountCategoryRows() (numOfRows int, err error) {
	o := orm.NewOrm()
	err = o.Raw("select count(*) from itemCategory").QueryRow(&numOfRows)
	return
}

// Insert single record into the item table at once
func (category *ItemCategory) InsertRecordIntoItemCategory() (err error) {
	o := orm.NewOrm()
	_, err = o.Insert(category)
	return
}

// Update Item Record
func (category *ItemCategory) UpdateRecord() (err error) {
	// orm.Debug = true
	var tempCategory ItemCategory
	tempCategory.CategoryName = category.CategoryName
	tempCategory.CategoryId = category.CategoryId
	o := orm.NewOrm()
	if o.Read(&tempCategory, "categoryId") == nil || o.Read(&tempCategory, "categoryName") == nil {
		tempCategory.SrNum = tempCategory.SrNum
		_, err = o.Update(category)
	}
	return
}

func (category *ItemCategory) DeleteRecord() (err error) {
	// orm.Debug = true
	o := orm.NewOrm()
	fmt.Println("Deleting the record with category id: ", category.CategoryId)
	_, err = o.QueryTable("itemCategory").Filter("categoryId", category.CategoryId).Delete()
	return
}
