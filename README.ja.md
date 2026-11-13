# qcore-mobile

[English](README.md) · [简体中文](README.zh-CN.md) · [日本語](README.ja.md)

qore の Android / iOS SDK プレビューです。コア v1.0.0 を使用します。

## Android · Gradle

settings.gradle.kts の dependencyResolutionManagement.repositories に Maven リポジトリを追加し、アプリモジュールに依存関係を追加します。

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

Android API 23+ · `armeabi-v7a` · `arm64-v8a` · `x86_64` · [Kotlin](example/DecodeGray.kt)

## iOS · Swift Package Manager

Xcode の **File → Add Package Dependencies** でリポジトリ URL を入力し、**Exact Version** `0.1.0-preview.2` を指定します。**QCore** をアプリターゲットに追加します。

`https://github.com/firesh/qcore-mobile`

iOS 15+ · arm64 (device) · arm64/x86_64 (simulator) · [Swift](example/DecodeGray.swift)

入力は 1 ピクセル 1 バイトのグレースケール画像です。幅、高さ、バイト単位の行ストライドを渡します。未検出時は `[]` を返します。座標は入力画像を基準とします。

[GitHub Releases](https://github.com/firesh/qcore-mobile/releases) · [Build and migration](README.md#frame-and-result-contract) · [MIT](LICENSE)
