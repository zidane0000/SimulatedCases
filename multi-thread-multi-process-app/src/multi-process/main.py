import multiprocessing
import time

class WorkerProcess(multiprocessing.Process):
    def __init__(self, task_queue):
        super().__init__()
        self.task_queue = task_queue

    def run(self):
        while True:
            task = self.task_queue.get()  # Get a task from the queue
            if task is None:  # Check for the sentinel value
                break  # Exit the loop if sentinel is received
            self.execute_task(task)

    def execute_task(self, task):
        # Simulate task execution
        print(f"{self.name} executing task: {task}")
        time.sleep(1)  # Simulate some work

def main():
    num_processes = 4  # Number of processes to create

    # Create a shared task queue using multiprocessing.Manager
    with multiprocessing.Manager() as manager:
        task_queue = manager.Queue()

        # Populate the task queue with tasks (example tasks)
        for i in range(10):
            task_queue.put(f"Task {i}")

        # Add sentinel values to signal the workers to stop
        for _ in range(num_processes):
            task_queue.put(None)

        processes = []

        for _ in range(num_processes):
            process = WorkerProcess(task_queue)
            processes.append(process)
            process.start()

        # Wait for all processes to complete
        for process in processes:
            process.join()

if __name__ == "__main__":
    main()