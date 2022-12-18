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
    let input = load(file: "input01_test", ofType: .txt)!
    var elves = getElves(from: input)
    return 0
}

func getElves(from input: String) -> [Elf] {
    var elvesCalories = input.split(separator: "\n\n")
    var elves: [Elf] = []
    for (index, cal) in elvesCalories.enumerated() {
        print(cal)
    }
    return elves
}

print(partOne())

//: [Next](@next)
