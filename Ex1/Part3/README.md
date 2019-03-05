# Reasons for concurrency and parallelism


To complete this exercise you will have to use git. Create one or several commits that adds answers to the following questions and push it to your groups repository to complete the task.

When answering the questions, remember to use all the resources at your disposal. Asking the internet isn't a form of "cheating", it's a way of learning.

 ### What is concurrency? What is parallelism? What's the difference?
 > *Concurrency: When an application can execute two or more tasks at the same time, and that these executions do not affect the ending result/final outcome. The tasks will start, run and complete in overlapping time periods and in no specific ordering. They do not have to be executed at the same time, but they may also do. Parallelism: When the tasks actually runs in the exactly same time. This is possible with the use of multi-core infrastructure of the CPU by assigning one task to each core. The difference is thus that parallelism is a type of concurrency, but in concurrency tasks do not have to run at the exactly same time.*
 
 ### Why have machines become increasingly multicore in the past decade?
 > *When the improvement of clock speed became more difficult, the use of multicore processors solved the problem of wanting faster processing capability. It was also a solution of the problem of high power consumption with the increased frequency scaling. And of course with the option of multicore prosessors it makes it possible for concurrent tasks!*
 
 ### What kinds of problems motivates the need for concurrent execution?
 (Or phrased differently: What problems do concurrency help in solving?)
 > *Firstly because most processors nowadays are multicore we need to be able to write programs that uses these cores and the advantage of the increased efficiency that the multicore offer.  *
 
 ### Does creating concurrent programs make the programmer's life easier? Harder? Maybe both?
 (Come back to this after you have worked on part 4 of this exercise)
 > *No, it makes it harder. This is because the programmer must be able to split his program into reasonable and cooperating parts. And when these different parts are executed they are nondeterministic in their ordering. This makes it again hard to debug and test the system. *
 
 ### What are the differences between processes, threads, green threads, and coroutines?
 > *Your answer here*
 
 ### Which one of these do `pthread_create()` (C/POSIX), `threading.Thread()` (Python), `go` (Go) create?
 > *Your answer here*
 
 ### How does pythons Global Interpreter Lock (GIL) influence the way a python Thread behaves?
 > *Your answer here*
 
 ### With this in mind: What is the workaround for the GIL (Hint: it's another module)?
 > *Your answer here*
 
 ### What does `func GOMAXPROCS(n int) int` change? 
 > *Your answer here*
