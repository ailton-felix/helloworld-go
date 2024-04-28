# Site Monitor

This is a simple website monitor developed in Go. It allows monitoring the status of specified sites in a text file and records the results in a log file.

## Features

- **Website Monitoring:** Periodically checks the status of sites listed in a text file.
- **Log Recording:** Records monitoring results in a log file for later reference.
- **Log Display:** Allows viewing stored monitoring records in the log file.

## How to Use

1. **Code Compilation:**
   Make sure you have Go installed on your system. Compile the code with the following command:

2. **Running the Program:**
Run the compiled program:

3. **Choose an Option:**
The program will display a menu with the following options:
- `1` to start monitoring the sites.
- `2` to display the monitoring records.
- `0` to exit the program.

4. **Website Monitoring:**
During monitoring, the program will periodically check the status of the specified sites. The results will be displayed in the console.

5. **Log Recording:**
The monitoring results will be recorded in a log file located at `./data/log.txt`.

6. **Log Display:**
By selecting option `2`, the monitoring records will be displayed in the console.
