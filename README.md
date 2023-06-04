# kubectl-openai
Solve Kubernetes problems using OpenAI

![demo](https://github.com/CirillaQL/kubectl-openai/blob/master/images/example.gif)

## Quick Start
1. Download the binary from [Github Releases](https://github.com/CirillaQL/kubectl-openai/releases)
2. If you want to use this as a kubectl plugin, then copy kubectl-openai binary to your $PATH. If not, you can also use the binary standalone.

## Usage
Set your ChatGPT API with ``token`` command

```kubectl openai token $(Your Key)```

Try to analyze pod or pods

```kubectl openai pod <pod-name> -n <pod-namespace>```

```kubectl openai pods <pods-namespace>```