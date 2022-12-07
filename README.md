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

Para utilização dos algoritmos de Lamport TimeStamp e Mutal Exclusion Suzuki-Kasami, é requerido Python3 e os seguintes pacotes:
```bash
sudo apt install libopenmpi-dev -y
sudo apt install python3-pip -y
pip3 install mpi4py
```

Para utilização do Bully Algorithm, apenas é requerido Go:
```bash
sudo apt install golang -y
```

### Local without Docker Containers

The complete list of flags is:

```bash
$ python3 run.py -h                                                     

usage: run.py [-h] [-v] [-d] [-a {ring,bully}] [-c CONFIG_FILE]  

Implementation of distributed election algorithms. Generic node.  

optional arguments:   
    -h, --help            show this help message and exit   
    -v, --verbose         increase output verbosity   
    -d, --delay           generate a random delay to forwarding messages   
    -a {ring,bully}, --algorithm {ring,bully}                         
                            ring by default   
    -c CONFIG_FILE, --config_file CONFIG_FILE
                            needed a config file in json format
```

The _local_config.json_ (in _SDCC/sdcc_) file has been defined to manage all network settings (i.e., IP addresses, port numbers). To display the options set:

```bash
python3 run.py -h
```

Firstly you make running the register node:

```bash
# to execute in SDCC/sdcc/register. The -v flag provides a verbose execution (i.e., all messages received and sent are shown)
python3 run.py -v -c ../local_config.json
```

A single node can be executed as:

```bash
# to execute in SDCC/sdcc/node. Without the '-a bully' option node runs the ring-based alg.
python3 run.py -v -a bully -c ../local_config.json
```

#### Tests

Test execution can be handled as:

```python
# to execute in SDCC/sdcc
sudo python3 run_tests.py
```
