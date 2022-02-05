# Wonderland (Olympus Pro fork) - Atomic Bonding
A tool for purchasing and redeeming Wonderland bonds atomically.\
\
Notes: Wonderland bonding is deprecated. For educational purposes only.

## Background
### Purchasing bond
1. Swap $TIME to $MIM
2. Bond $TIME using $MIM

### Redeeming bond
1. Redeem $TIME

## Issue (purchasing bond)
1. Rapid pair & bond price actions in two transactions.
2. Suboptimal ROI without pair & bond prices monitoring.

## Solution
1. No pair & bond price actions in one transaction.
2. Live pair & bond prices monitoring using on-chain data, slippage & fees inclusive.
3. redeem() theoretically safe to invoke with exposed accounts.

## Example (purchasing bond)
![AtomicBond](docs/AtomicBond.png)
