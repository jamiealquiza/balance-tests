# balance-tests

```
 % go test

[consistent-hashing]
         Empty nodes - 0
  Greatest imbalance - portion of keys: 6.25% / ratio: 1.28x
               Range - 22189 / highest value: 102246 / lowest value: 80057

node-0: 87221
node-1: 85459
node-2: 102246
node-3: 80057

[FNV1-mod]
         Empty nodes - 0
  Greatest imbalance - portion of keys: 0.11% / ratio: 1.00x
               Range - 407 / highest value: 88934 / lowest value: 88527

node-0: 88527
node-1: 88823
node-2: 88699
node-3: 88934

[FNV1a-mod]
         Empty nodes - 0
  Greatest imbalance - portion of keys: 0.11% / ratio: 1.00x
               Range - 407 / highest value: 88934 / lowest value: 88527

node-0: 88699
node-1: 88823
node-2: 88527
node-3: 88934

[vaporCH]
         Empty nodes - 0
  Greatest imbalance - portion of keys: 0.38% / ratio: 1.02x
               Range - 1350 / highest value: 89231 / lowest value: 87881

node-0: 87881
node-1: 88711
node-2: 89231
node-3: 89160
PASS
ok      github.com/jamiealquiza/balance-tests   0.665s
```