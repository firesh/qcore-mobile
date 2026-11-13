# qcore-mobile

[English](README.md) · [简体中文](README.zh-CN.md) · [日本語](README.ja.md)

qore 的 Android 和 iOS SDK，使用核心 v1.0.0。

## Android · Gradle

在 settings.gradle.kts 的 dependencyResolutionManagement.repositories 中加入 Maven 仓库，然后在 App 模块中添加依赖。

```kotlin
maven {
    url = uri("https://raw.githubusercontent.com/firesh/qcore-mobile/maven/")
    content { includeGroup("io.github.firesh") }
}
```

Add this to the app module's dependencies:

```kotlin
implementation("io.github.firesh:qcore-mobile:1.0.0")
```

Android API 23+ · `armeabi-v7a` · `arm64-v8a` · `x86_64` · [Kotlin](example/DecodeGray.kt)

## iOS · Swift Package Manager

在 Xcode 选择 **File → Add Package Dependencies**，输入仓库地址，选择 **Exact Version** `1.0.0`，将 **QCore** 产品添加到 App target。

`https://github.com/firesh/qcore-mobile`

iOS 15+ · arm64 (device) · arm64/x86_64 (simulator) · [Swift](example/DecodeGray.swift)

输入为每像素一字节的灰度数据，传入宽、高和以字节计的行步长。无识别结果时返回 `[]`。坐标对应输入图像，预览层需按方向和裁剪映射。

[GitHub Releases](https://github.com/firesh/qcore-mobile/releases) · [Build and migration](README.md#frame-and-result-contract) · [MIT](LICENSE)
