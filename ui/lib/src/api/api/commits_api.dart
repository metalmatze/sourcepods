part of swagger.api;



class CommitsApi {
  final ApiClient apiClient;

  CommitsApi([ApiClient apiClient]) : apiClient = apiClient ?? defaultApiClient;

  /// Get commits of a repository
  ///
  /// 
  Future<List<Commit>> listCommits(String owner, String name, { String ref }) async {
    Object postBody = null;

    // verify required params are set
    if(owner == null) {
     throw new ApiException(400, "Missing required param: owner");
    }
    if(name == null) {
     throw new ApiException(400, "Missing required param: name");
    }

    // create path and map variables
    String path = "/commits/{owner}/{name}".replaceAll("{format}","json").replaceAll("{" + "owner" + "}", owner.toString()).replaceAll("{" + "name" + "}", name.toString());

    // query params
    List<QueryParam> queryParams = [];
    Map<String, String> headerParams = {};
    Map<String, String> formParams = {};
    if(ref != null) {
      queryParams.addAll(_convertParametersForCollectionFormat("", "ref", ref));
    }
    
    List<String> contentTypes = ["application/json"];

    String contentType = contentTypes.length > 0 ? contentTypes[0] : "application/json";
    List<String> authNames = [];

    if(contentType.startsWith("multipart/form-data")) {
      bool hasFields = false;
      MultipartRequest mp = new MultipartRequest(null, null);
      
      if(hasFields)
        postBody = mp;
    }
    else {
          }

    var response = await apiClient.invokeAPI(path,
                                             'GET',
                                             queryParams,
                                             postBody,
                                             headerParams,
                                             formParams,
                                             contentType,
                                             authNames);

    if(response.statusCode >= 400) {
      throw new ApiException(response.statusCode, response.body);
    } else if(response.body != null) {
      return 
        (apiClient.deserialize(response.body, 'List<Commit>') as List).map((item) => item as Commit).toList();
    } else {
      return null;
    }
  }
}
