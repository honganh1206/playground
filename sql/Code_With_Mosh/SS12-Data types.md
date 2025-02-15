# Data types

- String/Numeric/Date and Time/Blob/Spatial (Geographical values)

## Strings

- `CHAR(x)` - Fixed length
- `VARCHAR(x)` - Max 65,646 characters
- `MEDIUMTEXT` - Max 16mb => JSON
- `LONGTEXT` - Max 4gb => Textbook/Large log files
- `TINYTEXT` - Max 255 bytes
- `TEXT` Max 64kb

> [!tip] Be consistent!
> 
> - `VARCHAR(50)` for short strings
> - `VARCHAR(255)` for medium-length strings


## Bytes

English: 1 byte / European or Middle-eastern: 2 bytes / Asian: 3 bytes


## Integer

- `TINYINT`: 1 byte `[-128, 127]`
- `UNSIGNED TINYINT`: `[0, 255]`
… and some more


### Zerofill

`INT(4)` => `0001`


> [!tip]
> 
> Use the smallest data type that suits your needs.

### Rationals

- `DECIMAL(p, s)`: `DECIMAL(9, 2)` => 1234567.89
- `DOUBLE` 4 bytes / `FLOAT` 8 bytes

## Booleans

```sql
SET is_published = 1 -- True
```

## Enums

`ENUM('small','large')` => Not really a good practice to use `ENUM`

## Date/Time

`DATE/TIME/DATETIME/TIMESTAMP/YEAR`

> [!warning]
> 
> `TIMESTAMP` only allows 4 bytes (up to 2038)

## Blobs

`TINYBLOB (255b) / BLOB (65kb) / MEDIUMBLOB / LONGBLOB`


> [!tip]
> 
> Not a good practice to store blobs inside a relational database.


## JSON

```sql
UPDATE sql_store.products p
SET p.properties = '
{
	"dimensions": [1,2,3],
	"weight": 10,
	"manufacturer": 
	{
		"name": "sony"
	} 
}
'
-- Or we use native functions
SET p.properties = JSON_OBJECT('weight', 10, 'dimensions', JSON_ARRAY(1,2,3), 'manufacturer', JSON_OBJECT('name', 'abc'))
WHERE p.product_id = 1;

-- Another way to update
UPDATE sql_store.products p
SET p.properties = JSON_SET(properties, '$.weight', 20) 
WHERE p.product_id = 1;

-- Deserialize the JSON obj
SELECT product_id, JSON_EXTRACT(properties, '$.weight') AS weight FROM sql_store.products p where product_id = 1;



```