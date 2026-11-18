// swift-tools-version:5.9
import PackageDescription

let package = Package(
    name: "QCore",
    platforms: [.iOS(.v15)],
    products: [.library(name: "QCore", targets: ["QCore"])],
    targets: [
        .binaryTarget(
            name: "QCore",
            url: "https://github.com/firesh/qcore-mobile/releases/download/v1.0.1/QCore.xcframework.zip",
            checksum: "44708629f2ef02a539e82148415d278eec8a5d1abc59b95f4a270ece795d4161"
        )
    ]
)
