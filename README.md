# Fiatconv
**Fiatconv** is a simple CLI app that let you convert one fiat currency to another.
Fiat money like USD, EUR, RUB etc.
**Fiatconv** uses exchangeratesapi.io as a backend for getting exchange rates. 
You can write your own exchange rate adapter. See ./pkg/exchangerate/exchangeratesapiio as an example.

# Usage
```console
fiatconv <amount> <src_symbol> <dst_symbol>
```
# Get started
```bash
git clone https://github.com/Maxfer4Maxfer/fiatconv.git
cd ./fiatconv
go install ./cmd/fiatconv
```

## Donations
 If you want to support this project, please consider donating:
 * PayPal: https://paypal.me/MaxFe
