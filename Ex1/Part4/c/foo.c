#include <pthread.h>
#include <stdio.h>

int i = 0;

// Note the return type: void*
void* incrementingThreadFunction(){
    for(int c = 0; c < 10000000; c++) {
       i++;
    }
    return NULL;
}

void* decrementingThreadFunction(){
    for(int c = 0; c < 10000000; c++) {
       i--;
    }
    return NULL;
}


int main(){
    printf("Start");
    pthread_t incrementingThread, decrementingThread;

    pthread_create(&incrementingThread, NULL, incrementingThreadFunction, NULL);
    pthread_create(&decrementingThread, NULL, decrementingThreadFunction, NULL);

    pthread_join(incrementingThread, NULL);
    pthread_join(decrementingThread, NULL);

    printf("The magic number is: %d\n", i);
    return 0;
}
