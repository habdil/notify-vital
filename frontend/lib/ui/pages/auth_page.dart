import 'package:flutter/material.dart';
import '../../shared/theme.dart';
import '../widgets/login_widget.dart';
import '../widgets/register_widget.dart';

class AuthPage extends StatefulWidget {
  final bool initialIsLogin;
  
  const AuthPage({
    Key? key, 
    this.initialIsLogin = true,
  }) : super(key: key);

  @override
  State<AuthPage> createState() => _AuthPageState();
}

class _AuthPageState extends State<AuthPage> {
  late bool isLogin;

  @override
  void initState() {
    super.initState();
    isLogin = widget.initialIsLogin;
  }

  void toggleAuthMode() {
    setState(() {
      isLogin = !isLogin;
    });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: backgroundColor,
      body: Container(
        width: double.infinity,
        height: double.infinity,
        decoration: const BoxDecoration(
          image: DecorationImage(
            image: AssetImage('assets/bg_auth.png'),
            fit: BoxFit.cover,
          ),
        ),
        child: SafeArea(
          child: SingleChildScrollView(
            child: Padding(
              padding: const EdgeInsets.symmetric(horizontal: 24.0),
              child: Column(
                crossAxisAlignment: CrossAxisAlignment.center,
                children: [
                  const SizedBox(height: 80),
                  Text(
                    isLogin ? 'LOGIN' : 'SIGN UP',
                    style: whiteTextStyle.copyWith(
                      fontSize: 28,
                      fontWeight: bold,
                      letterSpacing: 2.0,
                    ),
                  ),
                  const SizedBox(height: 30),
                  isLogin 
                    ? LoginWidget(onForgotPassword: () {
                        // Handle forgot password
                      }, onToggleAuthMode: toggleAuthMode)
                    : RegisterWidget(onToggleAuthMode: toggleAuthMode),
                ],
              ),
            ),
          ),
        ),
      ),
    );
  }
}