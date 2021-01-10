# Introduction
This repository is a test task implementation.

It is necessary to implement a system that is used by two types of users:
- An administrator who can influence the state of the system. It is assumed that only one administrator can work at a time.
- An observer who can watch the system change. The number of observers watching the system changes is not limited.

In his interface, the administrator sees a red square, the position of which he can change.
The observer interface displays the same red square and all the changes that the administrator makes in real time.
Requirements:
- the delay between changes made by the administrator and their display to observers should be minimal
- the service must be scalable