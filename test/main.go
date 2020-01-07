/*
  Copyright 2019 Whiteblock Inc.
  This file is a part of the definition.

  definition is free software: you can redistribute it and/or modify
  it under the terms of the GNU General Public License as published by
  the Free Software Foundation, either version 3 of the License, or
  (at your option) any later version.

  definition is distributed in the hope that it will be useful,
  but dock ANY WARRANTY; without even the implied warranty of
  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
  GNU General Public License for more details.

  You should have received a copy of the GNU General Public License
  along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

package main

import (
	"fmt"
	"github.com/whiteblock/definition"
	"github.com/whiteblock/go-prettyjson"
	"os"
	"strconv"
)

func main() {

	startingPoint := []byte(`services:
  - name: Quorum1
    image: quorumengineering/quorum:2.2.5
    shared-volumes:
      - source-path: /output.log
        name: eea-logs
    resources:
      cpus: 4
      memory: 4 GB
      storage: 5 GiB
    input-files:
      - source-path: genesis.json
        destination-path: /data/genesis.json
      - source-path: permissioned-nodes.json
        destination-path: /data/permissioned-nodes.json
      - source-path: permissioned-nodes.json
        destination-path: /data/static-nodes.json
      - source-path: key1
        destination-path: /data/keystore/key1
      - source-path: nodekey1
        destination-path: /data/nodekey
      - source-path: passwords.txt
        destination-path: /data/passwords.txt
    script:
      inline: geth --datadir /data init data/genesis.json && geth --datadir /data --unlock 0 --password /data/passwords.txt --ethstats Node1:eea_testnet_secret@eea:80 --syncmode full --mine --minerthreads 1 --rpc --rpcaddr 0.0.0.0 --rpcapi admin,db,eth,debug,miner,net,shh,txpool,personal,web3,quorum > /output.log 2>&1                      
  - name: Quorum2
    image: quorumengineering/quorum:2.2.5
    shared-volumes:
      - source-path: /output.log
        name: eea-logs
    resources:
      cpus: 4
      memory: 4 GB
      storage: 5 GiB
    input-files:
      - source-path: genesis.json
        destination-path: /data/genesis.json
      - source-path: permissioned-nodes.json
        destination-path: /data/permissioned-nodes.json
      - source-path: permissioned-nodes.json
        destination-path: /data/static-nodes.json
      - source-path: key2
        destination-path: /data/keystore/key2
      - source-path: nodekey2
        destination-path: /data/nodekey
      - source-path: passwords.txt
        destination-path: /data/passwords.txt
    script:
      inline: geth --datadir /data init data/genesis.json && geth --datadir /data --unlock 0 --password /data/passwords.txt --ethstats Node1:eea_testnet_secret@eea:80 --syncmode full --mine --minerthreads 1 --rpc --rpcaddr 0.0.0.0 --rpcapi admin,db,eth,debug,miner,net,shh,txpool,personal,web3,quorum > /output.log 2>&1                        
  - name: Quorum3
    image: quorumengineering/quorum:2.2.5
    shared-volumes:
      - source-path: /output.log
        name: eea-logs
    resources:
      cpus: 4
      memory: 4 GB
      storage: 5 GiB
    input-files:
      - source-path: genesis.json
        destination-path: /data/genesis.json
      - source-path: permissioned-nodes.json
        destination-path: /data/permissioned-nodes.json
      - source-path: permissioned-nodes.json
        destination-path: /data/static-nodes.json
      - source-path: key3
        destination-path: /data/keystore/key3
      - source-path: nodekey3
        destination-path: /data/nodekey
      - source-path: passwords.txt
        destination-path: /data/passwords.txt
    script:
      inline: geth --datadir /data init data/genesis.json && geth --datadir /data --unlock 0 --password /data/passwords.txt --ethstats Node1:eea_testnet_secret@eea:80 --syncmode full --mine --minerthreads 1 --rpc --rpcaddr 0.0.0.0 --rpcapi admin,db,eth,debug,miner,net,shh,txpool,personal,web3,quorum > /output.log 2>&1   
  - name: Quorum4
    image: quorumengineering/quorum:2.2.5
    shared-volumes:
      - source-path: /output.log
        name: eea-logs
    resources:
      cpus: 4
      memory: 4 GB
      storage: 5 GiB
    input-files:
      - source-path: genesis.json
        destination-path: /data/genesis.json
      - source-path: permissioned-nodes.json
        destination-path: /data/permissioned-nodes.json
      - source-path: permissioned-nodes.json
        destination-path: /data/static-nodes.json
      - source-path: key4
        destination-path: /data/keystore/key4
      - source-path: nodekey4
        destination-path: /data/nodekey
      - source-path: passwords.txt
        destination-path: /data/passwords.txt
    script:
      inline: geth --datadir /data init data/genesis.json && geth --datadir /data --unlock 0 --password /data/passwords.txt --ethstats Node1:eea_testnet_secret@eea:80 --syncmode full --mine --minerthreads 1 --rpc --rpcaddr 0.0.0.0 --rpcapi admin,db,eth,debug,miner,net,shh,txpool,personal,web3,quorum > /output.log 2>&1   
  - name: Quorum5
    image: quorumengineering/quorum:2.2.5
    shared-volumes:
      - source-path: /output.log
        name: eea-logs
    resources:
      cpus: 4
      memory: 4 GB
      storage: 5 GiB
    input-files:
      - source-path: genesis.json
        destination-path: /data/genesis.json
      - source-path: permissioned-nodes.json
        destination-path: /data/permissioned-nodes.json
      - source-path: permissioned-nodes.json
        destination-path: /data/static-nodes.json
      - source-path: key5
        destination-path: /data/keystore/key5
      - source-path: nodekey5
        destination-path: /data/nodekey
      - source-path: passwords.txt
        destination-path: /data/passwords.txt
    script:
      inline: geth --datadir /data init data/genesis.json && geth --datadir /data --unlock 0 --password /data/passwords.txt --ethstats Node1:eea_testnet_secret@eea:80 --syncmode full --mine --minerthreads 1 --rpc --rpcaddr 0.0.0.0 --rpcapi admin,db,eth,debug,miner,net,shh,txpool,personal,web3,quorum > /output.log 2>&1
  - name: Quorum6
    image: quorumengineering/quorum:2.2.5
    shared-volumes:
      - source-path: /output.log
        name: eea-logs
    resources:
      cpus: 4
      memory: 4 GB
      storage: 5 GiB
    input-files:
      - source-path: genesis.json
        destination-path: /data/genesis.json
      - source-path: permissioned-nodes.json
        destination-path: /data/permissioned-nodes.json
      - source-path: permissioned-nodes.json
        destination-path: /data/static-nodes.json
      - source-path: key6
        destination-path: /data/keystore/key6
      - source-path: nodekey6
        destination-path: /data/nodekey
      - source-path: passwords.txt
        destination-path: /data/passwords.txt
    script:
      inline: geth --datadir /data init data/genesis.json && geth --datadir /data --unlock 0 --password /data/passwords.txt --ethstats Node1:eea_testnet_secret@eea:80 --syncmode full --mine --minerthreads 1 --rpc --rpcaddr 0.0.0.0 --rpcapi admin,db,eth,debug,miner,net,shh,txpool,personal,web3,quorum > /output.log 2>&1
  - name: Quorum7
    image: quorumengineering/quorum:2.2.5
    shared-volumes:
      - source-path: /output.log
        name: eea-logs
    resources:
      cpus: 4
      memory: 4 GB
      storage: 5 GiB
    input-files:
      - source-path: genesis.json
        destination-path: /data/genesis.json
      - source-path: permissioned-nodes.json
        destination-path: /data/permissioned-nodes.json
      - source-path: permissioned-nodes.json
        destination-path: /data/static-nodes.json
      - source-path: key7
        destination-path: /data/keystore/key7
      - source-path: nodekey7
        destination-path: /data/nodekey
      - source-path: passwords.txt
        destination-path: /data/passwords.txt
    script:
      inline: geth --datadir /data init data/genesis.json && geth --datadir /data --unlock 0 --password /data/passwords.txt --ethstats Node1:eea_testnet_secret@eea:80 --syncmode full --mine --minerthreads 1 --rpc --rpcaddr 0.0.0.0 --rpcapi admin,db,eth,debug,miner,net,shh,txpool,personal,web3,quorum > /output.log 2>&1
  - name: ethstats
    image: gcr.io/whiteblock/ethstats:master
    env:
      HOST: "0.0.0.0"
    input-files:
      - source-path: ws_secret.json
        destination-path: /eth-netstats/ws_secret.json     
task-runners:
  - name: testnet-expiration
    script:
      inline: sleep 7200    
tests:
  - name: testnet
    timeout: infinite
    description: run an EEA testnet and execute some simple transactions
    system:
      - type: Quorum1
        count: 1
        port-mappings:
          - "30303:30303"
          - "8545:8545"
        resources: 
            networks:
              - name: quorum_network
      - type: Quorum2
        count: 1
        port-mappings:
          - "30304:30303"
          - "8546:8545"
        resources: 
            networks:
              - name: quorum_network
      - type: Quorum3
        count: 1
        port-mappings:
          - "30305:30303"
          - "8547:8545"
        resources: 
            networks:
              - name: quorum_network
      - type: Quorum4
        count: 1
        port-mappings:
          - "30306:30303"
          - "8548:8545"
        resources: 
            networks:
              - name: quorum_network
      - type: Quorum5
        count: 1
        port-mappings:
          - "30307:30303"
          - "8549:8545"
        resources: 
            networks:
              - name: quorum_network
      - type: Quorum6
        count: 1
        port-mappings:
          - "30308:30303"
          - "8550:8545"
        resources: 
            networks:
              - name: quorum_network
      - type: Quorum7
        count: 1
        port-mappings:
          - "30308:30303"
          - "8551:8545"
        resources: 
            networks:
              - name: quorum_network
      - type: ethstats
        count: 1
        port-mappings:
          - "80:3000"
        resources: 
            networks:
              - name: quorum_network`)

	def, err := definition.SchemaYAML(startingPoint)
	if err != nil {
		panic(err)
	}

	tests, err := definition.GetTests(def, definition.Meta{})
	if err != nil {
		panic(err)
	}
	cmds := tests[0].Commands
	out, _ := prettyjson.Marshal(cmds)
	if len(os.Args) > 1 {
		index, err := strconv.Atoi(os.Args[1])
		if err != nil {
			panic(err)
		}
		out, _ = prettyjson.Marshal(cmds[index])
	}
	fmt.Println(string(out))

}
