# qcore-mobile

Preview Android and iOS bindings for the [qore QR decoder](https://github.com/firesh/qcore), pinned to core v1.0.0.

## Install the preview

Download the archives and checksums from [Releases](https://github.com/firesh/qcore-mobile/releases).

- **Android:** add `qcore.aar` to `app/libs/` and add `implementation(files("libs/qcore.aar"))` to your app's Gradle dependencies. The preview targets API 23+ and includes arm64-v8a and x86_64. See the [Kotlin call site](example/DecodeGray.kt).
- **iOS:** unzip `QCore.xcframework.zip` and add `QCore.xcframework` to your app target in Xcode. It includes iOS arm64 and arm64/x86_64 simulator slices and targets iOS 15+. See the [Swift call site](example/DecodeGray.swift).

## Frame and result contract

Pass one 8-bit grayscale byte per pixel, with width, height, and row stride in bytes. The buffer must contain all pixels up to the end of the last visible row. The decoder returns a JSON array of QR results, or `[]` for no read. Invalid frames report an error. Coordinates are in the input buffer's orientation and pixel space. The JSON includes text, base64-encoded raw bytes, and four corner points.

### Migrating a camera pipeline

- **Android (ZXing / JourneyApps):** keep your CameraX preview and image analyzer, and call the binding with the copied Y plane of a YUV frame. Honor the row stride and confirm the Y plane has pixel stride 1, or repack it. Close the image after copying. If the existing library launches its own scan UI, wire up camera permission and preview before switching.
- **iOS (Vision / AVFoundation):** keep the capture session and preview. In the video data output delegate, use a bi-planar YUV pixel buffer and copy its luminance plane using its width, height, and bytes per row. Unlock the buffer after copying. Map returned corners through the existing orientation, crop, and preview transforms.

Decode away from the UI thread and bound the number of in-flight frames. Compare both engines on representative frames before switching traffic.

## Build

Install Go 1.23.6, the Android SDK and NDK, and Xcode for iOS. Install both `gomobile` and `gobind` at the revision in `go.mod`, and put them on `PATH`.

```sh
go test ./...
mkdir -p dist
gomobile bind -target=android/arm64,android/amd64 -androidapi=23 -javapkg io.github.firesh.qcore -o dist/qcore.aar ./mobile
gomobile bind -target=ios,iossimulator -iosversion=15.0 -prefix QCore -o dist/QCore.xcframework ./mobile
```

This preview ships binary artifacts and frame-level examples. Full camera sample apps, device latency measurements, and Maven/SwiftPM distribution are pending.

[MIT License](LICENSE)
