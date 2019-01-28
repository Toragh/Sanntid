# Python 3.3.3 and 2.7.6
# python fo.py

from threading import Thread
from threading import Lock


# Potentially useful thing:
#   In Python you "import" a global variable, instead of "export"ing it when you declare it
#   (This is probably an effort to make you feel bad about typing the word "global")
i = 0

def incrementingFunction(lock):
    global i
    for c in range(0,1000000):
        lock.acquire()
        i += 1
        lock.release()

def decrementingFunction(lock):
    global i
    for c in range(0,1000000):
        lock.acquire()
        i -= 1
        lock.release()


def main():
    global i
    lock = Lock()


    incrementing = Thread(target = incrementingFunction, args = (lock,),)
    decrementing = Thread(target = decrementingFunction, args = (lock,),)

    incrementing.start()
    decrementing.start()

    incrementing.join()
    decrementing.join()

    print("The magic number is %d" % (i))


main()
