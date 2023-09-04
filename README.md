# go-todos
A cli todo app with go using cobra, gorm and table

```sh
//commands help
todos -h
```

To make a new todo
```sh
todos new [Name] [Other name]
```
To view a todo
```sh
todos view

//By stated 
todos view --state="completed"
todos view --state="not completed"

//By id
todos view --id=1
```

To delete all completed todos
```sh
todos delete

//By id
todos delete --id=1
```
To toggle todo state
```sh
todos toggle [Id]
```
