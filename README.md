## Map Reduce

Map Reduce addresses the core architecture of how to process massive amounts of datasets across thousands of commodity, unreliable machines without exposing developers to the complexities of distributed coordination, fault tolerance, and network communication.

### Core architecture

Map Reduce is the backbone of Hadoop processing.

- It offers a framework that splits jobs into multiple tasks and executes them in parallel across a cluster and finishes by merging the results of the same.
- Its design ensures that we have:
  1. Parallelism
  2. Fault tolerance
  3. Data locality
  4. Scalability

It is ideal for operations like log analysis, machine learning, indexing, and recommendation systems, among others.

### Components

a) Client: entry point into Map Reduce, submits the job for processing by packing the mapper and reducer logic into a JAR file and specifying the input and output.

b) Job: Represents a complete processing request from the Client.

c) Master Node (Map Reduce Master): Coordinates job execution.

- Accepts jobs from the Client
- Splits jobs into parts
- Assigns these parts into the map and reduce tasks
- Tracks execution and reassigns failed tasks

d) Job Tasks: Every job is divided into smaller units called job parts.

  i) Map Job Parts - Process input data and split it into intermediary key-value pairs
  ii) Reduce Job Parts - Aggregate intermediary results into final parts

e) Map Phase:

Input data is split and given to Map tasks.
Each map task processes its own split and submits intermediary results in the form of key-value pairs.

f) Shuffle & Sort (between map and reduce)

