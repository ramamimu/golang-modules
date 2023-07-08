### **Add a Module**

Go lang is well integrated with Git. Use **tag** git to update and release module. The tag have to use **three dot**.

```go
go mod init <name_module>
```

### **Add dependency**

```go
go get <name_module>
```

### **List version**

```go
go list -m -versions <path_github>
```

### **Upgrade Module**

To upgrade module, we just make a new tag in Git

### **Upgrade Dependency**

To update to the new version. we possible to write in go.mod file then change the version to the spesific tag. Then give command.

```go
go get
```
