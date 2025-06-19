package service

import (
	"golang.org/x/net/context"
	"novel-go/config"
	"novel-go/model/pojo"
	"novel-go/model/resp"
)

type NewsService interface {
	ListLatestNews(ctx context.Context) ([]resp.NewsInfoRespDto, error)
	GetNews(ctx context.Context, id int64) (resp.NewsInfoRespDto, error)
}

type NewsServiceImpl struct{}

func NewNewsService() NewsService {
	return &NewsServiceImpl{}
}

func (n *NewsServiceImpl) ListLatestNews(ctx context.Context) ([]resp.NewsInfoRespDto, error) {
	var newsInfos []pojo.NewsInfo
	if err := config.DB.WithContext(ctx).Order("create_time desc").Limit(2).Find(&newsInfos).Error; err != nil {
		return nil, err
	}

	var respList []resp.NewsInfoRespDto
	for _, news := range newsInfos {
		respList = append(respList, resp.NewsInfoRespDto{
			ID:           news.ID,
			CategoryID:   news.CategoryID,
			CategoryName: news.CategoryName,
			SourceName:   news.SourceName,
			Title:        news.Title,
			UpdateTime:   news.UpdateTime,
		})
	}
	return respList, nil
}

func (n *NewsServiceImpl) GetNews(ctx context.Context, id int64) (resp.NewsInfoRespDto, error) {
	var news pojo.NewsInfo
	if err := config.DB.WithContext(ctx).First(&news, id).Error; err != nil {
		return resp.NewsInfoRespDto{}, err
	}

	var content pojo.NewsContent
	if err := config.DB.WithContext(ctx).Where("news_id = ?", id).Limit(1).Find(&content).Error; err != nil {
		return resp.NewsInfoRespDto{}, err
	}

	return resp.NewsInfoRespDto{
		ID:           news.ID,
		CategoryID:   news.CategoryID,
		CategoryName: news.CategoryName,
		SourceName:   news.SourceName,
		Title:        news.Title,
		UpdateTime:   news.UpdateTime,
		Content:      content.Content,
	}, nil
}
