import io.github.firesh.qcore.mobile.Mobile

fun decodeGray(pixels: ByteArray, width: Int, height: Int, stride: Int): String =
    Mobile.decodeGray(pixels, width.toLong(), height.toLong(), stride.toLong())
