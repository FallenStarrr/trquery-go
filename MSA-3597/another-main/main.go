package main
import(
	"fmt"
	"database/sql"
	_ "github.com/lib/pq"
)



func GetKLMIG() {
		type DocTypeAndId struct {
		DocType string
		Id string
		}

			doc_types := make([]DocTypeAndId, 0)
			
			docTypeQ := "select doc_type_id, id from files where to_tsvector(metadata) @@ to_tsquery('KLMIG')"

			// url := "postgres://postuser:Jv#/TWw3@10.4.110.3:5432/file_service"
			
			conn := "user=postuser host=10.4.110.3 sslmode=disable dbname=file_service password=Jv#/TWw3 port=5432"
			dbPool, err := sql.Open("postgres", conn)

			if err != nil {
				fmt.Println(err)
			}

			rows, err := dbPool.Query(docTypeQ)

			if err != nil {
				fmt.Println(err)
			}
			
			defer dbPool.Close()

			for rows.Next() {
				doc := DocTypeAndId{}
				err = rows.Scan(&doc.DocType, &doc.Id)
				if err != nil {
					fmt.Print(err)
				}
				doc_types = append(doc_types, doc)
			}

			fmt.Println(doc_types)

			fmt.Println(len(doc_types))
}			



func main() {

	GetKLMIG()
}