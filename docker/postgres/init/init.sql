CREATE TABLE sample_table (
    id SERIAL PRIMARY KEY,
    varchar_column VARCHAR(255),
    integer_column INT,
    bigint_column BIGINT,
    numeric_column NUMERIC(10, 2),
    boolean_column BOOLEAN,
    date_column DATE,
    timestamp_column TIMESTAMP,
    text_column TEXT
);

INSERT INTO sample_table (varchar_column, integer_column, bigint_column, numeric_column, boolean_column, date_column, timestamp_column, text_column)
VALUES
    ('John Doe', 25, 1234567890, 123.45, true, '2022-01-01', '2022-01-01 12:34:56', 'Sample text 1'),
    ('Jane Smith', 30, 9876543210, 456.78, false, '2022-02-15', '2022-02-15 18:45:30', 'Sample text 2'),
    ('Bob Johnson', 40, 111222333444, 789.12, true, '2022-03-20', '2022-03-20 09:15:00', 'Sample text 3'),
    ('Alice Brown', 22, 555666777888, 234.56, false, '2022-04-10', '2022-04-10 21:00:45', 'Sample text 4'),
    ('Charlie Davis', 35, 999888777666, 678.90, true, '2022-05-05', '2022-05-05 15:30:20', 'Sample text 5');