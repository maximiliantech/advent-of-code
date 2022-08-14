//: [Previous](@previous)

import Foundation

struct CrabAlignment {
    var totalFuel: Int
    var position: Int
    var horizontalPositions: [Int]
    
    init(for position: Int, horizontalPositions: [Int]) {
        self.position = position
        self.horizontalPositions = horizontalPositions
        var total = 0
        for curr_pos in horizontalPositions {
            total += abs(curr_pos - position)
        }
        self.totalFuel = total
    }
}

func partOne() -> String {
    // load input and remove \n
    let input = load(file: "input07_test", ofType: .txt)?.replacingOccurrences(of: "\n", with: "")
    // turn numbers in [Int]
    let horizontalPositions = input?.split(separator: ",").map { Int($0)! }
    
    var crabAlignmentArray: [CrabAlignment] = []
    for align_pos in horizontalPositions! {
        crabAlignmentArray.append(CrabAlignment(for: align_pos, horizontalPositions: horizontalPositions!))
    }
    // sort 
    crabAlignmentArray.sort { $0.totalFuel < $1.totalFuel }
    return crabAlignmentArray[0].totalFuel.description
}

print(partOne())

//: [Next](@next)
