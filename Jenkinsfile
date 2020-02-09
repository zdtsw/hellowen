node {
  if (env.GITHUB_REF != 'refs/heads/master') {
    checkoutSource()
    setenv()
    build()
    Test()
  }
}


def checkoutSource() {
  stage ('checkout') {
    sh('#!/bin/sh -e\n cp -r /github/workspace/* $WORKSPACE')
  }
}

def setenv(){
  stage('prepare'){
    sh("sudo add-apt-repository ppa:longsleep/golang-backports; sudo apt update; sudo apt install golang-go")
  }
}

def build () {
  stage ('Build') {
    sh('#!/bin/sh -e\n go fmt; go vet; go build')
  }
}


def Test() {
  stage ('Unit tests') {
    sh('#!/bin/sh -e\n go test')
  }
}
