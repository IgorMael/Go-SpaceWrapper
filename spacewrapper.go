/*
Package spacewrapper implements
an wrapper to access ocr.space OCR service
*/
package spacewrapper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"strings"
)

type config struct {
	apikey string
}

/*
File it's a file
*/
type File struct {
	FileName string //test.jpg
	Content  []byte //[]byte
}

/*
Params are the request options alias
*/
type Params map[string]string

/*
DefaultConfig API configuration
*/
var DefaultConfig config

/*
ProcessedDoc returns an processed OCR'ed document
*/
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

/*
Init creates a default configuration file
*/
func Init(key string) {
	DefaultConfig.apikey = key

}

func buildForm(paramTexts Params, file File) (*bytes.Buffer, string, error) {
	paramTexts["apikey"] = DefaultConfig.apikey
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)
	for k, v := range paramTexts {
		bodyWriter.WriteField(k, v)
	}
	if file.FileName != "" {
		fileWriter, err := bodyWriter.CreateFormFile("file", file.FileName)
		if err != nil {
			return nil, "", err
		}
		fileWriter.Write(file.Content)
	}
	bodyWriter.Close()
	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()
	return bodyBuf, contentType, nil
}

/*
Get makes a requisition with the informed URL
*/
func Get(imageURL string, args Params) (ProcessedDoc, error) {
	argString := ""
	doc := ProcessedDoc{}
	baseURL := fmt.Sprintf("https://api.ocr.space/parse/imageurl?apikey=%s&url=%s", DefaultConfig.apikey, imageURL)
	for k, v := range args {
		argString = strings.Join([]string{argString, fmt.Sprintf("&%s=%s", k, v)}, "")
	}
	resp, err := http.Get(strings.Join([]string{baseURL, argString}, ""))
	if err != nil {
		return ProcessedDoc{}, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&doc)
	if err != nil {
		return ProcessedDoc{}, err
	}
	return doc, nil
}

/*
PostFile submits a file
*/
func PostFile(paramTexts Params, paramFile File) (ProcessedDoc, error) {
	paramTexts["apikey"] = DefaultConfig.apikey
	doc := ProcessedDoc{}
	url := "https://api.ocr.space/parse/image"
	form, contentType, err := buildForm(paramTexts, paramFile)

	resp, err := http.Post(url, contentType, form)
	if err != nil {
		return doc, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		b, _ := ioutil.ReadAll(resp.Body)
		return doc, fmt.Errorf("[%d %s]%s", resp.StatusCode, resp.Status, string(b))
	}

	err = json.NewDecoder(resp.Body).Decode(&doc)
	if err != nil {
		return ProcessedDoc{}, err
	}
	return doc, nil
}

/*
PostBase64 submits a file
*/
func PostBase64(paramTexts Params, base64 string) (ProcessedDoc, error) {
	paramTexts["base64Image"] = base64
	doc := ProcessedDoc{}
	url := "https://api.ocr.space/parse/image"
	form, contentType, err := buildForm(paramTexts, File{})
	resp, err := http.Post(url, contentType, form)
	if err != nil {
		return doc, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		b, _ := ioutil.ReadAll(resp.Body)
		return doc, fmt.Errorf("[%d %s]%s", resp.StatusCode, resp.Status, string(b))
	}
	err = json.NewDecoder(resp.Body).Decode(&doc)
	if err != nil {
		return ProcessedDoc{}, err
	}
	return doc, nil
}

/*
PostURL submits a url
*/
func PostURL(paramTexts Params, URL string) (ProcessedDoc, error) {
	paramTexts["apikey"] = DefaultConfig.apikey
	paramTexts["url"] = URL
	doc := ProcessedDoc{}
	url := "https://api.ocr.space/parse/image"
	form, contentType, err := buildForm(paramTexts, File{})
	resp, err := http.Post(url, contentType, form)
	if err != nil {
		return doc, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		b, _ := ioutil.ReadAll(resp.Body)
		return doc, fmt.Errorf("[%d %s]%s", resp.StatusCode, resp.Status, string(b))
	}
	err = json.NewDecoder(resp.Body).Decode(&doc)
	if err != nil {
		return ProcessedDoc{}, err
	}
	return doc, nil
}
