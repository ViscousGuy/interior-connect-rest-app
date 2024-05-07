# Database Schema

This document outlines the database schema including tables, their columns, constraints, and relationships. It also includes sample data insertions for initial testing and Go struct definitions for ORM mapping.

## Tables

```sql
CREATE TABLE `material` (
  `id` INT  AUTO_INCREMENT,
  `name` VARCHAR(50),
  `display` BOOLEAN,
  PRIMARY KEY (`id`)
);

CREATE TABLE `contractor` (
  `id` INT AUTO_INCREMENT,
  `firstname` VARCHAR(255),
  `lastname` VARCHAR(255),
  `city` VARCHAR(50),
  `state` VARCHAR(50),
  `mobile` VARCHAR(15),
  `email` VARCHAR(255) UNIQUE,
  `slug` VARCHAR(255),
  `pincode` VARCHAR(10),
  `verified` BOOLEAN,
  `active` BOOLEAN,
  `display` BOOLEAN,
  PRIMARY KEY (`id`)
);

CREATE TABLE `furniture_type` (
  `id` INT AUTO_INCREMENT,
  `name` VARCHAR(20),
  `slug` VARCHAR(255),
  `display` BOOLEAN,
  PRIMARY KEY (`id`)
);

CREATE TABLE `room_type` (
  `id` INT AUTO_INCREMENT,
  `name` VARCHAR(50),
  `slug` VARCHAR(255),
  `display` BOOLEAN,
  PRIMARY KEY (`id`)
);

CREATE TABLE `color` (
  `id` INT AUTO_INCREMENT,
  `name` VARCHAR(50),
  `display` BOOLEAN,
  PRIMARY KEY (`id`)
);

CREATE TABLE `furniture` (
  `id` INT AUTO_INCREMENT,
  `furniture_type_id` INT,
  `room_type_id` INT,
  `name` VARCHAR(100),
  `description` TEXT,
  `dimensions` VARCHAR(50)  ,
  `price` DECIMAL(10,2),
  `contractor_id` INT,
  `slug` VARCHAR(255),
  `display` BOOLEAN,
  PRIMARY KEY (`id`),
  FOREIGN KEY (`contractor_id`) REFERENCES `contractor`(`id`),
  FOREIGN KEY (`furniture_type_id`) REFERENCES `furniture_type`(`id`),
  FOREIGN KEY (`room_type_id`) REFERENCES `room_type`(`id`)
);

CREATE TABLE `furniture_color` (
  `id` INT AUTO_INCREMENT,
  `furniture_id` INT,
  `color_id` INT,
  PRIMARY KEY (`id`),
  FOREIGN KEY (`furniture_id`) REFERENCES `furniture`(`id`),
  FOREIGN KEY (`color_id`) REFERENCES `color`(`id`)
);

CREATE TABLE `project` (
  `id` INT AUTO_INCREMENT,
  `contractor_id` INT,
  `project_name` VARCHAR(255),
  `description` VARCHAR(255),
  `city` VARCHAR(50),
  `slug` VARCHAR(255),
  `display` BOOLEAN,
  PRIMARY KEY (`id`),
  FOREIGN KEY (`contractor_id`) REFERENCES `contractor`(`id`)
);

CREATE TABLE `project_image` (
  `id` INT AUTO_INCREMENT,
  `project_id` INT,
  `image_path` VARCHAR(255),
  `display` BOOLEAN,
  PRIMARY KEY (`id`),
  FOREIGN KEY (`project_id`) REFERENCES `project`(`id`)
);

CREATE TABLE `furniture_material` (
  `id` INT AUTO_INCREMENT,
  `furniture_id` INT,
  `material_id` INT,
  PRIMARY KEY (`id`),
  FOREIGN KEY (`material_id`) REFERENCES `material`(`id`),
  FOREIGN KEY (`furniture_id`) REFERENCES `furniture`(`id`)
);

CREATE TABLE `furniture_image` (
  `id` INT AUTO_INCREMENT,
  `image_path` VARCHAR(255),
  `furniture_id` INT,
  PRIMARY KEY (`id`),
  FOREIGN KEY (`furniture_id`) REFERENCES `furniture`(`id`)
);
```

## Insert Data

```sql
INSERT INTO `material` (`name`, `display`) VALUES ('Wood', TRUE);
INSERT INTO `material` (`name`, `display`) VALUES ('Metal', TRUE);

INSERT INTO `contractor` (`firstname`, `lastname`, `city`, `state`, `mobile`, `email`, `slug`, `pincode`, `verified`, `active`, `display`) VALUES ('John', 'Doe', 'New York', 'NY', '1234567890', 'john.doe@example.com', 'john-doe', '10001', TRUE, TRUE, TRUE);

INSERT INTO `furniture_type` (`name`, `slug`, `display`) VALUES ('Chair', 'chair', TRUE);
INSERT INTO `furniture_type` (`name`, `slug`, `display`) VALUES ('Table', 'table', TRUE);

INSERT INTO `room_type` (`name`, `slug`, `display`) VALUES ('Living Room', 'living-room', TRUE);
INSERT INTO `room_type` (`name`, `slug`, `display`) VALUES ('Bedroom', 'bedroom', TRUE);

INSERT INTO `color` (`name`, `display`) VALUES ('Red', TRUE);
INSERT INTO `color` (`name`, `display`) VALUES ('Blue', TRUE);

INSERT INTO `furniture` (`furniture_type_id`, `room_type_id`, `name`, `description`, `dimensions`, `price`, `contractor_id`, `slug`, `display`) VALUES (1, 1, 'Wooden Chair', 'A comfortable wooden chair.', '20x20x30', 49.99, 1, 'wooden-chair', TRUE);

INSERT INTO `furniture_color` (`furniture_id`, `color_id`) VALUES (1, 1);
INSERT INTO `furniture_color` (`furniture_id`, `color_id`) VALUES (1, 2);

INSERT INTO `project` (`contractor_id`, `project_name`, `description`, `city`, `slug`, `display`) VALUES (1, 'Home Renovation', 'Complete home renovation project.', 'New York', 'home-renovation', TRUE);

INSERT INTO `project_image` (`project_id`, `image_path`, `display`) VALUES (1, '/path/to/image.jpg', TRUE);

INSERT INTO `furniture_material` (`furniture_id`, `material_id`) VALUES (1, 1);

INSERT INTO `furniture_image` (`furniture_id`, `image_path`) VALUES (1, '/path/to/image.jpg');

```

## Go Struct Definitions

```
type Material struct {
    Id      int     `orm:"auto"`
    Name    string  `orm:"size(50)"`
    Display bool
}

type Contractor struct {
    Id        int     `orm:"auto"`
    FirstName string  `orm:"size(255)"`
    LastName  string  `orm:"size(255)"`
    City      string  `orm:"size(50)"`
    State     string  `orm:"size(50)"`
    Mobile    string  `orm:"size(15)"`
    Email     string  `orm:"size(255);unique"`
    Slug      string  `orm:"size(255)"`
    Pincode   string  `orm:"size(10)"`
    Verified  bool
    Active    bool
    Display   bool
}

type FurnitureType struct {
    Id      int     `orm:"auto"`
    Name    string  `orm:"size(20)"`
    Slug    string  `orm:"size(255)"`
    Display bool
}

type RoomType struct {
    Id      int     `orm:"auto"`
    Name    string  `orm:"size(50)"`
    Slug    string  `orm:"size(255)"`
    Display bool
}

type Color struct {
    Id      int     `orm:"auto"`
    Name    string  `orm:"size(50)"`
    Display bool
}

type Furniture struct {
    Id               int     `orm:"auto"`
    FurnitureTypeId  int
    RoomTypeId       int
    Name             string  `orm:"size(100)"`
    Description      string  `orm:"type(text)"`
    Dimensions       string  `orm:"size(50)"`
    Price            float64 `orm:"digits(10);decimals(2)"`
    ContractorId     int
    Slug             string  `orm:"size(255)"`
    Display          bool
}

type FurnitureColor struct {
    Id           int `orm:"auto"`
    FurnitureId  int
    ColorId      int
}

type Project struct {
    Id            int     `orm:"auto"`
    ContractorId  int
    ProjectName   string  `orm:"size(255)"`
    Description   string  `orm:"size(255)"`
    City          string  `orm:"size(50)"`
    Slug          string  `orm:"size(255)"`
    Display       bool
}

type ProjectImage struct {
    Id         int     `orm:"auto"`
    ProjectId  int
    ImagePath  string  `orm:"size(255)"`
    Display    bool
}

type FurnitureMaterial struct {
    Id           int `orm:"auto"`
    FurnitureId  int
    MaterialId   int
}

type FurnitureImage struct {
    Id          int     `orm:"auto"`
    ImagePath   string  `orm:"size(255)"`
    FurnitureId int
}
```

## Initialization Function

```
func init() {
    orm.RegisterModel(new(Material), new(Contractor), new(FurnitureType), new(RoomType), new(Color), new(Furniture), new(FurnitureColor), new(Project), new(ProjectImage), new(FurnitureMaterial), new(FurnitureImage))
}

```
