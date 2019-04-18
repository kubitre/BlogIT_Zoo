import { LOGIN_USER_BY_CREDENTIALS, LOGIN_USER_LOADING, LOGIN_USER_FAILED, LOGIN_BY_LOCAL_STORAGE } from "../Constants/user";
import { ConfigureRequest } from "../API/requestBuilder";
import { LOGIN_URL } from "../API/main";

export function loginByCredentialsSuccess(data) {
    return {
        "type": LOGIN_USER_BY_CREDENTIALS,
        "payload": data
    }
}

export function fetchDataFromLS() {
    return dispatch => {
        let username = window.localStorage.getItem("username");
        let token = window.localStorage.getItem("token");
        let userid = window.localStorage.getItem("userid");

        dispatch(loginByLocalStorage({ username: username, token: token, userid: userid }))
    }
}

export function loginByLocalStorage(data) {
    return {
        "type": LOGIN_BY_LOCAL_STORAGE,
        "payload": data
    }
}

export function loginLoading(status) {
    return {
        "type": LOGIN_USER_LOADING,
        "payload": status
    }
}

export function loginFailed(status) {
    return {
        "type": LOGIN_USER_FAILED,
        "payload": status
    }
}

export function startFetchingData(data) {
    return dispatch => {
        dispatch(loginLoading(true));

        let request_packet = ConfigureRequest(LOGIN_URL, "POST", JSON.stringify({
            "username": data.username,
            "password": data.password,
        }))

        fetch(request_packet)
            .then(response => {
                if (!response.ok) {
                    dispatch(loginFailed(true));
                    throw Error("error request");
                }

                return response
            })
            .then(response => response.json())
            .then(respInJson => dispatch(loginByCredentialsSuccess(respInJson)))
            .catch(err => dispatch(loginFailed(true)))
    }
}