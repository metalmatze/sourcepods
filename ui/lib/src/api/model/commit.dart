part of swagger.api;

class Commit {
  
  String hash = null;
  

  String tree = null;
  

  String parent = null;
  

  String message = null;
  

  String body = null;
  

  CommitAuthor author = null;
  

  CommitAuthor committer = null;
  
  Commit();

  @override
  String toString() {
    return 'Commit[hash=$hash, tree=$tree, parent=$parent, message=$message, body=$body, author=$author, committer=$committer, ]';
  }

  Commit.fromJson(Map<String, dynamic> json) {
    if (json == null) return;
    hash =
        json['hash']
    ;
    tree =
        json['tree']
    ;
    parent =
        json['parent']
    ;
    message =
        json['message']
    ;
    body =
        json['body']
    ;
    author =
      
      
      new CommitAuthor.fromJson(json['author'])
;
    committer =
      
      
      new CommitAuthor.fromJson(json['committer'])
;
  }

  Map<String, dynamic> toJson() {
    return {
      'hash': hash,
      'tree': tree,
      'parent': parent,
      'message': message,
      'body': body,
      'author': author,
      'committer': committer
     };
  }

  static List<Commit> listFromJson(List<dynamic> json) {
    return json == null ? new List<Commit>() : json.map((value) => new Commit.fromJson(value)).toList();
  }

  static Map<String, Commit> mapFromJson(Map<String, Map<String, dynamic>> json) {
    var map = new Map<String, Commit>();
    if (json != null && json.length > 0) {
      json.forEach((String key, Map<String, dynamic> value) => map[key] = new Commit.fromJson(value));
    }
    return map;
  }
}

