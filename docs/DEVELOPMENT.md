# Development Guide

This document provides guidelines for setting up the development environment, running tasks, and contributing to the project.

## Prerequisites

* [VS Code](https://code.visualstudio.com/)
* [Task](https://taskfile.dev/)

## Setting up Development Environment

1. Clone and open the repository

    ```bash
    git clone git@github.com:Stellatarum/docker-machine-driver-pve.git
    ```

    ```bash
    code docker-machine-driver-pve
    ```

1. Install tools:

    ```bash
    task tools
    ```

## Running tasks

We use [Task](https://taskfile.dev/) for common development tasks. Run following command to list available tasks:

```bash
task --list
```
