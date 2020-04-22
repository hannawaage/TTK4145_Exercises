# Mutex and Channel basics

### What is an atomic operation?
> An operation in concurrent programming that runs independently of other processess. 

### What is a semaphore?
> A variable that controlls access to common resources by multiple processes in a concurrent system. 

### What is a mutex?
> A mutex is a locking mechanism which is aquired by the running operation. It constitues a buffer which decides which threads to be run in concurrent programming. Only the owner can release the lock. 

### What is the difference between a mutex and a binary semaphore?
> A binary semaphore is a signal from one of the threads signalling whether it is done or not, while a mutex is a locking mechanism that can only be aquired by one task at a time. 

### What is a critical section?
> A part of a program that is protected due to shared variables, that can only be executed by one process at a time. 

### What is the difference between race conditions and data races?
 > A race condition is when the program result is dependent on the sequence or timing of the operations. A data race is a situation, in which at least two threads access a shared variable at the same time, where at least one of the thread tries to modify the variable.

### List some advantages of using message passing over lock-based synchronization primitives.
> Message passing is safe and convienient for scaling, there is no common resources and it is easier to implement than lock-based.

### List some advantages of using lock-based synchronization primitives over message passing.
> Can be used for real-time programming as shared variable-operations are instantanious. Simpler for smaller systems that wont be scaled much.  
