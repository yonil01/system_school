package users

import (
	"database/sql"
	"errors"
	"foro-hotel/internal/logger"
	"foro-hotel/internal/models"
	"github.com/jmoiron/sqlx"
)

type psql struct {
	DB   *sqlx.DB
	user *models.User
	TxID string
}

func (s *psql) create(m *User) (*User, error) {
	const sqlQueryCreateUser = `INSERT INTO dbo.users(number_document, type_document, first_name, second_name, first_lastname, second_lastname, email, password) VALUES(:number_document, :type_document,:first_name,:second_name,:first_lastname,:second_lastname,:email,:password)`
	_, err := s.DB.NamedExec(sqlQueryCreateUser, &m)
	if err != nil {
		logger.Error.Println(s.TxID, "-cound't insert User: %V", err)
		return nil, err
	}

	return m, nil
}

func (s *psql) getDoctypeByUser(userId int) ([]*models.DoctypeUser, error) {
	var doctypeUser []*models.DoctypeUser
	query := `select doc.id, doc.name, doc.description, doc.url, doc.type, doc_us.user_id, doc.created_at, doc.updated_at,  doc.first_value, doc.second_value, doc.execution, doc.execution_garps from dbo.doctypes doc join dbo.user_doctype doc_us on doc.id = doc_us.doctype_id where doc_us.user_id = @user_id`
	err := s.DB.Select(&doctypeUser, query, sql.Named("user_id", userId))
	if err != nil {
		return doctypeUser, errors.New("El usuario ya existe!")
	}

	return doctypeUser, nil
}

func (s *psql) getUserByEmail(email string) (*User, int, error) {
	mdl := User{}
	query := `SELECT number_document, type_document, first_name, second_name, first_lastname, second_lastname, email, password FROM dbo.users WHERE email = @email`
	err := s.DB.Get(&mdl, query, sql.Named("email", email))
	if err != nil {
		return nil, 29, errors.New("El usuario ya existe!")
	}

	return &mdl, 1, nil
}

func (s *psql) createLogin(m *User) (int, error) {
	const sqlQueryCreateUser = `INSERT INTO dbo.login(number_document, type_document, first_name, second_name, first_lastname, second_lastname, email, real_ip) VALUES(:number_document, :type_document,:first_name,:second_name,:first_lastname,:second_lastname,:email,:real_ip)`
	_, err := s.DB.NamedExec(sqlQueryCreateUser, &m)
	if err != nil {
		logger.Error.Println(s.TxID, "-cound't insert User: %V", err)
		return 29, errors.New("Error en insertar el usuario!")
	}

	return 82, nil
}

func (s *psql) saveImagePerfil(document_id string, images *models.Image) (int, error) {

	data := changeValue{
		NumberDocument: document_id,
		B64:            images.B64,
		Name:           images.Name,
		Type:           images.Type,
		Url:            images.Url,
	}

	const sqlUpdate = `UPDATE dbo.users SET url = :url, type_url = :type  WHERE number_document = :number_document `
	_, err := s.DB.NamedExec(sqlUpdate, &data)
	if err != nil {
		logger.Error.Println(s.TxID, "-cound't insert User: %V", err)
		return 29, errors.New("Error en insertar el negocio!")
	}

	return 29, nil
}

func newUserPsqlRepository(db *sqlx.DB, user *models.User, txID string) *psql {
	return &psql{
		DB:   db,
		user: user,
		TxID: txID,
	}
}

type changeValue struct {
	NumberDocument string `json:"number_document" db:"number_document" valid:"required"`
	B64            string `json:"b64" db:"b64" valid:"required"`
	Name           string `json:"name" db:"name" valid:"required"`
	Type           string `json:"type" db:"type" valid:"_"`
	Url            string `json:"url" db:"url" valid:"_"`
}
