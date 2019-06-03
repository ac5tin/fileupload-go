# FileUpload 📄
### Very simple file uploader



## Build
### .env file
```
S3_REGION = ""
S3_BUCKET = ""
AWS_ACCESS_KEY = ""
AWS_SECRET_KEY = ""
REDIS_ADDR = "localhost:6379"
REDIS_PW = ""
```
### cd to project root and run
```
go build -o bin/fileupload
```

## Run built binary
```
./bin/fileupload -addr :3000
```