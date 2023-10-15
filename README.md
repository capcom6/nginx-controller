# Nginx Controller

[![License](https://img.shields.io/badge/license-Apache%202.0-blue)](LICENSE)

Nginx Controller is a service that allows you to control the configuration of an Nginx reverse-proxy through a JSON REST API. It provides endpoints to add new virtual hosts with multiple upstreams, generate Nginx configuration files based on templates, and reload Nginx to apply the changes.

## Table of contents

- [Nginx Controller](#nginx-controller)
  - [Table of contents](#table-of-contents)
  - [Features](#features)
  - [Getting Started](#getting-started)
    - [Prerequisites](#prerequisites)
    - [Quickstart](#quickstart)
  - [Usage](#usage)
    - [Configuration](#configuration)
    - [API](#api)
    - [Security](#security)
  - [Roadmap](#roadmap)
  - [Contributing](#contributing)
  - [License](#license)


## Features

- Add new virtual hosts with multiple upstreams
- Generate Nginx configuration files based on templates
- Reload Nginx to apply the configuration changes
- Remove virtual hosts from the configuration

## Getting Started

### Prerequisites

- Go 1.20 or higher
- Nginx installed on the system

### Quickstart

1. Clone the repository:

   ```shell
   git clone https://github.com/capcom6/nginx-controller.git
   ```
2. Change to the project directory:
   ```shell
   cd nginx-reverse-proxy-controller
   ```
3. Build the project:
   ```shell
   make build
   ```
4. Copy and edit example `config.yaml`:
   ```shell
   cp configs/config.example.yaml config.yaml
   vi config.yaml
   ```
5. Run the project:
   ```shell
   ./nginx-controller
   ```
6. Create first virtual host:
   ```
   curl http://localhost:3000/api/v1/hosts/example.local -X PUT -H "Content-Type: application/json" -d '{ \
        "upstream": [ \
            { \
                "host": "example.com", \
                "port": 80, \
                "weight": 1 \
            } \
        ] \
   }'
   ```

## Usage

### Configuration

The configuration file is stored in the `config.yaml` file. The file is a YAML document that describes the configuration of the controller. The following is an example of a valid configuration file:

```yaml
---
# nginx-related configuration
nginx:
    location: /etc/nginx/conf.d/ # where to put virtual hosts configs
    template:
        | # template for virtual host configs, seet `text/template` package description for syntax
        upstream {{ .UpstreamName }} {
            {{ range $server := .Servers }}
                server {{ $server.Host }}:{{ $server.Port }} {{ $server.Options }};
            {{ end }}
        }

        server {

            listen 80;
            server_name {{ .DomainName }};

            location / {
                proxy_pass http://{{ .UpstreamName }};

                include    proxy.conf;
            }
        }
```

By default the controller uses `config.yaml` file in local directory. You can change the configuration file location by setting the environment variable `CONFIG_PATH` to a custom path.

### API

You can find controller API documentation in OpenAPI format [here](./api/swagger.yaml). Some examples available in [requests.http](./api/requests.http) file.

### Security

The application does not include any built-in authorization or authentication mechanism. To enhance security, it is recommended to use an external service or tool, such as Nginx, to provide basic authentication and restrict access to the application. Alternatively, deploying the application in a private network can provide an additional layer of security and help prevent unauthorized access from external sources.

## Roadmap

- **TLS/SSL support:** Add support for managing TLS/SSL certificates and configuring HTTPS virtual hosts.
- **Configuration validation:** Add validation checks to ensure that the generated Nginx configuration files are valid and error-free.
- **Web-based administration interface:** Develop a web-based interface for managing virtual hosts and upstreams, providing a user-friendly alternative to the API.

## Contributing

Contributions are welcome! If you find any issues or have suggestions for improvement, please open an issue or submit a pull request.

## License

This project is licensed under the Apache License 2.0. See the LICENSE file for details.