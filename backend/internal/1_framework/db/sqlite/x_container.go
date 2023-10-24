package sqlite

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"backend/internal/2_adapter/gateway"
	domain "backend/internal/4_domain"
	"backend/pkg"
)

var (
	myErr *pkg.MyErr
)

type (
	// Sqlite ...
	Sqlite struct {
		Conn *gorm.DB
	}
)

// NewToSQLite ...
func NewToSQLite() gateway.ToSqlite {

	storeDBName, err := getDBName()
	if err != nil {
		return nil
	}

	sqliteFilePath := getSqlitePath() + "/" + storeDBName + ".sqlite3"

	conn, err := open(30, sqliteFilePath)
	if err != nil {
		myErr.Logging(err)
		panic(err)
	}

	sqlite := new(Sqlite)
	sqlite.Conn = conn

	return sqlite
}
func init() {
	myErr = pkg.NewMyErr("framework_driver", "sqlite")
	// DBのアップデート

	updateStoreDB()
}

func updateStoreDB() error {
	// local DB =======================
	storeDBName, err := getDBName()
	if err != nil {
		return err
	}

	storeDB, err := open(30, getSqlitePath()+"/"+storeDBName+".sqlite3")
	if err != nil {
		myErr.Logging(err)
		panic(err)
	}

	storeJANCodeList := []int{}

	storeDB.
		// Debug().
		Table("products").
		Select("jan_code").
		Order("jan_code desc").
		Find(&storeJANCodeList)

	// master DB =======================
	masterDB, err := open(30, getSqlitePath()+"/master.sqlite3")
	if err != nil {
		myErr.Logging(err)
		panic(err)
	}
	restJANCodeList := []int{}

	// storeDBにないjanCodeの検索
	masterDB.
		// Debug().
		Table("products").
		Select("jan_code").
		Order("jan_code desc").
		Not(map[string]interface{}{"jan_code": storeJANCodeList}).
		Find(&restJANCodeList)

	// 差分がなければ終了
	if len(restJANCodeList) == 0 {
		return nil
	}

	// storeDBにないデータ取得
	differentProductList := domain.AllProductList{}
	masterDB.
		// Debug().
		Where("jan_code IN (?)", restJANCodeList).
		Order("jan_code desc").
		Find(&differentProductList)

	// DB更新
	storeDB.
		// Debug().
		Create(&differentProductList)

	return nil
}

func getSqlitePath() string {

	currentPath, _ := os.Getwd()

	// localの場合は、db/engine/sqlite を確認
	sqlitePath := currentPath + "/db/engine/sqlite"
	// /home/yuji/Workspace/private_dev/template/backend/db/engine/sqlite/master.sqlite3
	if !exists(sqlitePath + "/master.sqlite3") {
		// dockerの場合は、/go/src/backend/sqlite を確認
		sqlitePath = filepath.Dir(currentPath) + "/backend/sqlite"
	}

	return sqlitePath
}

func exists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func open(
	count uint,
	dbPath string,
) (
	*gorm.DB,
	error,
) {

	db, err := gorm.Open(
		sqlite.Open(dbPath),
		&gorm.Config{},
	)
	if err != nil {
		if count == 0 {
			myErr.Logging(err)
			return nil, fmt.Errorf("retry count over")
		}
		time.Sleep(time.Second)
		// カウントダウンさせるようにする
		count--
		return open(count, dbPath)
	}

	return db, nil
}

func getDBName() (string, error) {
	currentPath, _ := os.Getwd()

	err := godotenv.Load(filepath.Join(currentPath, ".env"))
	if err != nil {
		log.Fatal("Error loading .env file")
		return "", nil
	}

	return os.Getenv("SQLITE_FILE"), nil
}
