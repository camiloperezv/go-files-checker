# Proof of concept 
## how to detect type of files on golang 

This proof of concept uses the free library github.com/h2non/filetype.

## Download
Clone this repo on yoyr go path src

## Run 
`go run main.go` 

## The test
In the `files` folder we have this files
- bin :heavy_multiplication_x: 
- bin.doc :heavy_multiplication_x: 
- bin.pdf :heavy_multiplication_x: 
- bin.png :heavy_multiplication_x: 
- doc.docx :heavy_check_mark:
- doc.pdf :heavy_check_mark:
- landscape.jpeg :heavy_check_mark:
- landscape.png :heavy_check_mark:
- thing.html :heavy_multiplication_x: 
- thing.html.png :heavy_multiplication_x: 

## Output
```
Sorry. Unable to upload: files/bin Unknown unknown 
---------------------------
Sorry. Unable to upload: files/bin.doc Unknown unknown 
---------------------------
Sorry. Unable to upload: files/bin.pdf Unknown unknown 
---------------------------
Sorry. Unable to upload: files/bin.png Unknown unknown 
---------------------------
We can trust this file: files/doc.docx
 Extension: docx
 Mimetype: application/vnd.openxmlformats-officedocument.wordprocessingml.document
---------------------------
We can trust this file: files/doc.pdf
 Extension: pdf
 Mimetype: application/pdf
---------------------------
We can trust this file: files/landscape.jpeg
 Extension: jpg
 Mimetype: image/jpeg
---------------------------
We can trust this file: files/landscape.png
 Extension: png
 Mimetype: image/png
---------------------------
Sorry. Unable to upload: files/thing.html Unknown unknown 
---------------------------
Sorry. Unable to upload: files/thing.html.png Unknown unknown 
```