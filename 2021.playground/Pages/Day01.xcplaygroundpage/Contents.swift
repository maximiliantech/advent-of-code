import Foundation

func partOne() -> String {
    let inputString = load(file: "input01")
    if let numbers = inputString?.split(separator: "\n").map({ Int($0)! }) {
        let result = zip(numbers, numbers.dropFirst())
            .lazy
            .filter { $0 < $1 }
            .count
            .description
        return result
    }
    return "No result!"
}

func partTwo() -> String {
    let inputString = load(file: "input01")
    if let numbers = inputString?.split(separator: "\n").map({ Int($0)! }) {
        let result = zip(numbers, numbers.dropFirst(3))
            .lazy
            .filter{ $0 < $1 }
            .count
            .description
        return result
    }
    return "No result!"
}

print(partOne())
print(partTwo())
