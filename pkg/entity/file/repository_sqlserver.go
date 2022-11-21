package file

import (
	"database/sql"
	"fmt"
	"foro-hotel/internal/models"
	"time"

	"github.com/jmoiron/sqlx"
)

// sqlServer estructura de conexión a la BD de mssql
type sqlserver struct {
	DB   *sqlx.DB
	user *models.User
	TxID string
}

func newFileSqlServerRepository(db *sqlx.DB, user *models.User, txID string) *sqlserver {
	return &sqlserver{
		DB:   db,
		user: user,
		TxID: txID,
	}
}

// Create registra en la BD
func (s *sqlserver) create(m *File) error {
	var id int
	date := time.Now()
	m.UpdatedAt = date
	m.CreatedAt = date
	const sqlInsert = `INSERT INTO entity.files (matricula_user, name, description, path, file_name, type_file, status, is_delete, created_at, updated_at) VALUES (@matricula_user, @name, @description, @path, @file_name, @type_file, @status, @is_delete, @created_at, @updated_at) SELECT ID = convert(bigint, SCOPE_IDENTITY()) `
	stmt, err := s.DB.Prepare(sqlInsert)
	if err != nil {
		return err
	}
	defer stmt.Close()
	err = stmt.QueryRow(
		sql.Named("matricula_user", m.MatriculaUser),
		sql.Named("name", m.Name),
		sql.Named("description", m.Description),
		sql.Named("path", m.Path),
		sql.Named("file_name", m.FileName),
		sql.Named("type_file", m.TypeFile),
		sql.Named("status", m.Status),
		sql.Named("is_delete", m.IsDelete),
		sql.Named("created_at", m.CreatedAt),
		sql.Named("updated_at", m.UpdatedAt),
	).Scan(&id)
	if err != nil {
		return err
	}
	m.ID = int(id)
	return nil
}

// Update actualiza un registro en la BD
func (s *sqlserver) update(m *File) error {
	date := time.Now()
	m.UpdatedAt = date
	const sqlUpdate = `UPDATE entity.files SET matricula_user = :matricula_user, name = :name, description = :description, path = :path, file_name = :file_name, type_file = :type_file, status = :status, is_delete = :is_delete, updated_at = :updated_at WHERE id = :id `
	rs, err := s.DB.NamedExec(sqlUpdate, &m)
	if err != nil {
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("ecatch:108")
	}
	return nil
}

// Delete elimina un registro de la BD
func (s *sqlserver) delete(id int) error {
	const sqlDelete = `DELETE FROM entity.files WHERE id = :id `
	m := File{ID: id}
	rs, err := s.DB.NamedExec(sqlDelete, &m)
	if err != nil {
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("ecatch:108")
	}
	return nil
}

// GetByID consulta un registro por su ID
func (s *sqlserver) getByID(id int) (*File, error) {
	const sqlGetByID = `SELECT convert(nvarchar(50), id) id , matricula_user, name, description, path, file_name, type_file, status, is_delete, created_at, updated_at FROM entity.files  WITH (NOLOCK)  WHERE id = @id `
	mdl := File{}
	err := s.DB.Get(&mdl, sqlGetByID, sql.Named("id", id))
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return &mdl, err
	}
	return &mdl, nil
}

// GetAll consulta todos los registros de la BD
func (s *sqlserver) getAll() ([]*File, error) {
	var ms []*File
	const sqlGetAll = `SELECT convert(nvarchar(50), id) id , matricula_user, name, description, path, file_name, type_file, status, is_delete, created_at, updated_at FROM entity.files  WITH (NOLOCK) `

	err := s.DB.Select(&ms, sqlGetAll)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return ms, err
	}
	return ms, nil
}
