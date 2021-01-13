package main

import(
	"fmt"
	"github.com/gin-gonic/gin"
	 "database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gomodule/redigo/redis"
	"log"
	"github.com/Shopify/sarama"
	"github.com/wvanbergen/kafka/consumergroup"
	"os"
	"time"
)

const (
	zookeeperConn = "localhost:9092"
	cgroup = "zgroup"
	topic = "quickstart-events"
)

type User struct {
	Name string `json:"name"`
}

type Book struct{
	title string
	author string
	category string
	price float64
}

func HomePage(c *gin.Context){
	c.JSON(200, gin.H{
		"message": "Hello world",
	})
}

func PostHomePage(c *gin.Context){
	c.JSON(200,gin.H{
		"message": "Post Home Page",
	})
}

func main() {
	r := gin.Default()
	db, err := sql.Open("mysql", "root:Aa123456@tcp(127.0.0.1:3306)/testdb")
	if err != nil{
		panic(err.Error())
	}
	defer db.Close()
	fmt.Println("Connect to MySQL successfully")
	r.GET("/", HomePage)
	//insert, err := db.Query("INSERT INTO users VALUES('ELLIOT')")
	//
	//if err!=nil{
	//	panic(err.Error())
	//}
	//
	//defer insert.Close()

	results, err:= db.Query("SELECT name from users")
	if err != nil{
		panic(err.Error())
	}

	for results.Next(){
		var user User

		err:=results.Scan(&user.Name)
		if err!=nil{
			panic(err.Error())
		}

		fmt.Println(user.Name)
	}

	defer results.Close()

	conn, err:= redis.Dial("tcp",":6379")
	if err != nil{
		log.Fatal(err)
	}

	_, err = conn.Do("HMSET",
		"book:1",
		"title",
		"Sherlock Holmes",
		"creator",
		"AC Doyle",
		"category",
		"detective",
		"price of book",
		59.99,)

	title, err := redis.String(conn.Do("HGET", "book:1","creator"))
	if err != nil {
		log.Fatal(err) // handle error
	}
	fmt.Println("Book :", title)

	defer conn.Close()

	r.POST("/", PostHomePage)

	// setup sarama log to stdout
	sarama.Logger = log.New(os.Stdout, "", log.Ltime)

	// init consumer
	cg, err := initConsumer()
	if err != nil {
		fmt.Println("Error consumer group: ", err.Error())
		os.Exit(1)
	}
	defer cg.Close()

	// run consumer
	consume(cg)


	r.Run()
}

func initConsumer()(*consumergroup.ConsumerGroup, error) {
	// consumer config
	config := consumergroup.NewConfig()
	config.Offsets.Initial = sarama.OffsetOldest
	config.Offsets.ProcessingTimeout = 10 * time.Second

	// join to consumer group
	cg, err := consumergroup.JoinConsumerGroup(cgroup, []string{topic}, []string{zookeeperConn}, config)
	if err != nil {
		return nil, err
	}

	return cg, err
}

func consume(cg *consumergroup.ConsumerGroup) {
	for {
		select {
		case msg := <-cg.Messages():
			// messages coming through chanel
			// only take messages from subscribed topic
			if msg.Topic != topic {
				continue
			}

			fmt.Println("Topic: ", msg.Topic)
			fmt.Println("Value: ", string(msg.Value))

			// commit to zookeeper that message is read
			// this prevent read message multiple times after restart
			err := cg.CommitUpto(msg)
			if err != nil {
				fmt.Println("Error commit zookeeper: ", err.Error())
			}
		}
	}
}
