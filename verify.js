// 安装依赖 npm install bech32

let { bech32, bech32m } = require("bech32");

function test(address, isMainnet) {
  if (address.length == 62) {
    let chkPrefix = isMainnet ? "bc" : "tb";
    try {
      let ret = bech32m.decode(address);
      if (ret.prefix === chkPrefix && ret.words.length === 53) {
        return true;
      }
    } catch (error) {
        return false
    }
  }

  return false;
}

console.log(test("bc1p2yzcv24v9tpw6ffhkqcq994y8p4ps2xfv65wx7nsmg4meuvzd0fqyesxg7", true));
console.log(test("1p2yzcv24v9tpw6ffhkqcq994y8p4ps2xfv65wx7nsmg4meuvzd0fqyesxg7", true))
console.log(test("tb1p2yzcv24v9tpw6ffhkqcq994y8p4ps2xfv65wx7nsmg4meuvzd0fqn3xfj3", false))
console.log(test("tbyzcv24v9tpw6ffhkqcq994y8p4ps2xfv65wx7nsmg4meuvzd0fqn3xfj3", false))
console.log(test("bc1q3fgz3ez545723aknffhkhr63y38fd5rxmk7mfud28c53", false));
