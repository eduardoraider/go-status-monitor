# Website Monitoring with CLI GO Application

This repository serves as an illustrative demonstration of utilizing a GO (Golang) application to perform website monitoring and register status logs via the command-line interface (CLI). Explore this codebase to gain insights into effective website monitoring practices.

## Learning Topics

By studying this application, you can learn the following topics related to Go programming:

1. **Working with Variables**: Discover how to declare and use variables to store data in Go.

2. **Controlling the Flow of the Script**: Understand how to use conditional statements and loops to control the flow of your Go script.

3. **Making Web Requests**: Learn how to make HTTP GET requests to monitor websites and retrieve their status.

4. **Go's Top Collections**: Explore Go's built-in data structures like slices and maps for efficient data management.

5. **Reading Files**: Understand how to read data from files, such as the list of websites to be monitored.

6. **Writing Files**: Learn how to write logs and other data to files for record-keeping.

## Changing Monitored Sites

To change the websites that are monitored by this application, you can edit the `urls.txt` file located in the project directory. Each line in this file should contain a single URL that you want to monitor. Simply add or remove URLs as needed, and the application will automatically update its monitoring list the next time it runs.

## Running the Application

To run the website monitoring application, follow these steps:

1. Ensure you have Go installed on your system. If not, you can download and install it from the official [Go website](https://golang.org/dl/).

2. Open a terminal or command prompt and navigate to the project directory.

3. Run the following command to start the application:

   ```bash
   go run hello.go
   ```

   This will execute the `hello.go` script, which contains the website monitoring logic. The application will start monitoring the websites listed in the `urls.txt` file and display status updates in the terminal.

4. After running the application, you will see a menu in the terminal:

   ```
   1 - Start Monitoring
   2 - Show logs
   0 - Exit
   ```

   You can select one of the options:

    - **Option 1**: Start Monitoring - This option will begin the website monitoring process, and the application will periodically check the status of the websites listed in `urls.txt`.

    - **Option 2**: Show Logs - This option will display the monitoring logs stored in the `log.txt` file. You can review the history of website status changes.

    - **Option 0**: Exit - Select this option to exit the application.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE.txt) file for details.

---

#### by Eduardo O Raider
🛠 🥋 **Software Engineer**