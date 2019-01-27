part of swagger.api;

class CommitAuthor {
  
  String name = null;
  

  String email = null;
  

  DateTime date = null;
  
  CommitAuthor();

  @override
  String toString() {
    return 'CommitAuthor[name=$name, email=$email, date=$date, ]';
  }

  CommitAuthor.fromJson(Map<String, dynamic> json) {
    if (json == null) return;
    name =
        json['name']
    ;
    email =
        json['email']
    ;
    date = json['date'] == null ? null : DateTime.parse(json['date']);
  }

  Map<String, dynamic> toJson() {
    return {
      'name': name,
      'email': email,
      'date': date == null ? '' : date.toUtc().toIso8601String()
     };
  }

  static List<CommitAuthor> listFromJson(List<dynamic> json) {
    return json == null ? new List<CommitAuthor>() : json.map((value) => new CommitAuthor.fromJson(value)).toList();
  }

  static Map<String, CommitAuthor> mapFromJson(Map<String, Map<String, dynamic>> json) {
    var map = new Map<String, CommitAuthor>();
    if (json != null && json.length > 0) {
      json.forEach((String key, Map<String, dynamic> value) => map[key] = new CommitAuthor.fromJson(value));
    }
    return map;
  }
}

