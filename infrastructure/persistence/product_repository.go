package persistence

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" //comment

	"github.com/koutsoumposval/ddd-playground-golang/domain/entity"
)

// ProductRepository implements IProductRepository
type ProductRepository struct {
	Connection *sql.DB
}

// Get returns a specific Product
func (r *ProductRepository) Get(id int64) (*entity.Product, error) {

	row, err := r.queryRow("select id, name, category_id from product where id=?", id)

	if err != nil {
		return nil, err
	}

	p := &entity.Product{}

	err = row.Scan(&p.ID, &p.Name, &p.CategoryID)

	if err != nil {
		return nil, err
	}

	return p, nil
}

// GetAll returns a list of all Products
func (r *ProductRepository) GetAll() ([]*entity.Product, error) {

	rows, err := r.query("select id, name , category_id from product")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	ps := make([]*entity.Product, 0)

	for rows.Next() {
		p := &entity.Product{}
		err = rows.Scan(&p.ID, &p.Name, &p.CategoryID)

		if err != nil {
			return nil, err
		}

		ps = append(ps, p)
	}

	return ps, nil
}

// Save saves Product
func (r *ProductRepository) Save(p *entity.Product) error {

	stmt, err := r.Connection.Prepare("insert into product (name, category_id) values (?, ?)")

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(p.Name, p.CategoryID)

	return err
}

func (r *ProductRepository) query(q string, args ...interface{}) (*sql.Rows, error) {

	stmt, err := r.Connection.Prepare(q)

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	return stmt.Query(args...)
}

func (r *ProductRepository) queryRow(q string, args ...interface{}) (*sql.Row, error) {

	stmt, err := r.Connection.Prepare(q)

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	return stmt.QueryRow(args...), nil
}
