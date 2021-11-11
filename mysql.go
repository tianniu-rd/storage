package storage

import (
    "encoding/base64"
    "errors"
    "fmt"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "time"
)

type (
    MysqlBackend struct {
        db *gorm.DB
    }

    Chart struct {
        ID           uint `gorm:"primaryKey"`
        Path         string `gorm:"index;unique"`
        Content string
        CreatedAt    time.Time
        UpdatedAt    time.Time
    }
)

func NewMysqlBackend(mysqlDsn string) (*MysqlBackend, error) {
    // refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
    // dsn eg: "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
    db, err := gorm.Open(mysql.Open(mysqlDsn), &gorm.Config{})
    if err != nil {
        return nil, err
    }
    return &MysqlBackend{
        db:       db,
    }, nil
}

func (mb *MysqlBackend)ListObjects(prefix string) ([]Object, error) {
    result := []Object{}
    //dbQueryRes := map[string]Chart{}
    charts := []Chart{}
    if prefix != "" {
        return result, fmt.Errorf("mysql backend donot support ListObjects with prefix: prefix=%s", prefix)
    }
    queryResult := mb.db.Select("path", "created_at").Find(&charts)
    if queryResult.Error != nil {
        return result, queryResult.Error
    }

    for _, content := range charts{
        result = append(result, Object{Path: content.Path, Content: []byte{}, LastModified: content.UpdatedAt})
    }
    return result, nil
}

func (mb *MysqlBackend)GetObject(path string) (Object, error){
    resultChart := Chart{}
    resultObject := Object{}
    queryRes := mb.db.Where("path = ?", path).Find(&resultChart)
    if queryRes.Error != nil {
        return resultObject, queryRes.Error
    }
    if queryRes.RowsAffected == 0 {
        return resultObject, fmt.Errorf("mysql backend donot find the chart: path=%s", path)
    }
    contentBytes, _ := base64.StdEncoding.DecodeString(resultChart.Content)
    resultObject = Object{
        Path: resultChart.Path,
        Content: contentBytes,
        LastModified: resultChart.UpdatedAt,
    }
    return resultObject, nil
}

func (mb *MysqlBackend)PutObject(path string, content []byte) error {
    contentStr := base64.StdEncoding.EncodeToString(content)
    newChart := Chart{
        Path:      path,
        Content:   contentStr,
    }

    oldChart := Chart{}
    dbResult := mb.db.Where("path = ?", path).First(&oldChart)
    if dbResult.Error != nil && !errors.Is(dbResult.Error, gorm.ErrRecordNotFound) {
        return dbResult.Error
    }
    if errors.Is(dbResult.Error, gorm.ErrRecordNotFound) {
        // add new chart
        dbResult := mb.db.Create(&newChart)
        if dbResult.Error != nil {
            return dbResult.Error
        }
        return nil
    } else {
        // update chart and index.yaml
        oldChart.Content = contentStr
        dbResult := mb.db.Save(&oldChart)
        if dbResult.Error != nil {
            return dbResult.Error
        }
        return nil
    }


}

func (mb *MysqlBackend) DeleteObject(path string) error {
    oldChart := Chart{}
    dbResult := mb.db.Where("path = ?", path).First(&oldChart)
    if dbResult.Error != nil && !errors.Is(dbResult.Error, gorm.ErrRecordNotFound) {
        return dbResult.Error
    }
    if errors.Is(dbResult.Error, gorm.ErrRecordNotFound) {
        return nil
    } else {
        // update chart and index.yaml
        dbResult := mb.db.Delete(&oldChart)
        if dbResult.Error != nil {
            return dbResult.Error
        }
        return nil
    }
}


