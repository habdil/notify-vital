import 'package:flutter/material.dart';
import 'package:google_fonts/google_fonts.dart';

// Color scheme for the Notify Vital app
Color backgroundColor = const Color(0xff2C3E50);
Color whiteColor = const Color(0xffFFFFFF);
Color greenColor = const Color(0xff27AE60);
Color greyColor = const Color(0xffD9D9D9);
Color orangeColor = const Color(0xffF39C12);

// Text styles with Fredoka font
TextStyle whiteTextStyle = GoogleFonts.fredoka(
  color: whiteColor,
);

TextStyle greenTextStyle = GoogleFonts.fredoka(
  color: greenColor,
);

TextStyle greyTextStyle = GoogleFonts.fredoka(
  color: greyColor,
);

TextStyle darkTextStyle = GoogleFonts.fredoka(
  color: backgroundColor,
);

// Font weights
FontWeight light = FontWeight.w300;
FontWeight regular = FontWeight.w400;
FontWeight medium = FontWeight.w500;
FontWeight semiBold = FontWeight.w600;
FontWeight bold = FontWeight.w700;
FontWeight extraBold = FontWeight.w800;
FontWeight black = FontWeight.w900;

// Theme data for the app
ThemeData lightTheme = ThemeData(
  scaffoldBackgroundColor: whiteColor,
  appBarTheme: AppBarTheme(
    backgroundColor: backgroundColor,
    elevation: 0,
    titleTextStyle: whiteTextStyle.copyWith(
      fontSize: 20,
      fontWeight: semiBold,
    ),
    iconTheme: IconThemeData(
      color: whiteColor,
    ),
  ),
  colorScheme: ColorScheme.light(
    primary: backgroundColor,
    secondary: greenColor,
  ),
);

// Dark theme variation
ThemeData darkTheme = ThemeData(
  scaffoldBackgroundColor: backgroundColor,
  appBarTheme: AppBarTheme(
    backgroundColor: backgroundColor,
    elevation: 0,
    titleTextStyle: whiteTextStyle.copyWith(
      fontSize: 20,
      fontWeight: semiBold,
    ),
    iconTheme: IconThemeData(
      color: whiteColor,
    ),
  ),
  colorScheme: ColorScheme.dark(
    primary: greenColor,
    secondary: whiteColor,
  ),
);