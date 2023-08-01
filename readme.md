# Uni - Text to Unix Command Converter

Uni is a command-line application that helps you convert natural language text into Unix commands effortlessly. It utilizes the power of OpenAI's GPT-3.5 language model to interpret the input text and generate the corresponding Unix commands.

## Features

- Convert plain English text into valid Unix commands.
- Interactive and user-friendly command-line interface.
- Optional `--exec` flag to execute the generated Unix commands.
- A safe and efficient way to interact with the GPT-3.5 API to handle language interpretation.

## Installation

1. Create a .env file
in the root directory of your project and add these variables:
```
API_KEY=your-api-key # replace with actual API key from OpenAI website
```
2. Next, install the uni binary to a directory in your system's PATH. For example, if you have /usr/local/bin in your PATH, you can install uni there:
```
sudo mv uni /usr/local/bin/
```
3. Hola! You are set to use the plugin. You can use the plugin in the following way. The -c stands for the command tag which specifies if you wish to execute the command. 

```
uni -c "create a new folder named new" --exec
```# text2unix
