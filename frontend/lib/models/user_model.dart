import 'dart:convert';
import 'package:shared_preferences/shared_preferences.dart';

class UserModel {
  final int? id;
  final String username;
  final String email;
  final String token;
  final String expiresAt;
  final String? createdAt;
  final String? lastLogin;
  final bool isActive;

  UserModel({
    this.id,
    required this.username,
    required this.email,
    required this.token,
    required this.expiresAt,
    this.createdAt,
    this.lastLogin,
    this.isActive = true,
  });

  factory UserModel.fromJson(Map<String, dynamic> json) {
    return UserModel(
      id: json['user']['user_id'],
      username: json['user']['username'],
      email: json['user']['email'],
      token: json['token'],
      expiresAt: json['expires_at'],
      createdAt: json['user']['created_at'],
      lastLogin: json['user']['last_login'],
      isActive: json['user']['is_active'] ?? true,
    );
  }

  Map<String, dynamic> toJson() {
    return {
      'user': {
        'user_id': id,
        'username': username,
        'email': email,
        'created_at': createdAt,
        'last_login': lastLogin,
        'is_active': isActive,
      },
      'token': token,
      'expires_at': expiresAt,
    };
  }

  // Save user data to SharedPreferences
  Future<bool> saveUserData() async {
    try {
      final SharedPreferences prefs = await SharedPreferences.getInstance();
      final userData = jsonEncode(toJson());
      
      // Save user data
      await prefs.setString('user_data', userData);
      
      // Save token separately for easy access
      await prefs.setString('token', token);
      
      return true;
    } catch (e) {
      print('Error saving user data: $e');
      return false;
    }
  }

  // Get user from SharedPreferences
  static Future<UserModel?> getUserFromPreferences() async {
    try {
      final SharedPreferences prefs = await SharedPreferences.getInstance();
      final userData = prefs.getString('user_data');
      
      if (userData != null) {
        final userMap = jsonDecode(userData);
        return UserModel.fromJson(userMap);
      }
      
      return null;
    } catch (e) {
      print('Error getting user data: $e');
      return null;
    }
  }

  // Get token from SharedPreferences
  static Future<String?> getToken() async {
    try {
      final SharedPreferences prefs = await SharedPreferences.getInstance();
      return prefs.getString('token');
    } catch (e) {
      print('Error getting token: $e');
      return null;
    }
  }

  // Clear user data from SharedPreferences
  static Future<bool> clearUserData() async {
    try {
      final SharedPreferences prefs = await SharedPreferences.getInstance();
      await prefs.remove('user_data');
      await prefs.remove('token');
      return true;
    } catch (e) {
      print('Error clearing user data: $e');
      return false;
    }
  }
}