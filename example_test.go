package poc_test

import (
	"fmt"

	"github.com/aereal/poc-protobuf-json-config/pb"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func Example_marshal() {
	def := &pb.Definition{Livers: make(map[string]*pb.Liver)}
	def.Livers["MitoTsukino"] = &pb.Liver{Name: "Tsukino Mito", Age: proto.Int32(16)}
	def.Livers["Lize_Helesta"] = &pb.Liver{Name: "Lize Helesta", Age: proto.Int32(17)}
	json, err := protojson.Marshal(def)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(json))
	// Output:
	// {"livers":{"Lize_Helesta":{"name":"Lize Helesta", "age":17}, "MitoTsukino":{"name":"Tsukino Mito", "age":16}}}
}
