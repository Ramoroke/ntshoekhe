# Ntsoekhe Database

Ntsoekhe is a distributed database system designed to provide users with the ability to create nodes using SQLite and Go as the primary tools. It employs a peer-to-peer distribution model, where the first node created establishes a cluster, and subsequent nodes join this cluster. Data stored in each node is fully replicated across all other nodes in the cluster.

## Features

- **Distributed Architecture**: Utilizes a peer-to-peer distribution model for enhanced scalability and fault tolerance.
- **SQLite Backend**: Data storage and management are handled using SQLite, a lightweight and efficient embedded database engine.
- **Go Language**: Developed primarily using Go, a versatile and powerful programming language known for its concurrency support.
- **Automatic Cluster Formation**: The first node created initiates a cluster, with subsequent nodes automatically joining the cluster upon creation.
- **Full Data Replication**: Ensures that data stored in each node is fully replicated across all other nodes, ensuring data consistency and availability.
- **Node Management**: Each node manages its own database file, creating it if it doesn't exist, and appropriately handles deletion events, notifying other nodes of its removal from the cluster.


## Notice

🚧 **This repository is under construction.** 🚧

The development of Ntsoekhe Database is ongoing. Feel free to explore the code and contribute to the project. However, please note that the current state of the repository may be incomplete or unstable.


## Getting Started

To get started with Ntsoekhe Database, follow these steps:

1. **Install Go**: Ensure that you have Go installed on your system. You can download and install it from the official [Go website](https://golang.org/).

2. **Install SQLite**: If SQLite is not already installed on your system, you can download and install it from the official [SQLite website](https://sqlite.org/download.html).

3. **Install Node.js**: If Node.js is not already installed on your system, you can download and install it from the official [Node.js website](https://nodejs.org/).

4. **Clone the Repository**: Clone the Ntsoekhe Database repository from Git:

   ```bash
   git clone https://github.com/yourusername/ntsoekhe.git


To run the code after downloading:

1. **Build and Run Backend**:
   - Open the terminal in the folder containing the code.
   - Run the command `go run main.go`.

2. **Setup Frontend**:
   - Go to the directory that has the front-end code.
   - Run the command `npm install`.
   - After completion of the previous command, run another command `npm start`.

3. **Access the Application**:
   - Copy the localhost link provided after starting the frontend (`npm start`) and paste it into your browser.

## To Create a Node

To run the program for the first time:

1. **Build the Program**:
   - Run `go build main.go`.

2. **Create a Node**:
   - Run `./main createNode "node-name" "node-port"`.
   - The above step can be repeated in a different terminal window or tab but in the same folder to simulate creating a second node.

## Contributing

Contributions to Ntsoekhe Database are welcome! If you find any bugs, have feature requests, or would like to contribute enhancements, please open an issue or submit a pull request on the [GitHub repository](https://github.com/yourusername/ntsoekhe).

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Acknowledgments

Ntsoekhe Database was inspired by the need for a distributed database solution with simplicity, scalability, and reliability in mind. We would like to thank the open-source community for their contributions and support.


