#coding:utf8

import math

def calcsize(count, contentType, imageSize):
    commitTxSize = 10.5 +  57.5 * 1 + 43 * (count + 1)
    revealTxSize = 10.5 +  57.5 * 1 + 43 * (count)
    ordinalsSize = 1 + 1 + (1+3) + 2 + 1 + len(contentType) + 2 + (3 * (imageSize / 520) + imageSize) + 1
    return  math.ceil( commitTxSize + revealTxSize + ordinalsSize)


size = calcsize(1, "image/png", 516)
print(size)


fee = size * 20

print(fee)
print(fee / 10**8)



# 参考unisat.io的计算方法
#
#    function Z(e) {
#         let {fileSize: s, mintRate: t, fileCount: i, inscriptionBalance: l, setInscriptionBalance: x, discountRate: u} = e
#           , [j,f] = (0,
#         a.useState)(!1);
#         u || (u = {
#             unisatCount: 0,
#             ogPassCount: 0,
#             unisatFeeCutPercent: 0,
#             ogPassFeeCutPercent: 0
#         });
#         let p = u.unisatFeeCutPercent + u.ogPassFeeCutPercent
#           , y = l * i
#           , N = Math.ceil((s / 4 + 175 * i) * t)
#           , F = 1999 * i
#           , C = Math.ceil(.0499 * N)
#           , S = Math.floor((F + C) * (1 - p / 100))
#           , w = y + N + S
#           , Z = 1e3 * Math.floor((y + N + S) / 1e3);
#         return (0,


def cccc( inscriptionBalance, count, fileSize, feeRate):
    p = 0
    y = inscriptionBalance * count    # 铭文utxo所占的费用
    N = math.ceil((fileSize / 4 + 175 * count) * feeRate)  # 矿工费,    fileSize 除以4, 4是witness的权重转为交易费用的权重,  175 是 10.5 +  57.5 * 1 + 43 + 64
    # F = 1999 * count    #  unisat的基础服务费 Service Base Fee
    # C = math.ceil(.0499 * N)   #  4.99% 的大小服务费 Fee by Size
    # S = math.floor((F + C) * (1 - p / 100))   # 总服务费
    w = y + N  # + S    # 用户需要支付的费用
    print(w)
    # Z = 1e3 * math.floor((y + N + S) / 1e3);
    Z = 1e3 * math.floor((y + N  ) / 1e3);    # 搞一下，去掉小数点
    print(Z)
    return Z


print(cccc(546, 1, 28828, 21))

