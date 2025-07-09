
# Raydium Launchpad Go

The task can be split into two parts:

1. Transaction parsing (Geyser)
2. Transaction sending (buy/sell before and after migration)

## 1. Parsing
You need to take the raw transaction data (its level-1 and level-2 instructions in the Geyser format)
and convert it into a structure we understand. This might be:

* A third-party token buy/sell/swap (the pre-migration and post-migration variants use different instructions)
* A token-creation transaction
* A migration transaction

For every instruction, you must parse both the instruction parameters and any data coming from
internal (inner) instructions. With Geyser you’ll have both on hand, so you need to collect all
used accounts (token, mint, any bonding-curve or pool PDAs), the token’s symbol, pool/curve size,
liquidity, etc. In a token-creation transaction you’ll often see a buy instruction as well, and
you must extract that, too. If one transaction includes creation, purchase, and sale, our code
should output those three distinct objects. If it includes three buys and one sell, we should
capture them all so we can calculate totals.

The parsed data should be represented by a single structure roughly like this:
```go
type Transaction struct {
  Signature solana.Signature
  Slot      uint64

  Create      []CreateInfo
  Trade       []TradeInfo
  TradeBuys   []int
  TradeSells  []int

  Migrate     []Migration
  SwapBuys    []SwapBuy
  SwapSells   []SwapSell
}
```

Also, tokens can be migrated, which changes the buy/sell instruction formats. You must parse
migrations and all trades that follow them.

## 2. Sending
You need to build and serialize buy/sell instructions. Ideally, also implement token-creation
instructions, since they’ll be useful too. You can:

* Implement instruction format (structure and functions) from scratch
* Or find the JSON schema and generate it via anchor-go

In either case, the output should be a Go struct with setters and a constructor function,
so it’s easy to assemble externally. And of course the instruction must compile to a valid
Solana instruction.

Testing
Your tests should include:
* Several submit-transaction cases (buy and sell)—use a wallet and token from environment variables
and actually run the test.
* Several parsing cases—either connect to a live Geyser stream or use provided sample
transactions—to verify that you map all relevant instructions (create, buy, sell, migrate, swap)
into our structures.

Requirements
* Must use `github.com/gagliardetto/solana-go v1.12.x`
* For binary-data helpers you may also use `github.com/gagliardetto/binary v0.8.0`
* `Go 1.24` (toolchain go1.24.4)
* Handle errors in true Go idiomatic style

No conversion between parsing structs and submission structs is required. They should remain separate.
