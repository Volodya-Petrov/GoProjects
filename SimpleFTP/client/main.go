package client

import "fmt"

func main() {
	ip, port, pathOnServer, pathLocal := "", "", "", ""
	fmt.Print("Введите ip сервера:")
	fmt.Scanf("%s", &ip)
	fmt.Print("Введите порт сервера:")
	fmt.Scanf("%s", &port)
	for {
		fmt.Print("Введите команду:")
		cm := ""
		fmt.Scanf("%s", &cm)
		if cm == "exit" {
			return
		}
		if cm == "1" {
			fmt.Print("Введите путь к каталогу на сервере:")
			fmt.Scanf("%s", &pathOnServer)
			result, err := List(ip, port, pathOnServer)
			if err != nil {
				fmt.Printf("Ошибка:%v", err)
				continue
			}
			for _, v := range result {
				fmt.Println("%v %v", v.path, v.isDir)
			}
		} else {
			fmt.Print("Введите путь к файлу на сервере:")
			fmt.Scanf("%s", &pathOnServer)
			fmt.Print("Введите локальный путь, куда сохранить файлу:")
			fmt.Scanf("%s", &pathLocal)
			result, err := Get(ip, port, pathOnServer, pathLocal)
			if err != nil {
				fmt.Printf("Ошибка:%v", err)
				continue
			}
			fmt.Printf("Размер скачанного файла:%v\n", result)
		}
	}
}
