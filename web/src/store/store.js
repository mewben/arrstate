import { createStore, applyMiddleware, compose } from "redux"
import promise from "redux-promise-middleware"
import thunk from "redux-thunk"

import rootReducer from "./reducers"
import errorMiddleware from "./error-middleware"

// Custom error middleware should go before the promise middleware
const middlewares = [thunk, errorMiddleware, promise]

if (process.env.NODE_ENV === `development`) {
  const { createLogger } = require(`redux-logger`)

  middlewares.push(createLogger({ collapsed: true }))
}

const composeEnhancers =
  (typeof window === "object" && window.__REDUX_DEVTOOLS_EXTENSION_COMPOSE__) ||
  compose
const store = createStore(
  rootReducer,
  composeEnhancers(applyMiddleware(...middlewares))
)

export default store
export const dispatch = store.dispatch
