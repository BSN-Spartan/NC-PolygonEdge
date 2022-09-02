## Spartan-III Chain (Powered by NC PolygonEdge)

# Introduction

A Non-Cryptocurrency Public Chain is a transformed public chain framework based on an existing public chain. Gas Credit transfers are not permitted between standard wallets. There will be no cryptocurrency incentives for mining or participating in consensus.

## 1. About Spartan-III Chain (Powered by NC PolygonEdge)

This document is a guide to install, configure and run a full node in the Non-Crypto Polygon Edge (NC-PolygonEdge) public blockchain.

NC-PolygonEdge networks have two identifiers, a network ID and a chain ID. Although they often have the same value, they have different uses. Peer-to-peer communication between nodes uses the network ID, while the transaction signature process uses the chain ID.

NC-PolygonEdge Network Id = Chain Id = 5566

Below is the instructions for Linux.

## 2. Hardware Requirement

It is recommended to build NC-PolygonEdge nodes on Linux Server with the following requirement.

#### Minimum Requirement

- 2CPU
- Memory: 2GB
- Disk: 100GB SSD
- OS: Ubuntu 16.04 LTS +
- Bandwidth: 20Mbps

#### Recommended Requirement

- 4 CPU
- Memory: 16GB
- Disk: 512GB SSD
- OS: Ubuntu 18.04 LTS +
- Bandwidth: 20Mbps

## 3. How to Install a Full Node

### 3.1 Prerequisites:

**Go 1.17** or above is recommended for building and installing the Spartan-PolygonEdge software. Install `go` by the following steps:

Download and untar the installation file

```
wget https://go.dev/dl/go1.18.5.linux-amd64.tar.gz

tar -C /usr/local -zxvf go1.18.5.linux-amd64.tar.gz
```

Modify environment variables, for example in bash

```
vim /etc/profile

# insert at the end 
export PATH=$PATH:/usr/local/go/bin

source /etc/profile
```

Check the installation result

```
go version
```



### 3.2 Installation

There are 3 methods to install `NC-PolygonEdge`: building from source and docker. Please refer to the installation method more applicable to you.

#### 3.2.1 Building from Source

The stable branch is `main`.

```
git clone https://github.com/BSN-Spartan/NC-PolygonEdge.git
cd NC-PolygonEdge/
go build -o polygon-edge main.go
sudo mv polygon-edge /usr/local/bin
```

#### 3.2.2 Using Docker Images


Before installing the node by Docker images, Docker 18 or later version should be installed in your server.

Run the following command to install the Docker image:

```
wget -qO- https://get.docker.com/ | sh
```

Grant your user permission to execute Docker commands:

```
sudo usermod -aG docker your-user
```


Official Docker images are hosted under the hub.docker.com registry. Run the following command to pull them to the server:

```
docker pull bsnspartan/nc-polygon-edge:latest
```

## 4. Run the Full Node

### 4.1 Configuration

Put `genesis.json` and `config.json` files into the run directory before starting the service.

`genesis.json` defines the genesis block data, which specifies the system parameters. Run the following command to get the latest BootNode:

```
{
    "name": "polygon-edge",
    "genesis": {
        "nonce": "0x0000000000000000",
        "timestamp": "0x0",
        "extraData": "0x0000000000000000000000000000000000000000000000000000000000000000f859f85494c60e7734a514f93ea7f6df5d089e43d83b34e843940ae9b2f34a51bf655e655b497b75de6479127e239477d43d3ca20038d19b1308b2fb6a838fc37512f69461ae43e5f52705343c1b77feba0d59b2dc4e9adbc080c0",
        "gasLimit": "0x3b9aca00",
        "difficulty": "0x1",
        "mixHash": "0x0000000000000000000000000000000000000000000000000000000000000000",
        "coinbase": "0x0000000000000000000000000000000000000000",
        "alloc": {
            "0x4161cAfFeeC110d34868205BB2d9013aFE3B4476": {
                "balance": "0xffffffffffffffffffffffffffffffff"
            },
            "0x6622bE37dAB1973428BcA3eDdabff90c6Ff9D809": {
                "balance": "0xffffffffffffffffffffffffffffffff"
            },
            "0x7AaBFE321f8ce80ca42cFFAf583fb73A5DA09ba5": {
                "balance": "0xffffffffffffffffffffffffffffffff"
            },
            "0x925fa6b05a98134563E08eBc8cE72223a327B9c7": {
                "balance": "0xffffffffffffffffffffffffffffffff"
            },
            "0xf28Ea1AA9e74F862217e7859F5a3CB320Fb2315B": {
                "balance": "0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff"
            }
        },
        "number": "0x0",
        "gasUsed": "0x70000",
        "parentHash": "0x0000000000000000000000000000000000000000000000000000000000000000"
    },
    "params": {
        "forks": {
            "homestead": 0,
            "byzantium": 0,
            "constantinople": 0,
            "petersburg": 0,
            "istanbul": 0,
            "EIP150": 0,
            "EIP158": 0,
            "EIP155": 0
        },
        "chainID": 5566,
        "engine": {
            "ibft": {
                "epochSize": 100000,
                "type": "PoA"
            }
        },
        "blockGasTarget": 0
    },
    "bootnodes": [
        "/ip4/18.143.84.30/tcp/10006/p2p/16Uiu2HAmRjLeEXdoWNyCTJ2QJahH1jUTHUP8xrywFzj8DaYUs539",
        "/ip4/18.143.89.236/tcp/10007/p2p/16Uiu2HAmUQHVk3GHjPLaUomyvkwJyQhEH8mVya2SP7Pq1akEUwur"
    ]
}

```

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

### 4.2 Start the Node
#### 4.2.1 Start by Commands

Make sure you have `NC-PolygonEdge` installed (refer to 3.2.1), and  `genesis.json` and `config.json`  are in the run directory.

Execute the following commands to check the version of Polygon Edge and confirm the installation status of `NC-PolygonEdge`:

```
polygon-edge version
//The following output indicates that the installation is complete
v1.0.0
```
Start your node with the command below:

```
polygon-edge server --config config.json
```
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

#### 4.2.2 Start by Docker Images

Make sure you have installed the node by Docker images (refer to 3.2.2), and `genesis.json` and `config.json` are in the run directory.

Execute the following command in your config file directory to start the node:

```
docker run -d -p 8545:8545 -p 1478:1478 -p 9632:9632 -v $(pwd):/opt/ --restart=always --name spartan-nc-poly-edge bsnspartan/nc-polygon-edge:latest
```

Confirm the node status:

```
docker exec spartan-nc-poly-edge polygon-edge status --grpc-address 127.0.0.1:9632
```


## 5. Generate the Node Signature

When joining the Spartan Network as a VDC, the VDC Owner will be rewarded a certain amount of NTT Incentives based on the quantity of the registered node. To achieve this, the VDC Owner should firstly provide the signature of the VDC node to verify the node's ownership.

####  Node Installed by Commands:

Execute the following command in the node's data directory after the node is started.

Remember to configure NAT for your node and enable node discovery if you haven't.

```
polygon-edge secrets validate --data-dir data --grpc-address 127.0.0.1:9632 --json
```

* `data-dir` is the data directory of the node. If you use local key management, you should specify this directory to store the data file of the node.

* `config ` is the key configuration file, if you use remote key management, you should specify this parameter.

* `grpc-address` is the node address that generates the signature, which is usually the current node address.



#### Node Installed by Docker

Execute below command:

```
docker exec spartan-nc-poly-edge polygon-edge secrets validate --data-dir data --grpc-address 127.0.0.1:9632
```



#### Node Signature

After executing the above commands，you will get the following information. Please submit it to the Spartan Governance System when registering the node .

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



