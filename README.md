# MC714-Trabalho2

Autores | RA
-----|----
Thiago Henrique da Costa |		206234 
Hitalo Cesar Alves |			217878 

Este projeto faz parte do curso MC714 - Sistemas Distribuídos da Unicamp, ministrado pelo PhD Luiz Bittencourt.

Em que abordamos três implementações de algoritmos referentes a sistemas distribuídos voltados para o deploy na nuvem. Sendo eles:

1. Lamport TimeStamp
2. Mutal Exclusion Suzuki-Kasami
3. Bully Algorithm

## Requerimentos

Para utilização dos algoritmos de Lamport TimeStamp e Mutual Exclusion Suzuki-Kasami, é requerido Python3 e os seguintes pacotes:
```bash
sudo apt install libopenmpi-dev -y
sudo apt install python3-pip -y
pip3 install mpi4py
```

Para utilização do Bully Algorithm, apenas é requerido Go:
```bash
sudo apt install golang -y
```

### Uso

Os algoritmos de Lamport e Suzuki-Kasami seguem o mesmo modelo de execução, aonde numProcess é o número de processos presente na dinâmica.

```bash
mpirun -mca btl ^openib -np numProcess python3 timestampLamport.py
mpirun -mca btl ^openib -np numProcess python3 meSuzukiKasami.py
```
Exemplificando:
Desejo executar o timestampLamport com 4 processos, ficará assim:
```bash
mpirun -mca btl ^openib -np 4 python3 timestampLamport.py
```

Já para execução do Bully, é simples:
```bash
go run bullyAlgorithm.go
```
Único ponto que devemos nos atentar é o seguinte trecho:
```
var bully = BullyAlgorithm{
	my_id: 		1,
	coordinator_id: 4,
	ids_ip: 	map[int]string{	1:"10.128.0.5:3030", 2:"10.128.0.6:3030", 3:"10.128.0.7:3030", 4:"10.128.0.8:3030"}}
```
Neste exemplo estabeleci a conversa entre os 4 sockets, se for desejável mudar a quantidade de IDs basta mudar a variável coordinator_id e também colocar os sockets referente a sua realidade.


Também é possível executar esse código de maneira local, apenas mudando será necessário identificar as portas, exemplifica-se:
```
var bully = BullyAlgorithm{
	my_id: 		1,
	coordinator_id: 4,
	ids_ip: 	map[int]string{	1:"127.0.0.1:3030", 2:"127.0.0.1:3031", 3:"127.0.0.1:3032", 4:"127.0.0.1:3033"}}
```

