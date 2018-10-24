package main

import (
	"bytes"
	"regexp"
	"strings"
	"testing"
)

/*
	Result benchmark:

	go version go1.11 linux/amd64
	
	run using: go test -bench=.

	BenchmarkRegexpBytesTable1-8    	  100000	     12680 ns/op
	BenchmarkRegexpStringTable1-8   	  100000	     12666 ns/op
	BenchmarkStringsTable1-8        	  100000	     16597 ns/op
	
	BenchmarkRegexpBytesTable2-8    	  200000	      9290 ns/op
	BenchmarkRegexpStringTable2-8   	  200000	      9108 ns/op
	BenchmarkStringsTable2-8        	  200000	      6360 ns/op

*/

// Benchmark
func BenchmarkRegexpBytesTable1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, value := range tableTest_1 {
			replaceRegexpByte(value.dataBytes)
		}
	}
}

func BenchmarkRegexpStringTable1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, value := range tableTest_1 {
			replaceRegexpString(value.dataString)
		}
	}
}

func BenchmarkStringsTable1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, value := range tableTest_1 {
			replaceStrings(value.dataString)
		}
	}
}

func BenchmarkRegexpBytesTable2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, value := range tableTest_2 {
			replaceRegexpByte(value.dataBytes)
		}
	}
}

func BenchmarkRegexpStringTable2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, value := range tableTest_2 {
			replaceRegexpString(value.dataString)
		}
	}
}

func BenchmarkStringsTable2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, value := range tableTest_2 {
			replaceStrings(value.dataString)
		}
	}
}

// Test
func TestRegexpBytes(t *testing.T) {
	for key, value := range tableTest_1 {
		if result := replaceRegexpByte(value.dataBytes); result != value.result {
			t.Errorf("Failed test pos [%d] expected [%s] result [%s]", key, value.result, result)
		}
	}

}

func TestRegexpString(t *testing.T) {
	for key, value := range tableTest_1 {
		if result := replaceRegexpString(value.dataString); result != value.result {
			t.Errorf("Failed test pos [%d] expected [%s] result [%s]", key, value.result, result)
		}
	}

}

func TestStrings(t *testing.T) {
	for key, value := range tableTest_1 {
		if result := replaceStrings(value.dataString); result != value.result {
			t.Errorf("Failed test pos [%d] expected [%s] result [%s]", key, value.result, result)
		}
	}

}

var (
	re = regexp.MustCompile("  +")

	//long text
	tableTest_1 = []struct {
		dataBytes  []byte
		dataString string
		result     string
	}{
		{dataBytes: []byte("   Hello,   World !   "), dataString: "   Hello,   World !   ", result: "Hello, World !"},
		{dataBytes: []byte("Lorem       Ipsum is    simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the      industry's    standard   dummy text ever since the 1500s"),
			dataString: "Lorem       Ipsum is    simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the      industry's    standard   dummy text ever since the 1500s",
			result:     "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s",
		},
		{dataBytes: []byte("     Sed ut    perspiciatis unde omnis iste natus error sit    voluptatem accusantium doloremque laudantium, totam rem aperiam, eaque ipsa quae ab illo inventore veritatis et quasi architecto beatae vitae dicta sunt explicabo. Nemo enim ipsam voluptatem quia voluptas sit aspernatur aut odit aut fugit, sed quia consequuntur magni dolores eos qui ratione voluptatem sequi nesciunt. Neque porro quisquam est, qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit, sed quia non numquam eius modi tempora incidunt ut labore et dolore magnam aliquam quaerat voluptatem. Ut enim ad minima veniam, quis nostrum exercitationem      ullam corporis suscipit laboriosam, nisi ut aliquid ex ea commodi consequatur? Quis autem vel eum iure reprehenderit qui in ea voluptate velit esse quam nihil molestiae     consequatur, vel illum qui dolorem eum fugiat quo voluptas nulla pariatur?"),
			dataString: "     Sed ut    perspiciatis unde omnis iste natus error sit    voluptatem accusantium doloremque laudantium, totam rem aperiam, eaque ipsa quae ab illo inventore veritatis et quasi architecto beatae vitae dicta sunt explicabo. Nemo enim ipsam voluptatem quia voluptas sit aspernatur aut odit aut fugit, sed quia consequuntur magni dolores eos qui ratione voluptatem sequi nesciunt. Neque porro quisquam est, qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit, sed quia non numquam eius modi tempora incidunt ut labore et dolore magnam aliquam quaerat voluptatem. Ut enim ad minima veniam, quis nostrum exercitationem      ullam corporis suscipit laboriosam, nisi ut aliquid ex ea commodi consequatur? Quis autem vel eum iure reprehenderit qui in ea voluptate velit esse quam nihil molestiae     consequatur, vel illum qui dolorem eum fugiat quo voluptas nulla pariatur?",
			result:     "Sed ut perspiciatis unde omnis iste natus error sit voluptatem accusantium doloremque laudantium, totam rem aperiam, eaque ipsa quae ab illo inventore veritatis et quasi architecto beatae vitae dicta sunt explicabo. Nemo enim ipsam voluptatem quia voluptas sit aspernatur aut odit aut fugit, sed quia consequuntur magni dolores eos qui ratione voluptatem sequi nesciunt. Neque porro quisquam est, qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit, sed quia non numquam eius modi tempora incidunt ut labore et dolore magnam aliquam quaerat voluptatem. Ut enim ad minima veniam, quis nostrum exercitationem ullam corporis suscipit laboriosam, nisi ut aliquid ex ea commodi consequatur? Quis autem vel eum iure reprehenderit qui in ea voluptate velit esse quam nihil molestiae consequatur, vel illum qui dolorem eum fugiat quo voluptas nulla pariatur?"},
		{dataBytes: []byte("   Hello World !   "), dataString: "   Hello World !   ", result: "Hello World !"},
		{dataBytes: []byte("TEST   Hello,   World !   "), dataString: "TEST   Hello,   World !   ", result: "TEST Hello, World !"},
		{dataBytes: []byte(" Lorem  ipsum dolor  ? !"), dataString: " Lorem  ipsum dolor  ? !", result: "Lorem ipsum dolor ? !"},
	}
	//short text
	tableTest_2 = []struct {
		dataBytes  []byte
		dataString string
		result     string
	}{
		{dataBytes: []byte("   Hello,   World !   "), dataString: "   Hello,   World !   ", result: "Hello, World !"},
		{dataBytes: []byte("Lorem       Ipsum is    simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the      industry's    standard   dummy"),
			dataString: "Lorem       Ipsum is    simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the      industry's    standard   dummy",
			result:     "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy",
		},
		{dataBytes: []byte("     Sed ut    perspiciatis unde omnis iste natus error sit    voluptatem accusantium doloremque laudantium, totam rem  aperiam, eaque ipsa quae ab illo inventore veritatis et"),
			dataString: "     Sed ut    perspiciatis unde omnis iste natus error sit    voluptatem accusantium doloremque laudantium, totam rem  aperiam, eaque ipsa quae ab illo inventore veritatis et",
			result:     "Sed ut perspiciatis unde omnis iste natus error sit voluptatem accusantium doloremque laudantium, totam rem aperiam, eaque ipsa quae ab illo inventore veritatis et"},
		{dataBytes: []byte("   Hello World !   "), dataString: "   Hello World !   ", result: "Hello World !"},
		{dataBytes: []byte("TEST   Hello,   World !   "), dataString: "TEST   Hello,   World !   ", result: "TEST Hello, World !"},
		{dataBytes: []byte(" Lorem  ipsum dolor  ? !"), dataString: " Lorem  ipsum dolor  ? !", result: "Lorem ipsum dolor ? !"},
	}
)

func replaceRegexpByte(s []byte) string {
	replaced := re.ReplaceAll(bytes.TrimSpace(s), []byte(" "))
	return string(replaced)
}

func replaceRegexpString(s string) string {
	return re.ReplaceAllString(strings.TrimSpace(s), " ")
}

func replaceStrings(s string) string {
	return strings.Join(strings.Fields(s), " ")
}
