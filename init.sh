#!/bin/bash
rm -r ~/.reputationCLI
rm -r ~/.reputationD
echo "1234567890" | reputationCLI keys add me
echo "1234567890" | reputationCLI keys add you
reputationD init mynode --chain-id reputation
reputationD add-genesis-account $(reputationCLI keys show me -a) 1000foo,100000000stake
reputationD add-genesis-account $(reputationCLI keys show you -a) 1foo
reputationCLI config chain-id reputation
reputationCLI config output json
reputationCLI config indent true
reputationCLI config trust-node true
echo "1234567890" | reputationD gentx --name me
reputationD collect-gentxs
