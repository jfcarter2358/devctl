```
     _                _   _
  __| | _____   _____| |_| |
 / _` |/ _ \ \ / / __| __| |
| (_| |  __/\ V / (__| |_| |
 \__,_|\___| \_/ \___|\__|_|
 ```

A command line utility for developing within Docker containers

# Getting Started

To begin using `devctl`, follow the below steps:
1. Download a binary from the latest GitHub release
2. Place said binary into your path/modify your path to include the location of the binary
3. Navigate to the directory containing the code you want to work on in your terminal
4. Run `devctl init` to generate a base `.devctl.json` file
5. Modify the `.devctl.json` file to suit your needs
6. Run `devctl up` to bring up the docker container and connect to it

# .devctl.json Format

The `.devctl.json` file contains a series of fields to control how the Docker container you want to develop in is brought up. An explanation of the fields along with some example values can be found below


| Field    | Default     | Example                     | Description                                                                             |
| -------- | ----------- | --------------------------- | --------------------------------------------------------------------------------------- |
| `name`   | `"foobar"`  | `"festive_kepler"`          | The name to give to the container that runs in this environment                         | 
| `image`  | `"busybox"` | `"my-favorite-image:1.0.0"` | The image to spawn a container of and connect to                                        |
| `shell`  | `"sh"`      | `"bash"`                    | What shell to launch when connecting to the container                                   |
| `user`   | `"root"`    | `"1001"`                    | What user to connect to the container as                                                |
| `mounts` | `[]`        | `["foobar:/dev/foobar"]`    | Any mounts other than the included current directory that you want available            |
| `args`   | `[]`        | `--network="none"`          | Any additional docker run arguments that should be included when spawning the container |
| `ports`  | `[]`        | `[80, 443]`                 | Ports to expose to the outside world from the container                                 |

