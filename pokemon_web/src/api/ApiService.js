import axios from "axios";
import { handleError, handleResponse } from "./ApiUtils";

// httpRequest is a wrapper over axios
const httpRequest = (method, url, request, _headers) => {
  return axios({
    method,
    url,
    data: request,
    headers: _headers,
  })
    .then((res) => {
      const result = handleResponse(res);
      return Promise.resolve(result);
    })
    .catch((err) => {
      //throw handleError(err);
      return Promise.reject(handleError(err));
    });
};

// get is a wrapper over axios get method
const get = (url, request, headers) => {
  let queryString = "";
  if (request && Object.keys(request).length > 0) {
    queryString += "?";
    let len = Object.keys(request).length,
      cnt = 0;
    for (let key in request) {
      cnt++;
      queryString += `${key}=${request[key].toString()}`;
      if (len > cnt) queryString += "&";
    }
  }
  return httpRequest("get", `${url}${queryString}`, null, headers);
};

// delete is a wrapper over axios delete method
const deleteRequest = (url, request, headers) => {
  return httpRequest("delete", url, request, headers);
};

// post is a wrapper over axios post method
const post = (url, request, headers) => {
  return httpRequest("post", url, request, headers);
};

// put is a wrapper over axios put method
const put = (url, request, headers) => {
  return httpRequest("put", url, request, headers);
};

// patch is a wrapper over axios patch method
const patch = (url, request, headers) => {
  return httpRequest("patch", url, request, headers);
};


const Api = {
  get,
  delete: deleteRequest,
  post,
  put,
  patch,
};

export default Api;
