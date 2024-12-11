import time

def record_time(fun):
    def inner_func(*args, **kwargs):
        start_time = time.perf_counter_ns()
        res = fun(*args, **kwargs)
        end_time = time.perf_counter_ns()
        elapsed = get_elapsed_time(start_time, end_time)
        fun_name = f"{fun.__name__}"  # Get the function's name directly
        print(f"Elapsed time for {fun_name}: {elapsed}")
        return res
    return inner_func

def get_elapsed_time(start_ns, end_ns):
    total_nanoseconds = end_ns - start_ns
    total_seconds = total_nanoseconds / 1_000_000_000  # Convert ns to seconds
    hours = int(total_seconds // 3600)
    minutes = int((total_seconds % 3600) // 60)
    seconds = int(total_seconds % 60)
    milliseconds = int((total_nanoseconds % 1_000_000_000) // 1_000_000)
    microseconds = int((total_nanoseconds % 1_000_000) // 1_000)
    nanoseconds = int(total_nanoseconds % 1_000)
    return f"{hours:02}:{minutes:02}:{seconds:02}.{milliseconds:03}{microseconds:03}{nanoseconds:03}"

