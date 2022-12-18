import Foundation

public enum FileExtension: String {
    case txt = "txt"
}

public func readInput(from fileURL: URL) -> String {
    do {
     // Get the saved data
     let savedData = try Data(contentsOf: fileURL)
     // Convert the data back into a string
     if let savedString = String(data: savedData, encoding: .utf8) {
        return savedString
     }
    } catch {
     // Catch any errors
     print("Unable to read the file")
    }
    return ""
}

public func load(file named: String, ofType fileType: FileExtension) -> String? {
    guard let fileUrl = Bundle.main.url(forResource: named, withExtension: fileType.rawValue) else {
        return nil
    }

    guard let content = try? String(contentsOf: fileUrl, encoding: .utf8) else {
        return nil
    }
    
    return content
}

