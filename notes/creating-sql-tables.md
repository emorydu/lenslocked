# Creating SQL Tables

```sql
CREATE TABLE table_name (
    field_name TYPE CONSTRAINTS,
    field_name TYPE(args) CONSTRAINTS
);
```

```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE
);
```