import 'dart:convert';
import 'package:http/http.dart' as http;
import 'package:flutter_dotenv/flutter_dotenv.dart';
import '../models/user_model.dart';

class AuthService {
  // Get base URL from .env file
  final String baseUrl = dotenv.env['API_URL'] ?? '';

  // User login
  Future<UserModel> login({
    required String email,
    required String password,
  }) async {
    try {
      final response = await http.post(
        Uri.parse('$baseUrl/auth/login'),
        headers: {
          'Content-Type': 'application/json',
        },
        body: jsonEncode({
          'email': email,
          'password': password,
        }),
      );

      if (response.statusCode == 200) {
        final data = jsonDecode(response.body);
        UserModel user = UserModel.fromJson(data);
        
        // Save user data and token
        await user.saveUserData();
        
        return user;
      } else {
        final data = jsonDecode(response.body);
        throw Exception(data['message'] ?? 'Login failed');
      }
    } catch (e) {
      throw Exception('Error connecting to server: $e');
    }
  }

  // User registration
  Future<UserModel> register({
    required String username,
    required String email,
    required String password,
  }) async {
    try {
      final response = await http.post(
        Uri.parse('$baseUrl/auth/register'),
        headers: {
          'Content-Type': 'application/json',
        },
        body: jsonEncode({
          'username': username,
          'email': email,
          'password': password,
        }),
      );

      if (response.statusCode == 201) {
        final data = jsonDecode(response.body);
        UserModel user = UserModel.fromJson(data);
        
        // Save user data and token
        await user.saveUserData();
        
        return user;
      } else {
        final data = jsonDecode(response.body);
        throw Exception(data['message'] ?? 'Registration failed');
      }
    } catch (e) {
      throw Exception('Error connecting to server: $e');
    }
  }

  // Check if user is logged in
  Future<bool> isLoggedIn() async {
    try {
      final user = await UserModel.getUserFromPreferences();
      return user != null;
    } catch (e) {
      return false;
    }
  }

  // Logout user
  Future<bool> logout() async {
    try {
      await UserModel.clearUserData();
      return true;
    } catch (e) {
      return false;
    }
  }
}