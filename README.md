<pre>
           _       _       _ 
  ___ ___ (_)_ __ | |_ ___| |
 / __/ _ \| | '_ \| __/ _ \ |
| (_| (_) | | | | | ||  __/ |
 \___\___/|_|_| |_|\__\___|_|
                             
</pre>

A very thin, straight-forward client for Coinbase that displays cryptocurrency balances in their native currency. Written in Go.

Displays the data as either a table or as JSON.

## Requirements

## Installation

## Usage

```
cointel --format=[json|table] --persist=[true|false]
```

### Options

#### format

* json - display the output as a JSON-encoded string
* table - display the output as an ASCII-art table

#### persist

* true - keep the program running, refreshing and rewriting the output every 15 minutes
* false - run the program once, write the output, and terminate
