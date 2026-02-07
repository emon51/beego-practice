## Project Setup
### 1. Create a project structure: 
```
bee new [appName]
```
```
cd [appName]
```
```
go mod tidy
```

### 2. Run the project
```
bee run
```

## Models
In Beego, a model:
- Represents a table in the database
- Maps Go structs → DB rows
- Is handled by Beego ORM
- One struct == one table
