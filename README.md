<pre>
           _       _       _ 
  ___ ___ (_)_ __ | |_ ___| |
 / __/ _ \| | '_ \| __/ _ \ |
| (_| (_) | | | | | ||  __/ |
 \___\___/|_|_| |_|\__\___|_|
                             
</pre>

A very thin, straight-forward client for Coinbase that displays cryptocurrency balances in their native currency. Written in Go.

Displays the data as either a table or as JSON.

### Examples

#### JSON

<pre>
{"Currencies":{"BTC":{"Quantity":0.00661205,"CashValue":36.971,"Symbol":"BTC","Updated":"2017-09-05 21:25:27"},"ETH":{"Quantity":0.04513654,"CashValue":18.52,"Symbol":"ETH","Updated":"2017-09-05 21:25:27"},"LTC":{"Quantity":0.34187566,"CashValue":32.57,"Symbol":"LTC","Updated":"2017-09-05 21:25:27"}},"Updated":"2017-09-05 21:25:27"}
</pre>

#### Table

<pre>
+----------|----------|---------+
| CURRENCY | QUANTITY |  VALUE  |
+----------|----------|---------+
| BTC      |   0.0061 |  $36.97 |
| ETH      |   0.0451 |  $18.52 |
| LTC      |   0.3418 |  $32.57 |
+----------|----------|---------+
|                        $88 06 |
+----------|----------|---------+
</pre>

## Requirements

## Installation

### Coinbase API

First, generate a Coinbase API key (`Settings -> API Access`). I recommend you creat a new key for `cointel` to use, so you can restrict the permissions as tightly as possible. This API key will need at least the following permissions:

```
* wallet:accounts:read 
* wallet:addresses:read
```

You *do not* need to give `cointel` any *write* permissions.

Next, you need to set two environment variables before running `cointel`:

#### COINBASE_KEY

Your Coinbase API key. Add this to your bash|zsh|fish profile:

```
export COINBASE_KEY='[your key here]'
```

#### COINBASE_SECRET

Your Coinbase API secret. Add this to your bash|zsh|fish profile:

```
export COINBASE_SECRET='[your secret here]'
```

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
