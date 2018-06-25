package courier

import (
	"fmt"
)

func ExampleMetadata() {
	metaA := Metadata{}
	metaA.Add("A", "valueA1")
	metaA.Add("A", "valueA2")

	metaB := Metadata{}
	metaB.Set("B", "valueB1")

	metaAll := FromMetas(metaA, metaB)

	results := []interface{}{
		metaAll.String(),
		metaAll.Has("A"),
		metaAll.Get("A"),
	}

	metaAll.Del("B")

	results = append(results,
		metaAll.Get("B"),
		metaAll.String(),
	)

	for _, r := range results {
		fmt.Printf("%v\n", r)
	}
	//Output:
	//A=valueA1&A=valueA2&B=valueB1
	//true
	//valueA1
	//
	//A=valueA1&A=valueA2
}
