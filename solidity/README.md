## Solidity

### 1 build the lexer files

```bash
java -jar 'antlr-4.9.3-complete.jar' -Dlanguage=Go -listener -visitor -o parser ./SolidityLexer.g4
```

### 2 build the grammar
```bash
antlr -Dlanguage=Go -listener -visitor -o parser -long-messages -Xforce-atn ./Solidity.g4
java -jar 'antlr-4.9.3-complete.jar' -Dlanguage=Go -listener -visitor -o parser ./Solidity.g4
```

## Download grammar

Grammar source: https://github.com/ethereum/solidity/tree/develop/docs/grammar

replace: `type=typeName` with `varType=typeName`
replace: `String` with `StringLiteral`