package model

import (
	"fmt"

	"github.com/beego/beego/client/orm"
)

type (
	Item struct {
		SrNum        int    `orm:"column(srNum);pk"`
		ItemId       string `orm:"column(itemId)"`
		ItemName     string `orm:"column(itemName); unique"`
		ItemCategory int    `orm:"column(itemCategory)"`
	}
	Items struct {
		ItemList []Item
	}
)

func init() {
	orm.RegisterModel(new(Item))
}

func (items *Items) ReadAllItemData() (err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable("Item").All(&items.ItemList)
	return
}

// Fetch all the items in a given category
func (items *Items) ReadAllItemsInACategory(categoryId int) (err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable("item").Filter("itemCategory", categoryId).All(&items)
	return
}

// item name is a unique value
func (ItemInst *Item) ReadItemByName() (err error) {
	// orm.Debug = true
	o := orm.NewOrm()
	err = o.QueryTable("item").Filter("itemName", ItemInst.ItemName).One(ItemInst)
	return
}

// Fetch Item by item id
func (item *Item) ReadItemByItemId() (err error) {
	o := orm.NewOrm()
	err = o.QueryTable("item").Filter("itemId", item.ItemId).One(&item)
	return
}

// Fetch the number of items in the `item` table
func CountItemRows() (numOfRows int, err error) {
	o := orm.NewOrm()
	err = o.Raw("select count(*) from item").QueryRow(&numOfRows)
	return
}

// Insert record(s) into the item table
func (items *Items) InsertIntoItem() (err error) {
	o := orm.NewOrm()
	// The 1st param is the number of records to insert in one bulk statement. The 2nd param is models slice.
	_, err = o.InsertMulti(10, items.ItemList)
	return
}

// Update Item Record
func (item *Item) UpdateRecord() (err error) {
	o := orm.NewOrm()
	if o.Read(item) == nil {
		_, err = o.Update(item)
	}
	return
}

func (item *Item) DeleteRecords() (err error) {
	o := orm.NewOrm()
	var errSlice map[string]string
	fmt.Println("Deleting the record with item id: ", item.ItemId)
	if _, err = o.Delete(item); err != nil {
		errSlice[item.ItemId] = err.Error()
	}

	return
}
