package main

import "net"
import "fmt"
import "bufio"
import "os"

type userModel struct {
	Connection net.Conn
	Name       string
}

func newUser(connection net.Conn) *userModel {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите имя: ")
	userName, _ := reader.ReadString('\n')
	usr := userModel{Connection: connection,
		Name: userName[:len(userName)-1]}
	return &usr
}

func (u *userModel) sendMessage(text string) error {
	_, err := fmt.Fprintf(u.Connection, u.Name+": "+text+"\n")
	return err

}
func (u *userModel) receiveMessage() {
	for {
		message, _ := bufio.NewReader(u.Connection).ReadString('\n')
		fmt.Printf("\n"+message+"%s: ", u.Name)
	}

}

func main() {

	// Подключаемся к сокету
	connection, err := net.Dial("tcp", "127.0.0.1:8090")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Подключено")
	//создали объект пользователя
	user1 := newUser(connection)
	go user1.receiveMessage()
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Printf("%s: ", user1.Name)
		text, _ := reader.ReadString('\n')
		user1.sendMessage(text)
	}
}
