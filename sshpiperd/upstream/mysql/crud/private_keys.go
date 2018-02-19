package crud

import (
	"database/sql"
	"time"

	"github.com/go-sql-driver/mysql"
)

type PrivateKeysRecord struct {
	Id          int64
	Name        string
	Data        string
	Type        string
	GmtCreate   time.Time
	GmtModified time.Time
}

type PrivateKeys struct {
	db *sql.DB
	tx *sql.Tx
}

func NewPrivateKeys(db *sql.DB) *PrivateKeys {
	return &PrivateKeys{
		db: db,
	}
}

// Function to help make the api feel cleaner
func (t *PrivateKeys) Commit() error {
	if t.tx == nil {
		return nil
	}

	err := t.tx.Commit()
	t.tx = nil
	return err
}

func (t *PrivateKeys) Rollback() error {
	if t.tx == nil {
		return nil
	}

	err := t.tx.Rollback()
	t.tx = nil
	return err
}

func (t *PrivateKeys) Post(u *PrivateKeysRecord) (int64, error) {
	var err error
	if t.tx == nil {
		// new transaction
		t.tx, err = t.db.Begin()
		if err != nil {
			return 0, err
		}
	}

	r, err := t.tx.Exec("insert into `private_keys` set `name`=?,`data`=?,`type`=?,  `gmt_modified` = now(), `gmt_create` = now()", u.Name, u.Data, u.Type)
	if err != nil {
		return 0, err
	}

	v, err := r.LastInsertId()
	if err != nil {
		return 0, err
	}

	return v, nil
}

func (t *PrivateKeys) Put(u *PrivateKeysRecord) (int64, error) {
	var err error
	if t.tx == nil {
		// new transaction
		t.tx, err = t.db.Begin()
		if err != nil {
			return 0, err
		}
	}

	r, err := t.tx.Exec("update `private_keys` set `id`=?,`name`=?,`data`=?,`type`=?, `gmt_modified` = now() where `id`=?", u.Id, u.Name, u.Data, u.Type, u.Id)
	if err != nil {
		return 0, err
	}

	v, err := r.RowsAffected()
	if err != nil {
		return 0, err
	}

	return v, nil
}

func (t *PrivateKeys) Delete(u *PrivateKeysRecord) (int64, error) {
	var err error
	if t.tx == nil {
		// new transaction
		t.tx, err = t.db.Begin()
		if err != nil {
			return 0, err
		}
	}

	r, err := t.tx.Exec("delete from private_keys where `id`=?", u.Id)
	if err != nil {
		return 0, err
	}

	v, err := r.RowsAffected()
	if err != nil {
		return 0, err
	}

	return v, nil
}

func (t *PrivateKeys) GetById(Id int64) ([]*PrivateKeysRecord, error) {
	r, err := t.db.Query("select * from private_keys where id=?", Id)
	if err != nil {
		return nil, err
	}
	res := make([]*PrivateKeysRecord, 0)
	for r.Next() {
		var colId sql.NullInt64
		var colName sql.NullString
		var colData sql.NullString
		var colType sql.NullString
		var colGmtCreate mysql.NullTime
		var colGmtModified mysql.NullTime
		err = r.Scan(&colId, &colName, &colData, &colType, &colGmtCreate, &colGmtModified)
		if err != nil {
			return nil, err
		}
		s := &PrivateKeysRecord{Id: colId.Int64, Name: colName.String, Data: colData.String, Type: colType.String, GmtCreate: colGmtCreate.Time, GmtModified: colGmtModified.Time}
		res = append(res, s)
	}
	return res, nil
}

func (t *PrivateKeys) GetByName(Name string) ([]*PrivateKeysRecord, error) {
	r, err := t.db.Query("select * from private_keys where name=?", Name)
	if err != nil {
		return nil, err
	}
	res := make([]*PrivateKeysRecord, 0)
	for r.Next() {
		var colId sql.NullInt64
		var colName sql.NullString
		var colData sql.NullString
		var colType sql.NullString
		var colGmtCreate mysql.NullTime
		var colGmtModified mysql.NullTime
		err = r.Scan(&colId, &colName, &colData, &colType, &colGmtCreate, &colGmtModified)
		if err != nil {
			return nil, err
		}
		s := &PrivateKeysRecord{Id: colId.Int64, Name: colName.String, Data: colData.String, Type: colType.String, GmtCreate: colGmtCreate.Time, GmtModified: colGmtModified.Time}
		res = append(res, s)
	}
	return res, nil
}

func (t *PrivateKeys) GetByData(Data string) ([]*PrivateKeysRecord, error) {
	r, err := t.db.Query("select * from private_keys where data=?", Data)
	if err != nil {
		return nil, err
	}
	res := make([]*PrivateKeysRecord, 0)
	for r.Next() {
		var colId sql.NullInt64
		var colName sql.NullString
		var colData sql.NullString
		var colType sql.NullString
		var colGmtCreate mysql.NullTime
		var colGmtModified mysql.NullTime
		err = r.Scan(&colId, &colName, &colData, &colType, &colGmtCreate, &colGmtModified)
		if err != nil {
			return nil, err
		}
		s := &PrivateKeysRecord{Id: colId.Int64, Name: colName.String, Data: colData.String, Type: colType.String, GmtCreate: colGmtCreate.Time, GmtModified: colGmtModified.Time}
		res = append(res, s)
	}
	return res, nil
}

func (t *PrivateKeys) GetByType(Type string) ([]*PrivateKeysRecord, error) {
	r, err := t.db.Query("select * from private_keys where type=?", Type)
	if err != nil {
		return nil, err
	}
	res := make([]*PrivateKeysRecord, 0)
	for r.Next() {
		var colId sql.NullInt64
		var colName sql.NullString
		var colData sql.NullString
		var colType sql.NullString
		var colGmtCreate mysql.NullTime
		var colGmtModified mysql.NullTime
		err = r.Scan(&colId, &colName, &colData, &colType, &colGmtCreate, &colGmtModified)
		if err != nil {
			return nil, err
		}
		s := &PrivateKeysRecord{Id: colId.Int64, Name: colName.String, Data: colData.String, Type: colType.String, GmtCreate: colGmtCreate.Time, GmtModified: colGmtModified.Time}
		res = append(res, s)
	}
	return res, nil
}

func (t *PrivateKeys) GetByGmtCreate(GmtCreate time.Time) ([]*PrivateKeysRecord, error) {
	r, err := t.db.Query("select * from private_keys where gmt_create=?", GmtCreate)
	if err != nil {
		return nil, err
	}
	res := make([]*PrivateKeysRecord, 0)
	for r.Next() {
		var colId sql.NullInt64
		var colName sql.NullString
		var colData sql.NullString
		var colType sql.NullString
		var colGmtCreate mysql.NullTime
		var colGmtModified mysql.NullTime
		err = r.Scan(&colId, &colName, &colData, &colType, &colGmtCreate, &colGmtModified)
		if err != nil {
			return nil, err
		}
		s := &PrivateKeysRecord{Id: colId.Int64, Name: colName.String, Data: colData.String, Type: colType.String, GmtCreate: colGmtCreate.Time, GmtModified: colGmtModified.Time}
		res = append(res, s)
	}
	return res, nil
}

func (t *PrivateKeys) GetByGmtModified(GmtModified time.Time) ([]*PrivateKeysRecord, error) {
	r, err := t.db.Query("select * from private_keys where gmt_modified=?", GmtModified)
	if err != nil {
		return nil, err
	}
	res := make([]*PrivateKeysRecord, 0)
	for r.Next() {
		var colId sql.NullInt64
		var colName sql.NullString
		var colData sql.NullString
		var colType sql.NullString
		var colGmtCreate mysql.NullTime
		var colGmtModified mysql.NullTime
		err = r.Scan(&colId, &colName, &colData, &colType, &colGmtCreate, &colGmtModified)
		if err != nil {
			return nil, err
		}
		s := &PrivateKeysRecord{Id: colId.Int64, Name: colName.String, Data: colData.String, Type: colType.String, GmtCreate: colGmtCreate.Time, GmtModified: colGmtModified.Time}
		res = append(res, s)
	}
	return res, nil
}

func (t *PrivateKeys) GetFirstById(Id int64) (*PrivateKeysRecord, error) {
	r, err := t.GetById(Id)
	if err != nil {
		return nil, err
	}

	if len(r) > 0 {
		return r[0], nil
	}

	return nil, nil
}

func (t *PrivateKeys) GetFirstByName(Name string) (*PrivateKeysRecord, error) {
	r, err := t.GetByName(Name)
	if err != nil {
		return nil, err
	}

	if len(r) > 0 {
		return r[0], nil
	}

	return nil, nil
}

func (t *PrivateKeys) GetFirstByData(Data string) (*PrivateKeysRecord, error) {
	r, err := t.GetByData(Data)
	if err != nil {
		return nil, err
	}

	if len(r) > 0 {
		return r[0], nil
	}

	return nil, nil
}

func (t *PrivateKeys) GetFirstByType(Type string) (*PrivateKeysRecord, error) {
	r, err := t.GetByType(Type)
	if err != nil {
		return nil, err
	}

	if len(r) > 0 {
		return r[0], nil
	}

	return nil, nil
}

func (t *PrivateKeys) GetFirstByGmtCreate(GmtCreate time.Time) (*PrivateKeysRecord, error) {
	r, err := t.GetByGmtCreate(GmtCreate)
	if err != nil {
		return nil, err
	}

	if len(r) > 0 {
		return r[0], nil
	}

	return nil, nil
}

func (t *PrivateKeys) GetFirstByGmtModified(GmtModified time.Time) (*PrivateKeysRecord, error) {
	r, err := t.GetByGmtModified(GmtModified)
	if err != nil {
		return nil, err
	}

	if len(r) > 0 {
		return r[0], nil
	}

	return nil, nil
}
