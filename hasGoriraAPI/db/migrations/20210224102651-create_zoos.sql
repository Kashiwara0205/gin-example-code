
-- +migrate Up
CREATE TABLE IF NOT EXISTS zoos (
    id int(15) AUTO_INCREMENT,
    pref_id int(15),
    name varchar(255),
    has_gorira tinyint(1),
    created_at datetime  default current_timestamp,
    updated_at timestamp default current_timestamp on update current_timestamp,
    PRIMARY KEY (id),
    FOREIGN KEY (pref_id) REFERENCES prefectures(id)
);

-- +migrate Down
DROP TABLE IF EXISTS zoos;