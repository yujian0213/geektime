package main
import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/hhxsv5/go-redis-memory-analysis"
)

var client redis.UniversalClient
var ctx context.Context

const  (
	ip string ="127.0.0.1"
	port uint16 = 6379
)

func init()  {
	client = redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%v:%v", ip, port),
		Password:     "",
		DB:           0,
		PoolSize:     128,
		MinIdleConns: 100,
		MaxRetries:   5,
	})
	ctx = context.Background()
}
func main()  {
	write(10000, "len10_10k", generateValue(10))
	write(50000, "len10_50k", generateValue(10))
	write(500000, "len10_500k", generateValue(10))
	analysis()
	
}
func write(num int,k,v string)  {
	for i := 0; i <num ; i++ {
		k = fmt.Sprintf("%s:%v",k,i)
		operate := client.Set(ctx,k,v,-1)
		err := operate.Err()
		if err !=nil {
			fmt.Println(operate.String())
		}
	}
}
func analysis()  {
	analysis,err := gorma.NewAnalysisConnection(ip,port,"")
	if err != nil {
		fmt.Println("something wrong:", err)
		return
	}
	defer analysis.Close()
	analysis.Start([]string{":"})
	err = analysis.SaveReports("./reports")
	if err == nil {
		fmt.Println("done")
	} else {
		fmt.Println("error:", err)
	}
}
func generateValue(size int) string {
	arr := make([]byte,size)
	for i := 0; i < size; i++ {
		arr[i] = 'a'
	}
	return string(arr)
}
