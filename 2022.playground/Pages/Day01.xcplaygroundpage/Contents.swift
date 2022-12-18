//: [Previous](@previous)

import Foundation

struct Elf {
    var name: String
    var calories: [Int]
    var caloriesSum: Int {
        calories.reduce(0, +)
    }
    
    init(name: String, calories: [Int]) {
        self.name = name
        self.calories = calories
    }
}

func partOne() -> Int {
    let input = load(file: "input01", ofType: .txt)!
    var elves = getElves(from: input)
    var highestCalorie = elves.sorted { $0.caloriesSum > $1.caloriesSum }
    return highestCalorie[0].caloriesSum
}

func partTwo() -> Int {
    let input = load(file: "input01", ofType: .txt)!
    var elves = getElves(from: input)
    var highestCalorie = elves.sorted { $0.caloriesSum > $1.caloriesSum }
    return highestCalorie[0].caloriesSum + highestCalorie[1].caloriesSum + highestCalorie[2].caloriesSum
}

func getElves(from input: String) -> [Elf] {
    let elvesCalories = input.split(separator: "\n\n")
    var elves: [Elf] = []
    for (index, cal) in elvesCalories.enumerated() {
        let calories = cal.split(separator: "\n").map { Int($0)! }
        elves.append(Elf(name: "\(index+1)", calories: calories))
    }
    return elves
}

print(partOne())
print(partTwo())

//: [Next](@next)
