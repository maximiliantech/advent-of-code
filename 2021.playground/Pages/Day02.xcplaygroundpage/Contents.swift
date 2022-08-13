//: [Previous](@previous)

import Foundation

func partOne() -> String {
    var x: Int = 0
    var y: Int = 0
    
    let input = load(file: "input02", ofType: .txt)
    
    input?.split(separator: "\n").forEach { line in
        let operation = line.split(separator: " ").first
        let number = Int(line.split(separator: " ").last ?? "0")!
        
        switch operation {
        case "forward": x += number
        case "down": y += number
        case "up": y -= number
        default:
            print("UNEXPECTED LINE \(line)")
            fatalError()
        }
    }
    
    return (x * y).description
}

print(partOne())

//: [Next](@next)
