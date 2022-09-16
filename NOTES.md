## Init Packages

go mod init github/com/Naithar01/{name...}

## Init Library

go get -u {Library url}

# How to use Fiber Query

## Create Struct

### First String must be capper

```
type Person struct {
Name string `query:"name"`
Pass string `query:"pass"`
Products []string `query:"products"`
}
```

## Create New Struct Keyword & use fiber Context QueryBarser & log querys

```
category_query := new(Person)

	if err := c.QueryParser(category_query); err != nil {
		return err
	}

	fmt.Println(category_query)

    // localhost:4000/hello?name=james&pass=good&products=apple&products=banana
    -- ( &{james good [apple banana]} )
```

# How to Search Data by Query

```
	db.Find(&user, "id = ?", "1b74413f-f3b8-409f-ac47-e8c062e3472a")
	// "&user" is save variable
	var user models.User
	// user is variable name, models.User is entitis, struct

	// mysql query
	// SELECT * FROM users WHERE id = "1b74413f-f3b8-409f-ac47-e8c062e3472a";

```

# How to Add ( Array ) Column

```
	Tags          pq.StringArray `json:"tags" gorm:"type:text"`
```

## if Data Type is Integer

```
	Tags          pq.Int64Array `json:"tags" gorm:"type:integer"`
```
