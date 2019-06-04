# FileUpload 📄
### Very simple file uploader


## Usage
### Upload
```
curl -X PUT http://localhost:8000/api/file/u -F file=@logo.png
```

### Download
```
curl http://localhost:8000/api/file/d/e4aef93c-f2f0-4c31-9bfb-0d01851fa077 -o logo.png
```



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
