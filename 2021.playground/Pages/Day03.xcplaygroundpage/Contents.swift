//: [Previous](@previous)

import Foundation

func partOne() -> String {
    let input = load(file: "input03", ofType: .txt)
    let bitList: [[Int]] = input?.split(separator: "\n").map { line in
        line.map { $0.wholeNumberValue! }
    } ?? [[0]]
    
    print(transform2DArray(array: bitList))
    return ""
}

func transform2DArray(array: [[Int]]) -> [[Int]] {
    var resultArray: [[Int]] = []
    for (j, line) in array.enumerated() {
        for (i, bit) in line.enumerated() {
            resultArray[i][j] = bit
        }
    }
    return resultArray
}

print(partOne())

//: [Next](@next)
