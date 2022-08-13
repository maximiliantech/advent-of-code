import Foundation

let directoryURL = FileManager.default.urls(for: .documentDirectory, in: .userDomainMask).first
let fileURL = URL(fileURLWithPath: "input", relativeTo: directoryURL).appendingPathExtension("txt")

func partOne() -> String {
    let inputString = load(file: "input01")
    if let numbers = inputString?.split(separator: "\n").map({ Int($0)! }) {
        let result = zip(numbers, numbers.dropFirst())
            .lazy
            .filter{ $0 < $1 }
            .count
            .description
        return result
    }
    return "No result!"
}

print(partOne())
