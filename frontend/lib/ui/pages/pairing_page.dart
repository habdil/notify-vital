import 'package:flutter/material.dart';
import '../../shared/theme.dart';
import '../../shared/notification.dart';

class PairingPage extends StatefulWidget {
  const PairingPage({super.key});

  @override
  State<PairingPage> createState() => _PairingPageState();
}

class _PairingPageState extends State<PairingPage> {
  bool _isScanning = false;
  List<Map<String, dynamic>> _devicesList = [];

  @override
  void initState() {
    super.initState();
    // Pre-populate with some sample devices for UI demonstration
    _devicesList = [
      {
        'name': 'Vital Watch SE',
        'address': 'A4:C1:38:8D:12:E5',
        'rssi': -65,
        'connected': false,
      },
      {
        'name': 'HeartGuard Pro',
        'address': 'B8:27:EB:F3:1A:C2',
        'rssi': -72,
        'connected': false,
      },
    ];
  }

  // Simulate scanning for devices
  void _startScan() {
    setState(() {
      _isScanning = true;
    });

    // Simulate device discovery with a delay
    Future.delayed(const Duration(seconds: 3), () {
      if (mounted) {
        setState(() {
          _isScanning = false;
          
          // Add a new "discovered" device
          _devicesList.add({
            'name': 'CardioSense GT',
            'address': 'C2:4A:7B:9D:5E:F1',
            'rssi': -58,
            'connected': false,
          });
        });
      }
    });
  }

  // Simulate pairing with a device
  void _connectToDevice(int index) {
    setState(() {
      // Reset all connections
      for (var device in _devicesList) {
        device['connected'] = false;
      }
      
      // Set the selected device as connected
      _devicesList[index]['connected'] = true;
    });

    // Show success notification
    NotificationHelper.showSuccessNotification(
      context,
      'Successfully paired with ${_devicesList[index]['name']}',
    );
  }

  // Get signal strength icon based on RSSI value
  Icon _getSignalIcon(int rssi) {
    if (rssi > -70) {
      return Icon(Icons.signal_cellular_4_bar, color: greenColor);
    } else if (rssi > -80) {
      return Icon(Icons.signal_cellular_alt_2_bar, color: Colors.orange);
    } else {
      return const Icon(Icons.signal_cellular_alt_1_bar, color: Colors.red);
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: backgroundColor,
      appBar: AppBar(
        title: Text(
          'Pair Your Device',
          style: whiteTextStyle.copyWith(
            fontSize: 20,
            fontWeight: semiBold,
          ),
        ),
        backgroundColor: backgroundColor,
        elevation: 0,
      ),
      body: Column(
        children: [
          // Header and explanation
          Container(
            padding: const EdgeInsets.all(16),
            margin: const EdgeInsets.all(16),
            decoration: BoxDecoration(
              color: whiteColor,
              borderRadius: BorderRadius.circular(16),
              boxShadow: [
                BoxShadow(
                  color: Colors.black.withOpacity(0.1),
                  blurRadius: 8,
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
                      Icons.watch,
                      color: greenColor,
                      size: 36,
                    ),
                    const SizedBox(width: 12),
                    Text(
                      'Connect Smartwatch',
                      style: darkTextStyle.copyWith(
                        fontSize: 18,
                        fontWeight: bold,
                      ),
                    ),
                  ],
                ),
                const SizedBox(height: 12),
                Text(
                  'Pairing your smartwatch allows Notify Vital to monitor your heart health in real-time and provide personalized alerts.',
                  style: darkTextStyle.copyWith(
                    fontSize: 14,
                  ),
                ),
                const SizedBox(height: 8),
                Text(
                  '1. Make sure your smartwatch is turned on and Bluetooth is enabled',
                  style: darkTextStyle.copyWith(
                    fontSize: 12,
                  ),
                ),
                Text(
                  '2. Keep your smartwatch within 10 meters of your phone',
                  style: darkTextStyle.copyWith(
                    fontSize: 12,
                  ),
                ),
                Text(
                  '3. Tap the "Scan for Devices" button below',
                  style: darkTextStyle.copyWith(
                    fontSize: 12,
                  ),
                ),
              ],
            ),
          ),
          
          // Scan button
          Padding(
            padding: const EdgeInsets.symmetric(horizontal: 16),
            child: ElevatedButton.icon(
              onPressed: _isScanning ? null : _startScan,
              style: ElevatedButton.styleFrom(
                backgroundColor: greenColor,
                foregroundColor: whiteColor,
                minimumSize: const Size(double.infinity, 50),
                shape: RoundedRectangleBorder(
                  borderRadius: BorderRadius.circular(12),
                ),
              ),
              icon: _isScanning 
                ? const SizedBox(
                    width: 20,
                    height: 20,
                    child: CircularProgressIndicator(
                      color: Colors.white,
                      strokeWidth: 2,
                    ),
                  )
                : const Icon(Icons.bluetooth_searching),
              label: Text(
                _isScanning ? 'Scanning...' : 'Scan for Devices',
                style: whiteTextStyle.copyWith(
                  fontSize: 16,
                  fontWeight: semiBold,
                ),
              ),
            ),
          ),
          
          // Available devices list
          Padding(
            padding: const EdgeInsets.all(16),
            child: Column(
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                Text(
                  'Available Devices',
                  style: whiteTextStyle.copyWith(
                    fontSize: 16,
                    fontWeight: semiBold,
                  ),
                ),
                const SizedBox(height: 8),
                if (_devicesList.isEmpty && !_isScanning)
                  Center(
                    child: Padding(
                      padding: const EdgeInsets.all(24),
                      child: Text(
                        'No devices found. Try scanning again.',
                        style: greyTextStyle,
                      ),
                    ),
                  )
                else
                  ListView.builder(
                    shrinkWrap: true,
                    physics: const NeverScrollableScrollPhysics(),
                    itemCount: _devicesList.length,
                    itemBuilder: (context, index) {
                      final device = _devicesList[index];
                      return Container(
                        margin: const EdgeInsets.only(bottom: 12),
                        decoration: BoxDecoration(
                          color: whiteColor,
                          borderRadius: BorderRadius.circular(12),
                        ),
                        child: ListTile(
                          leading: device['connected']
                            ? Icon(Icons.bluetooth_connected, color: greenColor, size: 28)
                            : _getSignalIcon(device['rssi']),
                          title: Text(
                            device['name'],
                            style: darkTextStyle.copyWith(
                              fontWeight: semiBold,
                            ),
                          ),
                          subtitle: Text(
                            device['address'],
                            style: greyTextStyle.copyWith(
                              fontSize: 12,
                            ),
                          ),
                          trailing: device['connected']
                            ? TextButton(
                                onPressed: () {},
                                child: Text(
                                  'Connected',
                                  style: greenTextStyle.copyWith(
                                    fontWeight: semiBold,
                                  ),
                                ),
                              )
                            : ElevatedButton(
                                onPressed: () => _connectToDevice(index),
                                style: ElevatedButton.styleFrom(
                                  backgroundColor: greenColor,
                                  foregroundColor: whiteColor,
                                  shape: RoundedRectangleBorder(
                                    borderRadius: BorderRadius.circular(20),
                                  ),
                                  padding: const EdgeInsets.symmetric(horizontal: 12),
                                ),
                                child: const Text('Connect'),
                              ),
                        ),
                      );
                    },
                  ),
              ],
            ),
          ),
          
          // Help text at bottom
          const Spacer(),
          Padding(
            padding: const EdgeInsets.all(16),
            child: Row(
              mainAxisAlignment: MainAxisAlignment.center,
              children: [
                Icon(
                  Icons.help_outline,
                  color: greyColor,
                  size: 20,
                ),
                const SizedBox(width: 8),
                Text(
                  'Need help with pairing?',
                  style: greyTextStyle.copyWith(
                    fontSize: 14,
                  ),
                ),
                TextButton(
                  onPressed: () {
                    // Show help dialog or navigate to help page
                  },
                  child: Text(
                    'Get Support',
                    style: greenTextStyle.copyWith(
                      fontWeight: semiBold,
                    ),
                  ),
                ),
              ],
            ),
          ),
        ],
      ),
    );
  }
}