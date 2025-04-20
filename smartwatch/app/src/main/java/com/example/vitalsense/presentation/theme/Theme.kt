package com.example.vitalsense.presentation.theme

import androidx.compose.runtime.Composable
import androidx.compose.ui.graphics.Color
import androidx.wear.compose.material.Colors
import androidx.wear.compose.material.MaterialTheme

// Primary dark blue color from your specification
val DarkBlue = Color(0xFF2C3E50)

// Complementary colors for a complete theme
val LightBlue = Color(0xFF3498DB)
val Teal = Color(0xFF1ABC9C)
val Red = Color(0xFFE74C3C)
val Orange = Color(0xFFE67E22)
val Yellow = Color(0xFFF1C40F)
val White = Color(0xFFFFFFFF)
val LightGray = Color(0xFFECF0F1)
val Gray = Color(0xFF95A5A6)
val DarkGray = Color(0xFF7F8C8D)

// Create a custom color palette for the Wear theme
private val VitalSenseColorPalette = Colors(
    primary = DarkBlue,
    primaryVariant = LightBlue,
    secondary = Teal,
    secondaryVariant = Teal,
    error = Red,
    onPrimary = White,
    onSecondary = White,
    onError = White,
    background = DarkBlue,
    onBackground = White,
    surface = DarkGray,
    onSurface = White
)

@Composable
fun VitalsenseTheme(
    content: @Composable () -> Unit
) {
    /**
     * Custom theme for the Notify Vital VitalSense smartwatch application.
     * Uses a dark blue primary color (2C3E50) with complementary colors
     * optimized for Wear OS display and readability.
     */
    MaterialTheme(
        colors = VitalSenseColorPalette,
        content = content
    )
}