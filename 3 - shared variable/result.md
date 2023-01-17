Both programs where never able to return 0. This is because both threads are operating on the same shared variable. When the program switches threads, problems can occur if one thread is trying to write to the variable, but its interrupted before it 
can finish.
