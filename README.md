# qcore-mobile

Preview Android and iOS bindings for the [qore QR decoder](https://github.com/firesh/qcore), pinned to core v1.0.0.

English · [简体中文](README.zh-CN.md) · [日本語](README.ja.md)

## Android · Gradle

Add the public Maven repository to `dependencyResolutionManagement.repositories` in `settings.gradle.kts`:

```kotlin
maven {
    url = uri("https://raw.githubusercontent.com/firesh/qcore-mobile/maven/")
    content { includeGroup("io.github.firesh") }
}
```

Add this to the app module's dependencies:

```kotlin
implementation("io.github.firesh:qcore-mobile:0.1.0-preview.2")
```

Android API 23+. The AAR includes `armeabi-v7a`, `arm64-v8a`, and `x86_64`. See the [Kotlin example](example/DecodeGray.kt).

## iOS · Swift Package Manager

In Xcode, choose **File → Add Package Dependencies**, enter `https://github.com/firesh/qcore-mobile`, and select **Exact Version** `0.1.0-preview.2`. Add the **QCore** product to your app target. Import `QCore` as in the [Swift example](example/DecodeGray.swift).

Requires iOS 15+. The package includes arm64 for devices and arm64/x86_64 for simulators. Direct downloads are available from [GitHub Releases](https://github.com/firesh/qcore-mobile/releases).

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
gomobile bind -target=android/arm,android/arm64,android/amd64 -androidapi=23 -javapkg io.github.firesh.qcore -o dist/qcore.aar ./mobile
gomobile bind -target=ios,iossimulator -iosversion=15.0 -prefix QCore -o dist/QCore.xcframework ./mobile
```

This preview ships binary artifacts and frame-level examples. Full camera sample apps and device latency measurements are pending.

[MIT License](LICENSE)
