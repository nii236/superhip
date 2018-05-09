import {
	DELETE_MANY,
	GET_LIST,
	GET_ONE,
	GET_MANY,
	GET_MANY_REFERENCE,
	CREATE,
	UPDATE,
	DELETE,
	fetchUtils
} from "react-admin"

const { fetchJson, queryParameters } = fetchUtils
const Provider = (apiUrl, httpClient = fetchJson) => {
	/**
	 * @param {String} type One of the constants appearing at the top if this file, e.g. 'UPDATE'
	 * @param {String} resource Name of the resource to fetch, e.g. 'posts'
	 * @param {Object} params The REST request params, depending on the type
	 * @returns {Object} { url, options } The HTTP request parameters
	 */
	const convertRESTRequestToHTTP = (type, resource, params) => {
		let url = ""
		const options = {}
		options.headers = new Headers()
		switch (type) {
			case GET_LIST: {
				options.method = "POST"
				options.body = JSON.stringify(params)
				url = `${apiUrl}/${resource}/get/list`
				break
			}
			case GET_ONE:
				options.method = "POST"
				options.body = JSON.stringify(params)
				url = `${apiUrl}/${resource}/get`
				break
			case GET_MANY: {
				options.method = "POST"
				options.body = JSON.stringify(params)
				url = `${apiUrl}/${resource}/get/many`
				break
			}
			case GET_MANY_REFERENCE: {
				options.method = "POST"
				options.body = JSON.stringify(params)
				url = `${apiUrl}/${resource}/get/many/reference`
				break
			}
			case UPDATE:
				options.method = "POST"
				options.body = JSON.stringify(params)
				url = `${apiUrl}/${resource}/update`
				break
			case CREATE:
				options.method = "POST"
				options.body = JSON.stringify(params)
				url = `${apiUrl}/${resource}/create`
				break
			case DELETE:
				options.method = "POST"
				options.body = JSON.stringify(params)
				url = `${apiUrl}/${resource}/delete`
				break
			case DELETE_MANY:
				options.method = "POST"
				options.body = JSON.stringify(params)
				url = `${apiUrl}/${resource}/delete/many`
				break
			default:
				throw new Error(`Unsupported fetch action type ${type}`)
		}
		return { url, options }
	}

	/**
	 * @param {Object} response HTTP response from fetch()
	 * @param {String} type One of the constants appearing at the top if this file, e.g. 'UPDATE'
	 * @param {String} resource Name of the resource to fetch, e.g. 'posts'
	 * @param {Object} params The REST request params, depending on the type
	 * @returns {Object} REST response
	 */
	const convertHTTPResponseToREST = (response, type, resource, params) => {
		const { headers, json } = response
		switch (type) {
			case GET_LIST:
				return { total: json.total, data: json.data }
			case GET_ONE:
				return { data: json.data[0] }
			case CREATE:
				return { data: json.data[0] }
			case UPDATE:
				return { data: json.data[0] }
			default:
				return json
		}
	}

	/**
	 * @param {string} type Request type, e.g GET_LIST
	 * @param {string} resource Resource name, e.g. "posts"
	 * @param {Object} payload Request parameters. Depends on the request type
	 * @returns {Promise} the Promise for a REST response
	 */
	return (type, resource, params) => {
		console.log("type", type)
		console.log("resource", resource)
		console.log("params", params)
		const { url, options } = convertRESTRequestToHTTP(type, resource, params)
		return httpClient(url, options).then(response => convertHTTPResponseToREST(response, type, resource, params))
	}
}

export default { Provider }
