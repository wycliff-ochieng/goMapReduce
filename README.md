## DISTRIBUTED SYSTERMS 
### CASE STUDY : MapReduce(Google internal data processing tool) - Open source version is Apache Spark

Distributed Systems focuses on three major themes **Fault Tolerance** , **Consistency** and **Performance**

### 1. The Big Picture: Why MapReduce?
The lecture notes introduce three major themes in distributed systems: **Fault Tolerance, Consistency, and Performance**. MapReduce is used as the perfect case study because it tackles all three:
*   **Performance (Scalability):** By splitting work into independent tasks, MapReduce allows hundreds or thousands of machines to process terabytes of data in parallel. 
*   **Fault Tolerance:** It assumes commodity hardware will fail. The framework hides these failures from the programmer by automatically re-running failed tasks.
*   **Simplicity (Hiding Complexity):** The programmer only writes simple, sequential `Map()` and `Reduce()` functions. The framework handles the messy details of networking, parallelization, and failure recovery.

### 2. Mapping the Lecture to the Paper's Details
The lecture highlights specific engineering choices from the paper that make MapReduce successful:

*   **Network Bottlenecks & Locality (Paper Section 3.4):** 
    *   *The Problem:* In 2004, network bandwidth was the biggest bottleneck (the lecture estimates only ~55 megabits/second/machine available at the root switch).
    *   *The Solution:* The Coordinator (Master) schedules Map tasks on the exact same machines where the input data is stored on disk via GFS (Google File System). This avoids sending massive amounts of raw input data over the network.
*   **Task Granularity & Load Balancing (Paper Section 3.5):** 
    *   *The Problem:* Machines compute at different speeds, and data chunks vary in size. 
    *   *The Solution:* The system creates far more tasks than there are worker machines (e.g., 200,000 tasks for 2,000 machines). Faster machines naturally churn through more tasks, preventing slow machines from holding up the entire job.
*   **Fault Tolerance & Determinism (Paper Section 3.3):** 
    *   *The Problem:* What happens when a machine crashes mid-job?
    *   *The Solution:* The Coordinator simply re-assigns the task to a new machine. However, the lecture emphasizes a crucial constraint: **Map and Reduce functions must be purely deterministic.** Because the system might run the same task twice (due to failures or timeouts), they cannot have side effects or rely on random numbers, ensuring consistent output regardless of crashes.
*   **Stragglers (Paper Section 3.6):** 
    *   *The Problem:* A single machine with a bad disk can slow down the end of the entire operation.
    *   *The Solution:* The Coordinator schedules "backup tasks" for the final remaining in-progress tasks. Whichever finishes first (the original or the backup) is accepted.
*   **Performance Analysis (Paper Section 5.2 & Figure 2):** 
    *   The lecture dissects the "Grep" benchmark, noting the massive 30 GB/s throughput. The notes point out that this is only possible because the data is being read locally from disk across 1,764 workers, completely bypassing the network bottleneck.

### 3. Built my own MapReduce in Go. 
implementation will need to handle:
1.  **A Coordinator** that tracks the state of all Map and Reduce tasks (idle, in-progress, completed).
2.  **RPCs (Remote Procedure Calls)** to hand out tasks to workers.
3.  **Intermediate Files:** Saving Map outputs to local disk, partitioned by a hash function so Reduce workers know which files to fetch.
4.  **Failure Detection:** Using timeouts to detect if a worker has crashed, and re-assigning that task to someone else.

