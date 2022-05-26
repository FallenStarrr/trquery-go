package main

import (
	"bytes"
	_"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	_"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	_"path/filepath"
	"strconv"
	_"strings"
	_"time"
	"database/sql"
	_ "github.com/lib/pq"
)


const signedToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2NvdW50X3R5cGUiOiJwYmI3dmZwQU1nM1AycG1vcFVWS21NdHQ1YkhCT1lWTHB6SEs2Z3AzT1pDeGJ5ZWluQT09IiwiY2l0eSI6IjU5dmVIa29nbDFUZCtSVFgrMzMyY2dpSnVXUXlURXVSTm9QeFhpeFdMYlZPV0ZIUTNWMXVheitSOTBJPSIsImNvbXBhbnkiOiJkc1krM1l5aTZPWFBvU2pzRXJidUFLNENPNEhDNHY1d0hiQkVGRy84anEyWCtxS2JUMmU1bCttUEJKMzRtUE1qemcydkptTWlBcU1iTlF1SG9BcW41dzcvQUZVSlZrQzBmQT09IiwiZGVwYXJ0bWVudCI6InpBSFd5ZWdvZ1k1MWdHSytiM2dMOGZ2b0NFc090MG5YOTZmcjlOeUpOZz09IiwiZGlzdGluZ3Vpc2hlZF9uYW1lIjoicDlENHhLZmh0VWFPb2k2OTRObjIxd3VhQ1gxeEFFK2JpRDR0WHhhbTc5ZG5oMFF2b2JHOWJ1Wmdvb1doMGNGc2Y5bGZMb1dYd25oSk43OEl5MXFTRHNlTUFUclciLCJleHBpcmF0aW9uIjoiMjE2OS0xMS0yNFQyMzoyNzoyNS4xNzMwOTgxODJaIiwiZ2VvIjoiRVVDUExsNDNsbjFaZXgyTlhpNzNvQmhMN3l3blNySE10UFRLTE16VjdUZ2Rodz09IiwiaW5zdGFuY2VfdHlwZSI6IlBNYzRvYVZUV3pkWE8reS9reUJlc2pWWk1OR0NFZ095aE80SFFwcz0iLCJtYWlsIjoiU01HNEhONHp0OFJHbzU5dTVQblJBaXdWY091T1V1eXhTdTl5cDNGVThiOUFZNEFoQ2tJeHMyd2xTa3BEZmVhWDhMdndHTkE9IiwibmFtZSI6Imo0NzZaL1o0NnZabTRSR2NlVXkyUXNqT2NsQ0FxVUFuV0p3ZjZscC9GeTl6NHpMSEtLeEFtN2c9IiwicHJpbWFyeV9ncm91cF9pZCI6IlBDTUJLcGdDMVlDSW5oKzRIWDRvcFBVSzZGd2ZYMjVTZFNhQ1hhdkRlUT09Iiwicm9sZXMiOiJBYm15MXU3aHI3b1Z3aHhTZlM4SFUrT1V0cSsxZDBySENlaDNqOG9xUVE4TnB2Z1l1TmJmR21ZdGxSN01xS1BxS0owSzN3PT0iLCJzY29wZXMiOiJpMzY1ZzJrczZ4WkNPMlN0WXAybE1IOHRhZnEwQUcwQmx6VjlUQT09Iiwic3lzdGVtcyI6IlFPdWt1N3Q4em1adzhIdC9XLzBPSnJZaXNDYXp4cXhYTjNoQkVnPT0iLCJ1c2VyIjoiR2J1cEg0SktWU2RkTkJuUkJSK2p0bXlkL3FnZVdqU08vOG9uc3lITGdUZE5rWWFGUkxzYlBlaz0ifQ.bIoZiJuae3dACE2O5M1DTHKdSFHOHLDFOtW79IyEKgc"

type Body struct {
   Type     string `json:"type"`
   Metadata struct {
      Iin       string `json:"iin"`
      Name      string `json:"name"`
      DbzNumber string `json:"dbzNumber"`
      City      string `json:"city"`
      Qtest     string `json:"q_test"`
   } `json:"metadata"`
}
var papka = "\\\\ftp.sberbank.kz\\share$\\UKA_Temp$\\Скан База ДБЗ для ДРПЗФЛ 2 пул"

type Files []struct {
   ID               string `json:"id"`
   FileName         string `json:"file_name"`
   Type             string `json:"type"`
   Checked          bool   `json:"checked"`
   UploadedUser     string `json:"uploaded_user"`
   CreatedAt        string `json:"created_at"`
   Signed           bool   `json:"signed"`
   Geo              string `json:"geo"`
   ExpiredAt        string `json:"expired_at"`
   OriginalFileName string `json:"original_file_name"`
   FileLink         string `json:"file_link"`
   Bin              string `json:"bin"`
   GeoName          string `json:"geo_name"`
   Deprecated       bool   `json:"deprecated"`
}

var fz int64

func main() {
				


//    var count int
//    dirFiles, err := ioutil.ReadDir(papka)
//    if err != nil {
//       log.Fatal("read directory of directories err")
//    }
//    content, err := ioutil.ReadFile("result.txt")
//    if err != nil {
//       log.Fatal(err)
//    }

//    badContent, err := ioutil.ReadFile("bad.txt")
//    if err != nil {
//       log.Fatal(err)
//    }

//    oldFile, err := os.OpenFile("result.txt", os.O_APPEND|os.O_RDWR|os.O_CREATE, 0600)
//    if err != nil {
//       panic(err)
//    }
//    defer oldFile.Close()

//    fileWithBadDocs, err := os.OpenFile("bad.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
//    if err != nil {
//       panic(err)
//    }
//    defer fileWithBadDocs.Close()


//    start := time.Now()

//    for _, dir := range dirFiles {

//       if dir.IsDir() {
//          files, err := ioutil.ReadDir(papka + "/" + dir.Name())
//          if err != nil {
//             log.Fatal("read directory of files err")
//          }
//          for _, file := range files {
//             if strings.Contains(file.Name(), "КЛ--") || strings.Contains(file.Name(), "КЛ-1") || strings.Contains(file.Name(), "КЛ-2"){
//                if strings.Contains(string(content),file.Name()){
//                   continue
//                }
//                fileName:= file.Name()
//                var extension = filepath.Ext(fileName)
//                var name = fileName[0 : len(fileName)-len(extension)]
//                s := strings.Split(name, "--")

//                if len(s) != 3 {
//                   if !strings.Contains(string(badContent),fileName){
//                      if _, err = fileWithBadDocs.WriteString(fileName+ ": Неправильное название"+ "\n"); err != nil {
//                         panic(err)
//                      }
//                   }
//                   continue
//                }


//                if file.Size() > 100000000 {
//                   if !strings.Contains(string(badContent),fileName){
//                      if _, err = fileWithBadDocs.WriteString(fileName+ ": Размер превышает 100мб"+ "\n"); err != nil {
//                         panic(err)
//                      }
//                   }
//                   continue
//                }



// 			    doc_types := make([]string, 0)
// 			    docTypeQ := "select doc_type_id from files where to_tsvector(metadata) @@ to_tsquery('KLMIG')"

// 			    url := "postgres://postuser:Jv#/TWw3@10.4.110.3:5432/file_service"

// 			    dbPool, err := pgxpool.Connect(context.Background(), url)

// 				if err != nil {
// 					fmt.Println(err)
// 				}

// 				rows, err := dbPool.Query(context.Background(), docTypeQ)

// 				if err != nil {
// 					fmt.Println(err)
// 				}
				
// 				defer dbPool.Close()

// 				for rows.Next() {
// 					var doc string
// 					err = rows.Scan(&doc)
// 					if err != nil {
// 						fmt.Print(err)
// 					}
// 					doc_types = append(doc_types, doc)
// 				}

// 				fmt.Println(doc_types)







//                var docType string
//                if strings.Contains(file.Name(), "КЛ--"){
//                   docType = "d5d16c1d-bb3d-4775-8ad2-271841bd55aa"
//                }else {
//                   docType = "4580b2b7-afff-4785-9b95-43fe04758324"
//                }



//                bodys := Body{}
//                bodys.Type = docType
//                bodys.Metadata.Iin = s[0]
//                bodys.Metadata.Name = s[2]
//                bodys.Metadata.DbzNumber = s[1]
//                bodys.Metadata.City = dir.Name()
//                bodys.Metadata.Qtest = "KLMIG"

//                fmt.Println("Start to upload file: "+fileName)
//                errMigrate := migrateFile(file, dir,bodys)
//                if errMigrate != nil {
//                   log.Fatal(errMigrate.Error())
//                }
//                fmt.Println("File uploaded: "+fileName)

//                if _, err = oldFile.WriteString(fileName+"\n"); err != nil {
//                   panic(err)
//                }
//                count++

//             }
//          }

//       }
//    }


//    duration := time.Since(start)
//    fmt.Println(duration.Minutes())
//    fmt.Println(count)

}

func migrateFile(file os.FileInfo, dir os.FileInfo , bodys Body) error {

   client := &http.Client{}

   jsonBody, err := json.Marshal(bodys)
   payload := &bytes.Buffer{}

   fileToUpload, errFile4 := os.Open(papka +"\\"+ dir.Name() + "\\" + file.Name())
   writer := multipart.NewWriter(payload)
   _ = writer.WriteField("body", string(jsonBody))

   part4, errFile4 := writer.CreateFormFile("file", file.Name())
   defer fileToUpload.Close()

   _, errFile4 = io.Copy(part4, fileToUpload)
   if errFile4 != nil {
      log.Fatal(errFile4.Error())
   }
   err = writer.Close()
   if err != nil {
      return err
   }

   reqst, err := http.NewRequest(http.MethodPost, "http://storage-haos.apps.ocp-t.sberbank.kz/api/v1/files", payload)
   if err != nil {
      return err
   }
   reqst.Header.Add("Content-Type", "multipart/form-data")
   reqst.Header.Add("Authorization", signedToken)
   reqst.Header.Add("Geo", "PUBLIC")
   reqst.Header.Set("Content-Type", writer.FormDataContentType())

   response, err := client.Do(reqst)
   if err != nil {
      return err
   }

   var result map[string]string
   decodeErr := json.NewDecoder(response.Body).Decode(&result)
   if decodeErr != nil {
      return decodeErr
   }

   if response.StatusCode != http.StatusCreated {
      msg, _ := result["msg"]
      return errors.New(strconv.Itoa(response.StatusCode) + " status from HAOS response with message :" + msg)
   }
   return nil
}

