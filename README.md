# 🐜 **Lemin-Ant Project**

In this project, we have designed an advanced **graph-based pathfinding algorithm** to guide ants from a **'start' room** to an **'end' room** in the **shortest number of turns**, while ensuring that ants do not collide or overlap on their paths. The simulation operates on a graph structure created from coordinates and room connections defined in an input file.

## 🧠 **Project Overview**

- The algorithm utilizes the **Breadth-First Search (BFS)** algorithm to identify **all possible paths** between the start and end rooms.  
- Once all valid paths are found, they are **filtered** to eliminate overlapping routes.  
- The ants are then **distributed across the optimal paths** in such a way that **maximum efficiency** is achieved in each round.  
- The program carefully handles ant movements to ensure that no two ants occupy the same room (except the start and end rooms) in the same turn.  
- The primary goal is to **minimize the number of rounds required** to move all ants from the start room to the end room.

## ⚠️ **Warning**
 - When you run the project with the command ' go run main.go coordinate txt file ',
 - if you get go mod error with version error
 - go mod init ant
 - go mod tidy
 - commands to run it.
 
## 🚀 **Running the Project**

To run the project, use the following command:

```bash
go run main.go <filename>

Example:  
go run main.go example01.txt
go run main.go badexample01.txt
```

## License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.


## Contact
For any questions or feedback, please contact:
- **Email**: ahatdemirezenn@gmail.com
- **GitHub**: [Ahat Demirezen](https://github.com/ahatdemirezen)
