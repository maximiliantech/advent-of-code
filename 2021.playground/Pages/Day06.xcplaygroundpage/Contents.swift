//: [Previous](@previous)

import Foundation

func partOne() -> String {
    // load input and remove \n
    let input = load(file: "input06", ofType: .txt)?.replacingOccurrences(of: "\n", with: "")
    // turn input to Int Array
    var numbers = input?.split(separator: ",").map({ Int($0)! })
    print("Initial state: ", numbers!)
    
    for day in 1...80 {
        for (index, lanternfish) in numbers!.enumerated() {
            numbers![index] -= 1
            if lanternfish == 0 {
                numbers![index] = 6
                numbers?.append(8)
            }
        }
    }
        
    return numbers?.count.description ?? "0"
}

func partTwo() -> String {
    // load input and remove \n
    let input = load(file: "input06", ofType: .txt)?.replacingOccurrences(of: "\n", with: "")
    // turn input to Int Array
    var numbers = input?.split(separator: ",").map({ Int($0)! })
    print("Initial state: ", numbers!)
    
    for day in 1...257 {
        for (index, lanternfish) in numbers!.enumerated() {
            numbers![index] -= 1
            if lanternfish == 0 {
                numbers![index] = 6
                numbers?.append(8)
            }
        }
    }
        
    return numbers?.count.description ?? "0"
}

print(partOne())
print(partTwo())
//: [Next](@next)
