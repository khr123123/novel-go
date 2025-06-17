package service

import (
	"context"
	"novel-go/config"
	"novel-go/model/pojo"
	"novel-go/model/resp"
)

// HomeService 首页服务接口
type HomeService interface {
	// ListHomeBooks 获取首页推荐小说列表
	// @Summary 首页小说推荐查询接口
	// @Description 获取首页推荐的小说列表
	// @Tags 首页
	// @Produce json
	// @Success 200 {array} resp.HomeBookRespDto
	// @Failure 500 {string} string "服务器错误"
	ListHomeBooks(ctx context.Context) ([]resp.HomeBookRespDto, error)

	// ListHomeFriendLinks 获取首页友情链接列表
	// @Summary 首页友情链接列表查询接口
	// @Description 获取首页展示的友情链接
	// @Tags 首页
	// @Produce json
	// @Success 200 {array} resp.HomeFriendLinkRespDto
	// @Failure 500 {string} string "服务器错误"
	ListHomeFriendLinks(ctx context.Context) ([]resp.HomeFriendLinkRespDto, error)
}

// homeServiceImpl 是 HomeService 的具体实现
type homeServiceImpl struct{}

// NewHomeService 构造函数
func NewHomeService() HomeService {
	return &homeServiceImpl{}
}

func (s *homeServiceImpl) ListHomeBooks(ctx context.Context) ([]resp.HomeBookRespDto, error) {
	var homeBooks []pojo.HomeBook
	if err := config.DB.WithContext(ctx).Order("sort asc").Find(&homeBooks).Error; err != nil {
		return nil, err
	}
	if len(homeBooks) == 0 {
		return []resp.HomeBookRespDto{}, nil
	}
	bookIds := make([]int64, 0, len(homeBooks))
	for _, hb := range homeBooks {
		bookIds = append(bookIds, hb.BookID)
	}
	var bookInfos []pojo.BookInfo
	if err := config.DB.WithContext(ctx).Where("id IN ?", bookIds).Find(&bookInfos).Error; err != nil {
		return nil, err
	}
	bookInfoMap := make(map[int64]pojo.BookInfo)
	for _, bi := range bookInfos {
		bookInfoMap[bi.ID] = bi
	}

	result := make([]resp.HomeBookRespDto, 0, len(homeBooks))
	for _, hb := range homeBooks {
		if bi, ok := bookInfoMap[hb.BookID]; ok {
			result = append(result, resp.HomeBookRespDto{
				Type:       hb.Type,
				BookID:     hb.BookID,
				PicUrl:     bi.PicURL,
				BookName:   bi.BookName,
				AuthorName: bi.AuthorName,
				BookDesc:   bi.BookDesc,
			})
		}
	}

	return result, nil
}

func (s *homeServiceImpl) ListHomeFriendLinks(ctx context.Context) ([]resp.HomeFriendLinkRespDto, error) {
	var links []pojo.HomeFriendLink
	if err := config.DB.WithContext(ctx).Order("sort asc").Find(&links).Error; err != nil {
		return nil, err
	}

	result := make([]resp.HomeFriendLinkRespDto, 0, len(links))
	for _, link := range links {
		result = append(result, resp.HomeFriendLinkRespDto{
			LinkName: link.LinkName,
			LinkUrl:  link.LinkUrl,
		})
	}
	return result, nil
}
