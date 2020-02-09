node {
  if (env.GITHUB_REF != 'refs/heads/master') {
    checkoutSource()
    build()
    Test()
  }
}


def checkoutSource() {
  stage ('checkout') {
    sh('#!/bin/sh -e\n cp -r /github/workspace/* $WORKSPACE")
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
