package main

import (
	"bufio"
	"city_game/src/game"
	"city_game/src/pb"
	"encoding/binary"
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"

	"google.golang.org/protobuf/proto"
)

func start_game(init_conn net.Conn, port int) {
	dialer := &net.Dialer{
		Timeout:   5 * time.Second,
		KeepAlive: 30 * time.Second,
	}

	conn, err := dialer.Dial("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	fmt.Println("Connected to client server")

	state := game.CreateInitialState()
	input := &pb.GameInput{}

	for range 5 {
		err := send_state(conn, state)
		if err != nil {
			fmt.Println("Ran into an error when sending state to client server: ", err)
			return
		}

		err = receive_input(conn, input)
		if err != nil {
			fmt.Println("Ran into an error when receiving input from client server: ", err)
			return
		}

		state = game.GetNextState(state, input)
	}
}

func send_state(client_conn net.Conn, state *pb.GameState) error {
	state_data, err := proto.Marshal(state)
	if err != nil {
		return err
	}

	length := uint32(len(state_data))
	err = binary.Write(client_conn, binary.BigEndian, length)
	if err != nil {
		return err
	}

	_, err = client_conn.Write(state_data)
	return err
}

func receive_input(client_conn net.Conn, input *pb.GameInput) error {
	fmt.Println("Getting length of buffer")

	lenBuf, err := read_full(client_conn, 4)
	if err != nil {
		return err
	}
	msgLen := binary.BigEndian.Uint32(lenBuf)

	fmt.Printf("Got length of %d, now reading data\n", msgLen)

	data, err := read_full(client_conn, int(msgLen))
	if err != nil {
		return err
	}

	fmt.Println("Got data, now parsing it")

	err = proto.Unmarshal(data, input)
	return err
}

func read_full(conn net.Conn, size int) ([]byte, error) {
	buf := make([]byte, size)
	total := 0
	for total < size {
		n, err := conn.Read(buf[total:])
		if err != nil {
			return nil, err
		}
		total += n
	}
	return buf, nil
}

func handle_connection(conn net.Conn) {
	defer conn.Close()

	if tcpConn, ok := conn.(*net.TCPConn); ok {
		tcpConn.SetKeepAlive(true)
		tcpConn.SetKeepAlivePeriod(30 * time.Second) // adjust as needed
	}

	reader := bufio.NewReader(conn)
	port_message, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Didn't receive initial port number:", err)
		return
	}

	port_message = strings.TrimRight(port_message, " \r\n")

	port_number, err := strconv.Atoi(port_message)
	if err != nil {
		fmt.Println("Initial message wasn't a port number:", err)
		return
	}

	if port_number <= 0 || port_number >= 65535 {
		fmt.Printf("Port number %d is invalid\n", port_number)
		return
	}

	start_game(conn, port_number)
}

func main() {
	listener, err := net.Listen("tcp", ":8002")
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error during accept: ", err)
		}

		go handle_connection(conn)
	}
}
