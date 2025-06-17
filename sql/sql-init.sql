create database if not exists `novel`;
-- auto-generated definition
create table book_info
(
    id                       bigint unsigned auto_increment comment '主键'
        primary key,
    work_direction           tinyint unsigned               null comment '作品方向;0-男频 1-女频',
    category_id              bigint unsigned                null comment '类别ID',
    category_name            varchar(50)                    null comment '类别名',
    pic_url                  varchar(200)                   not null comment '小说封面地址',
    book_name                varchar(50)                    not null comment '小说名',
    author_id                bigint unsigned                not null comment '作家id',
    author_name              varchar(50)                    not null comment '作家名',
    book_desc                varchar(2000)                  not null comment '书籍描述',
    score                    tinyint unsigned               not null comment '评分;总分:10 ，真实评分 = score/10',
    book_status              tinyint unsigned default '0'   not null comment '书籍状态;0-连载中 1-已完结',
    visit_count              bigint unsigned  default '103' not null comment '点击量',
    word_count               int unsigned     default '0'   not null comment '总字数',
    comment_count            int unsigned     default '0'   not null comment '评论数',
    last_chapter_id          bigint unsigned                null comment '最新章节ID',
    last_chapter_name        varchar(50)                    null comment '最新章节名',
    last_chapter_update_time datetime                       null comment '最新章节更新时间',
    is_vip                   tinyint unsigned default '0'   not null comment '是否收费;1-收费 0-免费',
    create_time              datetime                       null comment '创建时间',
    update_time              datetime                       null comment '更新时间',
    constraint pk_id
        unique (id),
    constraint uk_bookName_authorName
        unique (book_name, author_name)
)
    comment '小说信息';

create index idx_createTime
    on book_info (create_time);

create index idx_lastChapterUpdateTime
    on book_info (last_chapter_update_time);

-- auto-generated definition
create table user_info
(
    id              bigint unsigned auto_increment
        primary key,
    username        varchar(50)                  not null comment '登录名',
    password        varchar(100)                 not null comment '登录密码-加密',
    salt            varchar(8)                   not null comment '加密盐值',
    nick_name       varchar(50)                  null comment '昵称',
    user_photo      varchar(255)                 null comment '用户头像',
    user_sex        tinyint unsigned             null comment '用户性别;0-男 1-女',
    account_balance bigint unsigned  default '0' not null comment '账户余额',
    status          tinyint unsigned default '0' not null comment '用户状态;0-正常',
    create_time     datetime                     null comment '创建时间',
    update_time     datetime                     null comment '更新时间',
    constraint pk_id
        unique (id),
    constraint uk_username
        unique (username)
)
    comment '用户信息';

-- auto-generated definition
create table home_book
(
    id          bigint unsigned auto_increment
        primary key,
    type        tinyint unsigned not null comment '推荐类型;0-轮播图 1-顶部栏 2-本周强推 3-热门推荐 4-精品推荐',
    sort        tinyint unsigned not null comment '推荐排序',
    book_id     bigint unsigned  not null comment '推荐小说ID',
    create_time datetime         null comment '创建时间',
    update_time datetime         null comment '更新时间',
    constraint pk_id
        unique (id)
)
    comment '小说推荐';

-- auto-generated definition
create table home_friend_link
(
    id          bigint unsigned auto_increment
        primary key,
    link_name   varchar(50)                   not null comment '链接名',
    link_url    varchar(100)                  not null comment '链接url',
    sort        tinyint unsigned default '11' not null comment '排序号',
    is_open     tinyint unsigned default '1'  not null comment '是否开启;0-不开启 1-开启',
    create_time datetime                      null comment '创建时间',
    update_time datetime                      null comment '更新时间',
    constraint pk_id
        unique (id)
)
    comment '友情链接';

