import XCTest
import SwiftTreeSitter
import TreeSitterCoconut

final class TreeSitterCoconutTests: XCTestCase {
    func testCanLoadGrammar() throws {
        let parser = Parser()
        let language = Language(language: tree_sitter_coconut())
        XCTAssertNoThrow(try parser.setLanguage(language),
                         "Error loading Coconut grammar")
    }
}
