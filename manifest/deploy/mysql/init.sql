create database if not exists heygem;

use heygem;
-- auto-generated definition
create table if not exists model
(
    id         bigint auto_increment comment '主键'
        primary key,
    name       varchar(50) not null comment '模型名称',
    video_path text        null comment '视频地址',
    audio_path text        null comment '音频地址',
    voice_id   bigint      null comment '音频ID',
    created_at datetime    null comment '创建时间'
)
    comment '模特';

-- auto-generated definition
create table if not exists video
(
    id           bigint auto_increment comment '主键'
        primary key,
    file_path    text                    null comment '文件地址',
    status       varchar(20)             null comment '状态',
    message      text                    null comment '消息',
    model_id     bigint                  null comment '模特ID',
    audio_path   text                    null comment '音频地址',
    param        text                    null comment '参数',
    code         varchar(50)             null comment '编码',
    created_at   datetime                null comment '创建时间',
    progress     varchar(20) default '0' null comment '进度',
    name         varchar(50)             null comment '名称',
    duration     bigint                  null comment '视频时长',
    text_content text                    null comment '文本内容',
    voice_id     bigint                  null comment '语音ID'
)
    comment '视频';


-- auto-generated definition
create table if not exists voice
(
    id         bigint auto_increment comment '主键'
        primary key,
    name       varchar(50) not null comment '名称',
    audio_path text        null comment '语音路径',
    audio_text text        null comment '语音文本',
    created_at datetime    null comment '创建时间'
)
    comment '语音';

