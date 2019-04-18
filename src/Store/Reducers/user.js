import { LOGIN_USER_LOADING, LOGIN_USER_FAILED, LOGIN_USER_BY_CREDENTIALS, REGISTER_LOADING, LOGOUT_USER_LOADING, LOGOUT_USER_FAILED, LOGOUT_USER, LOGIN_BY_LOCAL_STORAGE } from "../Constants/user";

const initState = {
    token: null,
    username: "",
    userid: "",
    loading: false,
    failed: false,
    loaded: false,
    loadedByLS: false,
}

export function USI_userState (state = initState, action){
    const {type, payload} = action;

    switch(type){
        case LOGIN_USER_LOADING:
        case LOGOUT_USER_LOADING:
            return {...state, loading: true, loaded: false, failed: false}

        case LOGIN_USER_FAILED:
        case LOGOUT_USER_FAILED:
            return {...state, loading: false, loaded: false, failed: true}

        case LOGIN_USER_BY_CREDENTIALS:
            return {...state, loading: false, loaded: true, failed: false, username: payload.username, token: payload.token, userid: payload.userid}

        case LOGOUT_USER:
            return {...state, loading: false, loaded: true, failed: false, userName: "", token: null}
        

        case LOGIN_BY_LOCAL_STORAGE:
            return {...state, loadedByLS: true, username: payload.username, token: payload.token, userid: payload.userid, loaded: true}
        default:
            return state;
    }
}