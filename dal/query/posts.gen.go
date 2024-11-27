// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"pastebingo/dal/model"
)

func newPost(db *gorm.DB, opts ...gen.DOOption) post {
	_post := post{}

	_post.postDo.UseDB(db, opts...)
	_post.postDo.UseModel(&model.Post{})

	tableName := _post.postDo.TableName()
	_post.ALL = field.NewAsterisk(tableName)
	_post.ID = field.NewInt32(tableName, "id")
	_post.CreatedAt = field.NewTime(tableName, "created_at")
	_post.UpdatedAt = field.NewTime(tableName, "updated_at")
	_post.DeletedAt = field.NewField(tableName, "deleted_at")
	_post.Title = field.NewString(tableName, "title")
	_post.Content = field.NewString(tableName, "content")
	_post.UserID = field.NewInt32(tableName, "user_id")

	_post.fillFieldMap()

	return _post
}

type post struct {
	postDo

	ALL       field.Asterisk
	ID        field.Int32
	CreatedAt field.Time
	UpdatedAt field.Time
	DeletedAt field.Field
	Title     field.String
	Content   field.String
	UserID    field.Int32

	fieldMap map[string]field.Expr
}

func (p post) Table(newTableName string) *post {
	p.postDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p post) As(alias string) *post {
	p.postDo.DO = *(p.postDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *post) updateTableName(table string) *post {
	p.ALL = field.NewAsterisk(table)
	p.ID = field.NewInt32(table, "id")
	p.CreatedAt = field.NewTime(table, "created_at")
	p.UpdatedAt = field.NewTime(table, "updated_at")
	p.DeletedAt = field.NewField(table, "deleted_at")
	p.Title = field.NewString(table, "title")
	p.Content = field.NewString(table, "content")
	p.UserID = field.NewInt32(table, "user_id")

	p.fillFieldMap()

	return p
}

func (p *post) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *post) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 7)
	p.fieldMap["id"] = p.ID
	p.fieldMap["created_at"] = p.CreatedAt
	p.fieldMap["updated_at"] = p.UpdatedAt
	p.fieldMap["deleted_at"] = p.DeletedAt
	p.fieldMap["title"] = p.Title
	p.fieldMap["content"] = p.Content
	p.fieldMap["user_id"] = p.UserID
}

func (p post) clone(db *gorm.DB) post {
	p.postDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p post) replaceDB(db *gorm.DB) post {
	p.postDo.ReplaceDB(db)
	return p
}

type postDo struct{ gen.DO }

type IPostDo interface {
	gen.SubQuery
	Debug() IPostDo
	WithContext(ctx context.Context) IPostDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IPostDo
	WriteDB() IPostDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IPostDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IPostDo
	Not(conds ...gen.Condition) IPostDo
	Or(conds ...gen.Condition) IPostDo
	Select(conds ...field.Expr) IPostDo
	Where(conds ...gen.Condition) IPostDo
	Order(conds ...field.Expr) IPostDo
	Distinct(cols ...field.Expr) IPostDo
	Omit(cols ...field.Expr) IPostDo
	Join(table schema.Tabler, on ...field.Expr) IPostDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IPostDo
	RightJoin(table schema.Tabler, on ...field.Expr) IPostDo
	Group(cols ...field.Expr) IPostDo
	Having(conds ...gen.Condition) IPostDo
	Limit(limit int) IPostDo
	Offset(offset int) IPostDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IPostDo
	Unscoped() IPostDo
	Create(values ...*model.Post) error
	CreateInBatches(values []*model.Post, batchSize int) error
	Save(values ...*model.Post) error
	First() (*model.Post, error)
	Take() (*model.Post, error)
	Last() (*model.Post, error)
	Find() ([]*model.Post, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Post, err error)
	FindInBatches(result *[]*model.Post, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Post) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IPostDo
	Assign(attrs ...field.AssignExpr) IPostDo
	Joins(fields ...field.RelationField) IPostDo
	Preload(fields ...field.RelationField) IPostDo
	FirstOrInit() (*model.Post, error)
	FirstOrCreate() (*model.Post, error)
	FindByPage(offset int, limit int) (result []*model.Post, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IPostDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (p postDo) Debug() IPostDo {
	return p.withDO(p.DO.Debug())
}

func (p postDo) WithContext(ctx context.Context) IPostDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p postDo) ReadDB() IPostDo {
	return p.Clauses(dbresolver.Read)
}

func (p postDo) WriteDB() IPostDo {
	return p.Clauses(dbresolver.Write)
}

func (p postDo) Session(config *gorm.Session) IPostDo {
	return p.withDO(p.DO.Session(config))
}

func (p postDo) Clauses(conds ...clause.Expression) IPostDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p postDo) Returning(value interface{}, columns ...string) IPostDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p postDo) Not(conds ...gen.Condition) IPostDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p postDo) Or(conds ...gen.Condition) IPostDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p postDo) Select(conds ...field.Expr) IPostDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p postDo) Where(conds ...gen.Condition) IPostDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p postDo) Order(conds ...field.Expr) IPostDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p postDo) Distinct(cols ...field.Expr) IPostDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p postDo) Omit(cols ...field.Expr) IPostDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p postDo) Join(table schema.Tabler, on ...field.Expr) IPostDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p postDo) LeftJoin(table schema.Tabler, on ...field.Expr) IPostDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p postDo) RightJoin(table schema.Tabler, on ...field.Expr) IPostDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p postDo) Group(cols ...field.Expr) IPostDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p postDo) Having(conds ...gen.Condition) IPostDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p postDo) Limit(limit int) IPostDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p postDo) Offset(offset int) IPostDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p postDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IPostDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p postDo) Unscoped() IPostDo {
	return p.withDO(p.DO.Unscoped())
}

func (p postDo) Create(values ...*model.Post) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p postDo) CreateInBatches(values []*model.Post, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p postDo) Save(values ...*model.Post) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p postDo) First() (*model.Post, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Post), nil
	}
}

func (p postDo) Take() (*model.Post, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Post), nil
	}
}

func (p postDo) Last() (*model.Post, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Post), nil
	}
}

func (p postDo) Find() ([]*model.Post, error) {
	result, err := p.DO.Find()
	return result.([]*model.Post), err
}

func (p postDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Post, err error) {
	buf := make([]*model.Post, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p postDo) FindInBatches(result *[]*model.Post, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p postDo) Attrs(attrs ...field.AssignExpr) IPostDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p postDo) Assign(attrs ...field.AssignExpr) IPostDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p postDo) Joins(fields ...field.RelationField) IPostDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p postDo) Preload(fields ...field.RelationField) IPostDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p postDo) FirstOrInit() (*model.Post, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Post), nil
	}
}

func (p postDo) FirstOrCreate() (*model.Post, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Post), nil
	}
}

func (p postDo) FindByPage(offset int, limit int) (result []*model.Post, count int64, err error) {
	result, err = p.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = p.Offset(-1).Limit(-1).Count()
	return
}

func (p postDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p postDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p postDo) Delete(models ...*model.Post) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *postDo) withDO(do gen.Dao) *postDo {
	p.DO = *do.(*gen.DO)
	return p
}
