import 'package:flutter/material.dart';
import '../../shared/theme.dart';
import '../../shared/notification.dart';
import '../../services/auth_service.dart';

class RegisterWidget extends StatefulWidget {
  final VoidCallback onToggleAuthMode;

  const RegisterWidget({
    Key? key,
    required this.onToggleAuthMode,
  }) : super(key: key);

  @override
  State<RegisterWidget> createState() => _RegisterWidgetState();
}

class _RegisterWidgetState extends State<RegisterWidget> {
  final TextEditingController _usernameController = TextEditingController();
  final TextEditingController _emailController = TextEditingController();
  final TextEditingController _passwordController = TextEditingController();
  final TextEditingController _confirmPasswordController = TextEditingController();
  final AuthService _authService = AuthService();
  bool _obscurePassword = true;
  bool _obscureConfirmPassword = true;
  bool _isLoading = false;

  @override
  void dispose() {
    _usernameController.dispose();
    _emailController.dispose();
    _passwordController.dispose();
    _confirmPasswordController.dispose();
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
          // Username Field
          Text(
            'Username',
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
                controller: _usernameController,
                decoration: InputDecoration.collapsed(
                  hintText: 'Enter username',
                  hintStyle: greyTextStyle.copyWith(
                    fontSize: 14,
                  ),
                ),
                style: darkTextStyle.copyWith(
                  fontSize: 14,
                ),
              ),
            ),
          ),
          const SizedBox(height: 16),

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
                color: greenColor,
                width: 1,
              ),
            ),
            child: Center(
              child: TextField(
                controller: _emailController,
                decoration: InputDecoration.collapsed(
                  hintText: 'Enter email',
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
                color: Colors.blue,
                width: 1,
              ),
            ),
            child: Center(
              child: Row(
                children: [
                  Expanded(
                    child: TextField(
                      controller: _passwordController,
                      obscureText: _obscurePassword,
                      decoration: InputDecoration.collapsed(
                        hintText: 'Enter password',
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
                      _obscurePassword ? Icons.visibility_off : Icons.visibility,
                      color: greyColor,
                    ),
                    onPressed: () {
                      setState(() {
                        _obscurePassword = !_obscurePassword;
                      });
                    },
                  ),
                ],
              ),
            ),
          ),
          const SizedBox(height: 16),
          
          // Confirm Password Field
          Text(
            'Confirm Password',
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
                      controller: _confirmPasswordController,
                      obscureText: _obscureConfirmPassword,
                      decoration: InputDecoration.collapsed(
                        hintText: 'Confirm your password',
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
                      _obscureConfirmPassword ? Icons.visibility_off : Icons.visibility,
                      color: greyColor,
                    ),
                    onPressed: () {
                      setState(() {
                        _obscureConfirmPassword = !_obscureConfirmPassword;
                      });
                    },
                  ),
                ],
              ),
            ),
          ),
          const SizedBox(height: 24),
          
          // Register Button
          Center(
            child: SizedBox(
              width: 150,
              height: 45,
              child: ElevatedButton(
                onPressed: _isLoading 
                  ? null 
                  : () async {
                      // Validation
                      if (_usernameController.text.isEmpty ||
                          _emailController.text.isEmpty ||
                          _passwordController.text.isEmpty ||
                          _confirmPasswordController.text.isEmpty) {
                        NotificationHelper.showErrorNotification(
                          context,
                          'Please fill in all fields',
                        );
                        return;
                      }
                      
                      // Password matching validation
                      if (_passwordController.text != _confirmPasswordController.text) {
                        NotificationHelper.showErrorNotification(
                          context,
                          'Passwords do not match',
                        );
                        return;
                      }
                      
                      setState(() {
                        _isLoading = true;
                      });
                      
                      try {
                        // Call register API
                        
                        if (mounted) {
                          // Success registration notification
                          NotificationHelper.showSuccessNotification(
                            context,
                            'Sign up Successful',
                          );
                          
                          // Navigate to login after successful registration
                          widget.onToggleAuthMode();
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
                      'Sign Up',
                      style: whiteTextStyle.copyWith(
                        fontSize: 16,
                        fontWeight: semiBold,
                      ),
                    ),
              ),
            ),
          ),
          const SizedBox(height: 16),
          
          // Login Link
          Center(
            child: TextButton(
              onPressed: widget.onToggleAuthMode,
              child: Text(
                'Already have an account? Login',
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