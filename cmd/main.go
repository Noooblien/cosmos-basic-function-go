package main

import "github.com/Noooblien/cosmos-basic-function-go/pkg"

//How to Implement the function made in pkg 🫶 🫶 🫶
//* Author rsm050501@gmail.com  Rahul Singh Maraskole

func main() {

	// txn Decoder
	base64Tx := "Co0BCooBChwvY29zbW9zLmJhbmsudjFiZXRhMS5Nc2dTZW5kEmoKLWNvc21vczFsd3pjcmY5NWVqZjV5MHVoZmRhMzdkcnJlZ3d1eTlsbDJ2OHRudhItY29zbW9zMW56dzh2NzA1c3h2cWNkMzQ5bnczcDgzMHNsdHRxbWp4ZmY1ZjR1GgoKBXN0YWNrEgExElkKUQpGCh8vY29zbW9zLmNyeXB0by5zZWNwMjU2azEuUHViS2V5EiMKIQOePIw5ZTDJpuU+++f9yuevH6dYVFqwIyWkHHaVZW1mbBIECgIIARjJARIEEMCaDBpA8L7r9XK0tFC/yX0y+tFkwgk7AYSPQQ2M9ZWLevJkYvdt9iawn4S9/hGsnwRe/m8osfVlvD6zf5RDCElFrUR9qw=="
	txhash, err := pkg.ComputeTransactionHash(base64Tx)
	if err != nil {
		panic(err)
	}
	println(txhash, "This is the Decode Txn Hash ")

}
