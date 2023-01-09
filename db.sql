CREATE DATABASE IF NOT EXISTS shmily;

USE shmily;

#用户信息表(user)
CREATE TABLE IF NOT EXISTS user(
    password varchar(20) null,
    account varchar(30) null,
    nickname varchar(20) null,
    avatar varchar(255) null,
    birthday varchar(20) null,
    personalized_signature varchar(255) null,
    gender varchar(10) null,
    PRIMARY KEY (account)
);

#好友表(friend)
create table tbl_friend(
    `name` varchar(20) null,
    `gender` varchar(10) null,
    `relation` varchar(10) null
)

#纸条表(paper)
create table tbl_paper(

)

#邮局表(post_office)
create table tbl_post_office(

)

#纪念日表(anniversary)
create table tbl_anniversary(

)
