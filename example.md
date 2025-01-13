# Project Setup and Running

This project is a Node Driver implementation for the Container Storage Interface (CSI). You can set up and run the project by following the steps below.

## Requirements

- Go (1.16 or higher)
- Use of `go mod` for project dependencies

## Installation

1. **Navigate to the Project Directory:**
   ```bash
   cd /path/to/your/project
   ```

2. **Initialize the Go Module:**
   ```bash
   go mod init multicloud-custom-csi-driver
   ```

3. **Install Dependencies:**
   To install the project dependencies, run the following command:
   ```bash
   go mod tidy
   ```

4. **Build the Project:**
   Use the following command to build the project:
   ```bash
   go build -o csi-node-driver ./node
   ```

## Running

1. **Start the Server:**
   Start the server with the following command:
   ```bash
   ./csi-node-driver
   ```

2. **Check if the Server is Running Successfully:**
   Once the server starts, you should see the message "Starting CSI node server..." in the terminal.

## Usage

This project can be used as a CSI-compliant storage driver. It supports the necessary CSI protocols and requests.

## Contributing

If you'd like to contribute, please create a pull request or report any issues.

## License

This project is licensed under the [MIT License](LICENSE).