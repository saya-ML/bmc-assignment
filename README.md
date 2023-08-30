# bmc-assignment

Swagger
http://localhost:8080/docs/index.html

Passenger by key
http://localhost:8080/passenger/11

Passegner by key and filter
http://localhost:8080/passenger/25?fields=passenger_id&fields=fare

All the passengers
http://localhost:8080/passenger

Histogram
http://localhost:8080/histogram


Docker Environment:

      - SOURCE:   "CSV" or "sqlite3"
      - CSV_FILE: "/app/data/titanic.csv"
      - SQL_FILE: "/app/data/passengers.db"

