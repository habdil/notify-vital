import 'package:flutter/material.dart';
import '../../shared/theme.dart';
import 'dart:math' as math;

class StatusWidget extends StatelessWidget {
  final String status;
  final double value; // value between 0.0 - 1.0

  const StatusWidget({
    super.key,
    required this.status,
    required this.value,
  });

  @override
  Widget build(BuildContext context) {
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
                Icons.access_time,
                color: greenColor,
                size: 20,
              ),
              const SizedBox(width: 8),
              Text(
                'Status',
                style: darkTextStyle.copyWith(
                  fontSize: 16,
                  fontWeight: semiBold,
                ),
              ),
            ],
          ),
          Expanded(
            child: Column(
              mainAxisAlignment: MainAxisAlignment.center,
              children: [
                SizedBox(
                  height: 120,
                  child: CustomPaint(
                    size: Size(double.infinity, 120),
                    painter: GaugePainter(
                      value: value,
                      gaugeColor: getStatusColor(status),
                    ),
                  ),
                ),
                const SizedBox(height: 5),
                Text(
                  status,
                  style: darkTextStyle.copyWith(
                    fontSize: 14,
                    fontWeight: medium,
                  ),
                ),
              ],
            ),
          ),
        ],
      ),
    );
  }

  Color getStatusColor(String status) {
    switch (status) {
      case 'kurang gerak':
        return Colors.orange;
      case 'cukup gerak':
        return greenColor;
      case 'terlalu aktif':
        return Colors.red;
      default:
        return greenColor;
    }
  }
}

class GaugePainter extends CustomPainter {
  final double value; // 0.0 to 1.0
  final Color gaugeColor;

  GaugePainter({
    required this.value,
    required this.gaugeColor,
  });

  @override
  void paint(Canvas canvas, Size size) {
    final center = Offset(size.width / 2, size.height - 20);
    final radius = math.min(size.width, size.height * 2) * 0.7;
    final rect = Rect.fromCircle(center: center, radius: radius / 2);
    
    // Draw the colored gauge background with gradient
    final gradient = SweepGradient(
      center: Alignment.bottomCenter,
      startAngle: math.pi,
      endAngle: 2 * math.pi,
      colors: const [
        Colors.red,
        Colors.orange,
        Colors.green,
        Colors.orange,
        Colors.red,
      ],
      stops: const [0.0, 0.25, 0.5, 0.75, 1.0],
      transform: const GradientRotation(math.pi),
    );
    
    final gradientPaint = Paint()
      ..shader = gradient.createShader(rect)
      ..style = PaintingStyle.stroke
      ..strokeWidth = 12
      ..strokeCap = StrokeCap.round;

    // Draw the gauge arc (semi-circle)
    canvas.drawArc(
      rect,
      math.pi, // Start from the left side
      math.pi, // Draw a semi-circle
      false,
      gradientPaint,
    );
    
    // Draw the needle
    final needlePaint = Paint()
      ..color = Colors.black
      ..style = PaintingStyle.stroke
      ..strokeWidth = 2
      ..strokeCap = StrokeCap.round;
    
    // Calculate the needle position based on value
    final needleAngle = math.pi + value * math.pi;
    final needleLength = radius / 2 - 10;
    final needleEndPoint = Offset(
      center.dx + needleLength * math.cos(needleAngle),
      center.dy + needleLength * math.sin(needleAngle),
    );
    
    // Draw the needle
    canvas.drawLine(center, needleEndPoint, needlePaint);
    
    // Draw the needle base circle
    final circlePaint = Paint()
      ..color = gaugeColor
      ..style = PaintingStyle.fill;
    
    canvas.drawCircle(center, 6, circlePaint);
    
    // Draw the circle border
    final circleBorderPaint = Paint()
      ..color = Colors.black
      ..style = PaintingStyle.stroke
      ..strokeWidth = 2;
    
    canvas.drawCircle(center, 6, circleBorderPaint);
  }

  @override
  bool shouldRepaint(covariant CustomPainter oldDelegate) => true;
}