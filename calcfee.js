


function calcfee( count, averageFileSize, feeRate ) {

    utxoOutputValue = 10000
    commitTxSize = 68 + (43 + 1)
    commitTxSize += 64
    revealTxSize = 10.5 + (57.5 + 43)
    revealTxSize += 64
    feeSats = Math.ceil( ( averageFileSize / 4 + commitTxSize + revealTxSize ) * feeRate  )
    feeSats = 1e3 * Math.ceil(feeSats / 1e3)

    feeSats = feeSats * count

    baseService = 1000 * Math.ceil(feeSats * 0.1 / 1000)
    feeSats += baseService

    total = feeSats + utxoOutputValue * count
    return total
}


function test() {
    f = calcfee(1, 3)
    console.log( f )
    console.log( f / 1e8 )
}

test()