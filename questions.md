Exercise 1 - Theory questions
-----------------------------

### Concepts

What is the difference between *concurrency* and *parallelism*?
> Parallellism means running several tasks at the same time, concurrency means exploits the down-time of the system. Meaning that while the system is waiting for something, it is able to complete some other task.

What is the difference between a *race condition* and a *data race*? 
> Race condition: the order of tasks affects the result.
> Data race: One thread accesses the data, while another thread is writing to the same object.
> So a data race is a kind of race condition
 
*Very* roughly - what does a *scheduler* do, and how does it do it?
> A scheduler decides which threads to run, this is decided based on which threads a are to complete their tasks. The scheduler tries to regulary switch threads. 


### Engineering

Why would we use multiple threads? What kinds of problems do threads solve?
> It makes the code easier to read. In addition, it saves time when we have independent tasks, meaning they arent realient on eachother

Some languages support "fibers" (sometimes called "green threads") or "coroutines"? What are they, and why would we rather use them over threads?
> Fibers use cooperative multitasking, instead of preemptive. In addition, fibers start and stop in well defined places, meaning data integrity isnt an issue

Does creating concurrent programs make the programmer's life easier? Harder? Maybe both?
> It can be harder to implement. However, it can also create more readable and structured code.

What do you think is best - *shared variables* or *message passing*?
> Probably depends on the situation. However, using message passing allows for communincation between the threads. Shared variables seem easier to implement, and message passing seems to create a better structured program.


