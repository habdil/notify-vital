import 'package:flutter/material.dart';
import '../../shared/theme.dart';
import 'dart:math' as math;

class HeartWidget extends StatelessWidget {
  final int heartRate;

  const HeartWidget({
    Key? key,
    required this.heartRate,
  }) : super(key: key);

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
                Icons.favorite,
                color: greenColor,
                size: 20,
              ),
              const SizedBox(width: 8),
              Text(
                'Heart',
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
                  height: 100,
                  child: CustomPaint(
                    size: Size(double.infinity, 100),
                    painter: HeartRateGraphPainter(),
                  ),
                ),
                const SizedBox(height: 5),
                Row(
                  mainAxisAlignment: MainAxisAlignment.center,
                  children: [
                    Text(
                      '$heartRate',
                      style: darkTextStyle.copyWith(
                        fontSize: 24,
                        fontWeight: bold,
                      ),
                    ),
                    Text(
                      ' BPM',
                      style: greyTextStyle.copyWith(
                        fontSize: 14,
                        fontWeight: medium,
                      ),
                    ),
                  ],
                ),
              ],
            ),
          ),
        ],
      ),
    );
  }
}

class HeartRateGraphPainter extends CustomPainter {
  @override
  void paint(Canvas canvas, Size size) {
    final paint = Paint()
      ..color = greenColor
      ..style = PaintingStyle.stroke
      ..strokeWidth = 2.0;
    
    final path = Path();
    
    // Starting point at left center
    final double startY = size.height / 2;
    path.moveTo(0, startY);
    
    // Make simple ECG pattern repeating across width
    final double segmentWidth = size.width / 5; // 5 complete ECG patterns
    
    for (int i = 0; i < 5; i++) {
      final double start = i * segmentWidth;
      
      // First small bump
      path.lineTo(start + segmentWidth * 0.1, startY - 5);
      
      // Back to baseline
      path.lineTo(start + segmentWidth * 0.2, startY);
      
      // Sharp spike up
      path.lineTo(start + segmentWidth * 0.25, startY - 40);
      
      // Sharp drop below baseline
      path.lineTo(start + segmentWidth * 0.3, startY + 20);
      
      // Recovery spike up
      path.lineTo(start + segmentWidth * 0.35, startY - 15);
      
      // Back to baseline to complete the pattern
      path.lineTo(start + segmentWidth * 0.4, startY);
      path.lineTo(start + segmentWidth, startY);
    }
    
    canvas.drawPath(path, paint);
  }

  @override
  bool shouldRepaint(covariant CustomPainter oldDelegate) => false;
}