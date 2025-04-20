import 'package:flutter/material.dart';
import '../../shared/theme.dart';

class OverviewWidget extends StatelessWidget {
  final int heartRate;
  final double caloriesPercentage;
  final double stepsPercentage;
  final String activityStatus;

  const OverviewWidget({
    Key? key,
    required this.heartRate,
    required this.caloriesPercentage,
    required this.stepsPercentage,
    required this.activityStatus,
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
      child: Row(
        children: [
          Expanded(
            flex: 3,
            child: Column(
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                Center(
                  child: Text(
                    'status',
                    style: darkTextStyle.copyWith(
                      fontSize: 12,
                      fontWeight: medium,
                    ),
                  ),
                ),
                const SizedBox(height: 5),
                _buildStatusBar(),
                const SizedBox(height: 15),
                Text(
                  'kalori',
                  style: darkTextStyle.copyWith(
                    fontSize: 12,
                    fontWeight: medium,
                  ),
                ),
                const SizedBox(height: 5),
                _buildLinearProgress(caloriesPercentage, orangeColor),
                const SizedBox(height: 15),
                Text(
                  'langkah',
                  style: darkTextStyle.copyWith(
                    fontSize: 12,
                    fontWeight: medium,
                  ),
                ),
                const SizedBox(height: 5),
                _buildLinearProgress(stepsPercentage, Colors.blue),
              ],
            ),
          ),
          const SizedBox(width: 10),
          Container(
            padding: const EdgeInsets.symmetric(vertical: 10, horizontal: 10),
            decoration: BoxDecoration(
              border: Border.all(color: greenColor, width: 2),
              borderRadius: BorderRadius.circular(10),
            ),
            child: Column(
              mainAxisSize: MainAxisSize.min,
              children: [
                Icon(
                  Icons.favorite,
                  color: greenColor,
                  size: 20,
                ),
                const SizedBox(height: 5),
                Text(
                  '$heartRate',
                  style: darkTextStyle.copyWith(
                    fontSize: 20,
                    fontWeight: bold,
                  ),
                ),
                Text(
                  'BPM',
                  style: darkTextStyle.copyWith(
                    fontSize: 10,
                    fontWeight: regular,
                  ),
                ),
              ],
            ),
          ),
        ],
      ),
    );
  }

  Widget _buildStatusBar() {
    return Column(
      children: [
        Container(
          height: 15,
          decoration: BoxDecoration(
            borderRadius: BorderRadius.circular(10),
            gradient: const LinearGradient(
              colors: [
                Colors.orange,
                Colors.green,
                Colors.red,
              ],
              begin: Alignment.centerLeft,
              end: Alignment.centerRight,
            ),
          ),
          child: Row(
            children: [
              Expanded(
                flex: 1,
                child: Container(
                  decoration: const BoxDecoration(
                    borderRadius: BorderRadius.only(
                      topLeft: Radius.circular(10),
                      bottomLeft: Radius.circular(10),
                    ),
                  ),
                  child: Center(
                    child: activityStatus == 'kurang gerak' 
                        ? const CircleAvatar(
                            radius: 5,
                            backgroundColor: Colors.white,
                          )
                        : Container(),
                  ),
                ),
              ),
              Expanded(
                flex: 1,
                child: Container(
                  color: Colors.transparent,
                  child: Center(
                    child: activityStatus == 'cukup' 
                        ? const CircleAvatar(
                            radius: 5,
                            backgroundColor: Colors.white,
                          )
                        : Container(),
                  ),
                ),
              ),
              Expanded(
                flex: 1,
                child: Container(
                  decoration: const BoxDecoration(
                    borderRadius: BorderRadius.only(
                      topRight: Radius.circular(10),
                      bottomRight: Radius.circular(10),
                    ),
                  ),
                  child: Center(
                    child: activityStatus == 'terlalu aktif' 
                        ? const CircleAvatar(
                            radius: 5,
                            backgroundColor: Colors.white,
                          )
                        : Container(),
                  ),
                ),
              ),
            ],
          ),
        ),
        const SizedBox(height: 5),
        Row(
          mainAxisAlignment: MainAxisAlignment.spaceBetween,
          children: [
            Text(
              'kurang gerak',
              style: darkTextStyle.copyWith(
                fontSize: 10,
                fontWeight: medium,
              ),
            ),
            Text(
              'cukup',
              style: darkTextStyle.copyWith(
                fontSize: 10,
                fontWeight: medium,
              ),
            ),
            Text(
              'terlalu aktif',
              style: darkTextStyle.copyWith(
                fontSize: 10,
                fontWeight: medium,
              ),
            ),
          ],
        ),
      ],
    );
  }

  Widget _buildLinearProgress(double percentage, Color color) {
    return Stack(
      children: [
        Container(
          height: 12,
          width: double.infinity,
          decoration: BoxDecoration(
            color: greyColor,
            borderRadius: BorderRadius.circular(10),
          ),
        ),
        Container(
          height: 12,
          width: percentage * 100,
          decoration: BoxDecoration(
            color: color,
            borderRadius: BorderRadius.circular(10),
          ),
        ),
      ],
    );
  }
}