import 'package:flutter/material.dart';
import '../../shared/theme.dart';
import '../../shared/notification.dart';
import '../pages/pairing_page.dart';
import '../../services/auth_service.dart';

class LoginWidget extends StatefulWidget {
  final VoidCallback onForgotPassword;
  final VoidCallback onToggleAuthMode;

  const LoginWidget({
    Key? key,
    required this.onForgotPassword,
    required this.onToggleAuthMode,
  }) : super(key: key);

  @override
  State<LoginWidget> createState() => _LoginWidgetState();
}

class _LoginWidgetState extends State<LoginWidget> {
  final TextEditingController _emailController = TextEditingController();
  final TextEditingController _passwordController = TextEditingController();
  final AuthService _authService = AuthService();
  bool _obscureText = true;
  bool _isLoading = false;

  @override
  void dispose() {
    _emailController.dispose();
    _passwordController.dispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return Container(
      padding: const EdgeInsets.all(24),
      decoration: BoxDecoration(
        color: whiteColor,
        borderRadius: BorderRadius.circular(20),
      ),
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          // Email Field
          Text(
            'Email',
            style: darkTextStyle.copyWith(
              fontSize: 14,
              fontWeight: medium,
            ),
          ),
          const SizedBox(height: 8),
          Container(
            height: 50,
            padding: const EdgeInsets.symmetric(horizontal: 16),
            decoration: BoxDecoration(
              color: whiteColor,
              borderRadius: BorderRadius.circular(30),
              border: Border.all(
                color: Colors.blue,
                width: 1,
              ),
            ),
            child: Center(
              child: TextField(
                controller: _emailController,
                decoration: InputDecoration.collapsed(
                  hintText: 'Your Email',
                  hintStyle: greyTextStyle.copyWith(
                    fontSize: 14,
                  ),
                ),
                style: darkTextStyle.copyWith(
                  fontSize: 14,
                ),
                keyboardType: TextInputType.emailAddress,
              ),
            ),
          ),
          const SizedBox(height: 16),
          
          // Password Field
          Text(
            'Password',
            style: darkTextStyle.copyWith(
              fontSize: 14,
              fontWeight: medium,
            ),
          ),
          const SizedBox(height: 8),
          Container(
            height: 50,
            padding: const EdgeInsets.symmetric(horizontal: 16),
            decoration: BoxDecoration(
              color: whiteColor,
              borderRadius: BorderRadius.circular(30),
              border: Border.all(
                color: greenColor,
                width: 1,
              ),
            ),
            child: Center(
              child: Row(
                children: [
                  Expanded(
                    child: TextField(
                      controller: _passwordController,
                      obscureText: _obscureText,
                      decoration: InputDecoration.collapsed(
                        hintText: 'Your Password',
                        hintStyle: greyTextStyle.copyWith(
                          fontSize: 14,
                        ),
                      ),
                      style: darkTextStyle.copyWith(
                        fontSize: 14,
                      ),
                    ),
                  ),
                  IconButton(
                    icon: Icon(
                      _obscureText ? Icons.visibility_off : Icons.visibility,
                      color: greyColor,
                    ),
                    onPressed: () {
                      setState(() {
                        _obscureText = !_obscureText;
                      });
                    },
                  ),
                ],
              ),
            ),
          ),
          const SizedBox(height: 24),
          
          // Login Button
          Center(
            child: SizedBox(
              width: 150,
              height: 45,
              child: ElevatedButton(
                onPressed: _isLoading 
                  ? null 
                  : () async {
                      // Validation
                      if (_emailController.text.isEmpty || _passwordController.text.isEmpty) {
                        NotificationHelper.showErrorNotification(
                          context,
                          'Please fill in all fields',
                        );
                        return;
                      }
                      
                      setState(() {
                        _isLoading = true;
                      });
                      
                      try {
                        // Call login API
                        final user = await _authService.login(
                          email: _emailController.text,
                          password: _passwordController.text,
                        );
                        
                        if (mounted) {
                          // Success login
                          NotificationHelper.showSuccessNotification(
                            context,
                            'Login successful',
                          );
                          
                          // Navigate to pairing page after successful login
                          Navigator.pushReplacement(
                            context,
                            MaterialPageRoute(builder: (context) => const PairingPage()),
                          );
                        }
                      } catch (e) {
                        if (mounted) {
                          // Show error notification
                          NotificationHelper.showErrorNotification(
                            context,
                            e.toString().replaceAll('Exception: ', ''),
                          );
                        }
                      } finally {
                        if (mounted) {
                          setState(() {
                            _isLoading = false;
                          });
                        }
                      }
                    },
                style: ElevatedButton.styleFrom(
                  backgroundColor: Colors.blue,
                  foregroundColor: whiteColor,
                  shape: RoundedRectangleBorder(
                    borderRadius: BorderRadius.circular(30),
                  ),
                ),
                child: _isLoading
                  ? const SizedBox(
                      width: 20,
                      height: 20,
                      child: CircularProgressIndicator(
                        color: Colors.white,
                        strokeWidth: 2,
                      ),
                    )
                  : Text(
                      'Login',
                      style: whiteTextStyle.copyWith(
                        fontSize: 16,
                        fontWeight: semiBold,
                      ),
                    ),
              ),
            ),
          ),
          const SizedBox(height: 16),
          
          // Forgot Password
          Center(
            child: TextButton(
              onPressed: widget.onForgotPassword,
              child: Text(
                'Lupa password?',
                style: greenTextStyle.copyWith(
                  fontSize: 14,
                  fontWeight: medium,
                ),
              ),
            ),
          ),
          
          // Sign Up Link
          Center(
            child: TextButton(
              onPressed: widget.onToggleAuthMode,
              child: Text(
                'Don\'t have an account? Sign Up',
                style: darkTextStyle.copyWith(
                  fontSize: 14,
                  fontWeight: medium,
                ),
              ),
            ),
          ),
        ],
      ),
    );
  }
}