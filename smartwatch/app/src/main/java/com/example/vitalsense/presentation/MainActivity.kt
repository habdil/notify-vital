package com.example.vitalsense.presentation

import android.os.Bundle
import androidx.activity.ComponentActivity
import androidx.activity.compose.setContent
import androidx.compose.foundation.background
import androidx.compose.foundation.layout.*
import androidx.compose.runtime.*
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.platform.LocalConfiguration
import androidx.compose.ui.tooling.preview.Devices
import androidx.compose.ui.tooling.preview.Preview
import androidx.compose.ui.unit.dp
import androidx.core.splashscreen.SplashScreen.Companion.installSplashScreen
import androidx.wear.compose.material.MaterialTheme
import androidx.wear.compose.material.Scaffold
import androidx.wear.compose.material.TimeText
import androidx.wear.tooling.preview.devices.WearDevices
import com.example.vitalsense.presentation.components.SquareCompatibleProgressComponent
import com.example.vitalsense.presentation.theme.DarkBlue
import com.example.vitalsense.presentation.theme.VitalsenseTheme
import kotlinx.coroutines.delay
import kotlin.math.min

class MainActivity : ComponentActivity() {
    override fun onCreate(savedInstanceState: Bundle?) {
        // Apply splash screen
        installSplashScreen()

        super.onCreate(savedInstanceState)
        setTheme(android.R.style.Theme_DeviceDefault)

        setContent {
            VitalsenseTheme {
                VitalSenseApp()
            }
        }
    }
}

@Composable
fun VitalSenseApp() {
    // Dummy data that would come from sensors in a real app
    val stepsGoal = 10000
    val caloriesGoal = 500
    val maxHeartRate = 200

    // Using remember and mutableStateOf to hold our dummy values with initial values matching the reference image
    var steps by remember { mutableStateOf(149) }
    var heartRate by remember { mutableStateOf(65) }
    var calories by remember { mutableStateOf(79) }

    // Simulating changing values with LaunchedEffect
    LaunchedEffect(key1 = true) {
        while(true) {
            // Small random variations to simulate real-time updates
            steps = (steps + (-5..5).random()).coerceIn(140, 170)
            heartRate = (heartRate + (-2..2).random()).coerceIn(60, 80)
            calories = (calories + (-3..3).random()).coerceIn(70, 90)
            delay(2000) // Update every 2 seconds
        }
    }

    // Get screen dimensions for adaptive layout
    val configuration = LocalConfiguration.current
    val screenWidth = configuration.screenWidthDp
    val screenHeight = configuration.screenHeightDp
    val isSquare = screenWidth == screenHeight

    Scaffold(
        timeText = {
            // Only show TimeText if there's enough room (on round devices)
            if (!isSquare) {
                TimeText(
                    modifier = Modifier.padding(top = 9.dp)
                )
            }
        }
    ) {
        Box(
            modifier = Modifier
                .fillMaxSize()
                .background(DarkBlue),
            contentAlignment = Alignment.Center
        ) {
            // If it's a square device, add the time text at the top with more padding
            if (isSquare) {
                TimeText(
                    modifier = Modifier
                        .align(Alignment.TopCenter)
                        .padding(top = 10.dp)
                )
            }

            // Main circular progress display optimized for both square and round shapes
            SquareCompatibleProgressComponent(
                steps = steps,
                stepsGoal = stepsGoal,
                heartRate = heartRate,
                maxHeartRate = maxHeartRate,
                calories = calories,
                caloriesGoal = caloriesGoal,
                modifier = Modifier.align(Alignment.Center)
            )
        }
    }
}

// Preview for square devices (360x360)
@Preview(
    device = Devices.WEAR_OS_SMALL_ROUND,
    showSystemUi = true,
    widthDp = 360,
    heightDp = 360
)
@Composable
fun SquarePreview() {
    VitalsenseTheme {
        VitalSenseApp()
    }
}

// Preview for round devices
@Preview(device = WearDevices.SMALL_ROUND, showSystemUi = true)
@Composable
fun RoundPreview() {
    VitalsenseTheme {
        VitalSenseApp()
    }
}