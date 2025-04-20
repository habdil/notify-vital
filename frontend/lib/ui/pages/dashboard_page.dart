import 'package:flutter/material.dart';
import '../../shared/bottom_navbar.dart';
import '../../shared/notification.dart';
import '../widgets/overview_widget.dart';
import '../widgets/walk_widget.dart';
import '../widgets/heart_widget.dart';
import '../widgets/status_widget.dart';
import '../widgets/calories_widget.dart';

class DashboardPage extends StatefulWidget {
  final Map<String, dynamic>? connectedDevice;
  
  const DashboardPage({
    Key? key, 
    this.connectedDevice,
  }) : super(key: key);

  @override
  State<DashboardPage> createState() => _DashboardPageState();
}

class _DashboardPageState extends State<DashboardPage> {
  int _currentIndex = 1;
  int _heartRate = 90;
  int _steps = 935;
  int _calories = 1293;
  String _activityStatus = 'cukup gerak';
  
  @override
  void initState() {
    super.initState();
    
    // Update data if we have a connected device
    if (widget.connectedDevice != null) {
      // In a real app, this would initialize device communication
      // For demo, we just show a notification
      Future.delayed(const Duration(milliseconds: 500), () {
        if (mounted) {
          NotificationHelper.showSuccessNotification(
            context, 
            'Connected to ${widget.connectedDevice!['name']}'
          );
        }
      });
      
      // Simulate heart rate data after a delay
      Future.delayed(const Duration(seconds: 3), () {
        if (mounted) {
          setState(() {
            _heartRate = 92; // Slight change in heart rate
          });
          NotificationHelper.showInfoNotification(
            context,
            'Heart rate data synced from device'
          );
        }
      });
      
      // Simulate activity update after another delay
      Future.delayed(const Duration(seconds: 7), () {
        if (mounted) {
          setState(() {
            _steps = 941; // Increment steps
          });
          NotificationHelper.showSuccessNotification(
            context,
            'Activity data updated'
          );
        }
      });
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Stack(
        children: [
          // Background image
          Image.asset(
            'assets/bg_dashboard.png',
            width: double.infinity,
            height: double.infinity,
            fit: BoxFit.cover,
          ),
          
          // Main content
          SafeArea(
            child: Column(
              children: [
                // Overview widget at the top
                Padding(
                  padding: const EdgeInsets.all(15),
                  child: OverviewWidget(
                    heartRate: _heartRate,
                    caloriesPercentage: _calories / 2000, // Assuming 2000 is the daily goal
                    stepsPercentage: _steps / 10000, // Assuming 10000 is the daily goal
                    activityStatus: _activityStatus,
                  ),
                ),
                
                // Grid of widgets
                Expanded(
                  child: Padding(
                    padding: const EdgeInsets.symmetric(horizontal: 15),
                    child: GridView(
                      physics: const NeverScrollableScrollPhysics(),
                      gridDelegate: const SliverGridDelegateWithFixedCrossAxisCount(
                        crossAxisCount: 2,
                        childAspectRatio: 0.85,
                        crossAxisSpacing: 15,
                        mainAxisSpacing: 15,
                      ),
                      children: [
                        // First row
                        WalkWidget(steps: _steps),
                        HeartWidget(heartRate: _heartRate),
                        
                        // Second row
                        StatusWidget(
                          status: _activityStatus,
                          value: _activityStatus == 'kurang gerak' 
                              ? 0.2 
                              : (_activityStatus == 'cukup gerak' ? 0.5 : 0.8),
                        ),
                        CaloriesWidget(calories: _calories),
                      ],
                    ),
                  ),
                ),
              ],
            ),
          ),
        ],
      ),
      
      // Bottom navigation bar
      bottomNavigationBar: BottomNavBar(
        currentIndex: _currentIndex,
        onTap: (index) {
          setState(() {
            _currentIndex = index;
          });
        },
      ),
    );
  }
}