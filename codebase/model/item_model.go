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
func (items *Items) ReadAllItemsInACategory(CategoryId int) (err error) {
	// orm.Debug = true
	o := orm.NewOrm()
	_, err = o.QueryTable("item").Filter("itemCategory", CategoryId).All(&items.ItemList)
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

// Insert multiple records into the item table at once
func (items *Items) InsertRecordsIntoItem() (err error) {
	o := orm.NewOrm()
	// The 1st param is the number of records to insert in one bulk statement. The 2nd param is models slice.
	_, err = o.InsertMulti(10, items.ItemList)
	return
}

// Insert single record into the item table at once
func (item *Item) InsertRecordIntoItem() (err error) {
	o := orm.NewOrm()
	_, err = o.Insert(item)
	return
}

// Update Item Record
func (item *Item) UpdateRecord() (err error) {
	// orm.Debug = true
	var tempItem Item
	tempItem.ItemName = item.ItemName
	tempItem.ItemId = item.ItemId
	o := orm.NewOrm()
	if o.Read(&tempItem, "itemId") == nil || o.Read(&tempItem, "itemName") == nil {
		item.SrNum = tempItem.SrNum
		_, err = o.Update(item)
	}
	return
}

func (item *Item) DeleteRecords() (err error) {
	// orm.Debug = true
	o := orm.NewOrm()
	fmt.Println("Deleting the record with item id: ", item.ItemId)
	_, err = o.QueryTable("item").Filter("itemid", item.ItemId).Delete()
	return
}
