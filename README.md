# tcl (TCP Check List)

`tcl` is a simple command-line tool to check the health of a list of web services. It sends an HTTP GET request to each URL specified in a JSON file and displays their status in a table.

## Features

-   Checks the status of multiple URLs from a JSON file.
-   Displays results in a clean, tabular format.
-   Configurable timeout for HTTP requests (currently hardcoded at 2 seconds).
-   Simple and easy to use.

## Installation

To install `tcl`, you need to have Go installed on your system.

```bash
git clone https://github.com/komisan19/tcl.git
cd tcl
go build -o tcl .
```

This will create a `tcl` executable in the project directory.

## Usage

You can run `tcl` from the command line:

```bash
Usage of tcl:
  -f string
        target json file
  -v    show version
```

### Options

-   `-f`: (Required) Path to the JSON file containing the target URLs.
-   `-v`: (Optional) Show the version information.

## Configuration

You need to create a JSON file with a list of targets to check. The file should have the following format:

```json
{
  "targets": [
    {
      "name": "Service Name 1",
      "url": "https://example.com/"
    },
    {
      "name": "Service Name 2",
      "url": "https://example.org/"
    }
  ]
}
```

## Example

1.  Create a file named `targets.json` with the following content:

    ```json
    {
      "targets" : [
        {
          "name" : "yahoo",
          "url" : "https://www.yahoo.co.jp/"
        },
        {
          "name" : "google",
          "url" : "https://www.google.co.jp/"
        },
        {
          "name" : "404 Not Found",
          "url" : "https://example.com/404.html"
        }
      ]
    }
    ```

2.  Run `tcl` with the `-f` flag:

    ```bash
    ./tcl -f="targets.json"
    ```

3.  The output will be:

    ```
    +---------------+------------------------------+--------+
    |     NAME      |             URL              | STATUS |
    +---------------+------------------------------+--------+
    | yahoo         | https://www.yahoo.co.jp/     |    200 |
    | google        | https://www.google.co.jp/    |    200 |
    | 404 Not Found | https://example.com/404.html |    404 |
    +---------------+------------------------------+--------+
    ```

## License

This project is licensed under the MIT License.