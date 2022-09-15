## Init Packages

go mod init github/com/Naithar01/{name...}

## Init Library

go get -u {Library url}

# How to use Query

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
