## Spartan-III Chain (Powered by NC PolygonEdge)

# Introduction

A Non-Cryptocurrency Public Chain is a transformed public chain framework based on an existing public chain. Gas Credit transfers are not permitted between standard wallets. There are no cryptocurrency incentives for mining or participating in consensus. On Spartan Network, there are three Non-Cryptocurrency Public Chains at launch. We except to add more in the foreseeable future.

> As a clear demonstration, all commands in this document are run with root permission. These commands can also be run under normal user permissions, please set the file storage and configure the parameters properly.

## 1. About Spartan-III Chain (Powered by NC PolygonEdge)

This document is a guide to install, configure and run a full node in the Spartan-III Chain(Powered by NC-PolygonEdge) public blockchain.

Spartan-III Chain (Powered by NC PolygonEdge) network has two identifiers, a network ID and a chain ID. Although they often have the same value, they have different uses. Peer-to-peer communication between nodes uses the network ID, while the transaction signature process uses the chain ID.

**Spartan-III Chain Network Id = Chain Id = 5566**

Below is the instruction for Linux.

## 2. Hardware Requirements

It is recommended to build Spartan-III Chain full nodes on Linux Server with the following requirement.

#### Minimum Requirements:

- 2 CPU
- Memory: 2GB
- Disk: 100GB SSD
- Bandwidth: 20Mbps

#### Recommended Requirements

- 4 CPU
- Memory: 16GB
- Disk: 512GB SSD
- Bandwidth: 20Mbps

## 3. Node Installation

There are 2 methods to install NC PolygonEdge Node: building from source and installing by Docker. Please refer to the installation method that is most applicable in your specific case.

### 3.1 Building from Source

**Go 1.17** or above is recommended for building and installing the Spartan-III Chain node software. Install `go` by the following steps:

Download and untar the installation file

```
wget https://go.dev/dl/go1.18.5.linux-amd64.tar.gz
tar -C /usr/local -zxvf go1.18.5.linux-amd64.tar.gz
```

Change environment variables, for example in bash:

```
vi /etc/profile
```

Insert the parameter at the bottom of the file:

```shell
export PATH=$PATH:/usr/local/go/bin
```

Then, save the change and make the /etc/profile file take effect:

```
source /etc/profile
```

Now, check whether `go` is correctly installed:

```
go version
```

![](https://raw.githubusercontent.com/BSN-Spartan/NC-PolygonEdge/main/.github/images/1.go_version.jpg)

Before compiling the source code, make sure that `gcc` has been successfully installed. If not, please install `gcc` first. Check by the following command:

```shell
gcc -v
```

![](https://raw.githubusercontent.com/BSN-Spartan/NC-PolygonEdge/main/.github/images/2.%20gcc.jpg)

Download the source code of Spartan NC PolygonEdge from github (git has been installed):
```
git clone https://github.com/BSN-Spartan/NC-PolygonEdge.git
cd NC-PolygonEdge/
make build
cp polygon-edge /usr/bin
```

Execute the following commands to check the version of Spartan NC PolygonEdge:

```
polygon-edge version
```

![](https://raw.githubusercontent.com/BSN-Spartan/NC-PolygonEdge/main/.github/images/5.polygonversion.jpg)

### 3.2 Using Docker Images


Before installing the full node by Docker images, Docker 18 or later version should be installed in your server.

Run the following command to install the Docker image:

```
wget -qO- https://get.docker.com/ | sh
```

Grant user permission to execute Docker commands:

```
sudo usermod -aG docker your-user
```
Now, check the docker version:

```shell
docker version
```
![](https://raw.githubusercontent.com/BSN-Spartan/NC-PolygonEdge/main/.github/images/4.1dockerversion.jpg)


Official Docker images are hosted under the hub.docker.com registry. Run the following command to pull them to the server:

```
docker pull bsnspartan/nc-polygon-edge:latest
```

## 4. Running the Full Node

### 4.1 Configuration

Create a new directory `node1/`:

```
mkdir node1
```

Copy [genesis.json](https://github.com/BSN-Spartan/NC-PolygonEdge/blob/main/spartan/genesis.json) and [config.json](https://github.com/BSN-Spartan/NC-PolygonEdge/blob/main/spartan/config.json) files from `spartan/` directory to node1/ directory:

```
cp ./spartan/genesis.json ./spartan/config.json node1/
```

`genesis.json` defines the genesis block data and specifies the system parameters.

Refer to below template to configure your `config.json` file:

```
{
  "chain_config": "./genesis.json",
  "secrets_config": "",
  "data_dir": "./data",
  "block_gas_target": "0x0",
  "grpc_addr": "0.0.0.0:9632",
  "jsonrpc_addr": "0.0.0.0:8545",
  "telemetry": {
    "prometheus_addr": ""
  },
  "network": {
    "no_discover": false,
    "libp2p_addr": "0.0.0.0:1478",
    "nat_addr": "",
    "dns_addr": "",
    "max_peers": -1,
    "max_outbound_peers": -1,
    "max_inbound_peers": -1
  },
  "seal": true,
  "tx_pool": {
    "price_limit": 1000000000,
    "max_slots": 2500
  },
  "log_level": "INFO",
  "restore_file": "",
  "block_time_s": 5,
  "headers": {
    "access_control_allow_origins": [
      "*"
    ]
  },
  "log_to": "",
  "dev": true,
  "json_rpc_batch_request_limit":100
}

```

Edit below parameters of `config.json`:

`data_dir`is the directory used to store the ledger, secret key file and other information. To change the data storage configuration, please refer to: https://docs.polygon.technology/docs/edge/validator-hosting

`grpc_addr` is the gRPC port of the node. As the management function, it is not recommended to open this interface to the Internet.

`jsonrpc_addr` is the transaction API. You can set it public to your user or forward it to users.

`no_discover ` Set to false by default to make the node discoverable in the network. For VDCs, it is required to set this parameter to "false", otherwise the health check of the node cannot be performed.

`libp2p_addr `is the listening address of P2P. For VDCs, it is required to open this port, otherwise the health check of the node cannot be performed.

`nat_addr` is the public IP address of the server. For VDCs, it is required to set, otherwise the health check of the node cannot be performed.

`log_to`is the output file of the log. For detailed log configuration, please refer to:https://docs.polygon.technology/docs/edge/validator-hosting#log-files

To learn more about `config.json`, check out the following link:

https://docs.polygon.technology/docs/edge/configuration/sample-config

### 4.2 Starting the Node
#### 4.2.1 Starting by Commands

Start the node in `node1/` directory with the command below:

```
polygon-edge server --config config.json
```
You can see the blocks are synchronized to the node:
![](https://raw.githubusercontent.com/BSN-Spartan/NC-PolygonEdge/main/.github/images/6.startnode.jpg)

Or you can execute in the background via `nohup`:

```
nohup polygon-edge server --config config.json >/dev/null 2>&1 &
```
To stop the node in `nohup` mode, please refer to the below command:
```
pkill -INT polygon-edge
```

Confirm the node status:

```
polygon-edge status --grpc-address 127.0.0.1:9632
```
![](https://raw.githubusercontent.com/BSN-Spartan/NC-PolygonEdge/main/.github/images/7.status.jpg)

#### 4.2.2 Starting by Docker Images

Make sure you have installed the node by Docker images (refer to 3.2.2), and `genesis.json` and `config.json` are copied and configured into `node1/` directory.

Access to node1/ directory and start the node:

```
cd node1/
docker run -d -p 8545:8545 -p 1478:1478 -p 9632:9632 -v $PWD:/opt/ --restart=always --name spartan-nc-polygon-edge bsnspartan/nc-polygon-edge:latest server --config config.json
```

Confirm the node status:

```
docker exec spartan-nc-polygon-edge polygon-edge status --grpc-address 127.0.0.1:9632
```


## 5. Generate the Node Signature

When joining the Spartan Network as a Data Center, the Data Center Operator will be rewarded a certain amount of NTT Incentives based on the quantity of the registered node. To achieve this, the Data Center Operator should first provide the signature of the full node to verify the node's ownership.

####  Node Installed by Commands:

Execute the following command in the node's data directory after the node is started.

Remember to configure NAT for your node and enable node discovery.

```
polygon-edge secrets validate --data-dir data --grpc-address 127.0.0.1:9632 --json
```

* `data-dir` is the data directory of the node. If you use local key management, you should specify this directory to store the data file of the node.

* `config ` is the key configuration file, if you use remote key management, you should specify this parameter.

* `grpc-address` is the node address that generates the signature, which is usually the current node address.



#### Node Installed by Docker:

Execute below command:

```
docker exec spartan-nc-polygon-edge polygon-edge secrets validate --data-dir data --grpc-address 127.0.0.1:9632
```



### Node Signature

After executing the above commands, you will get the following information. Please submit it to the locally installed Data Center System when registering the node.

```
{
    "nodeId":"16Uiu2HAmTCgocz1Y25YDzQdqgHtBzh6UxnycXTwnhmxCFHy6xQPS",
    "address":"/ip4/10.0.51.109/tcp/1478/p2p/16Uiu2HAmTCgocz1Y25YDzQdqgHtBzh6UxnycXTwnhmxCFHy6xQPS",
    "signature":"AN1rKvtQYxuPqNffZ6Y3F5CZYAeRjWNKLZMB4FakqhJs3yp2GU4NcH6fgpRUpkxcDcQFPT8WvNRStCd5HJTbmCFqMqEUeGz2H"
}
```



## 6. Resources

### 6.1 JSON-RPC Commands

NC-PolygonEdge is compatible with ETH JSON RPC interface, please refer to the detailed interface list from below link:

https://docs.polygon.technology/docs/edge/get-started/json-rpc-commands

### 6.2 CLI Commands

NC-PolygonEdge provides a wealth of CLI commands for managing your nodes. For a detailed command list, please refer to the link below :

https://docs.polygon.technology/docs/edge/get-started/cli-commands

### 6.3 Prometheus Metrics

Polygon Edge can report and serve the Prometheus metrics, which in their turn can be consumed using Prometheus collector(s).

The following is a detailed description reference:

https://docs.polygon.technology/docs/edge/configuration/prometheus-metrics



### 6.4 Backup/Restore Node Instance

This guide goes into detail on how to back up and restore a Polygon Edge node instance. It covers the base folders and what they contain, as well as which files are critical for performing a successful backup and restore.

For detailed operation, please refer to the link below:

https://docs.polygon.technology/docs/edge/working-with-node/backup-restore

