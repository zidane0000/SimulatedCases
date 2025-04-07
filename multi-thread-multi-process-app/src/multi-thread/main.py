import threading
import time

class WorkerThread(threading.Thread):
    def __init__(self, thread_id, name, counter):
        threading.Thread.__init__(self)
        self.thread_id = thread_id
        self.name = name
        self.counter = counter

    def run(self):
        print(f"Starting {self.name}")
        self.perform_task()
        print(f"Exiting {self.name}")

    def perform_task(self):
        time.sleep(self.counter)
        print(f"{self.name} completed task after {self.counter} seconds")

def main():
    threads = []
    num_threads = 5

    for i in range(num_threads):
        thread = WorkerThread(i + 1, f"Thread-{i + 1}", i + 1)
        threads.append(thread)
        thread.start()

    for thread in threads:
        thread.join()

if __name__ == "__main__":
    main()