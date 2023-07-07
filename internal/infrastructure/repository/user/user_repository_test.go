package user

//
//import (
//	"context"
//	"github.com/oklog/ulid/v2"
//	"regexp"
//	"testing"
//
//	"github.com/DATA-DOG/go-sqlmock"
//	userDomain "github.com/hum2/backend/internal/domain/user"
//	"github.com/stretchr/testify/assert"
//	"gorm.io/driver/mysql"
//	"gorm.io/gorm"
//)
//
//func GetNewDbMock() (*gorm.DB, sqlmock.Sqlmock, error) {
//	db, mock, err := sqlmock.New()
//	if err != nil {
//		return nil, mock, err
//	}
//
//	gormDB, err := gorm.Open(mysql.Dialector{Config: &mysql.Config{DriverName: "mysql", Conn: db, SkipInitializeWithVersion: true}}, &gorm.Config{})
//	return gormDB, mock, err
//}
//
//func TestNew(t *testing.T) {
//	db, _, err := GetNewDbMock()
//	if err != nil {
//		t.Errorf("Failed to initialize mock db: %v", err)
//		return
//	}
//
//	repo := New(db)
//	assert.Equal(t, &Repository{db: db}, repo)
//}
//
//func TestFindAll(t *testing.T) {
//	db, mock, err := GetNewDbMock()
//	if err != nil {
//		t.Errorf("Failed to initialize mock db: %v", err)
//		return
//	}
//
//	rows := sqlmock.
//		NewRows([]string{"id", "name"}).
//		AddRow(1, "test1").
//		AddRow(2, "test2")
//	mock.
//		ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE `users`.`deleted_at` IS NULL")).
//		WillReturnRows(rows)
//
//	repo := New(db)
//	users, err := repo.FindAll()
//	assert.NoError(t, err)
//	assert.Len(t, users, 2)
//
//	expectUser1, err := userDomain.New(1, "test1")
//	assert.NoError(t, err)
//	expectUser2, err := userDomain.New(2, "test2")
//	assert.NoError(t, err)
//
//	assert.Equal(t, []*userDomain.User{expectUser1, expectUser2}, users)
//}
//
//func TestFindByID(t *testing.T) {
//	db, mock, err := GetNewDbMock()
//	if err != nil {
//		t.Errorf("Failed to initialize mock db: %v", err)
//		return
//	}
//
//	rows := sqlmock.
//		NewRows([]string{"id", "name"}).
//		AddRow(1, "test1")
//	mock.
//		ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE id = ? AND `users`.`deleted_at` IS NULL")).
//		WillReturnRows(rows)
//
//	repo := New(db)
//	user, err := repo.FindByID(context.Background(), ulid.ULID{})
//	assert.NoError(t, err)
//
//	expectUser, err := userDomain.New(ulid.ULID{}, "test1", "")
//	assert.NoError(t, err)
//	assert.Equal(t, expectUser, user)
//}
//
////func TestCreate(t *testing.T) {
////	db, mock, err := GetNewDbMock()
////	if err != nil {
////		t.Errorf("Failed to initialize mock db: %v", err)
////		return
////	}
////
////	id := 1
////	mock.
////		ExpectQuery("INSERT INTO `users` (.+) VALUES (.+)").
////		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(id))
////
////	repo := New(db)
////	user, err := repo.Create("test1")
////	assert.NoError(t, err)
////
////	expectUser, err := userDomain.New(1, "test1")
////	assert.NoError(t, err)
////	assert.Equal(t, expectUser, user)
////}
