package DataBases

import (
	"database/sql"
	"fmt"
	_ "github.com/snowflakedb/gosnowflake"
	"log"
)

func SnowflakeFetch() {
	// Set up the connection string with schema
	dsn := fmt.Sprintf("%s:%s@%s/%s?warehouse=%s&schema=%s",
		"thinkboxllc",            // replace with your Snowflake username
		"nQhsN@1ALXl@GCne$8t1O9", // replace with your Snowflake password
		"glfchfw-xr18961",        // replace with your Snowflake account identifier
		"FIGHTERSDB",             // replace with your Snowflake database name
		"COMPUTE_WH",             // replace with your Snowflake warehouse name
		"FIGHTERS",               // replace with your Snowflake schema name
	)

	// Open a connection to the Snowflake database
	db, err := sql.Open("snowflake", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Ping the database to verify the connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// Query the database
	rows, err := db.Query("SELECT * FROM FIGHTERS TABLESAMPLE BERNOULLI (1) LIMIT 1") // replace with your table name
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Iterate over the rows and print the results
	for rows.Next() {
		var name string
		// Add more variables as per your table schema

		err = rows.Scan(&name) // Adjust based on your table schema
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Fighter: %s\n", name)
	}

	// Check for errors from iterating over rows
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}
