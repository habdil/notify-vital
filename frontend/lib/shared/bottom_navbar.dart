import 'package:flutter/material.dart';
import '../shared/theme.dart';

class BottomNavBar extends StatelessWidget {
  final int currentIndex;
  final Function(int) onTap;

  const BottomNavBar({
    Key? key,
    required this.currentIndex,
    required this.onTap,
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Container(
      height: 65,
      width: double.infinity,
      color: Colors.transparent,
      child: Stack(
        alignment: Alignment.bottomCenter,
        children: [
          Container(
            height: 55,
            color: whiteColor,
          ),
          Row(
            mainAxisAlignment: MainAxisAlignment.spaceEvenly,
            children: [
              _buildNavItem(0, Icons.bar_chart, 'Stats'),
              _buildNavItem(1, Icons.home, 'Home'),
              _buildNavItem(2, Icons.person, 'Profile'),
            ],
          ),
        ],
      ),
    );
  }

  Widget _buildNavItem(int index, IconData icon, String label) {
    bool isSelected = currentIndex == index;
    
    return InkWell(
      onTap: () => onTap(index),
      child: Column(
        mainAxisAlignment: MainAxisAlignment.center,
        children: [
          SizedBox(
            width: 60,
            height: 60,
            child: Stack(
              alignment: Alignment.center,
              children: [
                // Selected indicator dot
                if (isSelected)
                  Positioned(
                    bottom: 8,
                    child: Container(
                      width: 5,
                      height: 5,
                      decoration: BoxDecoration(
                        color: index == 1 ? orangeColor : backgroundColor,
                        shape: BoxShape.circle,
                      ),
                    ),
                  ),
                
                // Icon
                Icon(
                  icon,
                  size: 24,
                  color: isSelected 
                      ? (index == 1 ? orangeColor : backgroundColor) 
                      : Colors.grey,
                ),
              ],
            ),
          ),
        ],
      ),
    );
  }
}