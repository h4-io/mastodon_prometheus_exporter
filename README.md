# mastodon_prometheus_exporter

![Github License](https://img.shields.io/badge/license-APACHE2.0-green)
[![CodeFactor](https://www.codefactor.io/repository/github/h4-io/mastodon_prometheus_exporter/badge)](https://www.codefactor.io/repository/github/h4-io/mastodon_prometheus_exporter)

#### mastodon_prometheus_exporter is a free and opensource simple exporter for mastodon business metrics

This project tries to fetch openly accessible business metrics for mastodon like number of active users, number of posts... 

## Table of content

- [**Getting Started**](#getting-started)
- [Built With](#built-with)
- [Contributing](#contributing)
- [License](#license)
- [Get Help](#get-help)
- [Motivation](#motivation)
- [Acknowledgments](#acknowledgements)

## Getting Started

This project is still at an early stage of development. Please report any bugs or suggestion by creating an issue. 

### Install
- Get the [latest release](https://github.com/h4-io/mastodon_prometheus_exporter/releases)
- `mv mastodon_prometheus_exporter /usr/local/bin`
- `chmod +x /usr/local/bin/mastodon_prometheus_exporter`

### Usage
Export the needed env var:  
````
export MASTODON_INSTANCE="https://h4.io"
export EXPORTER_PORT=8080
````
Then run the app: `/usr/local/bin/mastodon_prometheus_exporter`

Alternatively you can edit the [systemd service exemple](https://github.com/h4-io/mastodon_prometheus_exporter/blob/main/mastodon_prometheus_exporter.service) and then move it to `/etc/systemd/system/` and do:
````
sudo useradd -s /sbin/nologin -M mastodon_exporter
sudo systemctl daemon-reload 
sudo systemctl enable mastodon_prometheus_exporter.service
sudo systemctl start mastodon_prometheus_exporter.service
````

### Dashboard
The Dashboard is available [on the Grafana website](https://grafana.com/grafana/dashboards/20361).

## Built With

This repository is built with:
- [Victoria Metrics](https://github.com/VictoriaMetrics/metrics)

## Contributing

See [CONTRIBUTING.md](https://github.com/h4-io/mastodon_prometheus_exporter/blob/master/CONTRIBUTING.md)  

## License

This project is licensed under the [Apache2.0 License](https://github.com/h4-io/mastodon_prometheus_exporter/blob/master/LICENSE)

## Get Help
- Contact us on Matrix h4-dev:matrix.org
- If appropriate, [open an issue](https://github.com/h4-io/mastodon_prometheus_exporter/issues) on GitHub

See [the code of conduct](https://github.com/h4-io/mastodon_prometheus_exporter/blob/master/code_of_conduct.md) befor contacting us.

## Motivation
This project was launched to provide a simple and free prometheus export for mastodon business metrics.  
We created it for our own needs on the [H4](https://h4.io) project.

## Acknowledgements
This repository use [Victoria Metrics](https://github.com/VictoriaMetrics/metrics) licensed under the MIT License.  
Please refer to the other repositories to checkout acknowledgements for each services. 

