package service

import (
	"novel-go/config"
	"novel-go/model/pojo"
	"novel-go/model/req"
	"novel-go/model/resp"
	"strconv"
	"time"
)

type SearchService interface {
	SearchBooks(condition req.BookSearchReqDto) (resp.PageRespDto[resp.BookInfoRespDto], error)
}

type DBSearchServiceImpl struct {
}

func NewDBSearchServiceImpl() SearchService {
	return new(DBSearchServiceImpl)
}

func (D *DBSearchServiceImpl) SearchBooks(cond req.BookSearchReqDto) (resp.PageRespDto[resp.BookInfoRespDto], error) {
	var books []pojo.BookInfo
	var total int64
	query := config.DB.Model(&pojo.BookInfo{})

	if cond.Keyword != "" {
		query = query.Where("book_name LIKE ?", "%"+cond.Keyword+"%")
	}
	if cond.WorkDirection != nil {
		query = query.Where("work_direction = ?", *cond.WorkDirection)
	}
	if cond.Category != nil {
		query = query.Where("category_name = ?", *cond.Category)
	}
	if cond.IsVip != nil {
		query = query.Where("is_vip = ?", *cond.IsVip)
	}
	if cond.BookStatus != nil {
		query = query.Where("book_status = ?", *cond.BookStatus)
	}
	if cond.WordCountMin != nil {
		query = query.Where("word_count >= ?", *cond.WordCountMin)
	}
	if cond.WordCountMax != nil {
		query = query.Where("word_count <= ?", *cond.WordCountMax)
	}
	if cond.UpdateTime != nil && *cond.UpdateTime != "" {
		// 将字符串数字转换成整数，表示多少天内
		days, err := strconv.Atoi(*cond.UpdateTime)
		if err == nil && days > 0 {
			// 计算多少天前的时间
			cutoffTime := time.Now().AddDate(0, 0, -days)
			query = query.Where("update_time >= ?", cutoffTime)
		}
	}

	if cond.Sort != "" {
		query = query.Order(cond.Sort)
	}
	query.Count(&total)

	offset := (cond.PageNum - 1) * cond.PageSize
	err := query.Offset(int(offset)).Limit(int(cond.PageSize)).Find(&books).Error

	var result []resp.BookInfoRespDto
	for _, b := range books {
		result = append(result, resp.BookInfoRespDto{
			ID:              b.ID,
			BookName:        b.BookName,
			CategoryID:      b.CategoryID,
			CategoryName:    b.CategoryName,
			AuthorID:        b.AuthorID,
			AuthorName:      b.AuthorName,
			WordCount:       b.WordCount,
			LastChapterName: b.LastChapterName,
		})
	}

	return resp.PageRespDto[resp.BookInfoRespDto]{
		PageNum:  cond.PageNum,
		PageSize: cond.PageSize,
		Total:    total,
		List:     result,
	}, err
}

type ESSearchServiceImpl struct {
}

func NewESSearchServiceImpl() SearchService {
	return new(ESSearchServiceImpl)
}
func (D *ESSearchServiceImpl) SearchBooks(condition req.BookSearchReqDto) (resp.PageRespDto[resp.BookInfoRespDto], error) {
	//TODO implement me
	panic("implement me")
}
