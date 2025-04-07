# Multi-Threaded and Multi-Process Application

## Overview

This project demonstrates the implementation of a multi-threaded and multi-process application. It starts with a multi-thread architecture that utilizes multiple threads to perform tasks concurrently and expands to a multi-process architecture where each process handles tasks independently. This approach showcases the benefits of both threading and multiprocessing in Python.

## Project Structure

```
multi-thread-multi-process-app
├── src
│   ├── multi-thread
│   │   ├── main.py        # Entry point for the multi-thread application
│   ├── multi-process
│   │   ├── main.py        # Entry point for the multi-process application
├── requirements.txt        # Lists the dependencies required for the project
└── README.md               # Documentation for the project
```

## Getting Started

### Prerequisites

Make sure you have Python installed on your machine. You can download it from [python.org](https://www.python.org/downloads/).

### Installation

1. Clone the repository:
   ```
   git clone https://github.com/zidane0000/SimulatedCases.git
   cd multi-thread-multi-process-app
   ```

2. Install the required dependencies:
   ```
   pip install -r requirements.txt
   ```

### Running the Applications

#### Multi-Thread Application

To run the multi-thread application, navigate to the `src/multi-thread` directory and execute the following command:

```
python main.py
```

This will initialize the main thread and create multiple worker threads to perform tasks concurrently.

#### Multi-Process Application

To run the multi-process application, navigate to the `src/multi-process` directory and execute the following command:

```
python main.py
```

This will create multiple processes, each handling tasks independently, showcasing the benefits of multiprocessing.
