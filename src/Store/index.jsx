import { createStore, applyMiddleware } from 'redux';
import logger from 'redux-logger';
import rootReducer from './Reducers/index';
import thunk from 'redux-thunk';

export default function configureStore(initialState) {
    return createStore(
        rootReducer,
        initialState,
        applyMiddleware(thunk, logger)
    );
}