package repository

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"sesi4/model"
)

type BookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{
		db: db,
	}
}

func (br *BookRepository) Get() ([]model.ItemBook, error) {
	mapBook := make([]model.ItemBook, 0)

	tx := br.db.Find(&mapBook)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return mapBook, nil
}

func (br *BookRepository) GetID(id uint) (model.ItemBook, error) {
	book := model.ItemBook{}

	tx := br.db.First(&book, id)
	if tx.Error != nil {
		return model.ItemBook{}, tx.Error
	}
	return book, nil
}

func (br *BookRepository) Save(newBook model.ItemBook) (model.ItemBook, error) {
	tx := br.db.Create(&newBook)
	if tx.Error != nil {
		return model.ItemBook{}, tx.Error
	}
	return newBook, nil
	//rows := tx.Row()
	//err = rows.Scan(&newBook.ID, &newBook.Name, &newBook.Genre, &newBook.Author, &newBook.ShelfID)
}

func (br *BookRepository) Update(updapteBook model.ItemBook, id uint) (model.ItemBook, error) {
	tx := br.db.
		Clauses(clause.Returning{
			Columns: []clause.Column{
				{Name: "id"},
			},
		},
		).Where("id= ?", id).Updates(&updapteBook)
	if tx.Error != nil {
		return model.ItemBook{}, tx.Error
	}
	return updapteBook, nil
}

func (br *BookRepository) Delete(deletedBook model.ItemBook, id uint) error {
	tx := br.db.Delete(&deletedBook, id)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
