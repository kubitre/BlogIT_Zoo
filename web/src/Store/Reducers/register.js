const initState = {
    userName: "",
    password: "",
    email: "",
}

export function RSI_registerState(state = initState, action){
    const {type, payload} = action;

    switch(type){
        default:
            return state
    }
}