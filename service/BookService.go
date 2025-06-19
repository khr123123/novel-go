package service

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"novel-go/config"
	"novel-go/model/pojo"
	"novel-go/model/req"
	"novel-go/model/resp"
	"strconv"
	"time"
)

// BookService 小说模块服务接口
type BookService interface {
	ListVisitRankBooks(ctx context.Context) ([]resp.BookRankRespDto, error)
	ListNewestRankBooks(ctx context.Context) ([]resp.BookRankRespDto, error)
	ListUpdateRankBooks(ctx context.Context) ([]resp.BookRankRespDto, error)
	GetBookById(ctx context.Context, bookId string) (resp.BookInfoRespDto, error)
	GetBookContentAbout(ctx context.Context, chapterId string) (resp.BookContentAboutRespDto, error)
	GetLastChapterAbout(ctx context.Context, bookId string) (resp.BookChapterAboutRespDto, error)
	ListRecBooks(ctx context.Context, bookId string) ([]resp.BookInfoRespDto, error)
	AddVisitCount(ctx context.Context, bookId string) error
	GetPreChapterId(ctx context.Context, chapterId string) (string, error)
	GetNextChapterId(ctx context.Context, chapterId string) (string, error)
	ListChapters(ctx context.Context, bookId string) ([]resp.BookChapterRespDto, error)
	ListCategory(ctx context.Context, workDirection int) ([]resp.BookCategoryRespDto, error)
	SaveComment(ctx context.Context, dto req.UserCommentReqDto) error
	ListNewestComments(ctx context.Context, bookId string) (resp.BookCommentRespDto, error)
	DeleteComment(ctx context.Context, userId int64, commentId string) error
	UpdateComment(ctx context.Context, userId int64, id int64, content string) error
	SaveBook(ctx context.Context, dto req.BookAddReqDto) error
	SaveBookChapter(ctx context.Context, dto req.ChapterAddReqDto) error
	ListAuthorBooks(ctx context.Context, pageReq req.PageReqDto) (resp.PageRespDto[resp.BookInfoRespDto], error)
	ListBookChapters(ctx context.Context, bookId string, pageReq req.PageReqDto) (resp.PageRespDto[resp.BookChapterRespDto], error)
	ListComments(ctx context.Context, userId int64, pageReq req.PageReqDto) (resp.PageRespDto[resp.UserCommentRespDto], error)
	DeleteBookChapter(ctx context.Context, chapterId string) error
	GetBookChapter(ctx context.Context, chapterId string) (resp.ChapterContentRespDto, error)
	UpdateBookChapter(ctx context.Context, chapterId string, dto req.ChapterUpdateReqDto) error
}

type BookServiceImpl struct {
}

func (b *BookServiceImpl) ListVisitRankBooks(ctx context.Context) ([]resp.BookRankRespDto, error) {
	var books []pojo.BookInfo
	err := config.DB.WithContext(ctx).
		Order("visit_count DESC").
		Limit(20).
		Find(&books).Error
	if err != nil {
		return nil, err
	}
	return b.toBookRankRespDtoList(books), nil
}

func (b *BookServiceImpl) ListNewestRankBooks(ctx context.Context) ([]resp.BookRankRespDto, error) {
	var books []pojo.BookInfo
	err := config.DB.WithContext(ctx).
		Order("create_time DESC").
		Limit(20).
		Find(&books).Error
	if err != nil {
		return nil, err
	}
	return b.toBookRankRespDtoList(books), nil
}

func (b *BookServiceImpl) ListUpdateRankBooks(ctx context.Context) ([]resp.BookRankRespDto, error) {
	var books []pojo.BookInfo
	err := config.DB.WithContext(ctx).
		Order("last_chapter_update_time DESC").
		Limit(20).
		Find(&books).Error
	if err != nil {
		return nil, err
	}
	return b.toBookRankRespDtoList(books), nil
}

func (b *BookServiceImpl) toBookRankRespDtoList(books []pojo.BookInfo) []resp.BookRankRespDto {
	list := make([]resp.BookRankRespDto, 0, len(books))
	for _, bo := range books {
		list = append(list, resp.BookRankRespDto{
			ID:                    bo.ID,
			BookName:              bo.BookName,
			AuthorName:            bo.AuthorName,
			CategoryName:          bo.CategoryName,
			PicURL:                bo.PicURL,
			LastChapterName:       bo.LastChapterName,
			LastChapterUpdateTime: bo.LastChapterUpdateTime,
		})
	}
	return list
}

func (b *BookServiceImpl) GetBookById(ctx context.Context, bookId string) (resp.BookInfoRespDto, error) {
	var firstBookChapter pojo.BookChapter
	err := config.DB.WithContext(ctx).First(&firstBookChapter).Where("book_id", bookId).Error
	if err != nil {
		return resp.BookInfoRespDto{}, err
	}
	var book pojo.BookInfo
	if err = config.DB.WithContext(ctx).First(&book, bookId).Error; err != nil {
		return resp.BookInfoRespDto{}, err
	}
	return resp.BookInfoRespDto{
		ID:              book.ID,
		BookName:        book.BookName,
		AuthorName:      book.AuthorName,
		CategoryID:      book.CategoryID,
		CategoryName:    book.CategoryName,
		BookDesc:        book.BookDesc,
		PicURL:          book.PicURL,
		VisitCount:      book.VisitCount,
		FirstChapterId:  firstBookChapter.ID,
		LastChapterId:   book.LastChapterID,
		LastChapterName: book.LastChapterName,
		WordCount:       book.WordCount,
		UpdateTime:      book.UpdateTime,
	}, nil
}

func (b *BookServiceImpl) GetBookContentAbout(ctx context.Context, chapterId string) (resp.BookContentAboutRespDto, error) {
	// 查询章节信息
	var chapter pojo.BookChapter
	if err := config.DB.WithContext(ctx).First(&chapter, chapterId).Error; err != nil {
		return resp.BookContentAboutRespDto{}, fmt.Errorf("章节不存在: %w", err)
	}

	// 查询章节内容
	var content pojo.BookContent
	if err := config.DB.WithContext(ctx).Where("chapter_id = ?", chapterId).First(&content).Error; err != nil {
		return resp.BookContentAboutRespDto{}, fmt.Errorf("章节内容不存在: %w", err)
	}

	// 查询小说信息
	var book pojo.BookInfo
	if err := config.DB.WithContext(ctx).First(&book, chapter.BookID).Error; err != nil {
		return resp.BookContentAboutRespDto{}, fmt.Errorf("小说信息不存在: %w", err)
	}

	// 构造返回结构
	return resp.BookContentAboutRespDto{
		BookInfo: resp.BookInfoRespDto{
			ID:           book.ID,
			BookName:     book.BookName,
			AuthorName:   book.AuthorName,
			CategoryID:   book.CategoryID,
			CategoryName: book.CategoryName,
			PicURL:       book.PicURL,
			BookDesc:     book.BookDesc,
		},
		ChapterInfo: resp.BookChapterRespDto{
			ID:                chapter.ID,
			BookID:            chapter.BookID,
			ChapterNum:        chapter.ChapterNum,
			ChapterName:       chapter.ChapterName,
			ChapterWordCount:  chapter.WordCount,
			ChapterUpdateTime: chapter.UpdateTime,
			IsVip:             chapter.IsVip,
		},
		BookContent: content.Content,
	}, nil
}

func (b *BookServiceImpl) GetLastChapterAbout(ctx context.Context, bookId string) (resp.BookChapterAboutRespDto, error) {
	// 1. 查询小说信息
	var book pojo.BookInfo
	if err := config.DB.WithContext(ctx).First(&book, bookId).Error; err != nil {
		return resp.BookChapterAboutRespDto{}, fmt.Errorf("小说信息不存在: %w", err)
	}

	// 2. 查询最新章节信息
	var chapter pojo.BookChapter
	if err := config.DB.WithContext(ctx).
		Where("id = ?", book.LastChapterID).
		First(&chapter).Error; err != nil {
		return resp.BookChapterAboutRespDto{}, fmt.Errorf("章节信息不存在: %w", err)
	}

	// 3. 查询章节内容
	var content pojo.BookContent
	if err := config.DB.WithContext(ctx).
		Where("chapter_id = ?", book.LastChapterID).
		First(&content).Error; err != nil {
		return resp.BookChapterAboutRespDto{}, fmt.Errorf("章节内容不存在: %w", err)
	}

	// 4. 查询章节总数
	var chapterTotal int64
	if err := config.DB.WithContext(ctx).
		Model(&pojo.BookChapter{}).
		Where("book_id = ?", bookId).
		Count(&chapterTotal).Error; err != nil {
		return resp.BookChapterAboutRespDto{}, fmt.Errorf("章节总数获取失败: %w", err)
	}

	// 5. 提取前30字符摘要
	summary := content.Content
	if len(summary) > 30 {
		summary = summary[:30]
	}

	// 6. 构造章节信息响应
	chapterInfo := resp.BookChapterRespDto{
		ID:                chapter.ID,
		BookID:            chapter.BookID,
		ChapterNum:        chapter.ChapterNum,
		ChapterName:       chapter.ChapterName,
		ChapterWordCount:  chapter.WordCount,
		ChapterUpdateTime: chapter.UpdateTime,
		IsVip:             chapter.IsVip,
	}

	// 7. 返回封装结果
	return resp.BookChapterAboutRespDto{
		ChapterInfo:    chapterInfo,
		ChapterTotal:   chapterTotal,
		ContentSummary: summary,
	}, nil
}

func NewBookServiceImpl() BookService {
	return &BookServiceImpl{}
}

func (b *BookServiceImpl) ListRecBooks(ctx context.Context, bookId string) ([]resp.BookInfoRespDto, error) {
	var book pojo.BookInfo
	if err := config.DB.WithContext(ctx).First(&book, bookId).Error; err != nil {
		return nil, err
	}
	var books []pojo.BookInfo
	err := config.DB.WithContext(ctx).
		Where("category_id = ? AND id != ?", book.CategoryID, bookId).
		Order("visit_count DESC").
		Limit(10).
		Find(&books).Error
	if err != nil {
		return nil, err
	}
	var list []resp.BookInfoRespDto
	for _, bo := range books {
		list = append(list, resp.BookInfoRespDto{
			ID:           bo.ID,
			BookName:     bo.BookName,
			AuthorName:   bo.AuthorName,
			CategoryID:   bo.CategoryID,
			CategoryName: bo.CategoryName,
			PicURL:       bo.PicURL,
		})
	}
	return list, nil
}

func (b *BookServiceImpl) AddVisitCount(ctx context.Context, bookId string) error {
	return config.DB.WithContext(ctx).
		Model(&pojo.BookInfo{}).
		Where("id = ?", bookId).
		UpdateColumn("visit_count", gorm.Expr("visit_count + 1")).Error
}

func (b *BookServiceImpl) GetPreChapterId(ctx context.Context, chapterId string) (string, error) {
	var chapter pojo.BookChapter
	if err := config.DB.WithContext(ctx).First(&chapter, chapterId).Error; err != nil {
		return "0", err
	}
	var preChapter pojo.BookChapter
	err := config.DB.WithContext(ctx).
		Where("book_id = ? AND chapter_num < ?", chapter.BookID, chapter.ChapterNum).
		Order("chapter_num DESC").
		Limit(1).
		Find(&preChapter).Error
	if err != nil {
		return "0", err
	}
	return strconv.FormatInt(preChapter.ID, 10), nil
}

func (b *BookServiceImpl) GetNextChapterId(ctx context.Context, chapterId string) (string, error) {
	var chapter pojo.BookChapter
	if err := config.DB.WithContext(ctx).First(&chapter, chapterId).Error; err != nil {
		return "0", err
	}
	var nextChapter pojo.BookChapter
	err := config.DB.WithContext(ctx).
		Where("book_id = ? AND chapter_num > ?", chapter.BookID, chapter.ChapterNum).
		Order("chapter_num ASC").
		Limit(1).
		Find(&nextChapter).Error
	if err != nil {
		return "0", err
	}
	return strconv.FormatInt(nextChapter.ID, 10), nil
}

func (b *BookServiceImpl) ListChapters(ctx context.Context, bookId string) ([]resp.BookChapterRespDto, error) {
	var chapters []pojo.BookChapter
	err := config.DB.WithContext(ctx).
		Where("book_id = ?", bookId).
		Order("chapter_num ASC").
		Find(&chapters).Error
	if err != nil {
		return nil, err
	}
	var result []resp.BookChapterRespDto
	for _, c := range chapters {
		result = append(result, resp.BookChapterRespDto{
			ID:                c.ID,
			BookID:            c.BookID,
			ChapterNum:        c.ChapterNum,
			ChapterName:       c.ChapterName,
			ChapterWordCount:  c.WordCount,
			ChapterUpdateTime: c.UpdateTime,
			IsVip:             c.IsVip,
		})
	}
	return result, nil
}

func (b *BookServiceImpl) ListCategory(ctx context.Context, workDirection int) ([]resp.BookCategoryRespDto, error) {
	var categories []pojo.BookCategory
	err := config.DB.WithContext(ctx).
		Where("work_direction = ?", workDirection).
		Find(&categories).Error
	if err != nil {
		return nil, err
	}
	var result []resp.BookCategoryRespDto
	for _, c := range categories {
		result = append(result, resp.BookCategoryRespDto{
			ID:   c.ID,
			Name: c.Name,
		})
	}
	return result, nil
}

func (b *BookServiceImpl) SaveComment(ctx context.Context, dto req.UserCommentReqDto) error {
	lockKey := fmt.Sprintf("lock:comment:user:%d:book:%d", dto.UserId, dto.BookId)
	// 创建锁	expiry: 8 * time.Second,
	mutex := config.Redsync.NewMutex(lockKey)
	// 加锁
	if err := mutex.Lock(); err != nil {
		return fmt.Errorf("加锁失败: %v", err)
	}
	defer mutex.Unlock()

	// 检查是否已评论（可选逻辑）
	var count int64
	if err := config.DB.WithContext(ctx).Model(&pojo.BookComment{}).
		Where("user_id = ? AND book_id = ?", dto.UserId, dto.BookId).
		Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return fmt.Errorf("用户已评论，不能重复提交")
	}

	// 执行保存评论
	comment := pojo.BookComment{
		UserID:         dto.UserId,
		BookID:         dto.BookId,
		CommentContent: dto.CommentContent,
	}
	return config.DB.WithContext(ctx).Create(&comment).Error
}

func (b *BookServiceImpl) ListNewestComments(ctx context.Context, bookId string) (resp.BookCommentRespDto, error) {
	// 1. 查询评论总数
	var commentTotal int64
	if err := config.DB.WithContext(ctx).
		Model(&pojo.BookComment{}).
		Where("book_id = ?", bookId).
		Count(&commentTotal).Error; err != nil {
		return resp.BookCommentRespDto{}, err
	}

	result := resp.BookCommentRespDto{
		CommentTotal: commentTotal,
	}

	if commentTotal == 0 {
		result.Comments = []resp.CommentInfo{} // 空列表
		return result, nil
	}

	// 2. 查询最新评论（最多5条）
	var bookComments []pojo.BookComment
	if err := config.DB.WithContext(ctx).
		Where("book_id = ?", bookId).
		Order("create_time DESC").
		Limit(5).
		Find(&bookComments).Error; err != nil {
		return resp.BookCommentRespDto{}, err
	}

	// 3. 查询用户信息
	userIds := make([]int64, 0, len(bookComments))
	userIdSet := make(map[int64]struct{})
	for _, c := range bookComments {
		if _, exists := userIdSet[c.UserID]; !exists {
			userIdSet[c.UserID] = struct{}{}
			userIds = append(userIds, c.UserID)
		}
	}

	var userInfos []pojo.UserInfo
	if err := config.DB.WithContext(ctx).
		Where("id IN ?", userIds).
		Find(&userInfos).Error; err != nil {
		return resp.BookCommentRespDto{}, err
	}

	userInfoMap := make(map[int64]resp.UserInfoRespDto, len(userInfos))
	for _, u := range userInfos {
		userInfoMap[u.ID] = resp.UserInfoRespDto{
			NickName:  u.NickName,
			UserPhoto: u.UserPhoto,
			UserSex:   u.UserSex,
		}
	}

	// 4. 组装评论带用户信息
	commentInfos := make([]resp.CommentInfo, 0, len(bookComments))
	for _, c := range bookComments {
		userInfo, ok := userInfoMap[c.UserID]
		if !ok {
			userInfo = resp.UserInfoRespDto{
				NickName:  "未知用户",
				UserPhoto: "",
				UserSex:   0,
			}
		}
		commentInfos = append(commentInfos, resp.CommentInfo{
			ID:               c.ID,
			CommentUserId:    c.UserID,
			CommentUser:      userInfo.NickName,
			CommentUserPhoto: userInfo.UserPhoto,
			CommentContent:   c.CommentContent,
			CommentTime:      c.CreateTime,
		})
	}

	result.Comments = commentInfos
	return result, nil
}

func (b *BookServiceImpl) DeleteComment(ctx context.Context, userId int64, commentId string) error {
	return config.DB.WithContext(ctx).
		Where("id = ? AND user_id = ?", commentId, userId).
		Delete(&pojo.BookComment{}).Error
}

func (b *BookServiceImpl) UpdateComment(ctx context.Context, userId int64, id int64, content string) error {
	return config.DB.WithContext(ctx).
		Model(&pojo.BookComment{}).
		Where("id = ? AND user_id = ?", id, userId).
		Update("content", content).Error
}

func (b *BookServiceImpl) SaveBook(ctx context.Context, dto req.BookAddReqDto) error {
	// 1. 校验小说名是否已存在
	var count int64
	err := config.DB.WithContext(ctx).
		Model(&pojo.BookInfo{}).
		Where("book_name = ?", dto.BookName).
		Count(&count).Error
	if err != nil {
		return err
	}
	if count > 0 {
		return fmt.Errorf("小说名已存在") // 你可以自定义错误类型
	}

	// 2. 获取当前作者信息（假设有 UserHolder 获取当前用户ID）
	//userId := UserHolder.GetUserId(ctx) // 你需要自己实现这个方法或者传入用户ID
	//
	//author, err := authorInfoCacheManager.GetAuthor(ctx, userId)
	//if err != nil {
	//	return err
	//}
	//todo
	// 3. 创建小说实体并保存
	book := pojo.BookInfo{
		BookName: dto.BookName,
		//AuthorID:     author.ID,
		//AuthorName:   author.PenName,
		CategoryID:   dto.CategoryId,
		CategoryName: dto.CategoryName,
		BookDesc:     dto.BookDesc,
		PicURL:       dto.PicUrl,
		IsVip:        dto.IsVip,
		Score:        0,
		CreateTime:   time.Now(),
		UpdateTime:   time.Now(),
	}

	if err := config.DB.WithContext(ctx).Create(&book).Error; err != nil {
		return err
	}

	return nil
}

func (b *BookServiceImpl) SaveBookChapter(ctx context.Context, dto req.ChapterAddReqDto) error {
	// 1. 校验小说是否属于当前作者
	var bookInfo pojo.BookInfo
	if err := config.DB.WithContext(ctx).First(&bookInfo, dto.BookId).Error; err != nil {
		return err
	}
	//authorId := UserHolder.GetAuthorId(ctx) // 你自己实现
	//if bookInfo.AuthorID != authorId {
	//	return fmt.Errorf("无权限操作该小说")
	//}

	// 2. 查询当前最大章节号
	var maxChapter pojo.BookChapter
	err := config.DB.WithContext(ctx).
		Where("book_id = ?", dto.BookId).
		Order("chapter_num DESC").
		Limit(1).
		First(&maxChapter).Error
	chapterNum := 1
	if err == nil {
		chapterNum = maxChapter.ChapterNum + 1
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	// 3. 事务处理
	tx := config.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 4. 保存章节信息
	newChapter := pojo.BookChapter{
		BookID:      dto.BookId,
		ChapterName: dto.ChapterName,
		ChapterNum:  chapterNum,
		WordCount:   len(dto.ChapterContent),
		IsVip:       dto.IsVip,
		CreateTime:  time.Now(),
		UpdateTime:  time.Now(),
	}
	if err := tx.WithContext(ctx).Create(&newChapter).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 5. 保存章节内容
	newContent := pojo.BookContent{
		ChapterID:  newChapter.ID,
		Content:    dto.ChapterContent,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
	if err := tx.WithContext(ctx).Create(&newContent).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 6. 更新小说最新章节信息和总字数
	newWordCount := bookInfo.WordCount + newChapter.WordCount
	updateBook := pojo.BookInfo{
		ID:                    dto.BookId,
		LastChapterID:         newChapter.ID,
		LastChapterName:       newChapter.ChapterName,
		LastChapterUpdateTime: time.Now(),
		WordCount:             newWordCount,
	}
	if err := tx.WithContext(ctx).Model(&pojo.BookInfo{}).Where("id = ?", dto.BookId).Updates(updateBook).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	//// 7. 清除小说信息缓存（调用缓存管理器）
	//bookInfoCacheManager.EvictBookInfoCache(dto.BookID)
	//
	//// 8. 发送 MQ 消息通知小说信息更新
	//amqpMsgManager.SendBookChangeMsg(dto.BookID)

	return nil
}

func (b *BookServiceImpl) ListAuthorBooks(ctx context.Context, pageReq req.PageReqDto) (resp.PageRespDto[resp.BookInfoRespDto], error) {
	//todo
	return resp.PageRespDto[resp.BookInfoRespDto]{}, nil
}

func (b *BookServiceImpl) ListBookChapters(ctx context.Context, bookId string, pageReq req.PageReqDto) (resp.PageRespDto[resp.BookChapterRespDto], error) {
	var chapters []pojo.BookChapter

	offset := (pageReq.PageNum - 1) * pageReq.PageSize

	err := config.DB.WithContext(ctx).
		Where("book_id = ?", bookId).
		Order("chapter_num ASC").
		Offset(int(offset)).
		Limit(int(pageReq.PageSize)).
		Find(&chapters).Error
	if err != nil {
		return resp.PageRespDto[resp.BookChapterRespDto]{}, err
	}
	var count int64
	config.DB.Model(&pojo.BookChapter{}).Where("book_id = ?", bookId).Count(&count)
	var list []resp.BookChapterRespDto
	for _, c := range chapters {
		list = append(list, resp.BookChapterRespDto{
			ID:          c.ID,
			ChapterName: c.ChapterName,
			ChapterNum:  c.ChapterNum,
		})
	}
	return resp.PageRespDto[resp.BookChapterRespDto]{
		Total: count,
		List:  list,
	}, nil
}
func (b *BookServiceImpl) ListComments(ctx context.Context, userId int64, pageReq req.PageReqDto) (resp.PageRespDto[resp.UserCommentRespDto], error) {
	var comments []pojo.BookComment
	offset := (pageReq.PageNum - 1) * pageReq.PageSize
	// 1. 查询分页评论
	err := config.DB.WithContext(ctx).
		Where("user_id = ?", userId).
		Order("create_time DESC").
		Offset(int(offset)).
		Limit(int(pageReq.PageSize)).
		Find(&comments).Error
	if err != nil {
		return resp.PageRespDto[resp.UserCommentRespDto]{}, err
	}

	// 2. 查询评论总数
	var count int64
	err = config.DB.WithContext(ctx).
		Model(&pojo.BookComment{}).
		Where("user_id = ?", userId).
		Count(&count).Error
	if err != nil {
		return resp.PageRespDto[resp.UserCommentRespDto]{}, err
	}

	// 3. 收集 bookId 去批量查询书籍信息
	bookIdSet := make(map[int64]struct{})
	for _, c := range comments {
		bookIdSet[c.BookID] = struct{}{}
	}
	bookIds := make([]int64, 0, len(bookIdSet))
	for id := range bookIdSet {
		bookIds = append(bookIds, id)
	}

	var books []pojo.BookInfo
	if len(bookIds) > 0 {
		err = config.DB.WithContext(ctx).
			Where("id IN ?", bookIds).
			Find(&books).Error
		if err != nil {
			return resp.PageRespDto[resp.UserCommentRespDto]{}, err
		}
	}

	// 4. 构建 bookId -> 书籍信息 映射
	bookMap := make(map[int64]pojo.BookInfo, len(books))
	for _, book := range books {
		bookMap[book.ID] = book
	}

	// 5. 组装返回结果
	resultList := make([]resp.UserCommentRespDto, 0, len(comments))
	for _, c := range comments {
		bookName := ""
		bookPic := ""
		if b, ok := bookMap[c.BookID]; ok {
			bookName = b.BookName
			bookPic = b.PicURL
		}
		resultList = append(resultList, resp.UserCommentRespDto{
			CommentContent: c.CommentContent,
			CommentBook:    bookName,
			CommentBookPic: bookPic,
			CommentTime:    c.CreateTime,
		})
	}

	return resp.PageRespDto[resp.UserCommentRespDto]{
		PageNum:  pageReq.PageNum,
		PageSize: pageReq.PageSize,
		Total:    count,
		List:     resultList,
	}, nil
}

func (b *BookServiceImpl) DeleteBookChapter(ctx context.Context, chapterId string) error {
	tx := config.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 1. 查询章节信息
	var chapter pojo.BookChapter
	if err := tx.WithContext(ctx).First(&chapter, chapterId).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 2. 查询小说信息
	var bookInfo pojo.BookInfo
	if err := tx.WithContext(ctx).First(&bookInfo, chapter.BookID).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 3. 删除章节信息
	if err := tx.WithContext(ctx).Delete(&pojo.BookChapter{}, chapterId).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 4. 删除章节内容
	if err := tx.WithContext(ctx).Where("chapter_id = ?", chapterId).Delete(&pojo.BookContent{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 5. 更新小说信息
	bookInfo.WordCount -= chapter.WordCount
	bookInfo.UpdateTime = time.Now()

	// 如果删除的是最新章节，则更新最新章节信息
	if string(bookInfo.LastChapterID) == chapterId {
		var lastChapter pojo.BookChapter
		err := tx.WithContext(ctx).
			Where("book_id = ?", chapter.BookID).
			Order("chapter_num DESC").
			Limit(1).
			First(&lastChapter).Error
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			tx.Rollback()
			return err
		}

		if err == gorm.ErrRecordNotFound {
			// 没有剩余章节，重置最新章节信息
			bookInfo.LastChapterID = 0
			bookInfo.LastChapterName = ""
			bookInfo.LastChapterUpdateTime = time.Time{}
		} else {
			bookInfo.LastChapterID = lastChapter.ID
			bookInfo.LastChapterName = lastChapter.ChapterName
			bookInfo.LastChapterUpdateTime = lastChapter.UpdateTime
		}
	}

	if err := tx.WithContext(ctx).Save(&bookInfo).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 6. 清理缓存（假设有对应的缓存管理器方法）
	//bookChapterCacheManager.EvictBookChapterCache(chapterId)
	//bookContentCacheManager.EvictBookContentCache(chapterId)
	//bookInfoCacheManager.EvictBookInfoCache(chapter.BookID)
	//
	//// 7. 发送消息通知（假设有amqpMsgManager）
	//amqpMsgManager.SendBookChangeMsg(chapter.BookID)

	return tx.Commit().Error
}

func (b *BookServiceImpl) GetBookChapter(ctx context.Context, chapterId string) (resp.ChapterContentRespDto, error) {
	var chapter pojo.BookChapter
	if err := config.DB.WithContext(ctx).First(&chapter, chapterId).Error; err != nil {
		return resp.ChapterContentRespDto{}, err
	}
	var content pojo.BookContent
	if err := config.DB.WithContext(ctx).First(&content, chapterId).Error; err != nil {
		return resp.ChapterContentRespDto{}, err
	}
	return resp.ChapterContentRespDto{
		ChapterName:    chapter.ChapterName,
		ChapterContent: content.Content,
	}, nil
}

func (b *BookServiceImpl) UpdateBookChapter(ctx context.Context, chapterId string, dto req.ChapterUpdateReqDto) error {
	tx := config.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 1. 查询章节信息（假设有缓存管理器，先从数据库查）
	var chapter pojo.BookChapter
	if err := tx.WithContext(ctx).First(&chapter, chapterId).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 2. 查询小说信息
	var bookInfo pojo.BookInfo
	if err := tx.WithContext(ctx).First(&bookInfo, chapter.BookID).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 3. 更新章节信息
	newWordCount := len(dto.ChapterContent) // 章节内容长度
	i, _ := strconv.ParseInt(chapterId, 10, 64)
	updatedChapter := pojo.BookChapter{
		ID:          i,
		ChapterName: dto.ChapterName,
		WordCount:   newWordCount,
		IsVip:       dto.IsVip,
		UpdateTime:  time.Now(),
	}
	if err := tx.WithContext(ctx).Model(&pojo.BookChapter{}).
		Where("id = ?", chapterId).
		Updates(&updatedChapter).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 4. 更新章节内容
	updatedContent := pojo.BookContent{
		Content:    dto.ChapterContent,
		UpdateTime: time.Now(),
	}
	if err := tx.WithContext(ctx).
		Model(&pojo.BookContent{}).
		Where("chapter_id = ?", chapterId).
		Updates(&updatedContent).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 5. 更新小说信息的字数和最新章节信息（如果该章节是最新章节）
	bookInfo.WordCount = bookInfo.WordCount - chapter.WordCount + newWordCount
	bookInfo.UpdateTime = time.Now()
	id, err := strconv.ParseInt(chapterId, 10, 64)
	if err != nil {
		// 处理错误
	}
	if bookInfo.LastChapterID == id {
		bookInfo.LastChapterName = dto.ChapterName
		bookInfo.LastChapterUpdateTime = time.Now()
	}
	if err := tx.WithContext(ctx).Save(&bookInfo).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 6. 清理缓存（假设已实现对应方法）
	//bookChapterCacheManager.EvictBookChapterCache(chapterId)
	//bookContentCacheManager.EvictBookContentCache(chapterId)
	//bookInfoCacheManager.EvictBookInfoCache(chapter.BookID)
	//
	//// 7. 发送消息通知
	//amqpMsgManager.SendBookChangeMsg(chapter.BookID)

	return tx.Commit().Error
}
