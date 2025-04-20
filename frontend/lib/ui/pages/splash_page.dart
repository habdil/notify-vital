import 'dart:async';
import 'package:flutter/material.dart';
import '../../shared/theme.dart';
import 'onboarding_page.dart';

class SplashPage extends StatefulWidget {
  const SplashPage({super.key});

  @override
  State<SplashPage> createState() => _SplashPageState();
}

class _SplashPageState extends State<SplashPage> {
  @override
  void initState() {
    super.initState();
    
    // Add a timer to navigate to the onboarding screen after a delay
    Timer(const Duration(seconds: 3), () {
      Navigator.pushAndRemoveUntil(
        context,
        MaterialPageRoute(
          builder: (context) => const OnboardingPage(),
        ),
        (route) => false,
      );
    });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: backgroundColor,
      body: Center(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            Container(
              width: 150,
              height: 150,
              decoration: const BoxDecoration(
                image: DecorationImage(
                  image: AssetImage('assets/ic_logo.png'),
                ),
              ),
            ),
            const SizedBox(
              height: 24,
            ),
            Text(
              'NOTIFY VITAL',
              style: whiteTextStyle.copyWith(
                fontSize: 24,
                fontWeight: bold,
                letterSpacing: 2.0,
              ),
            ),
            const SizedBox(
              height: 8,
            ),
            Text(
              'Your Heart Health Companion',
              style: greyTextStyle.copyWith(
                fontSize: 16,
                fontWeight: medium,
              ),
            ),
          ],
        ),
      ),
    );
  }
}