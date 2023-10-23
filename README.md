# (WIP) madq - A high performance, message-oriented, distributed queue

## Overview

- Madq is a messaging queue built in Go. This project aims to provide a robust, scalable and custom-built message queue system. It uses a filesystem-like architecture internally for storing and managing messages. It is designed to be used as a library in other projects.

## Running the project

- Make sure to have Go installed
- go build
- Execute the compiled binary to start the messaging queue.

## TODO

- Message Publishing API: Build APIs to allow users to publish messages to the queue.
- Message Consumption API: Implement APIs to consume messages from the queue.
- Use of locks & atomic operations to make it concurrent-safe
- low level file system operation module  for high performance
- Expand madq.go: Flesh out the main program to provide more functionalities.

- Error Handling in flush.go: Implement robust error handling, particularly in the handleOps method.
- Unit Tests: Create unit tests for existing and new features to ensure reliability.
- Optimization in inode.go: Explore ways to optimize the inode system for better performance and error recovery.
- Build a CLI tool for this?
- Concurrency: Evaluate and optimize for concurrent operations.
- Monitoring & Logging: Add capabilities for monitoring queue health and message status.
- Message Retention Policies: Add support for setting message retention policies.
