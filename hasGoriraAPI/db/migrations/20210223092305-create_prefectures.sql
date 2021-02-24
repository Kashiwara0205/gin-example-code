
-- +migrate Up
CREATE TABLE IF NOT EXISTS prefectures (
    id int(15) AUTO_INCREMENT,
    name varchar(255),
    created_at datetime  default current_timestamp,
    updated_at timestamp default current_timestamp on update current_timestamp,
    PRIMARY KEY (id));

-- +migrate Down
DROP TABLE IF EXISTS prefectures;