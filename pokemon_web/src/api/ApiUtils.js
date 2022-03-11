// handleresponse function is used to handle the response from the server
export function handleResponse(response) {
  if (
    response.status === 200 ||
    response.status === 202 ||
    response.statusText === "OK" ||
    response.statusText === "Created"
  )
    return response.data;
  if (response.status === 400) {
    // So, a server-side validation error occurred.
    // Server side validation returns a string error message, so parse as text instead of json.
    const error = response.statusText();
    throw new Error(error);
  }
  throw new Error("Network response was not ok.");
}

// In a real app, would likely call an error logging service.
export function handleError(error) {
  // eslint-disable-next-line no-console
  console.error("API call failed. " + error);
  return error && error.response && error.response.data
    ? error.response.data
    : error;
}
