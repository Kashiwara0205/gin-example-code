
-- +migrate Up
CREATE TABLE IF NOT EXISTS prefectures (
    id int(15) AUTO_INCREMENT,
    name varchar(255),
    PRIMARY KEY (id));

-- +migrate Down
DROP TABLE IF EXISTS prefectures;