node {
  if (env.GITHUB_REF != 'refs/heads/master') {
    checkoutSource()
    TestSh()
  }
}


def checkoutSource() {
  stage ('checkout') {
    sh('#!/bin/sh -e\n cp -r /github/workspace/* $WORKSPACE')
  }
}


def setenvGo(){
  stage('prepare'){
    sh("sudo add-apt-repository ppa:longsleep/golang-backports; sudo apt update; sudo apt install golang-go")
  }
}

def buildGo () {
  stage ('Build') {
    sh('#!/bin/sh -e\n go fmt; go vet; go build')
  }
}


def TestGo() {
  stage ('Unit tests') {
    sh('#!/bin/sh -e\n go test')
  }
}

def TestSh(){
  stage ('Unit tests') {
    sh("./test.sh")
  }
}
