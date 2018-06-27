#!/bin/bash

function cleanup {
    kill -9 $ganache_pid ; true
    kill -9 $loom_pid ; true
}

REPO_ROOT=`pwd`
LOOM_DIR=`pwd`/tmp/loom-plasma-$BUILD_TAG
BUILD_NUMBER=201

rm -rf  $LOOM_DIR; true
mkdir -p $LOOM_DIR
cd $LOOM_DIR
if [[ "`uname`" == 'Darwin' ]]; then
wget https://private.delegatecall.com/loom/osx/build-$BUILD_NUMBER/loom
else 
wget https://private.delegatecall.com/loom/linux/build-$BUILD_NUMBER/loom
fi
chmod +x loom
export LOOM_BIN=`pwd`/loom
echo $REPO_ROOT
cp $REPO_ROOT/loom_test/loom-test.yml $LOOM_DIR/loom.yml
$LOOM_BIN init
echo 'Loom DAppChain initialized in ' $LOOM_DIR

trap cleanup EXIT

cd $REPO_ROOT/server
npm run --silent migrate:dev
sleep 1
echo 'Launched ganache' $ganache_pid

cd $LOOM_DIR
$LOOM_BIN run &
loom_pid=$!
echo 'Launched loom' $loom_pid

# Wait for Ganache & Loom to spin up
sleep 10

cd $REPO_ROOT/loom_test
./plasmascash_tester

kill -9 $loom_pid
rm -rf $LOOM_DIR
