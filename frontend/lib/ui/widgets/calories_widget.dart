import 'package:flutter/material.dart';
import '../../shared/theme.dart';
import 'dart:math' as math;

class CaloriesWidget extends StatelessWidget {
  final int calories;
  final int dailyGoal;

  const CaloriesWidget({
    Key? key,
    required this.calories,
    this.dailyGoal = 2000,
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    final double percentage = calories / dailyGoal;
    
    return Container(
      width: double.infinity,
      padding: const EdgeInsets.all(15),
      decoration: BoxDecoration(
        color: whiteColor,
        borderRadius: BorderRadius.circular(15),
        boxShadow: [
          BoxShadow(
            color: Colors.black.withOpacity(0.1),
            blurRadius: 10,
            offset: const Offset(0, 4),
          ),
        ],
      ),
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          Row(
            children: [
              Icon(
                Icons.local_fire_department,
                color: orangeColor,
                size: 20,
              ),
              const SizedBox(width: 8),
              Text(
                'Calories',
                style: darkTextStyle.copyWith(
                  fontSize: 16,
                  fontWeight: semiBold,
                ),
              ),
            ],
          ),
          Expanded(
            child: Center(
              child: SizedBox(
                width: 120,
                height: 120,
                child: Stack(
                  alignment: Alignment.center,
                  children: [
                    CustomPaint(
                      size: const Size(120, 120),
                      painter: CaloriesProgressPainter(
                        percentage: percentage > 1.0 ? 1.0 : percentage,
                        progressColor: orangeColor,
                      ),
                    ),
                    Column(
                      mainAxisAlignment: MainAxisAlignment.center,
                      children: [
                        Text(
                          calories.toString(),
                          style: darkTextStyle.copyWith(
                            fontSize: 24,
                            fontWeight: bold,
                          ),
                        ),
                      ],
                    ),
                  ],
                ),
              ),
            ),
          ),
        ],
      ),
    );
  }
}

class CaloriesProgressPainter extends CustomPainter {
  final double percentage;
  final Color progressColor;

  CaloriesProgressPainter({
    required this.percentage,
    required this.progressColor,
  });

  @override
  void paint(Canvas canvas, Size size) {
    final center = Offset(size.width / 2, size.height / 2);
    final radius = math.min(size.width, size.height) / 2;
    final strokeWidth = 12.0;

    // Draw background circle
    final backgroundPaint = Paint()
      ..color = greyColor.withOpacity(0.2)
      ..style = PaintingStyle.stroke
      ..strokeWidth = strokeWidth
      ..strokeCap = StrokeCap.round;

    canvas.drawCircle(center, radius - strokeWidth / 2, backgroundPaint);

    // Draw progress arc with gradient
    final rect = Rect.fromCircle(center: center, radius: radius - strokeWidth / 2);
    final gradient = SweepGradient(
      center: Alignment.center,
      startAngle: -math.pi / 2,
      endAngle: 3 * math.pi / 2,
      colors: [
        progressColor.withOpacity(0.7),
        progressColor,
      ],
      stops: const [0.0, 1.0],
      transform: GradientRotation(-math.pi / 2),
    );

    final progressPaint = Paint()
      ..shader = gradient.createShader(rect)
      ..style = PaintingStyle.stroke
      ..strokeWidth = strokeWidth
      ..strokeCap = StrokeCap.round;

    final startAngle = -math.pi / 2; // Start from the top
    final sweepAngle = 2 * math.pi * percentage;

    canvas.drawArc(
      rect,
      startAngle,
      sweepAngle,
      false,
      progressPaint,
    );
  }

  @override
  bool shouldRepaint(CustomPainter oldDelegate) => true;
}