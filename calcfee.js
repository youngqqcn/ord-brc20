


function calcfee( count, feeRate ) {

    // 每个铭文固定金额
    utxoSat = 10000
    averageFileSize = 2600  // 是 2600byte 不是 2600Byte

    utxoOutputValue = utxoSat * count
    commitTxSize = 68 + (43 + 1)*count
    commitTxSize += 64
    revealTxSize = 10.5 + (57.5 + 43) * count
    revealTxSize += 64
    feeSats = Math.ceil( ( averageFileSize / 4 + commitTxSize + revealTxSize ) * feeRate  )
    feeSats = 1e3 * Math.ceil(feeSats / 1e3)

    // base fee
    baseService = 1000 * Math.ceil(feeRate * 0.1 / 1000)
    feeSats += baseService

    total = feeSats + utxoOutputValue
    return total
}


function test() {
    f = calcfee(1, 21)
    console.log( f )
    console.log( f / 1e8 )
}

test()