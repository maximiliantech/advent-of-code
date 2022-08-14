//: [Previous](@previous)

import Foundation

struct CrabAlignment {
    var totalFuel: Int
    var position: Int
    var horizontalPositions: [Int]
    
    init(for position: Int, horizontalPositions: [Int], constantRate: Bool) {
        self.position = position
        self.horizontalPositions = horizontalPositions
        var total: Int = 0
        if constantRate {
            for curr_pos in horizontalPositions {
                total += abs(curr_pos - position)
            }
        } else {
            for curr_pos in horizontalPositions {
                total += ( ( abs(curr_pos - position) * (abs(curr_pos - position) + 1) ) / 2 )
            }
        }
        self.totalFuel = total
    }
}

func partOne() -> String {
    // load input and remove \n
    let input = load(file: "input07", ofType: .txt)?.replacingOccurrences(of: "\n", with: "")
    // turn numbers in [Int]
    let horizontalPositions = input?.split(separator: ",").map { Int($0)! }
    
    var crabAlignmentArray: [CrabAlignment] = []
    // Crabs can be aligned at any horizontal position
    for align_pos in horizontalPositions!.min()! ... horizontalPositions!.max()! {
        crabAlignmentArray.append(CrabAlignment(for: align_pos, horizontalPositions: horizontalPositions!, constantRate: true))
    }
    // sort array for the least total fuel
    crabAlignmentArray.sort { $0.totalFuel < $1.totalFuel }
    return crabAlignmentArray[0].totalFuel.description
}

func partTwo() -> String {
    // load input and remove \n
    let input = load(file: "input07", ofType: .txt)?.replacingOccurrences(of: "\n", with: "")
    // turn numbers in [Int]
    let horizontalPositions = input?.split(separator: ",").map { Int($0)! }
    
    var crabAlignmentArray: [CrabAlignment] = []
    // Crabs can be aligned at any horizontal position
    for align_pos in horizontalPositions!.min()! ... horizontalPositions!.max()! {
        crabAlignmentArray.append(CrabAlignment(for: align_pos, horizontalPositions: horizontalPositions!, constantRate: false))
    }
    // sort array for the least total fuel
    crabAlignmentArray.sort { $0.totalFuel < $1.totalFuel }
    return crabAlignmentArray[0].totalFuel.description
}

print(partOne())
print(partTwo())
//: [Next](@next)
