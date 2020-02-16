package persistence

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" //comment

	"github.com/koutsoumposval/polyglot-ddd-product/domain/entity"
	"github.com/koutsoumposval/polyglot-ddd-product/domain/repository"
	"github.com/koutsoumposval/polyglot-ddd-product/domain/value"
)

// productRepository implements repository.ProductRepository
type productRepository struct {
	conn *sql.DB
}

// ProductRepository returns initialized ProductRepositoryImp
func ProductRepository(conn *sql.DB) repository.ProductRepository {
	return &productRepository{conn: conn}
}

// Get returns entity.Product
func (r *productRepository) Get(id value.ProductID) (*entity.Product, error) {

	row, err := r.queryRow("select id, name, category_id from product where id=?", id.ID.ID())

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

// GetAll returns list of entity.Product
func (r *productRepository) GetAll() ([]*entity.Product, error) {

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

// Save saves domain.Product to storage
func (r *productRepository) Save(p *entity.Product) error {

	stmt, err := r.conn.Prepare("insert into product (name, category_id) values (?, ?)")

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(p.Name, p.CategoryID)

	return err
}

func (r *productRepository) query(q string, args ...interface{}) (*sql.Rows, error) {

	stmt, err := r.conn.Prepare(q)

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	return stmt.Query(args...)
}

func (r *productRepository) queryRow(q string, args ...interface{}) (*sql.Row, error) {

	stmt, err := r.conn.Prepare(q)

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	return stmt.QueryRow(args...), nil
}
