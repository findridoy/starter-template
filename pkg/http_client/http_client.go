package httpclient

import (
	"bytes"
	"context"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

func NewFormRequest() (*FormRequest, error) {
	fr := new(FormRequest)

	fr.body = &bytes.Buffer{}
	fr.multipartWriter = multipart.NewWriter(fr.body)
	fr.headers = make(map[string]string)

	return fr, nil
}

type FormRequest struct {
	body            *bytes.Buffer
	multipartWriter *multipart.Writer
	headers         map[string]string
}

func (fr *FormRequest) AttachFile(key, filePath string) (httpClient *FormRequest, err error) {
	httpClient = fr

	file, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer file.Close()

	formFile, err := fr.multipartWriter.CreateFormFile(key, filepath.Base(file.Name()))
	if err != nil {
		return
	}

	_, err = io.Copy(formFile, file)
	if err != nil {
		return
	}

	return
}

func (fr *FormRequest) AppendItem(key, value string) (*FormRequest, error) {
	err := fr.multipartWriter.WriteField(key, value)
	if err != nil {
		return nil, err
	}
	return fr, nil
}

func (fr *FormRequest) SendRequest(ctx context.Context, url, method string) (*http.Response, error) {
	err := fr.multipartWriter.Close()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, method, url, fr.body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", fr.multipartWriter.FormDataContentType())

	for k, v := range fr.headers {
		req.Header.Set(k, v)
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (fr *FormRequest) SetHeader(key, value string) {
	fr.headers[key] = value
}
