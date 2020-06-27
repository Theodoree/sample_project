package main

import (
    "bytes"
    "errors"
    "fmt"
    "github.com/streadway/amqp"
    "log"
    "math/rand"
    "strconv"
    "time"
)

var (
    isNotConnectError = errors.New("isNotConnectError")
)

type client struct {
    conn      *amqp.Connection
    ch        *amqp.Channel
    isConnect bool
    chName    string
}

func newClient() *client {
    return &client{}
}

func (c *client) Connect() error {
    conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
    if err != nil {
        return err
    }
    c.conn = conn
    c.ch, err = c.conn.Channel()
    if err != nil {
        c.conn.Close()
        return err
    }
    c.isConnect = true
    return nil
}

func (c *client) Close() {
    c.conn.Close()
    c.ch.Close()
}

func SendReceive() {
    client := newClient()

    err := client.Connect()
    if err != nil {
        fmt.Println(err)
    }
    defer client.Close()

    q, err := client.ch.QueueDeclare(
        "hello", // name
        true,    // 非持久化
        false,   // delete when unused
        false,   // exclusive
        false,   // no-wait
        nil,     // arguments
    )
    if err != nil {
        fmt.Println(err)
    }

    body := "Hello World!"
    err = client.ch.Publish(
        "",     // exchange
        q.Name, // routing key
        false,  // mandatory
        false,  // immediate
        amqp.Publishing{
            ContentType: "text/plain",
            Body:        []byte(body),
        })
    if err != nil {
        fmt.Println(err)
    }

}

func Receive() {
    client := newClient()

    err := client.Connect()
    if err != nil {
        log.Print(err)
    }

    defer client.Close()
    ch := client.ch

    q, err := ch.QueueDeclare(
        "hello", // name
        true,    // durable
        false,   // delete when unused
        false,   // exclusive
        false,   // no-wait
        nil,     // arguments
    )

    msgs, err := ch.Consume(
        q.Name, // queue
        "",     // consumer
        false,  // auto-ack
        false,  // exclusive
        false,  // no-local
        false,  // no-wait
        nil,    // args
    )
    forever := make(chan bool)

    go func() {
        for d := range msgs {

            err := d.Ack(false)
            if err != nil {
                fmt.Println(err)
            }
            log.Printf("Received a message: %s", d.Body)
        }
    }()

    log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
    <-forever

}

func newTask() {
    client := newClient()

    err := client.Connect()
    if err != nil {
        fmt.Println(err)
    }
    defer client.Close()
    ch := client.ch

    q, err := ch.QueueDeclare(
        "task_queue", // name
        true,         // 持久化
        false,        // delete when unused
        false,        // exclusive
        false,        // no-wait
        nil,          // arguments
    )
    if err != nil {
        fmt.Println(err)
    }

    body := `test.task`
    err = ch.Publish(
        "",     // exchange
        q.Name, // routing key
        false,  // mandatory
        false,
        amqp.Publishing{
            DeliveryMode: amqp.Persistent,
            ContentType:  "text/plain",
            Body:         []byte(body),
        })

    log.Printf(" [x] Sent %s", body)

}

func TaskWorker() {
    client := newClient()

    err := client.Connect()
    if err != nil {
        fmt.Println(err)
    }
    defer client.Close()
    ch := client.ch
    q, err := ch.QueueDeclare(
        "task_queue", // name
        true,         // 持久化
        false,        // delete when unused
        false,        // exclusive
        false,        // no-wait
        nil,          // arguments
    )
    if err != nil {
        fmt.Println(err)
    }

    // pre ack
    err = ch.Qos(
        1,     // 预计数量
        0,     // 预计大小(byte)
        false, // global
    )

    if err != nil {
        fmt.Println(err)
    }
    msgs, err := ch.Consume(
        q.Name, // queue
        "",     // consumer
        false,  // auto-ack
        false,  // exclusive
        false,  // no-local
        false,  // no-wait
        nil,    // args
    )
    if err != nil {
        fmt.Println(err)
    }
    forever := make(chan bool)

    go func() {
        for d := range msgs {
            log.Printf("Received a message: %s", d.Body)
            dot_count := bytes.Count(d.Body, []byte("."))
            t := time.Duration(dot_count)
            time.Sleep(t * time.Second)
            log.Printf("Done")
            d.Ack(false)
        }
    }()

    log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
    <-forever
}

func Publish() {
    client := newClient()

    err := client.Connect()
    if err != nil {
        fmt.Println(err)
    }
    defer client.Close()
    ch := client.ch

    err = ch.ExchangeDeclare( // 订阅模式
        "logs",   // name
        "fanout", // type
        true,     // 持久化
        false,    // auto-deleted
        false,    // internal
        false,    // no-wait
        nil,      // arguments
    )
    if err != nil {
        fmt.Println(err)
    }

    for i := 0; i <= 10; i++ {
        time.Sleep(time.Millisecond * 100)

        err = ch.Publish(
            "logs", // exchange
            "",     // routing key
            false,  // mandatory
            false,  // immediate
            amqp.Publishing{
                ContentType: "text/plain",
                Body:        []byte(strconv.Itoa(i)),
            })
    }

}

func Subscribe(i int) {
    client := newClient()

    err := client.Connect()
    if err != nil {
        fmt.Println(err)
    }
    defer client.Close()
    ch := client.ch
    err = ch.ExchangeDeclare(
        "logs",   // name
        "fanout", // type
        true,     // durable
        false,    // auto-deleted
        false,    // internal
        false,    // no-wait
        nil,      // arguments
    )
    if err != nil {
        fmt.Println(err)
    }

    q, err := ch.QueueDeclare(
        "",    // name
        false, // durable
        false, // delete when unused
        true,  // exclusive
        false, // no-wait
        nil,   // arguments
    )
    if err != nil {
        fmt.Println(err)
    }

    err = ch.QueueBind(
        q.Name, // queue name
        "",     // routing key
        "logs", // exchange
        false,
        nil)
    if err != nil {
        fmt.Println(err)
    }

    msgs, err := ch.Consume(
        q.Name, // queue
        "",     // consumer
        true,   // 自动回复
        false,  // exclusive
        false,  // no-local
        false,  // no-wait
        nil,    // args
    )
    if err != nil {
        fmt.Println(err)
    }

    for d := range msgs {
        log.Printf(" [x][%d] %s", i, d.Body)
    }
}

func RoutingPublish(key string) {
    client := newClient()

    err := client.Connect()
    if err != nil {
        fmt.Println(err)
    }
    defer client.Close()
    ch := client.ch

    err = ch.ExchangeDeclare( // 订阅模式
        "logs_direct", // name
        "direct",      // type
        true,          // 持久化
        false,         // auto-deleted
        false,         // internal
        false,         // no-wait
        nil,           // arguments
    )
    if err != nil {
        fmt.Println(err)
    }

    for i := 0; i <= 10; i++ {
        time.Sleep(time.Millisecond * 100)

        err = ch.Publish(
            "logs_direct", // exchange
            key,           // routing key
            false,         // mandatory
            false,         // immediate
            amqp.Publishing{
                ContentType: "text/plain",
                Body:        []byte(strconv.Itoa(i)),
            })
    }
}

func RoutingSubscribe(i int, key string) {
    client := newClient()

    err := client.Connect()
    if err != nil {
        fmt.Println(err)
    }
    defer client.Close()
    ch := client.ch
    err = ch.ExchangeDeclare(
        "logs_direct", // name
        "direct",      // type
        true,          // durable
        false,         // auto-deleted
        false,         // internal
        false,         // no-wait
        nil,           // arguments
    )
    if err != nil {
        fmt.Println(err)
    }

    // 初始化一个队列
    q, err := ch.QueueDeclare(
        "",    // name
        false, // durable
        false, // delete when unused
        true,  // exclusive
        false, // no-wait
        nil,   // arguments
    )
    if err != nil {
        fmt.Println(err)
    }

    // 队列绑定订阅类
    err = ch.QueueBind(
        q.Name,        // queue name
        key,           // routing key
        "logs_direct", // exchange
        false,
        nil)
    if err != nil {
        fmt.Println(err)
    }

    // 绑定队列
    msgs, err := ch.Consume(
        q.Name, // queue
        "",     // consumer
        true,   // 自动回复
        false,  // exclusive
        false,  // no-local
        false,  // no-wait
        nil,    // args
    )
    if err != nil {
        fmt.Println(err)
    }

    for d := range msgs {
        log.Printf(" [x][%d] %s", i, d.Body)
    }
}

func TopicsSender(key string) {
    client := newClient()

    err := client.Connect()
    if err != nil {
        fmt.Println(err)
    }
    defer client.Close()
    ch := client.ch

    err = ch.ExchangeDeclare(
        "logs_topic", // name
        "topic",      // type
        true,         // durable
        false,        // auto-deleted
        false,        // internal
        false,        // no-wait
        nil,          // arguments
    )
    if err != nil {
        fmt.Println(err)
    }

    for i := 0; i < 10; i++ {
        err = ch.Publish(
            "logs_topic", // exchange
            key,          // routing key
            false,        // mandatory
            false,        // immediate
            amqp.Publishing{
                ContentType: "text/plain",
                Body:        []byte(strconv.Itoa(i)),
            })
        if err != nil {
            fmt.Println(err)
        }
        log.Printf(" [x] Sent %s", strconv.Itoa(i))

    }

}

func TopicsReceive(key string) {
    client := newClient()

    err := client.Connect()
    if err != nil {
        fmt.Println(err)
    }
    defer client.Close()
    ch := client.ch

    err = ch.ExchangeDeclare(
        "logs_topic", // name
        "topic",      // type
        true,         // durable
        false,        // auto-deleted
        false,        // internal
        false,        // no-wait
        nil,          // arguments
    )
    if err != nil {
        fmt.Println(err)
    }

    q, err := ch.QueueDeclare(
        "",    // name
        false, // durable
        false, // delete when unused
        true,  // exclusive
        false, // no-wait
        nil,   // arguments
    )
    if err != nil {
        fmt.Println(err)
    }

    err = ch.QueueBind(
        q.Name,       // queue name
        key,          // routing key
        "logs_topic", // exchange
        false,
        nil)
    if err != nil {
        fmt.Println(err)
    }

    msgs, err := ch.Consume(
        q.Name, // queue
        "",     // consumer
        true,   // auto ack
        false,  // exclusive
        false,  // no local
        false,  // no wait
        nil,    // args
    )
    if err != nil {
        fmt.Println(err)
    }

    forever := make(chan bool)

    go func() {
        for d := range msgs {
            log.Printf(" [x] %s", d.Body)
        }
    }()

    log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
    <-forever
}

func rpcService() {
    client := newClient()

    err := client.Connect()
    if err != nil {
        fmt.Println(err)
    }
    defer client.Close()
    ch := client.ch

    q, err := ch.QueueDeclare(
        "rpc_queue", // name
        false,       // durable
        false,       // delete when unused
        false,       // exclusive
        false,       // no-wait
        nil,         // arguments
    )
    if err != nil {
        fmt.Println(err)
    }

    // pre ack
    err = ch.Qos(
        1,     // prefetch count
        0,     // prefetch size
        false, // global
    )
    if err != nil {
        fmt.Println(err)
    }

    msgs, err := ch.Consume(
        q.Name, // queue
        "",     // consumer
        false,  // auto-ack
        false,  // exclusive
        false,  // no-local
        false,  // no-wait
        nil,    // args
    )
    if err != nil {
        fmt.Println(err)
    }

    forever := make(chan bool)

    go func() {
        for d := range msgs {
            n, err := strconv.Atoi(string(d.Body))
            if err != nil {
                fmt.Println(err)
            }

            log.Printf(" [.] fib(%d)", n)
            response := fib(n)

            err = ch.Publish(
                "",        // exchange
                d.ReplyTo, // routing key
                false,     // mandatory
                false,     // immediate
                amqp.Publishing{
                    ContentType:   "text/plain",
                    CorrelationId: d.CorrelationId,
                    Body:          []byte(strconv.Itoa(response)),
                })
            if err != nil {
                fmt.Println(err)
            }

            d.Ack(false)
        }
    }()

    log.Printf(" [*] Awaiting RPC requests")
    <-forever
}

func rpcClient() {
    client := newClient()

    err := client.Connect()
    if err != nil {
        fmt.Println(err)
    }
    defer client.Close()
    ch := client.ch

    // 初始化一个队列
    q, err := ch.QueueDeclare(
        "",    // name
        false, // durable
        false, // delete when unused
        true,  // exclusive
        false, // noWait
        nil,   // arguments
    )
    if err != nil {
        fmt.Println(err)
    }

    msgs, err := ch.Consume(
        q.Name, // queue
        "",     // consumer
        true,   // auto-ack
        false,  // exclusive
        false,  // no-local
        false,  // no-wait
        nil,    // args
    )
    if err != nil {
        fmt.Println(err)
    }

    randomString := func(l int) string {
        buf := make([]byte, l)
        for i := 0; i < l; i++ {
            buf[i] = byte(65 + rand.Intn(90-65))
        }
        return string(buf)
    }

    corrId := randomString(32)

    err = ch.Publish(
        "",          // exchange
        "rpc_queue", // routing key
        false,       // mandatory
        false,       // immediate
        amqp.Publishing{
            ContentType:   "text/plain",
            CorrelationId: corrId,
            ReplyTo:       q.Name,
            Body:          []byte(strconv.Itoa(12)),
        })
    if err != nil {
        fmt.Println(err)
    }

    for d := range msgs {
        if corrId == d.CorrelationId {
            res, err := strconv.Atoi(string(d.Body))
            if err != nil {
                fmt.Println(err)
            }
            fmt.Println(res)
            break
        }
    }

}

func fib(n int) int {
    if n == 1 {
        return 1
    } else if n == 2 {
        return 1
    }
    var fibarr = make([]int, n+1)
    fibarr[0] = 0
    fibarr[1] = 1

    for i := 2; i <= n; i++ {
        fibarr[i] = fibarr[i-1] + fibarr[i-2]
    }

    return fibarr[n]
}

func main() {
    //c:=newClient()
    //c.Connect()
    //i,err:= c.ch.QueueDelete("hello",false,true,false)
    //fmt.Println(i,err)

    //SendReceive()
    //Receive()

    //newTask()
    //TaskWorker()

    //go Publish()
    //for i := 0; i < 3; i++ {
    //  go Subscribe(i)
    //}
    //time.Sleep(time.Second * 10)

    //go RoutingPublish("test_key")
    //for i := 0; i < 10; i++ {
    //go RoutingSubscribe(i,"test_key")
    //}
    //time.Sleep(time.Second * 10)

    //go TopicsReceive("")
    //for i := 0; i < 3; i++ {
    //   time.Sleep(time.Second)
    //   TopicsSender("")
    //}

    go rpcService()
    rpcClient()


}


