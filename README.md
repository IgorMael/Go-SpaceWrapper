# Go-SpaceWrapper
Golang Wrapper to OCR Space Api

## What it does?

This package adds a high level wrapper to the [SpaceOCR api](https://ocr.space/).

```go
spacewrapper.Init("d16ae8619488957");
newOcr, err := spacewrapper.Get(imageURL, spacewrapper.Params{})
if err != nil {
  fmt.Println(err)
}
fmt.Println(newOcr.ParsedResults[0].TextOverlay)
```


## Installation and Usage

### Installing

```
go get github.com/IgorMael/Go-SpaceWrapper
```

```go
import (
	"github.com/IgorMael/Go-SpaceWrapper"
)
```

### Usage
The wrapper allow you to Ocr images using get, post, or base64. 

#### Initializating
First you need get the API key, you can get one here on [Space api](https://ocr.space/ocrapi), the key will be sended to your e-mail
```go
spacewrapper.Init(ApiKey)
```

#### Using GET
The imageURL it's a string to an image file.
```go
func Get(imageURL string, args Params) (ProcessedDoc, error)
```
There are limitations to get info. Consult Space OCR api for informations about it.

#### Using POST 
You can make a POST request using a file or an url.
```go
func PostFile(paramTexts Params, paramFile File) (ProcessedDoc, error)
```
```go
func PostURL(paramTexts Params, URL string) (ProcessedDoc, error)
```

#### Using Base64
```go
func PostBase64(paramTexts Params, base64 string) (ProcessedDoc, error)
```

### Types
#### ProcessedDoc
```go
type ProcessedDoc struct {
	ParsedResults []struct {
		TextOverlay struct {
			Lines      []interface{} `json:"Lines"`
			HasOverlay bool          `json:"HasOverlay"`
			Message    string        `json:"Message"`
		} `json:"TextOverlay"`
		TextOrientation   string `json:"TextOrientation"`
		FileParseExitCode int    `json:"FileParseExitCode"`
		ParsedText        string `json:"ParsedText"`
		ErrorMessage      string `json:"ErrorMessage"`
		ErrorDetails      string `json:"ErrorDetails"`
	} `json:"ParsedResults"`
	OCRExitCode                  int    `json:"OCRExitCode"`
	IsErroredOnProcessing        bool   `json:"IsErroredOnProcessing"`
	ProcessingTimeInMilliseconds string `json:"ProcessingTimeInMilliseconds"`
	SearchablePDFURL             string `json:"SearchablePDFURL"`
}
```
#### File
```go
type File struct {
	FileName string //test.jpg
	Content  []byte //[]byte
}
```

#### Params
```go
type Params map[string]string
```
