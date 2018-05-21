package main

import (
	"net"
	"os"
	"fmt"
	"time"
)

const (

	MAX_CONN_NUM = 5

)


//echo server Goroutine

func EchoFunc(conn net.Conn, conn_close_flag chan int) {

	defer conn.Close()

	defer func() {

		conn_close_flag <- -1

	}()


	buf := make([]byte, 1024)

	for {

		_, err := conn.Read(buf)

		if err != nil {

			//println("Error reading:", err.Error())

			return

		}

		//send reply

		_, err = conn.Write(buf)

		if err != nil {

			//println("Error send reply:", err.Error())

			return

		}

	}

}


//initial listener and run

func main() {


	listener, err := net.Listen("tcp", "0.0.0.0:8088")


	if err != nil {

		println("error listening:", err.Error())

		os.Exit(1)

	}


	defer listener.Close()


	fmt.Printf("running ...\n")


	var cur_conn_num float64 = 0


	ch_conn_change := make(chan int, MAX_CONN_NUM)


	tick := time.Tick(1e8)


	for {

		//read all close flags berfor accept new connection

		//TODO: better code to handle batch close?

		readmore := 1

		for readmore > 0 {

			select {

			case conn_change := <-ch_conn_change:

				cur_conn_num = cur_conn_num + float64(conn_change)

			default:

				readmore = 0

			}

		}

		//FIXME: tick block by listener.Accept()

		select {

		case <-tick:

			fmt.Printf("cur conn num: %f\n", cur_conn_num)

		default:

		}

		if cur_conn_num >= MAX_CONN_NUM {

			//reach MAX_CONN_NUM, waiting for exist connection close

			time.Sleep(time.Second)

		} else {

			//accept new connetion

			conn, err := listener.Accept()

			if err != nil {

				println("Error accept:", err.Error())

				return

			}

			cur_conn_num++

			go EchoFunc(conn, ch_conn_change)

		}

	}

}
