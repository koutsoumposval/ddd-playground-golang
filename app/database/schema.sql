use polyglot_ddd_product;

CREATE TABLE IF NOT EXISTS product (
  `id`          int          NOT NULL AUTO_INCREMENT,
  `name`        varchar(256) NOT NULL,
  `category_id` int          NOT NULL,
  `created_at`  DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at`  DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
