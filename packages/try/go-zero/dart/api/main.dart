import 'api.dart';
import '../data/main.dart';

/// main

/// --/from/:name--
///
/// request: Request
/// response: Response
Future apiFromName(
    {Function(Response)? ok,
    Function(String)? fail,
    Function? eventually}) async {
  await apiGet('/from/:name',
  	 ok: (data) {
    if (ok != null) ok(Response.fromJson(data));
  }, fail: fail, eventually: eventually);
}

/// --/--
///
/// request:
/// response: Response
Future api(
    {Function(Response)? ok,
    Function(String)? fail,
    Function? eventually}) async {
  await apiGet('/',
  	 ok: (data) {
    if (ok != null) ok(Response.fromJson(data));
  }, fail: fail, eventually: eventually);
}

