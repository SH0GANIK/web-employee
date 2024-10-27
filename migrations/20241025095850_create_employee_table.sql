-- +goose Up
-- +goose StatementBegin
CREATE TABLE employees (
    Id SERIAL PRIMARY KEY,
    Name VARCHAR(255),
    Surname VARCHAR(255),
    Phone VARCHAR(50),
    CompanyId INT,
    Passport_Type VARCHAR(50),
    Passport_Number VARCHAR(50),
    Department_Name VARCHAR(255),
    Department_Phone VARCHAR(50)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS employees;
-- +goose StatementEnd
