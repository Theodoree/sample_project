package tag_service

import (
	"encoding/json"
	"io"
	"strconv"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/tealeg/xlsx"

	"github.com/Theodoree/sample_project/sample/gin-mvc-demo/models"
	"github.com/Theodoree/sample_project/sample/gin-mvc-demo/pkg/export"
	"github.com/Theodoree/sample_project/sample/gin-mvc-demo/pkg/gredis"
	"github.com/Theodoree/sample_project/sample/gin-mvc-demo/pkg/logging"
	"github.com/Theodoree/sample_project/sample/gin-mvc-demo/service/cache_service"
)

type Tag struct {
	ID         int
	Name       string
	CreatedBy  string
	ModifiedBy string
	State      int

	PageNum  int
	PageSize int
}

//是否存在
func (t *Tag) ExistByName() (bool, error) {
	return models.ExistTagByName(t.Name)
}

//存在通过id
func (t *Tag) ExistByID() (bool, error) {
	return models.ExistTagByID(t.ID)
}

//AddTag
func (t *Tag) Add() error {
	return models.AddTag(t.Name, t.State, t.CreatedBy)
}

func (t *Tag) Edit() error {
	data := make(map[string]interface{})
	data["modified_by"] = t.ModifiedBy
	data["name"] = t.Name
	if t.State >= 0 {
		data["state"] = t.State
	}

	return models.EditTag(t.ID, data)
}

func (t *Tag) Delete() error {
	return models.DeleteTag(t.ID)
}

func (t *Tag) Count() (int, error) {
	return models.GetTagTotal(t.getMaps())
}

func (t *Tag) GetAll() ([]models.Tag, error) {
	var (
		tags, cacheTags []models.Tag
	)
	cache := cache_service.Tag{
		State:    t.State,
		PageNum:  t.PageNum,
		PageSize: t.PageSize,
	}
	key := cache.GetTagsKey()
	if gredis.Exists(key) { //检查key是否存在
		data, err := gredis.Get(key)
		if err != nil {
			logging.Info(err)
		} else {
			json.Unmarshal(data, &cacheTags)
			return cacheTags, nil
		}
	}
	//如果从redis中搜索到key,则从数据库中取
	tags, err := models.GetTags(t.PageNum, t.PageSize, t.getMaps())
	if err != nil {
		return nil, err
	}

	gredis.Set(key, tags, 3600)
	return tags, nil
}

func (t *Tag) Export() (string, error) {
	tags, err := t.GetAll() //获取数据
	if err != nil {
		return "", err
	}

	file := xlsx.NewFile() //Create New File
	sheet, err := file.AddSheet("标签信息")
	if err != nil {
		return "", err
	}

	titles := []string{"ID", "名称", "创建人", "创建时间", "修改人", "修改时间"}
	row := sheet.AddRow()

	var cell *xlsx.Cell
	for _, title := range titles {
		cell = row.AddCell()
		cell.Value = title //添加标题
	}

	for _, v := range tags {
		values := []string{
			strconv.Itoa(v.ID),
			v.Name,
			v.CreatedBy,
			v.CreatedOn,
			v.ModifiedBy,
			v.ModifiedOn,
		}

		row = sheet.AddRow() //添加新一层
		for _, value := range values {
			cell = row.AddCell() //获取空格
			cell.Value = value   //添加value
		}
	}

	Time := time.Now().Format(time.Local.String())
	filename := "tags-" + Time + ".xlsx" //构建file Name

	fullPath := export.GetExcelFullPath() + filename //构建文件路径
	err = file.Save(fullPath)                        //储存
	if err != nil {
		return "", err
	}

	return filename, nil
}

func (t *Tag) Import(r io.Reader) error {
	Xlsx, err := excelize.OpenReader(r)
	if err != nil {
		return err
	}

	rows := Xlsx.GetRows("标签信息")
	for irow, row := range rows {
		if irow > 0 {
			var data []string
			for _, cell := range row {
				data = append(data, cell)
			}

			models.AddTag(data[1], 1, data[2])
		}
	}

	return nil
}

func (t *Tag) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	maps["deleted_on"] = 0

	if t.Name != "" {
		maps["name"] = t.Name
	}
	if t.State >= 0 {
		maps["state"] = t.State
	}

	return maps
}
