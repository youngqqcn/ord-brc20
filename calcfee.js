


function calcfee( count, averageFileSize, feeRate ) {

    // utxoOutputValue = 10000
    utxoOutputValue = 1000
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
    f = calcfee(1, 550, 10)
    console.log('每个: ', f, 'sats')
    console.log('每个: ', f * 30185/1e8 , 'USD')
    console.log('900个: ', f * 30185/1e8 *900, 'USD')
    console.log('900个: ', f * 900 , 'sats')
    console.log('900个: ', f * 900 /1e8, 'BTC')
}

test()