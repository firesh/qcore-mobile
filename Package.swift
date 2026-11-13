// swift-tools-version:5.9
import PackageDescription

let package = Package(
    name: "QCore",
    platforms: [.iOS(.v15)],
    products: [.library(name: "QCore", targets: ["QCore"])],
    targets: [
        .binaryTarget(
            name: "QCore",
            url: "https://github.com/firesh/qcore-mobile/releases/download/v1.0.0/QCore.xcframework.zip",
            checksum: "b710b71d8223ec6fca1454f7524b1635e5ef09faa2db6c5c40bbd1628ab7173d"
        )
    ]
)
