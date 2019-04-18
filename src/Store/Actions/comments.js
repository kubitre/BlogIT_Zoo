import { GET_COMMENTS_SUCCES, GET_COMMENTS_LOADING, GET_COMMENTS_FAILED, SET_COMMENTS } from "../Constants/comments";
import { ConfigureRequest, ConfigureRequestWithToken } from "../API/requestBuilder";
import { COMMENTS_URL, COMMENT_CREATE_URL } from "../API/main";

export function getCommentsSuccess(comments){
    return {
        "type": GET_COMMENTS_SUCCES,
        "payload": comments
    }
}

export function getCommentsLoading(status){
    return {
        "type": GET_COMMENTS_LOADING,
        "payload": status,
    }
}

export function getCommentsFailed(status){
    return{
        "type": GET_COMMENTS_FAILED,
        "payload": status,
    }
}

export function startFetchCommentsFromBackend(id){
    return dispatch => {
        console.log("START FETCHING: ", id);
        dispatch(getCommentsLoading(true));

        let request_packet = ConfigureRequest(COMMENTS_URL + "/" + id, "GET", null);

        fetch(request_packet)
        .then(response => {
            if (!response.ok){
                dispatch(getCommentsFailed(true));
                throw Error("error request");
            }
            return response;
        })
        .then(response => response.json())
        .then(respInJson => dispatch(getCommentsSuccess(respInJson)))
        .catch(err => dispatch(getCommentsFailed(true)))
    }
}

export function createNewComment(body, token){
    return dispatch => {
        let request_packet = ConfigureRequestWithToken(COMMENT_CREATE_URL, "POST", JSON.stringify(body), token)

        fetch(request_packet)
        .then(response=> {
            if (!response.ok){
                throw Error("error request");
            }

            return response;
        })
        .then(response => response.json())
        .then(response => dispatch(setComments(response)))
        .catch(err => console.log(err))
    }
}

export function setComments(comments){
    return{
        "type": SET_COMMENTS,
        "payload": comments   
    }
}