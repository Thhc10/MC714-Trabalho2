package main
import (
	"fmt"
	"net"
	"net/rpc"
	"log"
)

// Responder para obter respostas como mensagens OK das chamadas RPC.
type Reply struct{
	Data string
}

// Estrutura básica do Algoritmo Bully. Ele contém funções registradas para RPC.
// também contém informações sobre outros nós.
type BullyAlgorithm struct{
	my_id int
	coordinator_id int
	ids_ip map[int]string
}

// Se já invocou a eleição, não precisa iniciar a eleição novamente.
var no_election_invoked = true

// Esta é a função de eleição que é invocada quando um ID de host menor solicita uma eleição para este host
func (bully *BullyAlgorithm) Election(invoker_id int, reply *Reply) error{
	fmt.Println("Log: Recebendo eleição por ", invoker_id)
	if invoker_id < bully.my_id{
		fmt.Println("Log: Enviando OK para ", invoker_id)
		reply.Data = "OK"
		if no_election_invoked{
			no_election_invoked = false
			go invokeElection()			// evoca eleição para um host superior
		}
	}
	return nil
}


// Alterna quando algum host superior envia mensagem de OK
var superiorNodeAvailable = false

// Esta função evoca a eleição para seus hosts superiores. Ele envia seu Id como parâmetro ao chamar o RPC
func invokeElection(){
	for id,ip := range bully.ids_ip{
		reply := Reply{""}
		if id > bully.my_id{
			fmt.Println("Log: Enviando eleição para ", id)
			client, error := rpc.Dial("tcp",ip)
			if error != nil{
				fmt.Println("Log:", id, "não está disponível.")
				continue
			}
			err := client.Call("BullyAlgorithm.Election", bully.my_id, &reply)
			if err != nil{
				fmt.Println(err)
				fmt.Println("Log: Erro em evocar a função", id, " eleição.")
				continue
			}
			if reply.Data == "OK"{				// Means superior host exists
				fmt.Println("Log: Recebendo OK por ", id)
				superiorNodeAvailable = true
			}
		}
	}
	if !superiorNodeAvailable{
		makeYourselfCoordinator()
	}
	superiorNodeAvailable = false
	no_election_invoked = true
}

// Esta função é chamada pelo novo Coordenador para atualizar as informações do coordenador dos outros hosts
func (bully *BullyAlgorithm) NewCoordinator(new_id int, reply *Reply) error{
	bully.coordinator_id = new_id 
	fmt.Println("Log:", bully.coordinator_id, "é o novo coordenador.")
	return nil
}

func (bully *BullyAlgorithm) HandleCommunication(req_id int, reply *Reply) error{
	fmt.Println("Log: Recebida comunicação pelo ", req_id)
	reply.Data = "OK"
	return nil
}

func communicateToCoordinator(){
	coord_id := bully.coordinator_id
	coord_ip := bully.ids_ip[coord_id]
	fmt.Println("Log: Comunicando com o coordenador ", coord_id)
	my_id := bully.my_id
	reply := Reply{""}
	client, err := rpc.Dial("tcp", coord_ip)
	if err != nil{
		fmt.Println("Log: Coordenador ",coord_id, " comunicação falhou!")
        fmt.Println("Iniciando as eleições!")
		invokeElection()
		return
	}
	err = client.Call("BullyAlgorithm.HandleCommunication", my_id, &reply)
	if err != nil || reply.Data != "OK"{
		fmt.Println("Log: Comunicando com o coordenador ", coord_id, " falhou!")
        fmt.Println("Iniciando as eleições!")
        invokeElection()
		return
	}
	fmt.Println("Log: Comunicação recebido pela coordenador ", coord_id)
}

// Esta função é chamada quando o host decide que é o coordenador.
// transmite a mensagem para todos os outros hosts e atualiza as informações do líder, incluindo seu próprio host.
func makeYourselfCoordinator(){
	reply := Reply{""}
	for id, ip := range bully.ids_ip{
		client, error := rpc.Dial("tcp", ip)
		if error != nil{
			fmt.Println("Log:", id, "erro de comunicação.")
			continue
		}
		client.Call("BullyAlgorithm.NewCoordinator", bully.my_id, &reply)
	}
}

// Esse trecho de código é responsável por mudar os sockets e a quantidade de hosts.
var bully = BullyAlgorithm{
	my_id: 		1,
	coordinator_id: 4,
	ids_ip: 	map[int]string{	1:"10.128.0.5:3030", 2:"10.128.0.6:3030", 3:"10.128.0.7:3030", 4:"34.172.92.62:3030"}
	}


func main(){
	my_id := 0
	fmt.Printf("Entre com o host ID: ")
	fmt.Scanf("%d", &my_id)
	bully.my_id = my_id
	my_ip := bully.ids_ip[bully.my_id]
	address, err := net.ResolveTCPAddr("tcp", my_ip) 
	if err != nil{
		log.Fatal(err)
	}
	inbound, err := net.ListenTCP("tcp", address)
	if err != nil{
		log.Fatal(err)
	}
	rpc.Register(&bully)
	fmt.Println("Nó está executando com o socket: ", address)
	go rpc.Accept(inbound)

	fmt.Println("Iniciando as eleições!")
	invokeElection()
	random := ""
	for{
		fmt.Printf("Prossiga %d para comunicar com o coordenador.\n", bully.my_id)
		fmt.Scanf("%s", &random)
		communicateToCoordinator()
		fmt.Println("")
	}
	fmt.Scanf("%s", &random)
}
