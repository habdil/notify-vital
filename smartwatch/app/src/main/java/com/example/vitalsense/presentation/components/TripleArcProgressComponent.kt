package com.example.vitalsense.presentation.components

import androidx.compose.foundation.Canvas
import androidx.compose.foundation.layout.*
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.geometry.Offset
import androidx.compose.ui.geometry.Size
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.graphics.StrokeCap
import androidx.compose.ui.graphics.drawscope.Stroke
import androidx.compose.ui.platform.LocalConfiguration
import androidx.compose.ui.text.TextStyle
import androidx.compose.ui.text.font.FontWeight
import androidx.compose.ui.text.style.TextAlign
import androidx.compose.ui.unit.Dp
import androidx.compose.ui.unit.dp
import androidx.compose.ui.unit.sp
import androidx.wear.compose.material.Text
import com.example.vitalsense.presentation.theme.*
import kotlin.math.min

/**
 * A circular progress component optimized to work on both square and round Wear OS devices.
 * Displays metrics stacked vertically in the center with arc indicators on the sides.
 */
@Composable
fun SquareCompatibleProgressComponent(
    steps: Int,
    stepsGoal: Int,
    heartRate: Int,
    maxHeartRate: Int,
    calories: Int,
    caloriesGoal: Int,
    modifier: Modifier = Modifier
) {
    // Calculate progress percentages
    val stepsProgress = (steps.toFloat() / stepsGoal).coerceIn(0f, 1f)
    val heartRateProgress = (heartRate.toFloat() / maxHeartRate).coerceIn(0f, 1f)
    val caloriesProgress = (calories.toFloat() / caloriesGoal).coerceIn(0f, 1f)

    // Get screen dimensions to ensure proper sizing on different devices
    val configuration = LocalConfiguration.current
    val screenWidth = configuration.screenWidthDp.dp
    val screenHeight = configuration.screenHeightDp.dp

    // Use the smaller dimension to ensure the UI fits on both square and round displays
    val minDimension = min(screenWidth.value, screenHeight.value).dp

    // Calculate component size based on the screen dimensions
    val componentSize = minDimension * 0.85f

    Box(
        modifier = modifier
            .size(componentSize)
            .padding(8.dp),
        contentAlignment = Alignment.Center
    ) {
        // Background circle and arcs
        Canvas(modifier = Modifier.fillMaxSize()) {
            val center = Offset(size.width / 2, size.height / 2)
            val radius = min(size.width, size.height) / 2 - 16.dp.toPx()
            val arcWidth = 12.dp.toPx()

            // Draw outer background circle (light gray)
            drawCircle(
                color = Color.Gray.copy(alpha = 0.3f),
                radius = radius,
                center = center,
                style = Stroke(width = arcWidth)
            )

            // Steps arc (Teal) - TOP portion
            drawArc(
                color = Teal,
                startAngle = -60f,
                sweepAngle = 120f * stepsProgress,
                useCenter = false,
                topLeft = Offset(center.x - radius, center.y - radius),
                size = Size(radius * 2, radius * 2),
                style = Stroke(width = arcWidth, cap = StrokeCap.Round)
            )

            // Calories arc (Orange) - BOTTOM-LEFT portion
            drawArc(
                color = Orange,
                startAngle = 200f,
                sweepAngle = 60f * caloriesProgress,
                useCenter = false,
                topLeft = Offset(center.x - radius, center.y - radius),
                size = Size(radius * 2, radius * 2),
                style = Stroke(width = arcWidth, cap = StrokeCap.Round)
            )
        }

        // Stacked metrics in the center
        Column(
            modifier = Modifier.align(Alignment.Center),
            horizontalAlignment = Alignment.CenterHorizontally,
            verticalArrangement = Arrangement.Center
        ) {
            // First metric - Steps
            Text(
                text = "$steps",
                color = White,
                fontSize = 11.sp,
                fontWeight = FontWeight.Bold,
                textAlign = TextAlign.Center,
                style = TextStyle(fontWeight = FontWeight.Bold)
            )

            Text(
                text = "Steps",
                color = Teal,
                fontSize = 6.sp,
                textAlign = TextAlign.Center
            )

            Spacer(modifier = Modifier.height(4.dp))

            // Second metric - Heart Rate
            Text(
                text = "$heartRate",
                color = White,
                fontSize = 11.sp,
                fontWeight = FontWeight.Bold,
                textAlign = TextAlign.Center,
                style = TextStyle(fontWeight = FontWeight.Bold)
            )

            Text(
                text = "BPM",
                color = LightBlue,
                fontSize = 6.sp,
                textAlign = TextAlign.Center
            )

            Spacer(modifier = Modifier.height(4.dp))

            // Third metric - Calories
            Text(
                text = "$calories",
                color = White,
                fontSize = 11.sp,
                fontWeight = FontWeight.Bold,
                textAlign = TextAlign.Center,
                style = TextStyle(fontWeight = FontWeight.Bold)
            )

            Text(
                text = "Calories",
                color = Orange,
                fontSize = 6.sp,
                textAlign = TextAlign.Center
            )
        }
    }
}