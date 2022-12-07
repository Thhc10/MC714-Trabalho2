from mpi4py import MPI
import sys


def local_time(cc):
    return '(LAMPORT_TIME={})'.format(cc)


def calc_recv_timestamp(recv_time_stamp):
    return max(recv_time_stamp, counter) + 1


def event(rank, cc):
    cc += 1
    print('%s | Evento realizado pelo %d' % (local_time(cc), rank))
    return cc


def send_message(dest_rank, rank, cc, comm_a):
    cc += 1
    a = comm_a.isend(cc, dest=dest_rank, tag=11)
    a.wait()
    print('%s | Mensagem enviada de %d para %d' % (local_time(cc), rank, dest_rank))
    sys.stdout.flush()
    return cc


def recv_message(sender_rank, rank, cc, comm):
    a = comm.irecv(source=sender_rank, tag=11)
    timestamp = a.wait()
    cc = calc_recv_timestamp(timestamp)
    print('%s | Mensagem recebida pelo %d de %d' % (local_time(cc), rank, sender_rank))
    sys.stdout.flush()
    return cc


counter = 0
comm = MPI.COMM_WORLD
size = comm.Get_size()
myRank = comm.Get_rank()

print("Inicializando o processo %d" % myRank)

if myRank == 0:
    counter = event(myRank, counter)
    counter = send_message(1, myRank, counter, comm)
    counter = event(myRank, counter)
    counter = recv_message(1, myRank, counter, comm)
    counter = event(myRank, counter)

elif myRank == 1:
    counter = recv_message(0, myRank, counter, comm)
    counter = send_message(0, myRank, counter, comm)
    counter = send_message(2, myRank, counter, comm)
    counter = recv_message(2, myRank, counter, comm)

elif myRank == 2:
    counter = recv_message(1, myRank, counter, comm)
    counter = send_message(1, myRank, counter, comm)

elif myRank == 3:
    counter = 0
    counter = event(myRank, counter)
