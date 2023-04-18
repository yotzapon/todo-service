// Code generated for package migrations by go-bindata DO NOT EDIT. (@generated)
// sources:
// migrations/20230413143226_create_users_table.down.sql
// migrations/20230413143226_create_users_table.up.sql
// migrations/20230413143249_create_todos_table.down.sql
// migrations/20230413143249_create_todos_table.up.sql
// migrations/20230413143314_insert_users_table.down.sql
// migrations/20230413143314_insert_users_table.up.sql
package migrations

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var __20230413143226_create_users_tableDownSql = []byte(`DROP TABLE IF EXISTS ` + "`" + `users` + "`" + `;
`)

func _20230413143226_create_users_tableDownSqlBytes() ([]byte, error) {
	return __20230413143226_create_users_tableDownSql, nil
}

func _20230413143226_create_users_tableDownSql() (*asset, error) {
	bytes, err := _20230413143226_create_users_tableDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "20230413143226_create_users_table.down.sql", size: 30, mode: os.FileMode(420), modTime: time.Unix(1681396689, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __20230413143226_create_users_tableUpSql = []byte(`CREATE TABLE IF NOT EXISTS ` + "`" + `users` + "`" + ` (
     ` + "`" + `id` + "`" + ` int(11) NOT NULL AUTO_INCREMENT,
     ` + "`" + `username` + "`" + ` varchar(255) NOT NULL,
     ` + "`" + `password` + "`" + ` varchar(255) NOT NULL,
     PRIMARY KEY (` + "`" + `id` + "`" + `),
     UNIQUE KEY ` + "`" + `username` + "`" + ` (` + "`" + `username` + "`" + `)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;
`)

func _20230413143226_create_users_tableUpSqlBytes() ([]byte, error) {
	return __20230413143226_create_users_tableUpSql, nil
}

func _20230413143226_create_users_tableUpSql() (*asset, error) {
	bytes, err := _20230413143226_create_users_tableUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "20230413143226_create_users_table.up.sql", size: 281, mode: os.FileMode(420), modTime: time.Unix(1681436609, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __20230413143249_create_todos_tableDownSql = []byte(`DROP TABLE IF EXISTS ` + "`" + `todos` + "`" + `;
`)

func _20230413143249_create_todos_tableDownSqlBytes() ([]byte, error) {
	return __20230413143249_create_todos_tableDownSql, nil
}

func _20230413143249_create_todos_tableDownSql() (*asset, error) {
	bytes, err := _20230413143249_create_todos_tableDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "20230413143249_create_todos_table.down.sql", size: 30, mode: os.FileMode(420), modTime: time.Unix(1681396697, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __20230413143249_create_todos_tableUpSql = []byte(`CREATE TABLE IF NOT EXISTS ` + "`" + `todos` + "`" + ` (
     ` + "`" + `id` + "`" + ` int(11) NOT NULL AUTO_INCREMENT,
     ` + "`" + `title` + "`" + ` varchar(255) NOT NULL,
     ` + "`" + `description` + "`" + ` text,
     ` + "`" + `completed` + "`" + ` tinyint(1) NOT NULL DEFAULT '0',
     ` + "`" + `created_at` + "`" + ` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
     ` + "`" + `updated_at` + "`" + ` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
     ` + "`" + `user_id` + "`" + ` int(11) NOT NULL,
     PRIMARY KEY (` + "`" + `id` + "`" + `),
     KEY ` + "`" + `user_id` + "`" + ` (` + "`" + `user_id` + "`" + `),
     CONSTRAINT ` + "`" + `todos_ibfk_1` + "`" + ` FOREIGN KEY (` + "`" + `user_id` + "`" + `) REFERENCES ` + "`" + `users` + "`" + ` (` + "`" + `id` + "`" + `)
         ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;
`)

func _20230413143249_create_todos_tableUpSqlBytes() ([]byte, error) {
	return __20230413143249_create_todos_tableUpSql, nil
}

func _20230413143249_create_todos_tableUpSql() (*asset, error) {
	bytes, err := _20230413143249_create_todos_tableUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "20230413143249_create_todos_table.up.sql", size: 621, mode: os.FileMode(420), modTime: time.Unix(1681436609, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __20230413143314_insert_users_tableDownSql = []byte(``)

func _20230413143314_insert_users_tableDownSqlBytes() ([]byte, error) {
	return __20230413143314_insert_users_tableDownSql, nil
}

func _20230413143314_insert_users_tableDownSql() (*asset, error) {
	bytes, err := _20230413143314_insert_users_tableDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "20230413143314_insert_users_table.down.sql", size: 0, mode: os.FileMode(420), modTime: time.Unix(1681396394, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __20230413143314_insert_users_tableUpSql = []byte(`INSERT INTO ` + "`" + `users` + "`" + ` (` + "`" + `username` + "`" + `, ` + "`" + `password` + "`" + `) VALUES
     ('tester01', '1111'),
     ('tester02', '2222'),
     ('tester03', '3333');
`)

func _20230413143314_insert_users_tableUpSqlBytes() ([]byte, error) {
	return __20230413143314_insert_users_tableUpSql, nil
}

func _20230413143314_insert_users_tableUpSql() (*asset, error) {
	bytes, err := _20230413143314_insert_users_tableUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "20230413143314_insert_users_table.up.sql", size: 133, mode: os.FileMode(420), modTime: time.Unix(1681396752, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"20230413143226_create_users_table.down.sql": _20230413143226_create_users_tableDownSql,
	"20230413143226_create_users_table.up.sql":   _20230413143226_create_users_tableUpSql,
	"20230413143249_create_todos_table.down.sql": _20230413143249_create_todos_tableDownSql,
	"20230413143249_create_todos_table.up.sql":   _20230413143249_create_todos_tableUpSql,
	"20230413143314_insert_users_table.down.sql": _20230413143314_insert_users_tableDownSql,
	"20230413143314_insert_users_table.up.sql":   _20230413143314_insert_users_tableUpSql,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//
//	data/
//	  foo.txt
//	  img/
//	    a.png
//	    b.png
//
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"20230413143226_create_users_table.down.sql": &bintree{_20230413143226_create_users_tableDownSql, map[string]*bintree{}},
	"20230413143226_create_users_table.up.sql":   &bintree{_20230413143226_create_users_tableUpSql, map[string]*bintree{}},
	"20230413143249_create_todos_table.down.sql": &bintree{_20230413143249_create_todos_tableDownSql, map[string]*bintree{}},
	"20230413143249_create_todos_table.up.sql":   &bintree{_20230413143249_create_todos_tableUpSql, map[string]*bintree{}},
	"20230413143314_insert_users_table.down.sql": &bintree{_20230413143314_insert_users_tableDownSql, map[string]*bintree{}},
	"20230413143314_insert_users_table.up.sql":   &bintree{_20230413143314_insert_users_tableUpSql, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
