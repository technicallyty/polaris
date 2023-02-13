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

	"github.com/berachain/stargazer/wasp/models"
)

func newEthBlockModel(db *gorm.DB, opts ...gen.DOOption) ethBlockModel {
	_ethBlockModel := ethBlockModel{}

	_ethBlockModel.ethBlockModelDo.UseDB(db, opts...)
	_ethBlockModel.ethBlockModelDo.UseModel(&models.EthBlockModel{})

	tableName := _ethBlockModel.ethBlockModelDo.TableName()
	_ethBlockModel.ALL = field.NewAsterisk(tableName)
	_ethBlockModel.ID = field.NewInt64(tableName, "id")
	_ethBlockModel.ParentHash = field.NewBytes(tableName, "parent_hash")
	_ethBlockModel.UncleHash = field.NewBytes(tableName, "uncle_hash")
	_ethBlockModel.Coinbase = field.NewBytes(tableName, "coinbase")
	_ethBlockModel.Root = field.NewBytes(tableName, "root")
	_ethBlockModel.TxHash = field.NewBytes(tableName, "tx_hash")
	_ethBlockModel.ReceiptHash = field.NewBytes(tableName, "receipt_hash")
	_ethBlockModel.Bloom = field.NewBytes(tableName, "bloom")
	_ethBlockModel.Difficulty = field.NewString(tableName, "difficulty")
	_ethBlockModel.BlockNumber = field.NewString(tableName, "block_number")
	_ethBlockModel.GasLimit = field.NewString(tableName, "gas_limit")
	_ethBlockModel.GasUsed = field.NewString(tableName, "gas_used")
	_ethBlockModel.Time = field.NewString(tableName, "time")
	_ethBlockModel.Extra = field.NewBytes(tableName, "extra")
	_ethBlockModel.MixDigest = field.NewBytes(tableName, "mix_digest")
	_ethBlockModel.Nonce = field.NewUint64(tableName, "nonce")
	_ethBlockModel.BaseFee = field.NewBytes(tableName, "base_fee")
	_ethBlockModel.Hash = field.NewBytes(tableName, "hash")
	_ethBlockModel.Size = field.NewString(tableName, "size")
	_ethBlockModel.ReceivedAt = field.NewTime(tableName, "received_at")
	_ethBlockModel.Txs = ethBlockModelHasManyTxs{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Txs", "models.TransactionModel"),
		Receipt: struct {
			field.RelationField
			Logs struct {
				field.RelationField
			}
		}{
			RelationField: field.NewRelation("Txs.Receipt", "models.EthTxnReceipt"),
			Logs: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("Txs.Receipt.Logs", "models.EthLog"),
			},
		},
	}

	_ethBlockModel.fillFieldMap()

	return _ethBlockModel
}

type ethBlockModel struct {
	ethBlockModelDo ethBlockModelDo

	ALL         field.Asterisk
	ID          field.Int64
	ParentHash  field.Bytes
	UncleHash   field.Bytes
	Coinbase    field.Bytes
	Root        field.Bytes
	TxHash      field.Bytes
	ReceiptHash field.Bytes
	Bloom       field.Bytes
	Difficulty  field.String
	BlockNumber field.String
	GasLimit    field.String
	GasUsed     field.String
	Time        field.String
	Extra       field.Bytes
	MixDigest   field.Bytes
	Nonce       field.Uint64
	BaseFee     field.Bytes
	Hash        field.Bytes
	Size        field.String
	ReceivedAt  field.Time
	Txs         ethBlockModelHasManyTxs

	fieldMap map[string]field.Expr
}

func (e ethBlockModel) Table(newTableName string) *ethBlockModel {
	e.ethBlockModelDo.UseTable(newTableName)
	return e.updateTableName(newTableName)
}

func (e ethBlockModel) As(alias string) *ethBlockModel {
	e.ethBlockModelDo.DO = *(e.ethBlockModelDo.As(alias).(*gen.DO))
	return e.updateTableName(alias)
}

func (e *ethBlockModel) updateTableName(table string) *ethBlockModel {
	e.ALL = field.NewAsterisk(table)
	e.ID = field.NewInt64(table, "id")
	e.ParentHash = field.NewBytes(table, "parent_hash")
	e.UncleHash = field.NewBytes(table, "uncle_hash")
	e.Coinbase = field.NewBytes(table, "coinbase")
	e.Root = field.NewBytes(table, "root")
	e.TxHash = field.NewBytes(table, "tx_hash")
	e.ReceiptHash = field.NewBytes(table, "receipt_hash")
	e.Bloom = field.NewBytes(table, "bloom")
	e.Difficulty = field.NewString(table, "difficulty")
	e.BlockNumber = field.NewString(table, "block_number")
	e.GasLimit = field.NewString(table, "gas_limit")
	e.GasUsed = field.NewString(table, "gas_used")
	e.Time = field.NewString(table, "time")
	e.Extra = field.NewBytes(table, "extra")
	e.MixDigest = field.NewBytes(table, "mix_digest")
	e.Nonce = field.NewUint64(table, "nonce")
	e.BaseFee = field.NewBytes(table, "base_fee")
	e.Hash = field.NewBytes(table, "hash")
	e.Size = field.NewString(table, "size")
	e.ReceivedAt = field.NewTime(table, "received_at")

	e.fillFieldMap()

	return e
}

func (e *ethBlockModel) WithContext(ctx context.Context) IEthBlockModelDo {
	return e.ethBlockModelDo.WithContext(ctx)
}

func (e ethBlockModel) TableName() string { return e.ethBlockModelDo.TableName() }

func (e ethBlockModel) Alias() string { return e.ethBlockModelDo.Alias() }

func (e *ethBlockModel) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := e.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (e *ethBlockModel) fillFieldMap() {
	e.fieldMap = make(map[string]field.Expr, 21)
	e.fieldMap["id"] = e.ID
	e.fieldMap["parent_hash"] = e.ParentHash
	e.fieldMap["uncle_hash"] = e.UncleHash
	e.fieldMap["coinbase"] = e.Coinbase
	e.fieldMap["root"] = e.Root
	e.fieldMap["tx_hash"] = e.TxHash
	e.fieldMap["receipt_hash"] = e.ReceiptHash
	e.fieldMap["bloom"] = e.Bloom
	e.fieldMap["difficulty"] = e.Difficulty
	e.fieldMap["block_number"] = e.BlockNumber
	e.fieldMap["gas_limit"] = e.GasLimit
	e.fieldMap["gas_used"] = e.GasUsed
	e.fieldMap["time"] = e.Time
	e.fieldMap["extra"] = e.Extra
	e.fieldMap["mix_digest"] = e.MixDigest
	e.fieldMap["nonce"] = e.Nonce
	e.fieldMap["base_fee"] = e.BaseFee
	e.fieldMap["hash"] = e.Hash
	e.fieldMap["size"] = e.Size
	e.fieldMap["received_at"] = e.ReceivedAt

}

func (e ethBlockModel) clone(db *gorm.DB) ethBlockModel {
	e.ethBlockModelDo.ReplaceConnPool(db.Statement.ConnPool)
	return e
}

func (e ethBlockModel) replaceDB(db *gorm.DB) ethBlockModel {
	e.ethBlockModelDo.ReplaceDB(db)
	return e
}

type ethBlockModelHasManyTxs struct {
	db *gorm.DB

	field.RelationField

	Receipt struct {
		field.RelationField
		Logs struct {
			field.RelationField
		}
	}
}

func (a ethBlockModelHasManyTxs) Where(conds ...field.Expr) *ethBlockModelHasManyTxs {
	if len(conds) == 0 {
		return &a
	}

	exprs := make([]clause.Expression, 0, len(conds))
	for _, cond := range conds {
		exprs = append(exprs, cond.BeCond().(clause.Expression))
	}
	a.db = a.db.Clauses(clause.Where{Exprs: exprs})
	return &a
}

func (a ethBlockModelHasManyTxs) WithContext(ctx context.Context) *ethBlockModelHasManyTxs {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a ethBlockModelHasManyTxs) Model(m *models.EthBlockModel) *ethBlockModelHasManyTxsTx {
	return &ethBlockModelHasManyTxsTx{a.db.Model(m).Association(a.Name())}
}

type ethBlockModelHasManyTxsTx struct{ tx *gorm.Association }

func (a ethBlockModelHasManyTxsTx) Find() (result []*models.TransactionModel, err error) {
	return result, a.tx.Find(&result)
}

func (a ethBlockModelHasManyTxsTx) Append(values ...*models.TransactionModel) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a ethBlockModelHasManyTxsTx) Replace(values ...*models.TransactionModel) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a ethBlockModelHasManyTxsTx) Delete(values ...*models.TransactionModel) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a ethBlockModelHasManyTxsTx) Clear() error {
	return a.tx.Clear()
}

func (a ethBlockModelHasManyTxsTx) Count() int64 {
	return a.tx.Count()
}

type ethBlockModelDo struct{ gen.DO }

type IEthBlockModelDo interface {
	gen.SubQuery
	Debug() IEthBlockModelDo
	WithContext(ctx context.Context) IEthBlockModelDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IEthBlockModelDo
	WriteDB() IEthBlockModelDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IEthBlockModelDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IEthBlockModelDo
	Not(conds ...gen.Condition) IEthBlockModelDo
	Or(conds ...gen.Condition) IEthBlockModelDo
	Select(conds ...field.Expr) IEthBlockModelDo
	Where(conds ...gen.Condition) IEthBlockModelDo
	Order(conds ...field.Expr) IEthBlockModelDo
	Distinct(cols ...field.Expr) IEthBlockModelDo
	Omit(cols ...field.Expr) IEthBlockModelDo
	Join(table schema.Tabler, on ...field.Expr) IEthBlockModelDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IEthBlockModelDo
	RightJoin(table schema.Tabler, on ...field.Expr) IEthBlockModelDo
	Group(cols ...field.Expr) IEthBlockModelDo
	Having(conds ...gen.Condition) IEthBlockModelDo
	Limit(limit int) IEthBlockModelDo
	Offset(offset int) IEthBlockModelDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IEthBlockModelDo
	Unscoped() IEthBlockModelDo
	Create(values ...*models.EthBlockModel) error
	CreateInBatches(values []*models.EthBlockModel, batchSize int) error
	Save(values ...*models.EthBlockModel) error
	First() (*models.EthBlockModel, error)
	Take() (*models.EthBlockModel, error)
	Last() (*models.EthBlockModel, error)
	Find() ([]*models.EthBlockModel, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.EthBlockModel, err error)
	FindInBatches(result *[]*models.EthBlockModel, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*models.EthBlockModel) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IEthBlockModelDo
	Assign(attrs ...field.AssignExpr) IEthBlockModelDo
	Joins(fields ...field.RelationField) IEthBlockModelDo
	Preload(fields ...field.RelationField) IEthBlockModelDo
	FirstOrInit() (*models.EthBlockModel, error)
	FirstOrCreate() (*models.EthBlockModel, error)
	FindByPage(offset int, limit int) (result []*models.EthBlockModel, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IEthBlockModelDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (e ethBlockModelDo) Debug() IEthBlockModelDo {
	return e.withDO(e.DO.Debug())
}

func (e ethBlockModelDo) WithContext(ctx context.Context) IEthBlockModelDo {
	return e.withDO(e.DO.WithContext(ctx))
}

func (e ethBlockModelDo) ReadDB() IEthBlockModelDo {
	return e.Clauses(dbresolver.Read)
}

func (e ethBlockModelDo) WriteDB() IEthBlockModelDo {
	return e.Clauses(dbresolver.Write)
}

func (e ethBlockModelDo) Session(config *gorm.Session) IEthBlockModelDo {
	return e.withDO(e.DO.Session(config))
}

func (e ethBlockModelDo) Clauses(conds ...clause.Expression) IEthBlockModelDo {
	return e.withDO(e.DO.Clauses(conds...))
}

func (e ethBlockModelDo) Returning(value interface{}, columns ...string) IEthBlockModelDo {
	return e.withDO(e.DO.Returning(value, columns...))
}

func (e ethBlockModelDo) Not(conds ...gen.Condition) IEthBlockModelDo {
	return e.withDO(e.DO.Not(conds...))
}

func (e ethBlockModelDo) Or(conds ...gen.Condition) IEthBlockModelDo {
	return e.withDO(e.DO.Or(conds...))
}

func (e ethBlockModelDo) Select(conds ...field.Expr) IEthBlockModelDo {
	return e.withDO(e.DO.Select(conds...))
}

func (e ethBlockModelDo) Where(conds ...gen.Condition) IEthBlockModelDo {
	return e.withDO(e.DO.Where(conds...))
}

func (e ethBlockModelDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IEthBlockModelDo {
	return e.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (e ethBlockModelDo) Order(conds ...field.Expr) IEthBlockModelDo {
	return e.withDO(e.DO.Order(conds...))
}

func (e ethBlockModelDo) Distinct(cols ...field.Expr) IEthBlockModelDo {
	return e.withDO(e.DO.Distinct(cols...))
}

func (e ethBlockModelDo) Omit(cols ...field.Expr) IEthBlockModelDo {
	return e.withDO(e.DO.Omit(cols...))
}

func (e ethBlockModelDo) Join(table schema.Tabler, on ...field.Expr) IEthBlockModelDo {
	return e.withDO(e.DO.Join(table, on...))
}

func (e ethBlockModelDo) LeftJoin(table schema.Tabler, on ...field.Expr) IEthBlockModelDo {
	return e.withDO(e.DO.LeftJoin(table, on...))
}

func (e ethBlockModelDo) RightJoin(table schema.Tabler, on ...field.Expr) IEthBlockModelDo {
	return e.withDO(e.DO.RightJoin(table, on...))
}

func (e ethBlockModelDo) Group(cols ...field.Expr) IEthBlockModelDo {
	return e.withDO(e.DO.Group(cols...))
}

func (e ethBlockModelDo) Having(conds ...gen.Condition) IEthBlockModelDo {
	return e.withDO(e.DO.Having(conds...))
}

func (e ethBlockModelDo) Limit(limit int) IEthBlockModelDo {
	return e.withDO(e.DO.Limit(limit))
}

func (e ethBlockModelDo) Offset(offset int) IEthBlockModelDo {
	return e.withDO(e.DO.Offset(offset))
}

func (e ethBlockModelDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IEthBlockModelDo {
	return e.withDO(e.DO.Scopes(funcs...))
}

func (e ethBlockModelDo) Unscoped() IEthBlockModelDo {
	return e.withDO(e.DO.Unscoped())
}

func (e ethBlockModelDo) Create(values ...*models.EthBlockModel) error {
	if len(values) == 0 {
		return nil
	}
	return e.DO.Create(values)
}

func (e ethBlockModelDo) CreateInBatches(values []*models.EthBlockModel, batchSize int) error {
	return e.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (e ethBlockModelDo) Save(values ...*models.EthBlockModel) error {
	if len(values) == 0 {
		return nil
	}
	return e.DO.Save(values)
}

func (e ethBlockModelDo) First() (*models.EthBlockModel, error) {
	if result, err := e.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*models.EthBlockModel), nil
	}
}

func (e ethBlockModelDo) Take() (*models.EthBlockModel, error) {
	if result, err := e.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*models.EthBlockModel), nil
	}
}

func (e ethBlockModelDo) Last() (*models.EthBlockModel, error) {
	if result, err := e.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*models.EthBlockModel), nil
	}
}

func (e ethBlockModelDo) Find() ([]*models.EthBlockModel, error) {
	result, err := e.DO.Find()
	return result.([]*models.EthBlockModel), err
}

func (e ethBlockModelDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.EthBlockModel, err error) {
	buf := make([]*models.EthBlockModel, 0, batchSize)
	err = e.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (e ethBlockModelDo) FindInBatches(result *[]*models.EthBlockModel, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return e.DO.FindInBatches(result, batchSize, fc)
}

func (e ethBlockModelDo) Attrs(attrs ...field.AssignExpr) IEthBlockModelDo {
	return e.withDO(e.DO.Attrs(attrs...))
}

func (e ethBlockModelDo) Assign(attrs ...field.AssignExpr) IEthBlockModelDo {
	return e.withDO(e.DO.Assign(attrs...))
}

func (e ethBlockModelDo) Joins(fields ...field.RelationField) IEthBlockModelDo {
	for _, _f := range fields {
		e = *e.withDO(e.DO.Joins(_f))
	}
	return &e
}

func (e ethBlockModelDo) Preload(fields ...field.RelationField) IEthBlockModelDo {
	for _, _f := range fields {
		e = *e.withDO(e.DO.Preload(_f))
	}
	return &e
}

func (e ethBlockModelDo) FirstOrInit() (*models.EthBlockModel, error) {
	if result, err := e.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*models.EthBlockModel), nil
	}
}

func (e ethBlockModelDo) FirstOrCreate() (*models.EthBlockModel, error) {
	if result, err := e.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*models.EthBlockModel), nil
	}
}

func (e ethBlockModelDo) FindByPage(offset int, limit int) (result []*models.EthBlockModel, count int64, err error) {
	result, err = e.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = e.Offset(-1).Limit(-1).Count()
	return
}

func (e ethBlockModelDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = e.Count()
	if err != nil {
		return
	}

	err = e.Offset(offset).Limit(limit).Scan(result)
	return
}

func (e ethBlockModelDo) Scan(result interface{}) (err error) {
	return e.DO.Scan(result)
}

func (e ethBlockModelDo) Delete(models ...*models.EthBlockModel) (result gen.ResultInfo, err error) {
	return e.DO.Delete(models)
}

func (e *ethBlockModelDo) withDO(do gen.Dao) *ethBlockModelDo {
	e.DO = *do.(*gen.DO)
	return e
}
