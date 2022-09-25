class Tokens {
  /// 用于访问的token, 每次请求都必须带在Header里面
  final String accessToken;
  final int accessExpire;

  /// 用于刷新token
  final String refreshToken;
  final int refreshExpire;
  final int refreshAfter;
  Tokens(
      {required this.accessToken,
      required this.accessExpire,
      required this.refreshToken,
      required this.refreshExpire,
      required this.refreshAfter});
  factory Tokens.fromJson(Map<String, dynamic> m) {
    return Tokens(
        accessToken: m['access_token'],
        accessExpire: m['access_expire'],
        refreshToken: m['refresh_token'],
        refreshExpire: m['refresh_expire'],
        refreshAfter: m['refresh_after']);
  }
  Map<String, dynamic> toJson() {
    return {
      'access_token': accessToken,
      'access_expire': accessExpire,
      'refresh_token': refreshToken,
      'refresh_expire': refreshExpire,
      'refresh_after': refreshAfter,
    };
  }
}
