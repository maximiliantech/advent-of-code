//: [Previous](@previous)

import Foundation

func partOne() -> String {
    // load input and remove \n
    let input = load(file: "input06", ofType: .txt)?.replacingOccurrences(of: "\n", with: "")
    // turn input to Int Array
    let numbers = input?.split(separator: ",").map({ Int($0)! })
    return ""
}

print(partOne())
//: [Next](@next)
