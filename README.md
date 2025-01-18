# Kubernetes course

## Exercises

- [Part 1](part1)
  - [Exercise 1.01](part1/exercise1.01.md)
  - [Exercise 1.02](part1/exercise1.02.md)
  - [Exercise 1.03](part1/exercise1.03.md)
  - [Exercise 1.04](part1/exercise1.04.md)
  - [Exercise 1.05](part1/exercise1.05.md)
  - [Exercise 1.06](part1/exercise1.06.md)
  - [Exercise 1.07](part1/exercise1.07.md)
  - [Exercise 1.08](part1/exercise1.08.md)
  - [Exercise 1.09](part1/exercise1.09.md)
  - [Exercise 1.10](part1/exercise1.10.md)
  - [Exercise 1.11](part1/exercise1.11.md)
  - [Exercise 1.12](part1/exercise1.12.md)
  - [Exercise 1.13 (Image 1)](part1/exercise1.13.png)
- [Part 2](part2)
  - [Exercise 2.01](part2/exercise2.01.md)
  - [Exercise 2.02 (Image 1)](part2/exercise2.02_1.png)
  - [Exercise 2.02 (Image 2)](part2/exercise2.02_2.png)
  - [Exercise 2.02 (Image 3)](part2/exercise2.02_3.png)
  - [Exercise 2.03](part2/exercise2.03.md)
  - [Exercise 2.04](part2/exercise2.04.md)
  - [Exercise 2.06](part2/exercise2.06.md)
  - [Exercise 2.07](part2/exercise2.07.md)
  - [Exercise 2.08 (Image 1)](part2/exercise2.08_01.png)
  - [Exercise 2.08 (Image 2)](part2/exercise2.08_02.png)
  - [Exercise 2.09](part2/exercise2.09.md)
  - [Exercise 2.09 (Image 1)](part2/exercise2.09.png)
  - [Exercise 2.10 (Image 1)](part2/exercise2.10.png)
- [Part 3](part3)
  - [Exercise 3.01](part3/exercise3.01.md)
  - [Exercise 3.02](part3/exercise3.02.md)
  - [Exercise 3.03](part3/exercise3.03.md)
  - [Exercise 3.04](part3/exercise3.04.md)
  - [Exercise 3.05](part3/exercise3.05.md)
  - [Exercise 3.08](part3/exercise3.08.md)
  - [Exercise 3.09](part3/exercise3.09.md)
  - [Exercise 3.10 (Image 1)](part3/exercise3.10.png)

## DBaaS vs DIY

### DBaaS

DBaaS stands for Database as a service. It lets users access and use database software without purcahsing and setting up infrastucture, software, or managing system themselves.

Benefits include cost savings on managing databases.
DBaaS scale up and down allowing users quickly and easily provision extra storage and computing capacity.
DBaaS can speed up time-to-market by allowing the developer teams manage databases cababilities themselves.
DBaas typically offer enterprice-grade security out of the box.[0] Using DBaaS users can typically off load tasks like data backup and recovery to the vendor or cloud platform.[3]

Cons of DBaaS include lack of access to the server and servers depending on cloud provider, and unexpected costs if pre-paid usage limits are exceeded.[1]
Some DBaaS solutions might lock users into a vendor or cloud platform. There might me reafactoring costs moving vendors or cloud platforms[2].

### DIY

DIY hosting a Database means users managing the software, the infrastucture and the system themselfes, similar what we have been doing in this course so far.
DIY hosting a Database benefits include transparent costs and full control of the performance by having direct access to the infrastructure.[3]

Cons of the DIY hosting a Database are almost 1-to-1 with the Pros of the DBaaS, managing a DIY Database takes resources[0].
Even though users have control over the performance of the infrastucture themselves, horizontal scalability of systems might be difficult to implement, compated to DBaaS where this usually done automatically[3]. DIY hosting a Database users need to manage the backup and software updates them selves, which is a hidden cost.[1]

### Verdict

For a new software startup I would recommend a DBaaS solution, this allows product development teams be more effective building features quickly. For bigger enterprice companies having full control over infrastucture might be beneficial in terms of cost savings.

### Sources

[0] https://www.ibm.com/think/topics/dbaas
[1] https://scalegrid.io/blog/dbaas-vs-self-managed-cloud-databases/
[2] https://aiven.io/blog/5-tips-for-choosing-a-dbaas-vendor
[3] https://www.geeksforgeeks.org/overview-of-database-as-a-service/
