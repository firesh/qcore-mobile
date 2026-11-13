import Foundation
import QCore

func decodeGray(_ pixels: Data, width: Int, height: Int, stride: Int) throws -> String {
    var error: NSError?
    let result = QCoreMobileDecodeGray(pixels, width, height, stride, &error)
    if let error { throw error }
    return result
}
